package v1service

import (
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/database"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/models"
)

type MembershipService struct {
	Membership models.Membership
}

func (membershipService MembershipService) CreateMemberShip() map[string]interface{} {
	response := make(map[string]interface{})
	result := database.Database.DB.Create(&membershipService.Membership)
	if result.Error != nil {
		response["error"] = result.Error
	} else {
		response["data"] = membershipService.Membership
	}
	return response
}

func (membershipService MembershipService) GetMemberships() map[string]interface{} {
	response := make(map[string]interface{})
	var memberships []models.Membership
	result := database.Database.DB.Preload("MembershipBenefits").Find(&memberships)
	if result.Error != nil {
		response["error"] = result.Error
	} else {
		response["data"] = memberships
	}
	return response
}
