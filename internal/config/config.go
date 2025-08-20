package config

import (
	"fmt"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

const (
	defaultHTTPPort               = "8000"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
	defaultLimiterRPS             = 10
	defaultLimiterBurst           = 20
	defaultLimiterTTL             = 10 * time.Minute
	defaultCacheTTL               = 60 * time.Second
	defaultLoggerLevel            = "info"

	EnvLocal = "local"
	Prod     = "prod"
)

type Config struct {
	Environment string `envconfig:"APP_ENV" default:"local"`
	PostgreSQL  PostgreSQLConfig
	HTTP        HTTPConfig
	Auth        AuthConfig
	Limiter     LimiterConfig
	Cache       CacheConfig
	Logger      LoggerConfig
	Telegram    TelegramConfig
}

type PostgreSQLConfig struct {
	Host              string        `envconfig:"POSTGRESQL_HOST" default:"localhost"`
	Port              int           `envconfig:"POSTGRESQL_PORT" default:"5432"`
	User              string        `envconfig:"POSTGRESQL_USER" default:"root"`
	Password          string        `envconfig:"POSTGRESQL_PASSWORD" default:"fake_password"`
	DBName            string        `envconfig:"POSTGRESQL_NAME" default:"upserv"`
	SSLMode           string        `mapstructure:"ssl_mode"`
	MaxConnections    int           `mapstructure:"max_connections"`
	MinConnections    int           `mapstructure:"min_connections"`
	MaxConnLifeTime   time.Duration `mapstructure:"max_conn_life_time"`
	MaxConnIdleTime   time.Duration `mapstructure:"max_conn_idle_time"`
	HealthCheckPeriod time.Duration `mapstructure:"health_check_period"`
}

type HTTPConfig struct {
	Host               string        `envconfig:"HTTP_HOST" default:"localhost"`
	Port               string        `mapstructure:"port"`
	ReadTimeout        time.Duration `mapstructure:"readTimeout"`
	WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
}

type TelegramConfig struct {
	Token  string `envconfig:"TELEGRAM_TOKEN" default:"localhost"`
	ChatID string `envconfig:"TELEGRAM_CHATID" default:"localhost"`
}

type LimiterConfig struct {
	RPS   int           `mapstructure:"rps"`
	Burst int           `mapstructure:"burst"`
	TTL   time.Duration `mapstructure:"ttl"`
}

type AuthConfig struct {
	JWT JWTConfig `mapstructure:",squash"`
}

type JWTConfig struct {
	AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
	RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
	SigningKey      string        `envconfig:"JWT_SIGNING_KEY" default:"sign_key"`
}

type CacheConfig struct {
	TTL time.Duration `mapstructure:"ttl"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

func Init(configsDir string) (*Config, error) {
	const op = "config.Init"

	populateDefaults()

	if err := parseConfigFile(configsDir, os.Getenv("APP_ENV")); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var cfg Config

	if err := unmarshalFromViper(&cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := setFromEnv(&cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &cfg, nil
}

func unmarshalFromViper(cfg *Config) error {
	const op = "config.unmarshalFromViper"

	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func setFromEnv(cfg *Config) error {
	const op = "config.setFromEnv"

	err := envconfig.Process("", cfg)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func parseConfigFile(folder, env string) error {
	const op = "config.parseConfigFile"

	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if env == EnvLocal || env == "" {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.readTimeout", defaultHTTPRWTimeout)
	viper.SetDefault("http.writeTimeout", defaultHTTPRWTimeout)
	viper.SetDefault("http.maxHeaderBytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("auth.accessTokenTTL", defaultAccessTokenTTL)
	viper.SetDefault("auth.refreshTokenTTL", defaultRefreshTokenTTL)
	viper.SetDefault("limiter.rps", defaultLimiterRPS)
	viper.SetDefault("limiter.burst", defaultLimiterBurst)
	viper.SetDefault("limiter.ttl", defaultLimiterTTL)
	viper.SetDefault("cache.ttl", defaultCacheTTL)
	viper.SetDefault("logger.level", defaultCacheTTL)
}

func (p *PostgreSQLConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode)
}

func (c *Config) String() string {
	return fmt.Sprintf(`
Environment: %s

Database:
  Host: %s
  Port: %d
  User: %s
  DB Name: %s
  SSL Mode: %s
  Max Connections: %d
  Min Connections: %d
  Max Conn IDLE time: %v
  Max Conn life time: %v
  Health Check period: %v

HTTP Server:
  Host: %s
  Port: %s
  Read Timeout: %v
  Write Timeout: %v
  Max Header Bytes: %d

Authentication:
  Access Token TTL: %v
  Refresh Token TTL: %v

Rate Limiter:
  RPS: %d
  Burst: %d
  TTL: %v

Cache:
  TTL: %v

Logger:
  Level: %s
`,
		c.Environment,
		c.PostgreSQL.Host, c.PostgreSQL.Port, c.PostgreSQL.User,
		c.PostgreSQL.DBName, c.PostgreSQL.SSLMode, c.PostgreSQL.MaxConnections,
		c.PostgreSQL.MinConnections, c.PostgreSQL.MaxConnIdleTime, c.PostgreSQL.MaxConnLifeTime,
		c.PostgreSQL.HealthCheckPeriod,
		c.HTTP.Host, c.HTTP.Port, c.HTTP.ReadTimeout,
		c.HTTP.WriteTimeout, c.HTTP.MaxHeaderMegabytes,
		c.Auth.JWT.AccessTokenTTL, c.Auth.JWT.RefreshTokenTTL,
		c.Limiter.RPS, c.Limiter.Burst, c.Limiter.TTL,
		c.Cache.TTL, c.Logger.Level,
	)
}
