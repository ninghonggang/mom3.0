package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/ninghonggang/mom-platform/pkg/eventbus"
	"mom-platform/services/wms-service/internal/config"
	"mom-platform/services/wms-service/internal/handler"
	"mom-platform/services/wms-service/internal/model"
	"mom-platform/services/wms-service/internal/repository"
	"mom-platform/services/wms-service/internal/server"
	"mom-platform/services/wms-service/internal/service"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	// --- Logger ---
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to init logger: " + err.Error())
	}
	defer logger.Sync()

	// --- Config ---
	cfg, err := config.Load(*configPath)
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}
	logger.Info("config loaded",
		zap.Int("grpc_port", cfg.Server.Port),
		zap.String("db", cfg.Database.DBName),
		zap.String("nats", cfg.NATS.URL))

	// --- Database ---
	db, err := initDB(cfg, logger)
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}

	// Auto-migrate (safe to run repeatedly; use migrations/ for production).
	if err := db.AutoMigrate(model.AllModels()...); err != nil {
		logger.Fatal("auto-migrate failed", zap.Error(err))
	}
	logger.Info("database migrated")

	// --- NATS Event Publisher ---
	pub, err := eventbus.NewEventPublisher(cfg.NATS.URL, logger)
	if err != nil {
		logger.Fatal("failed to connect NATS", zap.Error(err))
	}
	defer pub.Close()

	// Ensure WMS stream exists.
	ctx := context.Background()
	if err := pub.EnsureStream(ctx, "WMS", []string{
		eventbus.SubjectWMSReceiveCompleted,
		eventbus.SubjectWMSShipped,
	}); err != nil {
		logger.Fatal("failed to ensure WMS stream", zap.Error(err))
	}

	// --- Repository ---
	repo := repository.New(db)

	// --- Service ---
	svc := service.New(repo, logger, pub)

	// --- Handler ---
	h := handler.New(svc, logger)

	// --- Server ---
	srv := server.New(cfg.Server.Port, h, logger)
	if err := srv.Start(); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
	logger.Info("WMS service started", zap.String("name", cfg.Server.Name))

	// --- Graceful shutdown ---
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh
	logger.Info("received signal, shutting down", zap.String("signal", sig.String()))

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Stop(shutdownCtx); err != nil {
		logger.Error("graceful shutdown error", zap.Error(err))
	}
	logger.Info("WMS service stopped")
}

// initDB creates and configures the GORM DB connection.
func initDB(cfg *config.Config, log *zap.Logger) (*gorm.DB, error) {
	level := gormlogger.Warn
	if os.Getenv("DB_LOG_LEVEL") == "debug" {
		level = gormlogger.Info
	}
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Info("database connected",
		zap.String("host", cfg.Database.Host),
		zap.Int("port", cfg.Database.Port))
	return db, nil
}
