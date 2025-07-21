package dto

import (
	"errors"
)

const (
	// SUCCESS
	MESSAGE_SUCCESS_SUBMIT_SCORE    = "success submit score"
	MESSAGE_SUCCESS_UPDATE_SCORE    = "success update score"
	MESSAGE_SUCCESS_GET_RANK        = "success get rank"
	MESSAGE_SUCCESS_GET_LEADERBOARD = "success get leaderboard"

	// FAILED
	MESSAGE_FAILED_SUBMIT_SCORE    = "failed submit score"
	MESSAGE_FAILED_UPDATE_SCORE    = "failed update score"
	MESSAGE_FAILED_GET_RANK        = "failed get rank"
	MESSAGE_FAILED_GET_LEADERBOARD = "failed get leaderboard"
)

var (
	ErrUserGameExists = errors.New("failed record already create")
	ErrFailedGetGame  = errors.New("failed get game from db")
	ErrSubmitScore    = errors.New("failed submit score")
	ErrGetUserGame    = errors.New("failed get user_game")
	ErrUpdateScore    = errors.New("failed update score")
	ErrGetRank        = errors.New("failed get rank")
	ErrGetUsername    = errors.New("failed get username")
	ErrGetUserScore   = errors.New("failed get user score")
)

type (
	SubmitScoreRequest struct {
		GameID string  `json:"game_id" form:"game_id" binding:"required"`
		Score  float64 `json:"score" form:"score" binding:"required"`
	}

	SubmitScoreResponse struct {
		Name   string  `json:"name" form:"name"`
		UserID string  `json:"user_id" form:"user_id"`
		Score  float64 `json:"score" form:"score"`
	}

	GetRankResponse struct {
		Name  string  `json:"name" form:"name"`
		Score float64 `json:"score" form:"score"`
		Rank  int64   `json:"rank" form:"rank"`
	}
)
