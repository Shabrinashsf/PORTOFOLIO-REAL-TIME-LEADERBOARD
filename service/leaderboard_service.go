package service

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/dto"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/entity"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/repository"
	"context"

	"github.com/google/uuid"
)

type (
	LeaderboardService interface {
		SubmitScore(ctx context.Context, userID string, req dto.SubmitScoreRequest) (dto.SubmitScoreResponse, error)
		UpdateScore(ctx context.Context, userID string, req dto.SubmitScoreRequest) (dto.SubmitScoreResponse, error)
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

	ID, err := uuid.Parse(userID)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrParsingUUID
	}

	_, err = s.leaderboardRepo.GetUserGameByUserIDAndGameID(ctx, nil, ID, GID)
	if err == nil {
		return dto.SubmitScoreResponse{}, dto.ErrUserGameExists
	}

	scoreGame := entity.UserGame{
		UserID: ID,
		GameID: GID,
		Score:  req.Score,
	}

	entry, err := s.leaderboardRepo.SubmitScore(ctx, nil, scoreGame)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrSubmitScore
	}

	return dto.SubmitScoreResponse{
		Name:  game.Name,
		Score: entry.Score,
	}, nil
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

	ID, err := uuid.Parse(userID)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrParsingUUID
	}

	userGame, err := s.leaderboardRepo.GetUserGameByUserIDAndGameID(ctx, nil, ID, GID)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrGetUserGame
	}

	userGame.Score = req.Score

	result, err := s.leaderboardRepo.UpdateScore(ctx, nil, userGame)
	if err != nil {
		return dto.SubmitScoreResponse{}, dto.ErrUpdateScore
	}

	return dto.SubmitScoreResponse{
		Name:  game.Name,
		Score: result.Score,
	}, nil
}
