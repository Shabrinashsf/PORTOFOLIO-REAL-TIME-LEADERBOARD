package service

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/dto"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/repository"
	"context"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type (
	LeaderboardService interface {
		SubmitScore(ctx context.Context, userID string, req dto.SubmitScoreRequest) (dto.SubmitScoreResponse, error)
		UpdateScore(ctx context.Context, userID string, req dto.SubmitScoreRequest) (dto.SubmitScoreResponse, error)
		GetRankByGame(ctx context.Context, userID string, gameID string) (dto.GetRankResponse, error)
		GetLeaderboard(ctx context.Context, limit int64, gameID string) ([]redis.Z, error)
	}

	leaderboardService struct {
		leaderboardRepo repository.LeaderboardRepository
	}
)

func NewLeaderboardService(leaderboardRepo repository.LeaderboardRepository) LeaderboardService {
	return &leaderboardService{
		leaderboardRepo: leaderboardRepo,
	}
}

func (s *leaderboardService) SubmitScore(ctx context.Context, userID string, req dto.SubmitScoreRequest) (dto.SubmitScoreResponse, error) {
	GID, err := uuid.Parse(req.GameID)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrParsingUUID
	}

	game, err := s.leaderboardRepo.GetGameByGameID(ctx, nil, GID)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrFailedGetGame
	}

	result, err := s.leaderboardRepo.SubmitScore(ctx, nil, req, userID, game.Name)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrSubmitScore
	}

	return result, nil
}

func (s *leaderboardService) UpdateScore(ctx context.Context, userID string, req dto.SubmitScoreRequest) (dto.SubmitScoreResponse, error) {
	GID, err := uuid.Parse(req.GameID)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrParsingUUID
	}

	game, err := s.leaderboardRepo.GetGameByGameID(ctx, nil, GID)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrFailedGetGame
	}

	result, err := s.leaderboardRepo.UpdateScore(ctx, nil, req, userID, game.Name)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrUpdateScore
	}

	return result, nil
}

func (s *leaderboardService) GetLeaderboard(ctx context.Context, limit int64, gameID string) ([]redis.Z, error) {
	return s.leaderboardRepo.GetLeaderboard(ctx, nil, limit, gameID)
}

func (s *leaderboardService) GetRankByGame(ctx context.Context, userID string, gameID string) (dto.GetRankResponse, error) {
	rank, err := s.leaderboardRepo.GetRankByGame(ctx, nil, gameID, userID)
	if err != nil {
		return dto.GetRankResponse{}, dto.ErrGetRank
	}

	ID, err := uuid.Parse(userID)
	if err != nil {
		return dto.GetRankResponse{}, dto.ErrParsingUUID
	}

	user, err := s.leaderboardRepo.GetUsernameByUserID(ctx, nil, ID)
	if err != nil {
		return dto.GetRankResponse{}, dto.ErrGetUsername
	}

	skor, err := s.leaderboardRepo.GetUserScore(ctx, nil, userID, gameID)
	if err != nil {
		return dto.GetRankResponse{}, dto.ErrGetUserScore
	}

	return dto.GetRankResponse{
		Name:  user.Username,
		Score: skor,
		Rank:  rank,
	}, nil
}
