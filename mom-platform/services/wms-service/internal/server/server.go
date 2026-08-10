package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	wms "github.com/ninghonggang/mom-platform/gen/wms"
	"mom-platform/services/wms-service/internal/handler"
)

// Server runs the WMS gRPC service.
type Server struct {
	grpcServer *grpc.Server
	httpServer *http.Server
	port       int
	log        *zap.Logger
}

// New creates a new Server with the WMS gRPC service registered.
func New(port int, h *handler.Handler, log *zap.Logger) *Server {
	grpcSrv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(),
		),
	)

	// Register the WMS gRPC service.
	wms.RegisterWmsServiceServer(grpcSrv, h)

	// Register gRPC health check service.
	healthSrv := health.NewServer()
	healthSrv.SetServingStatus("mom.wms.WmsService", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcSrv, healthSrv)

	// HTTP health endpoint (for k8s liveness/readiness probes).
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"healthy"}`))
	})

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port+1),
		Handler: mux,
	}

	return &Server{
		grpcServer: grpcSrv,
		httpServer: httpSrv,
		port:       port,
		log:        log,
	}
}

// Start begins serving. The gRPC server listens on the configured port;
// the HTTP health endpoint listens on port+1.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("grpc listen on %d: %w", s.port, err)
	}

	go func() {
		s.log.Info("gRPC server starting", zap.Int("port", s.port))
		if err := s.grpcServer.Serve(lis); err != nil {
			s.log.Error("gRPC server stopped", zap.Error(err))
		}
	}()

	go func() {
		httpPort := s.port + 1
		s.log.Info("HTTP health endpoint starting", zap.Int("port", httpPort))
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Error("HTTP server stopped", zap.Error(err))
		}
	}()

	return nil
}

// Stop gracefully shuts down both servers.
func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("shutting down servers")

	stopped := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
	case <-ctx.Done():
		s.log.Warn("grpc graceful stop timed out, forcing stop")
		s.grpcServer.Stop()
	}

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("http shutdown: %w", err)
	}
	return nil
}

// GRPCServer returns the underlying gRPC server.
func (s *Server) GRPCServer() *grpc.Server {
	return s.grpcServer
}
