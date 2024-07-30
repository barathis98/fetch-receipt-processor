package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"receipt-processor/internal/model"
	"receipt-processor/internal/service"
	"receipt-processor/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func ProcessReceipt(c *gin.Context) {
	// var receipt model.Receipt
	fmt.Println("Processing receipt")
	data, err := utils.ParseFields(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var receipt model.Receipt
	updatedData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode data into JSON"})
		return
	}
	if err := json.Unmarshal(updatedData, &receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode JSON into struct"})
		return
	}

	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase date format"})
		return
	}

	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase time format"})
		return
	}

	id, err := service.SaveReceipt(receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save receipt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, err := service.GetPoints(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
