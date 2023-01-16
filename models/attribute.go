package models

type Attribute struct {
	ID                   uint             `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	AttributeName        string           `json:"attribute_name" gorm:"not null"`
	AttributeDescription string           `json:"attribute_description" gorm:"not null"`
	AtributeValues       []AttributeValue `gorm:"foreignKey:AttributeId;references:ID"  gorm:"not null"`
	Profiles             []*Profile       `gorm:"many2many:profile_attributes;"`
}
