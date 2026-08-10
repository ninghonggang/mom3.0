package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"mom-platform/services/andon-service/internal/config"
	"mom-platform/services/andon-service/internal/model"
	"mom-platform/services/andon-service/internal/server"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	// Logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
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
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Auto-migrate for development
	if err := db.AutoMigrate(
		&model.AndonCall{},
		&model.AndonAction{},
		&model.AlertConfig{},
		&model.Alert{},
		&model.AlertEscalation{},
	); err != nil {
		logger.Fatal("failed to auto migrate", zap.Error(err))
	}

	// Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})
	logger.Info("andon-service initialized successfully")

	// Server
	srv := server.NewServer(cfg, db, rdb, logger)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Start(); err != nil {
			logger.Fatal("server failed", zap.Error(err))
		}
	}()

	<-quit
	logger.Info("shutting down...")
	srv.Stop()
}
