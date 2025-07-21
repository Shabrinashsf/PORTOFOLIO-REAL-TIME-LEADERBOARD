package repository

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/dto"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/entity"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type (
	LeaderboardRepository interface {
		GetGameByGameID(ctx context.Context, tx *gorm.DB, gameID uuid.UUID) (entity.Game, error)
		SubmitScore(ctx context.Context, redisClient *redis.Client, req dto.SubmitScoreRequest, userID string, gameName string) (dto.SubmitScoreResponse, error)
		UpdateScore(ctx context.Context, redisClient *redis.Client, req dto.SubmitScoreRequest, userID string, gameName string) (dto.SubmitScoreResponse, error)
		GetLeaderboard(ctx context.Context, redisClient *redis.Client, limit int64, gameID string) ([]redis.Z, error)
		GetRankByGame(ctx context.Context, redisClient *redis.Client, gameID string, userID string) (int64, error)
		GetUsernameByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) (entity.User, error)
		GetUserScore(ctx context.Context, redisClient *redis.Client, userID string, gameID string) (float64, error)
	}

	leaderboardRepository struct {
		redisClient *redis.Client
		db          *gorm.DB
	}
)

func NewLeaderboardRepository(redisClient *redis.Client, db *gorm.DB) LeaderboardRepository {
	return &leaderboardRepository{
		redisClient: redisClient,
		db:          db,
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

func (r *leaderboardRepository) SubmitScore(ctx context.Context, redisClient *redis.Client, req dto.SubmitScoreRequest, userID string, gameName string) (dto.SubmitScoreResponse, error) {
	if redisClient == nil {
		redisClient = r.redisClient
	}

	key := fmt.Sprintf("leaderboard:%s", req.GameID)

	// ZAdd with NX (only if user not exists)
	addedCount, err := redisClient.ZAddArgs(ctx, key, redis.ZAddArgs{
		NX: true,
		Members: []redis.Z{
			{
				Member: userID,
				Score:  req.Score,
			},
		},
	}).Result()

	if err != nil {
		return dto.SubmitScoreResponse{}, err
	}

	if addedCount == 0 {
		return dto.SubmitScoreResponse{}, fmt.Errorf("score for user '%s' already exists in leaderboard", userID)
	}

	entry := dto.SubmitScoreResponse{
		Name:   gameName,
		UserID: userID,
		Score:  req.Score,
	}

	return entry, nil
}

func (r *leaderboardRepository) UpdateScore(ctx context.Context, redisClient *redis.Client, req dto.SubmitScoreRequest, userID string, gameName string) (dto.SubmitScoreResponse, error) {
	if redisClient == nil {
		redisClient = r.redisClient
	}

	key := fmt.Sprintf("leaderboard:%s", req.GameID)

	// ZAdd with NX (only update if user already exists)
	addedCount, err := redisClient.ZAddArgs(ctx, key, redis.ZAddArgs{
		XX: true,
		Members: []redis.Z{
			{
				Member: userID,
				Score:  req.Score,
			},
		},
	}).Result()

	if err != nil {
		return dto.SubmitScoreResponse{}, err
	}

	if addedCount == 0 {
		return dto.SubmitScoreResponse{}, fmt.Errorf("cannot update score: user '%s' does not exist in leaderboard", userID)
	}

	entry := dto.SubmitScoreResponse{
		Name:   gameName,
		UserID: userID,
		Score:  req.Score,
	}

	return entry, nil
}

func (r *leaderboardRepository) GetLeaderboard(ctx context.Context, redisClient *redis.Client, limit int64, gameID string) ([]redis.Z, error) {
	if redisClient == nil {
		redisClient = r.redisClient
	}

	key := fmt.Sprintf("leaderboard:%s", gameID)

	return r.redisClient.ZRevRangeWithScores(ctx, key, 0, limit-1).Result()
}

func (r *leaderboardRepository) GetRankByGame(ctx context.Context, redisClient *redis.Client, gameID string, userID string) (int64, error) {
	if redisClient == nil {
		redisClient = r.redisClient
	}

	key := fmt.Sprintf("leaderboard:%s", gameID)

	rank, err := redisClient.ZRevRank(ctx, key, userID).Result()
	if err != nil {
		return 0, err
	}

	return rank + 1, nil
}

func (r *leaderboardRepository) GetUsernameByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) (entity.User, error) {
	var user entity.User

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *leaderboardRepository) GetUserScore(ctx context.Context, redisClient *redis.Client, userID string, gameID string) (float64, error) {
	key := fmt.Sprintf("leaderboard:%s", gameID)

	score, err := redisClient.ZScore(ctx, key, userID).Result()
	if err != nil {
		return 0, err
	}

	return score, nil
}
