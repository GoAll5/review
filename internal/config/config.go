package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"time"
)

type Config struct {
	Env    string       `yaml:"env" env:"ENV" env-default:"local"`
	HTTP   HTTPConfig   `yaml:"http"`
	Public PublicConfig `yaml:"public"`
}

type PublicConfig struct {
	BaseURL string `yaml:"baseUrl" env:"PUBLIC_BASE_URL" env-default:"/"`
}

type HTTPConfig struct {
	Port    int           `yaml:"port" env:"HTTP_PORT" env-default:"8083"`
	Addr    string        `yaml:"addr" env:"HTTP_ADDR" env-default:"localhost"`
	Timeout time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"5s"`
}

func MustLoad() *Config {
	var cfg Config
	_ = godotenv.Load()

	//path := os.Getenv("CONFIG_PATH")
	//if path != "" {
	//	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
	//		panic(fmt.Errorf("read env: %w", err))
	//	}
	//}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(fmt.Errorf("read env: %w", err))
	}

	return &cfg
}
