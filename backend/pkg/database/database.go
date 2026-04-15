package database

import (
	"fmt"
	"time"

	"github.com/lingojack/proj_template/config"
	"github.com/lingojack/proj_template/model"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func New(cfg *config.Config, log zerolog.Logger) (*gorm.DB, func(), error) {
	gormCfg := &gorm.Config{
		Logger: gormlogger.Default.LogMode(parseLogMode(cfg.Database.LogMode)),
	}

	var dialector gorm.Dialector
	switch cfg.Database.Driver {
	case "mysql":
		dialector = mysql.Open(cfg.Database.DSN)
	default:
		return nil, nil, fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}

	db, err := gorm.Open(dialector, gormCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetimeMinutes) * time.Minute)

	log.Info().Str("driver", cfg.Database.Driver).Msg("database connected")

	// 自动迁移模型
	if err := db.AutoMigrate(&model.Post{}); err != nil {
		return nil, nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close database")
		}
	}
	return db, cleanup, nil
}

func parseLogMode(mode string) gormlogger.LogLevel {
	switch mode {
	case "silent":
		return gormlogger.Silent
	case "error":
		return gormlogger.Error
	case "warn":
		return gormlogger.Warn
	default:
		return gormlogger.Info
	}
}
