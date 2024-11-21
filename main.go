package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{"message": "Hello, You created a web app!"})
    })

    router.Run(":8080")
}