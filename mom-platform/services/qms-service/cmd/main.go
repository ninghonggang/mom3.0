package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"mom-platform/services/qms-service/internal/config"
	"mom-platform/services/qms-service/internal/handler"
	"mom-platform/services/qms-service/internal/repository"
	"mom-platform/services/qms-service/internal/server"
	"mom-platform/services/qms-service/internal/service"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	// --- Logger ---
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// --- Config ---
	cfg, err := config.Load(*configPath)
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err), zap.String("path", *configPath))
	}
	logger.Info("configuration loaded",
		zap.Int("port", cfg.Server.Port),
		zap.String("db", cfg.Database.DBName),
	)

	// --- Database ---
	db, err := initDB(cfg, logger)
	if err != nil {
		logger.Fatal("failed to init database", zap.Error(err))
	}

	// --- Repository ---
	repo := repository.NewGormRepository(db)
	if err := repo.AutoMigrate(); err != nil {
		logger.Fatal("auto-migrate failed", zap.Error(err))
	}
	logger.Info("database migration complete")

	// --- Service ---
	svc := service.NewService(repo, logger, service.AQLConfig{
		AQLThreshold: 0.01, // 1% defect rate triggers auto-fail
	})

	// --- Handler ---
	h := handler.NewHandler(svc, logger)

	// --- gRPC Server ---
	srv := server.New(cfg, logger, h)

	// --- Start with graceful shutdown ---
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := srv.Start(ctx); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}

	logger.Info("service stopped")
}

// initDB connects to PostgreSQL and returns a *gorm.DB instance.
func initDB(cfg *config.Config, log *zap.Logger) (*gorm.DB, error) {
	dsn := cfg.Database.DSN()
	log.Info("connecting to database",
		zap.String("host", cfg.Database.Host),
		zap.Int("port", cfg.Database.Port),
		zap.String("dbname", cfg.Database.DBName),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get *sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Verify connectivity.
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	log.Info("database connected")
	return db, nil
}
