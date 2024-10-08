package routes

import (
	"paundraP/rest-api-with-go/controllers"
	"paundraP/rest-api-with-go/middleware"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.Engine) {
	r := router.Group("/")
	r.GET("", controllers.HomePage)
	r.Use(middleware.AuthMiddleware())
	r.POST("/tickets", controllers.OrderTicket)
	r.GET("/tickets", controllers.FindTicket)
	r.GET("/ticket/:id", middleware.AdminOnly(), controllers.FindTicket)
	r.PATCH("/ticket/:id", middleware.AdminOnly(), controllers.UpdateTicket)
	r.DELETE("/ticket/:id", middleware.AdminOnly(), controllers.DeleteTicket)
}
