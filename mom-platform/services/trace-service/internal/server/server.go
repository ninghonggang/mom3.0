package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	trace "github.com/ninghonggang/mom-platform/gen/trace"

	"mom-platform/services/trace-service/internal/config"
	"mom-platform/services/trace-service/internal/handler"
	"mom-platform/services/trace-service/internal/repository"
	"mom-platform/services/trace-service/internal/service"
)

type Server struct {
	cfg     *config.Config
	db      *gorm.DB
	redis   *redis.Client
	logger  *zap.Logger
	grpcSrv *grpc.Server
	pub     *eventbus.EventPublisher
}

func NewServer(cfg *config.Config, db *gorm.DB, redis *redis.Client, logger *zap.Logger) *Server {
	return &Server{
		cfg:    cfg,
		db:     db,
		redis:  redis,
		logger: logger,
	}
}

func (s *Server) Start() error {
	repo := repository.NewGormRepository(s.db, s.logger)
	svc := service.NewTraceService(repo, s.redis, s.logger)

	// NATS EventPublisher
	var pub *eventbus.EventPublisher
	if s.cfg.NATS.URL != "" {
		var err error
		pub, err = eventbus.NewEventPublisher(s.cfg.NATS.URL, s.logger)
		if err != nil {
			s.logger.Warn("failed to connect NATS, events disabled", zap.Error(err))
		} else {
			// Ensure the stream exists
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			_ = pub.EnsureStream(ctx, "mom-events", []string{"mom.trace.*", "mom.andon.*"})
		}
	}

	h := handler.NewTraceHandler(svc, pub, s.logger)

	s.grpcSrv = grpc.NewServer()
	reflection.Register(s.grpcSrv)

	// Register gRPC service
	trace.RegisterTraceServiceServer(s.grpcSrv, h)

	// Register health check
	healthSrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.grpcSrv, healthSrv)
	healthSrv.SetServingStatus("trace-service", grpc_health_v1.HealthCheckResponse_SERVING)

	addr := fmt.Sprintf(":%d", s.cfg.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	s.pub = pub

	s.logger.Info("trace-service starting",
		zap.String("name", s.cfg.Server.Name),
		zap.Int("port", s.cfg.Server.Port),
		zap.String("db", s.cfg.Database.DBName),
	)

	return s.grpcSrv.Serve(lis)
}

func (s *Server) Stop() {
	if s.pub != nil {
		s.pub.Close()
	}
	if s.grpcSrv != nil {
		s.grpcSrv.GracefulStop()
	}
	s.logger.Info("trace-service stopped")
}
