package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ijul/be-monggo/bootstrap"
	domain "github.com/ijul/be-monggo/domain/request"
	response "github.com/ijul/be-monggo/domain/response"
	"github.com/ijul/be-monggo/internal/expection"
	"github.com/ijul/be-monggo/internal/password"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) SignUp(ctx *gin.Context) {
	var request domain.SignupRequest
	if err := ctx.ShouldBind(&request); expection.ErrorResponse(ctx, err, http.StatusBadRequest) {
		return
	}

	_, err := sc.SignupUsecase.GetUserByEmail(ctx, request.Email)
	if err == nil {
		err = fmt.Errorf("user already exists with the given email")
		expection.ErrorResponse(ctx, err, http.StatusConflict)
		return
	}

	encryptedPassword, err := password.HashPassword(request.Password)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	argUser := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: encryptedPassword,
	}

	err = sc.SignupUsecase.Create(ctx, &argUser)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&argUser, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&argUser, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	rsp := response.SuccessResponse{
		Status:  true,
		Message: "success signup",
		Data: domain.SignupResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	ctx.JSON(http.StatusOK, rsp)
}