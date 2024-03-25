package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	domain "github.com/ijul/be-monggo/domain/request"
	response "github.com/ijul/be-monggo/domain/response"
	"github.com/ijul/be-monggo/internal/expection"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) Create(ctx *gin.Context) {
	var request domain.Task
	if err := ctx.ShouldBind(&request); expection.ErrorResponse(ctx, err, http.StatusBadRequest) {
		return
	}

	userID := ctx.GetString("x-user-id")
	request.ID = primitive.NewObjectID()

	userIDHex, err := primitive.ObjectIDFromHex(userID)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	request.UserID = userIDHex

	err = tc.TaskUsecase.Create(ctx, &request)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	rsp := response.SuccessResponse{
		Status:  true,
		Message: "Success create task",
		Data:    request,
	}

	ctx.JSON(http.StatusCreated, rsp)
}

func (u *TaskController) Fetch(ctx *gin.Context) {
	userID := ctx.GetString("x-user-id")

	tasks, err := u.TaskUsecase.FetchByUserID(ctx, userID)
	if expection.ErrorResponse(ctx, err, http.StatusInternalServerError) {
		return
	}

	rsp := response.SuccessResponse{
		Status:  true,
		Message: "Success fetch tasks",
		Data:    tasks,
	}

	ctx.JSON(http.StatusOK, rsp)
}
