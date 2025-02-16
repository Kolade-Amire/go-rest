package services

import (
	"errors"
	"fmt"

	"GoREST/models"
	"GoREST/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("artist not found")
)

type ArtistService struct {
	repo *repositories.ArtistRepository
}

func NewArtistService(repo *repositories.ArtistRepository) *ArtistService {
	return &ArtistService{repo: repo}
}

func (service *ArtistService) RetrieveAllArtists() ([]models.Artist, error) {
	artists, err := service.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get artists: %w", err)
	}
	return artists, nil
}

func (service *ArtistService) CreateArtist(artist *models.Artist) (*models.Artist, error) {
	createdArtist, err := service.repo.Save(artist)
	if err != nil {
		return nil, fmt.Errorf("failed to create artist: %w", err)
	}
	return createdArtist, nil
}

func (service *ArtistService) RetrieveArtistByID(id uuid.UUID) (*models.Artist, error) {
	artist, err := service.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get artist by ID: %w", err)
	}
	return artist, nil
}

func (service *ArtistService) GetArtistByName(name string) (*models.Artist, error) {
	artist, err := service.repo.FindArtistByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get artist by name: %w", err)
	}
	return artist, nil
}

func (service *ArtistService) RetrieveArtistAlbums(id uuid.UUID) ([]models.Album, error) {
	albums, err := service.repo.FindAllArtistAlbums(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get artist albums: %w", err)
	}
	return albums, nil
}

func (service *ArtistService) UpdateArtist(id uuid.UUID, updates *models.Artist) (*models.Artist, error) {
	existingArtist, err := service.RetrieveArtistByID(id)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve artist: %w", err)
	}

	existingArtist.Name = updates.Name
	existingArtist.Country = updates.Country
	existingArtist.Albums = append(existingArtist.Albums, updates.Albums...)

	return service.repo.Save(existingArtist)
}
