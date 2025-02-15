package repositories

import (
	"GoREST/models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlbumRepository struct {
	database *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) *AlbumRepository {
	return &AlbumRepository{database: db}
}
func (repository *AlbumRepository) FindAll() ([]models.Album, error) {
	var albums []models.Album
	err := repository.database.Find(&albums).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return albums, err
}

func (repository *AlbumRepository) Save(album *models.Album) (*models.Album, error) {
	err := repository.database.Create(album).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return album, err
}

func (repository *AlbumRepository) FindById(id uuid.UUID) (*models.Album, error) {
	var album models.Album
	err := repository.database.First(&album, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &album, err
}

func (repository *AlbumRepository) FindAllByAlbumNameContains(name string) ([]models.Album, error) {
	var albums []models.Album
	err := repository.database.Where("LOWER(name) Like LOWER(?)", "%"+name+"%").Find(&albums).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return albums, err
}

func (repository *AlbumRepository) FindSingleByAlbumName(name string) (*models.Album, error) {
	var album models.Album
	err := repository.database.Where("LOWER(name) Like LOWER(?)", name).First(&album).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &album, err
}
