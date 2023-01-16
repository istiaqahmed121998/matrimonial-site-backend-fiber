package models

type Contact struct {
	ID            uint   `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	ContactNumber string `json:"contact_number" gorm:"index"`
	Country       string `json:"country" gorm:"default:bangladesh"`
	OwnerID       int
	OwnerType     string
}
