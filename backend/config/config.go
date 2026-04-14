package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Database   DatabaseConfig   `mapstructure:"database"`
	Log        LogConfig        `mapstructure:"log"`
	Middleware MiddlewareConfig `mapstructure:"middleware"`
	API        APIConfig        `mapstructure:"api"`
}

type ServerConfig struct {
	Host                           string `mapstructure:"host"`
	Port                           int    `mapstructure:"port"`
	ReadTimeoutSeconds             int    `mapstructure:"read_timeout_seconds"`
	WriteTimeoutSeconds            int    `mapstructure:"write_timeout_seconds"`
	GracefulShutdownTimeoutSeconds int    `mapstructure:"graceful_shutdown_timeout_seconds"`
}

type DatabaseConfig struct {
	Driver                 string `mapstructure:"driver"`
	DSN                    string `mapstructure:"dsn"`
	MaxOpenConns           int    `mapstructure:"max_open_conns"`
	MaxIdleConns           int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetimeMinutes int    `mapstructure:"conn_max_lifetime_minutes"`
	LogMode                string `mapstructure:"log_mode"`
}

type LogConfig struct {
	Level    string `mapstructure:"level"`
	Format   string `mapstructure:"format"`
	Output   string `mapstructure:"output"`
	FilePath string `mapstructure:"file_path"`
}

type MiddlewareConfig struct {
	CORS      CORSConfig      `mapstructure:"cors"`
	RateLimit RateLimitConfig `mapstructure:"rate_limit"`
	Auth      AuthConfig      `mapstructure:"auth"`
	RequestID RequestIDConfig `mapstructure:"request_id"`
	Recover   RecoverConfig   `mapstructure:"recover"`
}

type CORSConfig struct {
	Enabled          bool     `mapstructure:"enabled"`
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
	AllowedMethods   []string `mapstructure:"allowed_methods"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
}

type RateLimitConfig struct {
	Enabled            bool    `mapstructure:"enabled"`
	RequestsPerSecond  float64 `mapstructure:"requests_per_second"`
	Burst              int     `mapstructure:"burst"`
}

type AuthConfig struct {
	Enabled        bool   `mapstructure:"enabled"`
	JWTSecret      string `mapstructure:"jwt_secret"`
	JWTExpiryHours int    `mapstructure:"jwt_expiry_hours"`
}

type RequestIDConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

type RecoverConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

type APIConfig struct {
	Prefix  string `mapstructure:"prefix"`
	Version string `mapstructure:"version"`
}

func Load(cfgPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(cfgPath)

	// 允许用 APP_ 前缀的环境变量覆盖任意配置项
	// 例如 APP_DATABASE_DSN 覆盖 database.dsn
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
