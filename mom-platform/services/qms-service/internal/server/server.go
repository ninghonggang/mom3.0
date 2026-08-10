package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	qms "github.com/ninghonggang/mom-platform/gen/qms"
	"mom-platform/services/qms-service/internal/config"
	"mom-platform/services/qms-service/internal/handler"
)

// loggingInterceptor logs each gRPC unary call.
func loggingInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		log.Info("gRPC call started",
			zap.String("method", info.FullMethod),
		)

		resp, err = handler(ctx, req)

		code := status.Code(err)
		fields := []zap.Field{
			zap.String("method", info.FullMethod),
			zap.Duration("duration", time.Since(start)),
			zap.Stringer("code", code),
		}
		if err != nil {
			fields = append(fields, zap.Error(err))
			log.Error("gRPC call failed", fields...)
		} else {
			log.Info("gRPC call completed", fields...)
		}
		return resp, err
	}
}

// Server wraps the gRPC server with graceful-shutdown support.
type Server struct {
	cfg     *config.Config
	log     *zap.Logger
	handler *handler.Handler
	grpc    *grpc.Server
}

// New creates a new Server.
func New(cfg *config.Config, log *zap.Logger, h *handler.Handler) *Server {
	return &Server{
		cfg:     cfg,
		log:     log,
		handler: h,
	}
}

// Start initialises the gRPC server, registers services, and blocks
// until a shutdown signal is received. It performs graceful shutdown.
func (s *Server) Start(ctx context.Context) error {
	// Register the unary logging interceptor.
	s.grpc = grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor(s.log)),
	)

	// Register the QMS gRPC service via generated proto code.
	qms.RegisterQmsServiceServer(s.grpc, s.handler)

	// Register the gRPC health check service.
	healthServer := health.NewServer()
	healthServer.SetServingStatus("qms-service", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s.grpc, healthServer)

	// Build the listen address.
	addr := fmt.Sprintf(":%d", s.cfg.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", addr, err)
	}

	// Channel to capture serve errors.
	serveErr := make(chan error, 1)
	go func() {
		s.log.Info("gRPC server starting",
			zap.String("addr", addr),
			zap.String("name", s.cfg.Server.Name),
		)
		if err := s.grpc.Serve(lis); err != nil {
			serveErr <- err
		}
	}()

	// Wait for interrupt/termination signal or serve error.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serveErr:
		s.log.Error("gRPC server failed", zap.Error(err))
		return err
	case sig := <-sigCh:
		s.log.Info("shutdown signal received", zap.Stringer("signal", sig))
	case <-ctx.Done():
		s.log.Info("context cancelled, shutting down")
	}

	// Graceful shutdown.
	s.log.Info("graceful shutdown starting...")
	stopped := make(chan struct{})
	go func() {
		s.grpc.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		s.log.Info("graceful shutdown complete")
	case <-time.After(15 * time.Second):
		s.log.Warn("graceful shutdown timed out, forcing stop")
		s.grpc.Stop()
	}

	// Update health status.
	healthServer.SetServingStatus("qms-service", healthpb.HealthCheckResponse_NOT_SERVING)
	return nil
}

// StopForced immediately stops the gRPC server (for error paths).
func (s *Server) StopForced() {
	if s.grpc != nil {
		s.grpc.Stop()
	}
}
