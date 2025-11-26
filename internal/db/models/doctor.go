package models

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Foreign Key to User (1:1 Relationship)
	UserID uint `gorm:"unique;not null" json:"userId"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`

	// Foreign Key to Specialty (Many-to-One Relationship)
	SpecialtyID uint      `gorm:"not null" json:"specialtyId"`
	Specialty   Specialty `json:"specialty,omitempty"`

	// Doctor specific fields
	Bio             string  `gorm:"type:text" json:"bio"`
	LicenseNumber   string  `gorm:"type:varchar(50); not null; unique" json:"licenseNumber"`
	ConsultationFee float64 `gorm:"type:decimal(10,2); default:0.00" json:"consultationFee"`

	IsAvailable bool `gorm:"default:true" json:"isAvailable"`
	IsVerified  bool `gorm:"default:false" json:"isVerified"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
