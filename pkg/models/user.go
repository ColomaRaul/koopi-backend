package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"unique;not null;size:255"`
	Password  string         `json:"-" gorm:"not null;size:255"`
	Name      string         `json:"name" gorm:"not null;size:255"`
	Role      string         `json:"role" gorm:"default:'user';size:50"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relaciones
	Organizations []Organization `json:"organizations,omitempty" gorm:"foreignKey:OwnerID"`
	Bookings      []Booking      `json:"bookings,omitempty" gorm:"foreignKey:UserID"`
}

// TableName especifica el nombre de la tabla
func (User) TableName() string {
	return "users"
}

// BeforeCreate hook para GORM
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Role == "" {
		u.Role = "user"
	}
	return nil
} 