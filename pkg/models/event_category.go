package models

import (
	"time"

	"gorm.io/gorm"
)

type EventCategory struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:100;unique"`
	Description string         `json:"description" gorm:"type:text"`
	Icon        string         `json:"icon" gorm:"size:100"`
	Color       string         `json:"color" gorm:"size:7"` // Hex color code
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relaciones
	Events []Event `json:"events,omitempty" gorm:"foreignKey:CategoryID"`
}

// TableName especifica el nombre de la tabla
func (EventCategory) TableName() string {
	return "event_categories"
} 