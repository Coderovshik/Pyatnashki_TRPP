package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

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

	for range 10 {
		res, _ := http.Get("http://backend:9999/actuator")
		if res.StatusCode == 200 {
			break
		}
		time.Sleep(3 * time.Second)
	}

	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("migrated successfully")
}
