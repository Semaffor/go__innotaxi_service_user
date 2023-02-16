package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	defaultServerPort             = "8000"
	defaultServerHost             = "localhost"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour
	defaultRefreshTokenLength     = 20

	configDir      = "configs"
	configFileName = "config"
)

type (
	Config struct {
		Server     ServerConfig
		Mongo      DBConfig
		Postgres   DBConfig
		Redis      DBConfig
		AuthConfig AuthConfig
	}

	ServerConfig struct {
		Host           string        `mapstructure:"host"`
		Port           int           `mapstructure:"port"`
		MaxHeaderBytes int           `mapstructure:"maxHeaderBytes" `
		ReadTimeout    time.Duration `mapstructure:"readTimeout" `
		WriteTimeout   time.Duration `mapstructure:"writeTimeout"`
	}

	DBConfig struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		DBName   string `mapstructure:"dbName"`
		SslMode  string `mapstructure:"sslMode"`
		Password string `mapstructure:"password"`
	}

	AuthConfig struct {
		JWT          JWTConfig
		PasswordSalt string
	}

	JWTConfig struct {
		AccessTokenTTL     time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL    time.Duration `mapstructure:"refreshTokenTTL"`
		RefreshTokenLength int           `mapstructure:"refreshTokenLength"`
		SigningKey         string
	}
)

func InitConfig() (*Config, error) {
	populateDefaults()

	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	setFromEnv(&cfg)

	return &cfg, nil
}

func populateDefaults() {
	viper.SetDefault("server.port", defaultServerPort)
	viper.SetDefault("server.host", defaultServerHost)
	viper.SetDefault("server.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("server.writeTimeout", defaultHTTPRWTimeout)
	viper.SetDefault("auth.accessTokenTTL", defaultAccessTokenTTL)
	viper.SetDefault("auth.refreshTokenTTL", defaultRefreshTokenTTL)
	viper.SetDefault("auth.refreshTokenLength", defaultRefreshTokenLength)
}

func parseConfigFile() error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName(configFileName)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("server", &cfg.Server); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("mongodb", &cfg.Mongo); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("redis", &cfg.Redis); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("auth", &cfg.AuthConfig.JWT); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.Postgres.Password = os.Getenv("DB_POSTGRES_PASSWORD")
	cfg.Mongo.Password = os.Getenv("DB_MONGO_PASSWORD")
	cfg.AuthConfig.PasswordSalt = os.Getenv("PASSWORD_SALT")
	cfg.AuthConfig.JWT.SigningKey = os.Getenv("JWT_SIGNATURE")
}
