package controller

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/dto"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/service"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	LeaderboardController interface {
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
