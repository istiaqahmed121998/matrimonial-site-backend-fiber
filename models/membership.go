package models

type Membership struct {
	ID                 uint                `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	MembershipType     string              `json:"membership_type" validate:"required" gorm:"not null;unique"`
	Price              float32             `json:"price"`
	Validity           string              `json:"validity"  validate:"required"`
	AccessLimit        string              `json:"access_limit"  validate:"required"`
	MembershipBenefits []MembershipBenefit `json:"membership_benefits" gorm:"foreignKey:MembershipId;references:ID"`
}
