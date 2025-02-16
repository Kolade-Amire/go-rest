package routes

import (
	"GoREST/handlers"
	"GoREST/repositories"
	"GoREST/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupArtistRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize repo and service
	artistRepo := repositories.NewArtistRepository(db)
	artistService := services.NewArtistService(artistRepo)
	artistHandler := handlers.NewArtistHandler(artistService)

	// Group routes under /artists
	artistRoutes := router.Group("/artists")
	{
		artistRoutes.GET("/", artistHandler.RetrieveAllArtists)             // GET /artists
		artistRoutes.POST("/", artistHandler.CreateArtist)                  // POST /artists
		artistRoutes.GET("/:id", artistHandler.RetrieveArtistByID)          // GET /artists/:id
		artistRoutes.GET("/name/:name", artistHandler.GetArtistByName)      // GET /artists/name/:name
		artistRoutes.GET("/:id/albums", artistHandler.RetrieveArtistAlbums) // GET /artists/:id/albums
		artistRoutes.PUT("/:id", artistHandler.UpdateArtist)                // PUT /artists/:id
	}
}
