package repositories

import (
	"GoREST/models"
	"errors"
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

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return artists, err
}

func (repository *ArtistRepository) Save(artist *models.Artist) (*models.Artist, error) {
	err := repository.database.Create(artist).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return artist, err
}

func (repository *ArtistRepository) FindById(id uuid.UUID) (*models.Artist, error) {
	var artist models.Artist
	err := repository.database.First(&artist, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &artist, err
}

func (repository *ArtistRepository) FindArtistByName(name string) (*models.Artist, error) {
	var artist models.Artist
	err := repository.database.Where("LOWER(name) = LOWER(?)", name).First(&artist).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &artist, err
}

func (repository *ArtistRepository) FindAllArtistAlbums(id uuid.UUID) ([]models.Album, error) {
	var artist models.Artist
	err := repository.database.
		Preload("Albums").
		Where("id = ?", id).
		First(&artist).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return artist.Albums, err
}
