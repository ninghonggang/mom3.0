package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	apsPb "github.com/ninghonggang/mom-platform/gen/aps"
	"mom-platform/services/aps-service/internal/config"
	"mom-platform/services/aps-service/internal/handler"
)

// Server wraps the gRPC server with lifecycle management.
type Server struct {
	cfg       *config.Config
	logger    *zap.Logger
	grpcSrv   *grpc.Server
	healthSrv *health.Server
}

// New creates a new gRPC server with interceptors and health check.
func New(cfg *config.Config, h *handler.APSHandler, logger *zap.Logger) *Server {
	// Logging interceptor
	logInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handlerFn grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		resp, err := handlerFn(ctx, req)
		duration := time.Since(start)

		fields := []zap.Field{
			zap.String("method", info.FullMethod),
			zap.Duration("duration", duration),
		}
		if err != nil {
			fields = append(fields, zap.Error(err))
			logger.Error("RPC completed with error", fields...)
		} else {
			logger.Info("RPC completed", fields...)
		}
		return resp, err
	}

	// Recovery interceptor
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) error {
			logger.Error("panic recovered", zap.Any("panic", p))
			return status.Errorf(codes.Internal, "internal error: %v", p)
		}),
	}

	grpcSrv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recoveryOpts...),
			logInterceptor,
		),
	)

	// Register APS services
	apsPb.RegisterMPSServiceServer(grpcSrv, h)
	apsPb.RegisterMRPServiceServer(grpcSrv, h)
	apsPb.RegisterScheduleServiceServer(grpcSrv, h)
	apsPb.RegisterWorkCenterServiceServer(grpcSrv, h)

	// Register health check
	healthSrv := health.NewServer()
	healthSrv.SetServingStatus("aps-service", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcSrv, healthSrv)

	return &Server{
		cfg:       cfg,
		logger:    logger,
		grpcSrv:   grpcSrv,
		healthSrv: healthSrv,
	}
}

// Start starts the gRPC server and blocks until shutdown.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.cfg.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	s.logger.Info("gRPC server starting",
		zap.String("addr", addr),
		zap.String("name", s.cfg.Server.Name),
	)

	// Graceful shutdown on signal
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigCh
		s.logger.Info("received signal, shutting down", zap.String("signal", sig.String()))

		s.healthSrv.SetServingStatus("aps-service", healthpb.HealthCheckResponse_NOT_SERVING)
		stopped := make(chan struct{})
		go func() {
			s.grpcSrv.GracefulStop()
			close(stopped)
		}()

		select {
		case <-stopped:
			s.logger.Info("gRPC server stopped gracefully")
		case <-time.After(10 * time.Second):
			s.logger.Warn("graceful shutdown timeout, forcing stop")
			s.grpcSrv.Stop()
		}
	}()

	if err := s.grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("grpc serve: %w", err)
	}
	return nil
}

// Stop stops the gRPC server.
func (s *Server) Stop() {
	s.logger.Info("stopping APS service")
	s.healthSrv.SetServingStatus("aps-service", healthpb.HealthCheckResponse_NOT_SERVING)
	s.grpcSrv.GracefulStop()
}
