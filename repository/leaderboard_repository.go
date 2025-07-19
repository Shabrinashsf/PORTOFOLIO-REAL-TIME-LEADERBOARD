package repository

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/entity"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	LeaderboardRepository interface {
		GetGameByGameID(ctx context.Context, tx *gorm.DB, gameID uuid.UUID) (entity.Game, error)
		GetUserGameByUserIDAndGameID(ctx context.Context, tx *gorm.DB, userID uuid.UUID, gameId uuid.UUID) (entity.UserGame, error)
		SubmitScore(ctx context.Context, tx *gorm.DB, userGame entity.UserGame) (entity.UserGame, error)
		UpdateScore(ctx context.Context, tx *gorm.DB, userGame entity.UserGame) (entity.UserGame, error)
	}

	leaderboardRepository struct {
		db *gorm.DB
	}
)

func NewLeaderboardRepository(db *gorm.DB) LeaderboardRepository {
	return &leaderboardRepository{
		db: db,
	}
}

func (r *leaderboardRepository) GetGameByGameID(ctx context.Context, tx *gorm.DB, gameID uuid.UUID) (entity.Game, error) {
	var game entity.Game

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("id = ?", gameID).Take(&game).Error; err != nil {
		return entity.Game{}, err
	}

	return game, nil
}

func (r *leaderboardRepository) GetUserGameByUserIDAndGameID(ctx context.Context, tx *gorm.DB, userID uuid.UUID, gameId uuid.UUID) (entity.UserGame, error) {
	var userGame entity.UserGame

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("user_id = ? AND game_id = ?", userID, gameId).Take(&userGame).Error; err != nil {
		return entity.UserGame{}, err
	}

	return userGame, nil
}

func (r *leaderboardRepository) SubmitScore(ctx context.Context, tx *gorm.DB, userGame entity.UserGame) (entity.UserGame, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&userGame).Error; err != nil {
		return entity.UserGame{}, err
	}

	return userGame, nil
}

func (r *leaderboardRepository) UpdateScore(ctx context.Context, tx *gorm.DB, userGame entity.UserGame) (entity.UserGame, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("id = ?", userGame.ID).Updates(&userGame).Error; err != nil {
		return entity.UserGame{}, err
	}

	return userGame, nil
}
