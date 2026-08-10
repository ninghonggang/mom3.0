package main

import (
	"flag"
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"mom-platform/services/mdm-service/internal/config"
	"mom-platform/services/mdm-service/internal/model"
	"mom-platform/services/mdm-service/internal/repository"
	"mom-platform/services/mdm-service/internal/server"
	"mom-platform/services/mdm-service/internal/service"
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
		&model.Material{}, &model.Bom{}, &model.BomItem{},
		&model.Workshop{}, &model.ProductionLine{}, &model.Workstation{},
		&model.Customer{}, &model.Supplier{},
	); err != nil {
		logger.Fatal("auto-migrate failed", zap.Error(err))
	}

	// Repositories
	materialRepo := repository.NewMaterialRepo(db)
	bomRepo := repository.NewBomRepo(db)
	workshopRepo := repository.NewWorkshopRepo(db)
	lineRepo := repository.NewProductionLineRepo(db)
	workstationRepo := repository.NewWorkstationRepo(db)
	customerRepo := repository.NewCustomerRepo(db)
	supplierRepo := repository.NewSupplierRepo(db)

	// Service
	svc := service.NewMDMService(
		logger, db,
		materialRepo, bomRepo,
		workshopRepo, lineRepo, workstationRepo,
		customerRepo, supplierRepo,
	)

	// Server (creates EventPublisher, registers gRPC services, health check)
	srv := server.New(cfg, svc, logger)

	logger.Info("starting MDM service",
		zap.String("name", cfg.Server.Name),
		zap.Int("port", cfg.Server.Port),
	)

	if err := srv.Start(); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}
}
