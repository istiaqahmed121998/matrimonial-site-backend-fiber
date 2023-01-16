package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	MemberUUID    uuid.UUID `json:"profile_uuid" gorm:"type:uuid;default:gen_random_uuid()"`
	FullName      string    `json:"full_name" gorm:"not null" validate:"required,min=3,max=50"`
	Email         string    `json:"email" gorm:"not null;uniqueIndex" validate:"required,email,min=6,max=52"`
	Admin         bool      `json:"admin" gorm:"default:false"`
	ContactNumber string    `json:"contact_number" gorm:"not null;uniqueIndex" validate:"required,min=6,max=20"`
	Password      string    `json:"password" validate:"required"`
	OnBehalfOf    string    `json:"on_behalf_of" gorm:"default:'myself'" validate:"required"`
	LookingFor    string    `json:"looking_for" validate:"required"`
	verified      bool      `gorm:"default:false"`
	Profile       Profile
	CreatedAt     time.Time
}
