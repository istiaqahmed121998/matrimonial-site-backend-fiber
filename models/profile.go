package models

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID                  uint             `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	ProfileUUID         uuid.UUID        `json:"profile_uuid" gorm:"type:uuid;default:gen_random_uuid()"`
	FullName            string           `json:"full_name" gorm:"not null; index"`
	Complexion          string           `json:"complexion"`
	Height              int              `json:"height" gorm:"not null"`
	Weight              int              `json:"weight" gorm:"not null"`
	Religion            string           `json:"religion" gorm:"not null; default:islam"`
	MaritalStatus       string           `json:"marital_status" gorm:"not null; default:unmarried"`
	BloodGroup          string           `json:"blood_group"  gorm:"not null"`
	DateOfBirth         time.Time        `json:"date_of_birth"  gorm:"not null"`
	Hometown            string           `json:"hometown"  gorm:"not null"`
	PresentAddress      string           `json:"present_address"  gorm:"not null"`
	PermanentAddress    string           `json:"permanent_address"  gorm:"not null"`
	CurrentInBangladesh bool             `json:"current_in_bangladesh"  gorm:"not null;default:true"`
	Contacts            []Contact        `gorm:"polymorphic:Owner;polymorphicValue:profile"`
	Educations          []Education      `gorm:"foreignKey:ProfileId;references:ID"  gorm:"not null"`
	WorkExperience      []WorkExperience `gorm:"polymorphic:Employee;polymorphicValue:profile"`
	Family              []Family         `gorm:"foreignKey:ProfileId;references:ID"`
	Images              []ProfileImage   `gorm:"foreignKey:ProfileId;references:ID"`
	Attributes          []*Attribute     `gorm:"many2many:profile_attributes;"`
	UserID              uint
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
