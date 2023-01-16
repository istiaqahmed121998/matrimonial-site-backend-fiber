package models

type ContactInfo struct {
	ID            uint   `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	OfficeAddress string `json:"office_address"`
	PhoneNumber   string `json:"phonenumber"`
	Email         string `json:"email_address"`
	Iframe        string `json:"iframe"`
}
