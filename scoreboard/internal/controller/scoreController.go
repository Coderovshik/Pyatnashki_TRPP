package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"scoreboard/internal/service"
	"strconv"
)

type ScoreController struct {
	log     *slog.Logger
	service *service.ScoreService
}

func NewScoreController(log *slog.Logger, service *service.ScoreService) *ScoreController {
	return &ScoreController{
		log:     log,
		service: service,
	}
}

func (s *ScoreController) GetTopScores(w http.ResponseWriter, r *http.Request) {
	const op = "ScoreController.GetTopScores"
	log := s.log.With(slog.String("op", op))

	limitParam := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil && len(limitParam) != 0 {
		log.Error("invalid limit param", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	offsetParam := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offsetParam)
	if err != nil && len(offsetParam) != 0 {
		log.Error("invalid offset param", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	scores, err := s.service.GetTopScores(r.Context(), limit, offset)
	if err != nil && len(limitParam) != 0 {
		log.Error("failed to get top scores", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	jsonBytes, _ := json.Marshal(scores)
	w.Write(jsonBytes)
}

func (s *ScoreController) SetScore(w http.ResponseWriter, r *http.Request) {
	const op = "ScoreController.SetScore"
	log := s.log.With(slog.String("op", op))

	idParam := r.URL.Query().Get("user_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error("invalid id param", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req struct {
		Score int `json:"score"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error("failed to decode request", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.service.SetScore(r.Context(), id, req.Score)
	if err != nil {
		log.Error("failed to set user score", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
