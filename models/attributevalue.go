package models

type AttributeValue struct {
	ID                        uint   `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	AttributeValueName        string `json:"attribute_value_name" gorm:"not null"`
	AttributeValueDescription string `json:"attribute_value_description"`
	AttributeId               uint
}
