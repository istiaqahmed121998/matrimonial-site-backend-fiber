package v1service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/database"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/models"
	v1r "github.com/istiaqahmed121998/matrimonial-site-backend-fiber/resources/api/v1"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/utils"
)

type UserService struct {
	User models.User
}

// UserList function returns the list of users
func (userService *UserService) UserList() map[string]interface{} {
	user := userService.User
	result := database.Database.DB.Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	if result.RowsAffected == 1 {
		fmt.Println(result)
	}
	response := make(map[string]interface{})

	return response
}
func CreateUser(user *models.User) map[string]interface{} {
	response := make(map[string]interface{})
	findUser := getUser(user)
	if findUser != nil {
		response["error"] = errors.New("Duplicate email existed").Error()
	} else {
		result := database.Database.DB.Create(&user)

		if result.Error != nil {
			response["error"] = result.Error
		}
		if result.RowsAffected == 1 {
			response["data"] = user
		}
	}
	return response
}

func getUser(user *models.User) *models.User {
	result := database.Database.DB.Where("email = ?", &user.Email).First(&user)
	if result.RowsAffected == 1 {
		return user
	}
	return nil
}

func AccessToken(input v1r.LoginInput) map[string]interface{} {
	response := make(map[string]interface{})
	user := getUser((&models.User{Email: input.Email}))
	if user == nil {
		response["error"] = "User is not found"
		return response
	}
	if user != nil && utils.DoPasswordsMatch(user.Password, input.Password) {
		claims := jwt.MapClaims{
			"id":       user.ID,
			"email":    user.Email,
			"fullname": user.FullName,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		}
		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			response["error"] = err
		}
		response["token"] = t
		response["user"] = map[string]interface{}{"full_name": user.FullName, "admin": user.Admin}
		return response
	}
	response["error"] = "User password is not matched"
	return response
}
