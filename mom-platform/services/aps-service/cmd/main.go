package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"mom-platform/services/aps-service/internal/config"
	"mom-platform/services/aps-service/internal/handler"
	"mom-platform/services/aps-service/internal/model"
	"mom-platform/services/aps-service/internal/repository"
	"mom-platform/services/aps-service/internal/server"
	"mom-platform/services/aps-service/internal/service"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	// Logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// Config
	cfg, err := config.Load(*configPath)
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	// Database
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("failed to get sql.DB", zap.Error(err))
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)

	// Auto-migrate
	if err := db.AutoMigrate(
		&model.MpsPlan{}, &model.MrpPlan{}, &model.WorkCenter{},
		&model.ScheduleJob{}, &model.ScheduleConstraint{}, &model.Changeover{},
	); err != nil {
		logger.Fatal("auto-migrate failed", zap.Error(err))
	}

	// NATS EventPublisher
	pub, err := eventbus.NewEventPublisher(cfg.NATS.URL, logger)
	if err != nil {
		logger.Fatal("failed to create event publisher", zap.Error(err))
	}
	defer pub.Close()

	ctx := context.Background()
	if err := pub.EnsureStream(ctx, "APS", []string{
		eventbus.SubjectAPSMPSReleased,
		eventbus.SubjectAPSMRPCompleted,
		eventbus.SubjectAPSSchedulePublished,
	}); err != nil {
		logger.Fatal("failed to ensure APS stream", zap.Error(err))
	}

	// Repositories
	mpsRepo := repository.NewMpsPlanRepo(db)
	mrpRepo := repository.NewMrpPlanRepo(db)
	workCenterRepo := repository.NewWorkCenterRepo(db)
	jobRepo := repository.NewScheduleJobRepo(db)
	changeoverRepo := repository.NewChangeoverRepo(db)

	// Service
	svc := service.NewAPSService(
		logger, db,
		mpsRepo, mrpRepo,
		workCenterRepo, jobRepo, changeoverRepo,
		pub,
	)

	// Handler
	h := handler.NewAPSHandler(logger, svc)

	// Server (Start blocks until signal; graceful shutdown is handled internally)
	srv := server.New(cfg, h, logger)

	logger.Info("starting APS service",
		zap.String("name", cfg.Server.Name),
		zap.Int("port", cfg.Server.Port),
	)

	if err := srv.Start(); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}
}
