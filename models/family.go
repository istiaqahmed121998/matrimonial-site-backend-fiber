package models

type Family struct {
	ID             uint             `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	Relation       string           `json:"relation" gorm:"not null"`
	WorkExperience []WorkExperience `gorm:"polymorphic:Employee;polymorphicValue:family"`
	ProfileId      uint
}
