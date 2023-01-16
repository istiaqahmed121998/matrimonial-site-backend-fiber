package models

type Degree struct {
	ID                   uint        `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	Certificate          string      `json:"certificate" gorm:"not null"`
	CertificateShortFrom string      `json:"certificate_short_from" gorm:"not null"`
	Educations           []Education `gorm:"foreignKey:DegreeId;references:ID"  gorm:"not null"`
}
