package models

type MembershipBenefit struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	Benefit      string `json:"benefit" gorm:"unique"`
	Has          bool   `json:"has"`
	MembershipId uint
}
