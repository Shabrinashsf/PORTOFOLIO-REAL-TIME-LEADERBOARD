package controller

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/dto"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/service"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	LeaderboardController interface {
		SubmitScore(ctx *gin.Context)
		UpdateScore(ctx *gin.Context)
		GetLeaderboard(ctx *gin.Context)
		GetRankByGame(ctx *gin.Context)
	}

	leaderboardController struct {
		leaderboardService service.LeaderboardService
	}
)

func NewLeaderboardController(lc service.LeaderboardService) LeaderboardController {
	return &leaderboardController{
		leaderboardService: lc,
	}
}

func (c *leaderboardController) SubmitScore(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)
	if userID == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var req dto.SubmitScoreRequest

	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := c.leaderboardService.SubmitScore(ctx.Request.Context(), userID, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_SUBMIT_SCORE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_SUBMIT_SCORE, response)
	ctx.JSON(http.StatusOK, res)
}

func (c *leaderboardController) UpdateScore(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)
	if userID == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var req dto.SubmitScoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := c.leaderboardService.UpdateScore(ctx, userID, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_SCORE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_SCORE, response)
	ctx.JSON(http.StatusOK, res)
}

func (c *leaderboardController) GetLeaderboard(ctx *gin.Context) {
	limitStr := ctx.Query("limit")
	gameIdStr := ctx.Query("game")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		limit = 10
	}

	players, err := c.leaderboardService.GetLeaderboard(ctx, limit, gameIdStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LEADERBOARD, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LEADERBOARD, players)
	ctx.JSON(http.StatusOK, res)
}

func (c *leaderboardController) GetRankByGame(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)
	if userID == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	gameID := ctx.Param("game")
	if gameID == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "game_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := c.leaderboardService.GetRankByGame(ctx, userID, gameID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_RANK, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_RANK, response)
	ctx.JSON(http.StatusOK, res)
}
