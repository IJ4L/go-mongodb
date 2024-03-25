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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
