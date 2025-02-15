package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Artist struct {
	gorm.Model
	Id      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Country string    `json:"country"`
	Albums  []Album   `gorm:"foreignKey:ArtistId"`
}

type Album struct {
	gorm.Model
	Id             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name           string    `json:"name"`
	ArtistId       uuid.UUID `json:"artistId" gorm:"type:uuid"`
	NumberOfTracks int       `json:"numberOfTracks"`
	DateReleased   time.Time `json:"dateReleased"`
}
