package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Id struct {
	Id uuid.UUID `json:"id" gorm:"type:text;primaryKey"`
}

// BeforeCreate Hook for UUID Generation
func (id *Id) BeforeCreate(db *gorm.DB) (err error) {
	if id.Id == uuid.Nil {
		id.Id = uuid.New()
	}
	return
}

type Artist struct {
	gorm.Model
	Id
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Country string  `json:"country"`
	Albums  []Album `gorm:"foreignKey:ArtistId"`
}

type Album struct {
	gorm.Model
	Id
	Name           string    `json:"name"`
	ArtistId       uuid.UUID `json:"artistId" gorm:"type:uuid"`
	NumberOfTracks int       `json:"numberOfTracks"`
	DateReleased   time.Time `json:"dateReleased"`
}
