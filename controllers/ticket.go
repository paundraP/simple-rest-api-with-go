package controllers

import (
	"log"
	"net/http"
	"paundraP/rest-api-with-go/database"
	"paundraP/rest-api-with-go/models"

	"github.com/gin-gonic/gin"
)

type CreateTicketInput struct {
	Name   string `json:"name" binding:"required"`
	NIK    int64  `json:"nik" binding:"required"`
	Amount int64  `json:"amount" binding:"required"`
}

type UpdateTicketInput struct {
	Name   string `json:"name"`
	NIK    int64  `json:"nik"`
	Amount int64  `json:"amount"`
	Status bool   `json:"status"`
}

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "welcome!"})
}

func FindTickets(c *gin.Context) {
	var user models.User
	if val, exists := c.Get("user"); exists {
		user = val.(models.User)
	}

	if user.Role == "admin" {
		var tickets []models.Ticket
		if err := database.DB.Find(&tickets).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve tickets"})
			return
		}
		c.JSON(http.StatusOK, tickets)
	}
	if user.Role == "user" {
		var tickets []models.Ticket
		if err := database.DB.Where("id = ?", user.ID).Find(&tickets).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve tickets"})
			return
		}
		c.JSON(http.StatusOK, tickets)
	}
}

func OrderTicket(c *gin.Context) {
	var input CreateTicketInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if val, exists := c.Get("user"); exists {
		user = val.(models.User)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var existingTicket models.Ticket
	if err := database.DB.Where("id = ?", user.ID).First(&existingTicket).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have already ordered a ticket"})
		return
	}
	if input.Amount > 5 {
		input.Amount = 5
	}
	ticket := models.Ticket{
		Name:   input.Name,
		NIK:    input.NIK,
		Amount: input.Amount,
		ID:     user.ID,
		Status: false,
	}

	result := database.DB.Create(&ticket)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ticket})
}

func FindTicket(c *gin.Context) {
	var ticket models.Ticket

	if err := database.DB.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Ticket tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ticket})
}

func UpdateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := database.DB.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Ticket tidak ditemukan!"})
		return
	}

	var input UpdateTicketInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&ticket).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": ticket})
}

func DeleteTicket(c *gin.Context) {
	var ticket models.Ticket

	if err := database.DB.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
		log.Println("Error finding record:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&ticket)
	c.JSON(http.StatusOK, gin.H{"data": "terhapus"})
}
