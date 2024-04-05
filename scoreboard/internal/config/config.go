package config

import (
	"fmt"
	"log"
	"net"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres PostgresConfig `env-prefix:"POSTGRES_"`
	Http     HttpConfig     `env-prefix:"HTTP_"`
	Auth     AuthConfig
}

type PostgresConfig struct {
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	Database string `env:"DATABASE"`
}

func (p PostgresConfig) Addr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.Database,
	)
}

type HttpConfig struct {
	Host string `env:"HOST"`
	Port string `env:"PORT"`
}

func (h HttpConfig) Addr() string {
	return net.JoinHostPort(h.Host, h.Port)
}

type AuthConfig struct {
	SigningKey string
}

func New() *Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}
