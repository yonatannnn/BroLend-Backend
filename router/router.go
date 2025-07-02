package routers

import (
	middleware "brolend/infrastructure"
	"brolend/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(uc *controller.UserController) *gin.Engine {
	r := gin.Default()
	r.POST("/register", uc.Register)
	r.POST("/login", uc.Login)

	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware())
	{
		
	}

	return r
}
