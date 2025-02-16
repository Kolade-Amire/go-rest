package services

import (
	"GoREST/models"
	"GoREST/repositories"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlbumService struct {
	repo *repositories.AlbumRepository
}

func NewAlbumService(repository *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{repo: repository}
}

func (service *AlbumService) RetrieveAllAlbums() ([]models.Album, error) {
	albums, err := service.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get albums: %w", err)
	}
	return albums, nil
}

func (service *AlbumService) RetrieveById(id uuid.UUID) (*models.Album, error) {
	album, err := service.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
	}
	return album, nil
}

func (service *AlbumService) SearchAlbumsByName(name string) ([]models.Album, error) {
	albums, err := service.repo.FindAllByAlbumNameContains(name)

	if err != nil {
		return nil, fmt.Errorf("failed to get albums: %w", err)
	}
	return albums, nil
}

func (service *AlbumService) RetrieveByName(name string) (*models.Album, error) {
	album, err := service.repo.FindSingleByAlbumName(name)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get album: %w", err)
	}
	return album, nil
}

func (service *AlbumService) Create(album *models.Album) (*models.Album, error) {
	createdAlbum, err := service.repo.Save(album)
	if err != nil {
		return nil, fmt.Errorf("failed to create album: %w", err)
	}
	return createdAlbum, nil
}

func (service *AlbumService) Update(id uuid.UUID, updates *models.Album) (*models.Album, error) {
	existingAlbum, err := service.RetrieveById(id)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve album: %w", err)
	}

	// Apply updates
	existingAlbum.Name = updates.Name
	existingAlbum.ArtistId = updates.ArtistId
	existingAlbum.NumberOfTracks = updates.NumberOfTracks
	existingAlbum.DateReleased = updates.DateReleased

	return service.repo.Save(existingAlbum)
}
