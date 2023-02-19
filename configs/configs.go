package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfigs struct {
	HTTPIP           string `yaml:"http_ip" env:"HTTP_IP"`
	HTTPPort         string `yaml:"http_port" env:"HTTP_PORT"`
	HTTPReadTimeout  string `yaml:"http_read_timeout" env:"HTTP_READ_TIMEOUT"`
	HTTPWriteTimeout string `yaml:"http_write_timeout" env:"HTTP_WRITE_TIMEOUT"`

	JWTSecretKey  string `json:"jwt_secret_key" env:"JWT_SECRET_KEY"`
	JWTExpiration string `json:"jwt_expiration" env:"JWT_EXPIRATION"`

	DatabaseIP   string `yaml:"database_host" env:"DATABASE_ADDR"`
	DatabasePort string `yaml:"database_port" env:"DATABASE_PORT"`
	DatabaseName string `yaml:"database_name" env:"DATABASE_NAME"`
	DatabaseUser string `yaml:"database_user" env:"DATABASE_USER"`
	DatabasePass string `yaml:"database_pass" env:"DATABASE_PASS"`

	RedisAddr string `yaml:"redis_address" env:"REDIS_ADDRESS"`
	RedisPass string `yaml:"redis_password" env:"REDIS_PASSWORD"`
	RedisTTL  string `yaml:"redis_ttl" env:"REDIS_TTL"`

	LogPath  string `yaml:"log_path" env:"APP_LOG_PATH"`
	LogLevel string `yaml:"log_level" env:"FILE_LOG_LEVEL"`
}

func Load() (*AppConfigs, error) {
	var cfg AppConfigs

	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
