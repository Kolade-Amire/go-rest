package services

import (
	"GoREST/models"
	"GoREST/repositories"
	"github.com/google/uuid"
)

type AlbumService struct {
	repo *repositories.AlbumRepository
}

func NewAlbumService(repository *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{repo: repository}
}

func (service *AlbumService) RetrieveAllAlbums() ([]models.Album, error) {
	return service.repo.FindAll()
}

func (service *AlbumService) RetrieveById(id uuid.UUID) (*models.Album, error) {
	return service.repo.FindById(id)
}

func (service *AlbumService) SearchAlbumsByName(name string) ([]models.Album, error) {
	return service.repo.FindAllByAlbumNameContains(name)
}

func (service *AlbumService) RetrieveByName(name string) (*models.Album, error) {
	return service.repo.FindSingleByAlbumName(name)
}

func (service *AlbumService) Create(album *models.Album) (*models.Album, error) {
	return service.repo.Save(album)
}

func (service *AlbumService) Update(id uuid.UUID, updatedAlbum *models.Album) (*models.Album, error) {
	existingAlbum, err := service.RetrieveById(id)

	if err != nil {
		return nil, err
	}

	// Apply updates
	existingAlbum.Name = updatedAlbum.Name
	existingAlbum.ArtistId = updatedAlbum.ArtistId
	existingAlbum.NumberOfTracks = updatedAlbum.NumberOfTracks
	existingAlbum.DateReleased = updatedAlbum.DateReleased

	return service.repo.Save(existingAlbum)
}
