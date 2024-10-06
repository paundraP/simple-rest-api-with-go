package routes

import (
	"paundraP/rest-api-with-go/controllers"
	"paundraP/rest-api-with-go/middleware"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.Engine) {
	r := router.Group("/")
	r.Use(middleware.AuthMiddleware())
	r.POST("/tickets", controllers.CreateTicket)
	r.GET("/tickets", controllers.FindTicket)
	r.GET("/ticket/:id", middleware.AdminOnly(), controllers.FindTicket)
	r.PATCH("/ticket/:id", middleware.AdminOnly(), controllers.UpdateTicket)
	r.DELETE("/ticket/:id", middleware.AdminOnly(), controllers.DeleteTicket)
}
