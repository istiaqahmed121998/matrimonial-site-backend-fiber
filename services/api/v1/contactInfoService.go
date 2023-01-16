package v1service

import (
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/database"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/models"
)

type ContactInfoService struct {
	ContactInfo *models.ContactInfo
}

func (contactInfoService ContactInfoService) CreateContactInfo() map[string]interface{} {
	response := make(map[string]interface{})
	result := database.Database.DB.Create(contactInfoService.ContactInfo)
	if result.Error != nil {
		response["error"] = result.Error
	} else {
		response["data"] = *contactInfoService.ContactInfo
	}
	return response
}

func (contactInfoService ContactInfoService) GetContactInfo() map[string]interface{} {
	response := make(map[string]interface{})
	var contactinfos []models.ContactInfo
	result := database.Database.DB.Find(&contactinfos)
	if result.Error != nil {
		response["error"] = result.Error
	} else {
		response["data"] = contactinfos
	}
	return response
}
