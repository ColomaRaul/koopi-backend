package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	EventID   uint           `json:"event_id" gorm:"not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	Status    string         `json:"status" gorm:"default:'pending';size:50"`
	Quantity  int            `json:"quantity" gorm:"default:1"`
	TotalPrice float64       `json:"total_price" gorm:"type:decimal(10,2)"`
	Notes     string         `json:"notes" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relaciones
	Event Event `json:"event,omitempty" gorm:"foreignKey:EventID"`
	User  User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName especifica el nombre de la tabla
func (Booking) TableName() string {
	return "bookings"
}

// BeforeCreate hook para GORM
func (b *Booking) BeforeCreate(tx *gorm.DB) error {
	if b.Status == "" {
		b.Status = "pending"
	}
	if b.Quantity == 0 {
		b.Quantity = 1
	}
	return nil
} 