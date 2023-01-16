package models

type ProfileImage struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	FileName  string `json:"file_name" gorm:"not null"`
	FilePath  string `json:"file_path" gorm:"not null"`
	Extension string `json:"extension" gorm:"not null"`
	FileSize  string `json:"file_size" gorm:"not null"`
	ProfileId uint
}
