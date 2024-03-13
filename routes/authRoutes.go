package routes

import (
	"RishabhBajpai97/whatsapp_backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.POST("/login", controllers.LoginController())
	incomingRoutes.POST("/signout", controllers.SignOutController())
	incomingRoutes.POST("/send-otp", controllers.SendOtp())
	incomingRoutes.POST("/verify-otp", controllers.VerifyOtp())

}