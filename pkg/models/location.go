package models

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null;size:255"`
	Address   string         `json:"address" gorm:"not null;size:500"`
	City      string         `json:"city" gorm:"not null;size:100"`
	State     string         `json:"state" gorm:"size:100"`
	Country   string         `json:"country" gorm:"size:100"`
	PostalCode string        `json:"postal_code" gorm:"size:20"`
	Latitude  float64        `json:"latitude" gorm:"not null"`
	Longitude float64        `json:"longitude" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relaciones
	Events []Event `json:"events,omitempty" gorm:"foreignKey:LocationID"`
}

// TableName especifica el nombre de la tabla
func (Location) TableName() string {
	return "locations"
} 