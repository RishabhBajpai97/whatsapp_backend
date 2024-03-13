package main

import (
	"RishabhBajpai97/whatsapp_backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message":"heartbeeat"})
	})
	routes.AuthRoutes(r)
	err := r.Run()
	if err!=nil{
		log.Fatal("Error Ocurred while starting server")
	}
}