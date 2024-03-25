package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ijul/be-monggo/api/middleware"
	"github.com/ijul/be-monggo/bootstrap"
	"github.com/ijul/be-monggo/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	NewTaskRouter(env, timeout, db, protectedRouter)
}
