package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

//go:embed migrations/*.sql
var schemaFS embed.FS

type PostgresConfig struct {
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	Database string `env:"POSTGRES_DATABASE"`
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

func main() {
	godotenv.Load("config/.env")

	var cfg PostgresConfig
	cleanenv.ReadEnv(&cfg)

	d, err := iofs.New(schemaFS, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, cfg.Addr())
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("migrated successfully")
}
