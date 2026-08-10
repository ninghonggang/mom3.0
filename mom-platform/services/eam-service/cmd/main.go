package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"mom-platform/services/eam-service/internal/config"
	"mom-platform/services/eam-service/internal/handler"
	"mom-platform/services/eam-service/internal/repository"
	"mom-platform/services/eam-service/internal/server"
	"mom-platform/services/eam-service/internal/service"
)

func main() {
	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// Load configuration
	cfgPath := os.Getenv("EAM_CONFIG_PATH")
	if cfgPath == "" {
		cfgPath = "config.yaml"
	}
	cfg := config.MustLoad(logger, cfgPath)

	logger.Info("configuration loaded",
		zap.Int("port", cfg.Server.Port),
		zap.String("db_host", cfg.Database.Host),
	)

	// Connect to database
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}

	// Auto-migrate
	if err := repository.AutoMigrate(db); err != nil {
		logger.Fatal("failed to run migrations", zap.Error(err))
	}
	logger.Info("database migrations completed")

	// Wire dependencies
	repo := repository.NewRepository(db)
	svc := service.New(repo, db, logger)
	h := handler.New(svc, logger, nil)

	// Start gRPC server
	grpcServer := server.New(cfg, h, logger)
	if err := grpcServer.Start(); err != nil {
		logger.Fatal("gRPC server error", zap.Error(err))
	}
}
