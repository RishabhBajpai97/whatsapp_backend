package controllers

import "github.com/gin-gonic/gin"

func LoginController() gin.HandlerFunc  {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"login route"})
	}
}
func SignOutController() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"Signout"})
	}
}
func SendOtp() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"send otp"})
	}
}
func VerifyOtp() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"verify otp"})
	}
}