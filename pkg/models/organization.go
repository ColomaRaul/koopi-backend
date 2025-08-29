package models

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:255"`
	Description string         `json:"description" gorm:"type:text"`
	OwnerID     uint           `json:"owner_id" gorm:"not null"`
	IsVerified  bool           `json:"is_verified" gorm:"default:false"`
	Website     string         `json:"website" gorm:"size:255"`
	Phone       string         `json:"phone" gorm:"size:50"`
	Email       string         `json:"email" gorm:"size:255"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relaciones
	Owner  User    `json:"owner,omitempty" gorm:"foreignKey:OwnerID"`
	Events []Event `json:"events,omitempty" gorm:"foreignKey:OrganizerID"`
}

// TableName especifica el nombre de la tabla
func (Organization) TableName() string {
	return "organizations"
} 