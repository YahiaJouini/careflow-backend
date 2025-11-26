package models

import (
	"time"

	"gorm.io/gorm"
)

// medical category (Dentist, Cardiologist, etc.)
type Specialty struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(100); not null; unique" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Icon        string         `gorm:"type:varchar(500); not null" json:"icon"`
	Doctors     []Doctor       `json:"-"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
