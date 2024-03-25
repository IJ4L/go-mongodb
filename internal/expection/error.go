package expection

import (
	"github.com/gin-gonic/gin"

	domain "github.com/ijul/be-monggo/domain/response"
)

func ErrorResponse(ctx *gin.Context, err error, statusCode int) bool {
	if err != nil {
		ctx.JSON(statusCode, domain.NewErrorResponse(err))
		return true
	}
	return false
}
