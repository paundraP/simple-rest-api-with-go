package routes

import (
	"paundraP/rest-api-with-go/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
}
