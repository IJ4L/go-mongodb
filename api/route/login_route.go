package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ijul/be-monggo/api/controller"
	"github.com/ijul/be-monggo/bootstrap"
	domain "github.com/ijul/be-monggo/domain/request"
	"github.com/ijul/be-monggo/mongo"
	"github.com/ijul/be-monggo/repository"
	"github.com/ijul/be-monggo/usecase"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}