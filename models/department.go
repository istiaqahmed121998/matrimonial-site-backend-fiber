package models

type Department struct {
	ID            uint        `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	DeptName      string      `json:"dept_name"`
	DeptShortFrom string      `json:"dept_short_from"`
	Educations    []Education `gorm:"foreignKey:DepartmentId;references:ID"  gorm:"not null"`
}
