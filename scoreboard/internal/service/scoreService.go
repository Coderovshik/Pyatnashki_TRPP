package service

import (
	"context"
	"log/slog"
	"scoreboard/internal/domain"
	"scoreboard/internal/repository"
)

type ScoreService struct {
	repo *repository.ScoreRepository
	log  *slog.Logger
}

func NewScoreService(log *slog.Logger, repo *repository.ScoreRepository) *ScoreService {
	return &ScoreService{
		repo: repo,
		log:  log,
	}
}

func (s *ScoreService) GetTopScores(ctx context.Context, limit int, offset int) ([]domain.Score, error) {
	const op = "ScoreService.GetTopScores"
	log := s.log.With(slog.String("op", op))

	if limit == 0 {
		limit = 3
	}

	scores, err := s.repo.GetTopScores(ctx, limit, offset)
	if err != nil {
		log.Error("failed to get top scores")
		return nil, err
	}

	if scores == nil {
		scores = make([]domain.Score, 0)
	}
	return scores, nil
}

func (s *ScoreService) SetScore(ctx context.Context, id int, scoreValue int) error {
	const op = "ScoreService.GetTopScores"
	log := s.log.With(slog.String("op", op))

	err := s.repo.SetScore(ctx, id, scoreValue)
	if err != nil {
		log.Error("failed to set user score")
		return err
	}

	return nil
}
