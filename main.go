package main

import (
	"log"
	"paundraP/rest-api-with-go/models"
	"paundraP/rest-api-with-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	routes.TicketRoutes(r)
	routes.RegisterAuthRoutes(r)
	err := r.Run()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
