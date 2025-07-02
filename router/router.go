package routers

import (
	"brolend/controller"
	middleware "brolend/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(uc *controller.UserController) *gin.Engine {
	r := gin.Default()
	r.POST("/register", uc.Register)
	r.POST("/login", uc.Login)
	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.PUT("/user", uc.Update)
		auth.DELETE("/user/:user_id", uc.DeleteUser)
		auth.GET("/user/:username", uc.FindUserByUsername)
	}

	return r
}
