package main

import (
	"GoREST/config"
	"GoREST/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//initialize database
	config.InitializeDB()

	//create gin router
	router := gin.Default()

	routes.SetupArtistRoutes(router, config.DB)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
