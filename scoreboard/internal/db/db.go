package db

import (
	"database/sql"
	"log"
	"scoreboard/internal/config"

	_ "github.com/lib/pq"
)

type Databse struct {
	DB *sql.DB
}

func New(cfg *config.Config) *Databse {
	db, err := sql.Open("postgres", cfg.Postgres.Addr())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database connection established")

	return &Databse{
		DB: db,
	}
}

func (d *Databse) Close() error {
	return d.DB.Close()
}
