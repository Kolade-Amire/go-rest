package handlers

import (
	"errors"
	"net/http"

	"GoREST/models"
	"GoREST/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ArtistHandler struct {
	service *services.ArtistService
}

func NewArtistHandler(service *services.ArtistService) *ArtistHandler {
	return &ArtistHandler{service: service}
}

// RetrieveAllArtists handles GET /artists
func (handler *ArtistHandler) RetrieveAllArtists(context *gin.Context) {
	artists, err := handler.service.RetrieveAllArtists()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve artists"})
		return
	}
	context.JSON(http.StatusOK, artists)
}

// CreateArtist handles POST /artists
func (handler *ArtistHandler) CreateArtist(context *gin.Context) {
	var artist models.Artist
	if err := context.ShouldBindJSON(&artist); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	createdArtist, err := handler.service.CreateArtist(&artist)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create artist"})
		return
	}
	context.JSON(http.StatusCreated, createdArtist)
}

// RetrieveArtistByID handles GET /artists/:id
func (handler *ArtistHandler) RetrieveArtistByID(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	artist, err := handler.service.RetrieveArtistByID(uuidID)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve artist"})
		}
		return
	}
	c.JSON(http.StatusOK, artist)
}

// GetArtistByName handles GET /artists/name/:name
func (handler *ArtistHandler) GetArtistByName(c *gin.Context) {
	name := c.Param("name")
	artist, err := handler.service.GetArtistByName(name)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve artist"})
		}
		return
	}
	c.JSON(http.StatusOK, artist)
}

// RetrieveArtistAlbums handles GET /artists/:id/albums
func (handler *ArtistHandler) RetrieveArtistAlbums(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	albums, err := handler.service.RetrieveArtistAlbums(uuidID)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve albums"})
		}
		return
	}
	c.JSON(http.StatusOK, albums)
}

// UpdateArtist handles PUT /artists/:id
func (handler *ArtistHandler) UpdateArtist(context *gin.Context) {
	id := context.Param("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	var updates models.Artist
	if err := context.ShouldBindJSON(&updates); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	updatedArtist, err := handler.service.UpdateArtist(uuidID, &updates)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update artist"})
		return
	}
	context.JSON(http.StatusOK, updatedArtist)
}
