package models

type WorkExperience struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	WorkPlace    string `json:"work_place"`
	Postion      string `json:"position"`
	Duration     string `json:"duration"`
	EmployeeID   int
	EmployeeType string
}
