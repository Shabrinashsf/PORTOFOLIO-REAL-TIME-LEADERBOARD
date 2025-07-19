package dto

import (
	"errors"
)

const (
	// SUCCESS
	MESSAGE_SUCCESS_SUBMIT_SCORE = "success submit score"
	MESSAGE_SUCCESS_UPDATE_SCORE = "success update score"

	// FAILED
	MESSAGE_FAILED_SUBMIT_SCORE = "failed submit score"
	MESSAGE_FAILED_UPDATE_SCORE = "failed update score"
)

var (
	ErrUserGameExists = errors.New("failed record already create")
	ErrFailedGetGame  = errors.New("failed get game from db")
	ErrSubmitScore    = errors.New("failed submit score")
	ErrGetUserGame    = errors.New("failed get user_game")
	ErrUpdateScore    = errors.New("failed update score")
)

type (
	SubmitScoreRequest struct {
		GameID string `json:"game_id" form:"game_id" binding:"required"`
		Score  int    `json:"score" form:"score" binding:"required"`
	}

	SubmitScoreResponse struct {
		Name  string `json:"name" form:"name"`
		Score int    `json:"score" form:"score"`
	}
)
