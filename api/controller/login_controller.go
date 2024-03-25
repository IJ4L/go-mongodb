package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ijul/be-monggo/bootstrap"
	domain "github.com/ijul/be-monggo/domain/request"
	response "github.com/ijul/be-monggo/domain/response"
	"github.com/ijul/be-monggo/internal/expection"
	"github.com/ijul/be-monggo/internal/password"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(ctx *gin.Context) {
	var request domain.LoginRequest
	if err := ctx.ShouldBindJSON(&request); expection.ErrorResponse(ctx, err, http.StatusBadRequest) {
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(ctx, request.Email)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	err = password.CheckPassword(request.Password, user.Password)
	if expection.ErrorResponse(ctx, err, http.StatusUnauthorized) {
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	rsp := response.SuccessResponse{
		Status:  true,
		Message: "login success",
		Data: domain.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	ctx.JSON(http.StatusOK, rsp)
}
