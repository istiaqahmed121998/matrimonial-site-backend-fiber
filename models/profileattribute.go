package models

type ProfileAttribute struct {
	PersonID           uint
	AddressID          uint
	ShowInFrontProfile bool `gorm:"default:true"`
}
