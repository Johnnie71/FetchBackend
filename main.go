package main

import (
	"backend-service/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{"message": "Hello, Welcome to the receipts app!"})
    })

    router.POST("/receipts/process", controllers.ProcessReciept)
    router.GET("/receipts/:id/points", controllers.GetRecieptsPoints)

    router.Run(":8080")
}