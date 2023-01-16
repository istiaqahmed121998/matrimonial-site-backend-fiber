package models

type Education struct {
	ID            uint   `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	FromYear      string `json:"from_year"`
	ToYear        string `json:"to_year"`
	Result        int    `json:"result"`
	OutOf         int    `json:"out_of"`
	Completed     bool   `json:"completed" gorm:"default:true"`
	ProfileId     uint
	DegreeId      uint
	InstitutionId uint
	DepartmentId  uint
}
