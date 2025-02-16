package repositories

import (
	"GoREST/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArtistRepository struct {
	database *gorm.DB
}

func NewArtistRepository(db *gorm.DB) *ArtistRepository {
	return &ArtistRepository{database: db}
}

func (repository *ArtistRepository) FindAll() ([]models.Artist, error) {
	var artists []models.Artist
	err := repository.database.Find(&artists).Error

	return artists, err
}

func (repository *ArtistRepository) Save(artist *models.Artist) (*models.Artist, error) {
	err := repository.database.Create(artist).Error

	return artist, err
}

func (repository *ArtistRepository) FindById(id uuid.UUID) (*models.Artist, error) {
	var artist models.Artist
	err := repository.database.First(&artist, "id = ?", id).Error

	return &artist, err
}

func (repository *ArtistRepository) FindArtistByName(name string) (*models.Artist, error) {
	var artist models.Artist
	err := repository.database.Where("LOWER(name) = LOWER(?)", name).First(&artist).Error

	return &artist, err
}

func (repository *ArtistRepository) FindAllArtistAlbums(id uuid.UUID) ([]models.Album, error) {
	var artist models.Artist
	err := repository.database.
		Preload("Albums").
		Where("id = ?", id).
		First(&artist).Error

	return artist.Albums, err
}
