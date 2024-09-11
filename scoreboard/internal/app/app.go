package app

import (
	"net/http"
	"scoreboard/internal/config"
	"scoreboard/internal/controller"
	"scoreboard/internal/db"
	"scoreboard/internal/logger"
	"scoreboard/internal/repository"
	"scoreboard/internal/service"
)

type App struct {
	cfg    *config.Config
	router *http.ServeMux
}

func NewApp(cfg *config.Config) *App {
	r := http.NewServeMux()

	log := logger.New()
	database := db.New(cfg)

	repo := repository.NewScoreRepository(log, database.DB)
	s := service.NewScoreService(log, repo)
	c := controller.NewScoreController(log, s)

	r.HandleFunc("OPTIONS /top-scores", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.WriteHeader(http.StatusOK)
	})
	r.HandleFunc("GET /top-scores", c.GetTopScores)

	r.HandleFunc("OPTIONS /scores", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.WriteHeader(http.StatusOK)
	})
	r.HandleFunc("POST /scores", c.SetScore)

	return &App{
		cfg:    cfg,
		router: r,
	}
}

func (a *App) Run() {
	http.ListenAndServe(a.cfg.Http.Addr(), a.router)
}
