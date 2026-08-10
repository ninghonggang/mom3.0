package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ninghonggang/mom-platform/gen/mes"
	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/config"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/handler"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/model"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/repository"
	"github.com/ninghonggang/mom-platform/services/mes-service/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm"
)

type Server struct {
	cfg       *config.Config
	logger    *zap.Logger
	db        *gorm.DB
	grpc      *grpc.Server
	publisher *eventbus.EventPublisher
	h         *handler.MESHandler
}

func New(cfg *config.Config) (*Server, error) {
	logCfg := zap.NewProductionConfig()
	if cfg.Logging.Level == "debug" {
		logCfg = zap.NewDevelopmentConfig()
	}
	if cfg.Logging.Format == "json" {
		logCfg.Encoding = "json"
	}
	logger, err := logCfg.Build()
	if err != nil {
		return nil, fmt.Errorf("build logger: %w", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	// Auto-migrate for development
	if cfg.Server.Mode == "debug" {
		if err := db.AutoMigrate(
			&model.ProductionOrder{},
			&model.MobileJobReport{},
			&model.Dispatch{},
			&model.ProductionComplete{},
		); err != nil {
			return nil, fmt.Errorf("auto migrate: %w", err)
		}
		logger.Info("database migrated")
	}

	// NATS EventPublisher
	publisher, err := eventbus.NewEventPublisher(cfg.NATS.URL, logger)
	if err != nil {
		return nil, fmt.Errorf("create event publisher: %w", err)
	}
	ctx := context.Background()
	if err := publisher.EnsureStream(ctx, "MES_EVENTS", []string{
		eventbus.SubjectMESOrderCreated,
		eventbus.SubjectMESOrderCompleted,
		eventbus.SubjectMESOrderHold,
		eventbus.SubjectMESReportAudited,
		eventbus.SubjectMESException,
	}); err != nil {
		return nil, fmt.Errorf("ensure mes stream: %w", err)
	}

	// Repositories
	orderRepo := repository.NewOrderRepository(db)
	reportRepo := repository.NewReportRepository(db)
	dispatchRepo := repository.NewDispatchRepository(db)
	completeRepo := repository.NewCompleteRepository(db)

	// Services
	orderService := service.NewOrderService(orderRepo, publisher)
	reportService := service.NewReportService(reportRepo, orderRepo)
	dispatchService := service.NewDispatchService(dispatchRepo, orderRepo)
	completeService := service.NewCompleteService(completeRepo, orderRepo, publisher)

	// Handler
	h := handler.NewMESHandler(orderService, reportService, dispatchService, completeService)

	// gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor(logger)),
	)
	reflection.Register(grpcServer)

	// Register all 4 MES gRPC services
	mes.RegisterProductionOrderServiceServer(grpcServer, h)
	mes.RegisterDispatchServiceServer(grpcServer, h)
	mes.RegisterJobReportServiceServer(grpcServer, h)
	mes.RegisterProductionCompleteServiceServer(grpcServer, h)

	// Health check
	healthSrv := health.NewServer()
	healthSrv.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcServer, healthSrv)

	return &Server{
		cfg:       cfg,
		logger:    logger,
		db:        db,
		grpc:      grpcServer,
		publisher: publisher,
		h:         h,
	}, nil
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Server.Port))
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		s.logger.Info("shutting down...")
		s.grpc.GracefulStop()
		if s.publisher != nil {
			s.publisher.Close()
		}
		cancel()
	}()

	s.logger.Info("MES service starting", zap.Int("port", s.cfg.Server.Port))
	return s.grpc.Serve(lis)
}

func loggingInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.Info("gRPC request", zap.String("method", info.FullMethod))
		return handler(ctx, req)
	}
}
