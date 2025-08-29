package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null;size:255"`
	Description string         `json:"description" gorm:"type:text"`
	CategoryID  uint           `json:"category_id" gorm:"not null"`
	OrganizerID uint           `json:"organizer_id" gorm:"not null"`
	LocationID  uint           `json:"location_id" gorm:"not null"`
	StartDate   time.Time      `json:"start_date" gorm:"not null"`
	EndDate     time.Time      `json:"end_date" gorm:"not null"`
	Price       float64        `json:"price" gorm:"default:0;type:decimal(10,2)"`
	Capacity    int            `json:"capacity"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	IsFree      bool           `json:"is_free" gorm:"default:true"`
	ImageURL    string         `json:"image_url" gorm:"size:500"`
	Website     string         `json:"website" gorm:"size:255"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relaciones
	Category   EventCategory `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Organizer  Organization  `json:"organizer,omitempty" gorm:"foreignKey:OrganizerID"`
	Location   Location      `json:"location,omitempty" gorm:"foreignKey:LocationID"`
	Bookings   []Booking     `json:"bookings,omitempty" gorm:"foreignKey:EventID"`
}

// TableName especifica el nombre de la tabla
func (Event) TableName() string {
	return "events"
}

// BeforeCreate hook para GORM
func (e *Event) BeforeCreate(tx *gorm.DB) error {
	if e.Price == 0 {
		e.IsFree = true
	} else {
		e.IsFree = false
	}
	return nil
} 