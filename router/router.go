package routers

import (
	"brolend/controller"
	middleware "brolend/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(uc *controller.UserController, debtController *controller.DebtController) *gin.Engine {
	r := gin.Default()
	r.POST("/register", uc.Register)
	r.POST("/login", uc.Login)
	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.PUT("/user", uc.Update)
		auth.DELETE("/user/:user_id", uc.DeleteUser) 
		auth.GET("/user/:username", uc.FindUserByUsername)
		auth.GET("/user/id/:id", uc.FindUserByID)
		auth.POST("/debt", debtController.CreateDebt)
		auth.POST("/debt/:id/accept", debtController.AcceptDebt)
		auth.POST("/debt/:id/reject", debtController.RejectDebt)
		auth.POST("/debt/:id/request-paid", debtController.RequestPaidApproval)
		auth.POST("/debt/:id/approve-payment", debtController.ApprovePayment)
		auth.POST("/debt/:id/reject-payment", debtController.RejectPaymentRequest)
		auth.GET("/debt/net", debtController.GetNetAmounts)
		auth.GET("/debt/history", debtController.GetHistory)
		auth.GET("/debt/active-incoming", debtController.GetActiveIncoming)
		auth.GET("/debt/active-outgoing", debtController.GetActiveOutgoing)
		auth.GET("/debt/incoming-requests", debtController.GetIncomingRequests)
	}

	return r
}
