package main

import (
	"scoreboard/internal/app"
	"scoreboard/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("config/.env")

	cfg := config.New()
	app := app.NewApp(cfg)

	app.Run()
}
