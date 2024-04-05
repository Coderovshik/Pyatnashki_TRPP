package repository

import (
	"context"
	"database/sql"
	"log/slog"
	"scoreboard/internal/domain"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type ScoreRepository struct {
	log *slog.Logger
	db  DBTX
}

func NewScoreRepository(log *slog.Logger, db DBTX) *ScoreRepository {
	return &ScoreRepository{
		log: log,
		db:  db,
	}
}

func (s *ScoreRepository) GetTopScores(ctx context.Context, limit int, offset int) ([]domain.Score, error) {
	const op = "ScoreRepository.GetTopScores"
	log := s.log.With(slog.String("op", op))

	const query = `
		SELECT u.email, s.created_at, s.score_value
		FROM users AS u
		INNER JOIN score AS s ON u.user_id = s.user_id
		WHERE s.score_value = (SELECT MAX(score_value) FROM score AS s2 WHERE s2.user_id = s.user_id)
		ORDER BY s.score_value DESC
		LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Error("query failed")
		return nil, err
	}

	var scores []domain.Score
	for rows.Next() {
		var score domain.Score
		err := rows.Scan(&score.Email, &score.CreatedAt, &score.Score)
		if err != nil {
			log.Error("scan failed")
			return nil, err
		}

		scores = append(scores, score)
	}

	return scores, nil
}

func (s *ScoreRepository) GetTopScore(ctx context.Context, id int) {
	const op = "ScoreRepository.GetTopScore"
	log := s.log.With(slog.String("op", op))

	const query = `
		SELECT u.email, s.created_at, s.score_value
		FROM users AS u
		INNER JOIN score AS s ON u.user_id = s.user_id
		WHERE s.score_value = (SELECT MAX(score_value) FROM score AS s2 WHERE s2.user_id = s.user_id)`

	_ = log

	// TODO: implement
}

func (s *ScoreRepository) SetScore(ctx context.Context, id int, scoreValue int) error {
	const op = "ScoreRepository.SetScore"
	log := s.log.With(slog.String("op", op))

	const query = `INSERT INTO score(user_id, created_at, score_value) VALUES ($1, NOW()::timestamp, $2)`
	_, err := s.db.ExecContext(ctx, query, id, scoreValue)
	if err != nil {
		log.Error("insert failed")
		return err
	}

	return nil
}
