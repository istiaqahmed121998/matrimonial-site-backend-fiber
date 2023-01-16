package models

type Institution struct {
	ID              uint        `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	InstitutionName string      `json:"institution_name" gorm:"not null"`
	Country         string      `json:"country" gorm:"default:bangladesh"`
	Educations      []Education `gorm:"foreignKey:InstitutionId;references:ID"  gorm:"not null"`
}
