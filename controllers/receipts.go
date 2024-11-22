package controllers

import (
	"backend-service/models"
	"backend-service/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var points = make(map[string]int)

func ProcessReciept(c *gin.Context) {
	var receipt models.Reciept

	if err := c.ShouldBindJSON(&receipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	total_points, err := services.CalculatePoints(receipt); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
	}
	
	receiptId := uuid.New().String()
	
	points[receiptId] = total_points

	c.JSON(http.StatusOK, gin.H{"id": receiptId})
}

func GetRecieptsPoints(c *gin.Context) {
	receiptID := c.Param("id")

	if _, ok := points[receiptID]; ok {
			c.JSON(http.StatusOK, gin.H{"points": points[receiptID]})
	} else {
			fmt.Println("ID not found:", receiptID)
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Receipt ID %s not found", receiptID)})
	}

}