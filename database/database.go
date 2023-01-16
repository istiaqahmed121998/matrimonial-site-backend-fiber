package database

import (
	"log"
	"os"

	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInstance struct {
	DB *gorm.DB
}

var Database dbInstance

func ConnectDb() {

	dsn := "host=localhost user=istiaqahmed password='' dbname=go_matrimonial port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		log.Fatal("Failed to connect the database", err.Error())
		os.Exit(1)
	}
	log.Println("Connected to database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migration")
	db.Debug().AutoMigrate(&models.ContactInfo{}, &models.User{}, &models.Institution{}, &models.Department{}, &models.Degree{}, &models.Profile{}, &models.Contact{}, &models.WorkExperience{}, &models.Family{}, &models.ProfileImage{}, &models.Education{}, &models.Membership{}, &models.MembershipBenefit{}, &models.Attribute{}, &models.AttributeValue{})
	//db.Debug().AutoMigrate(&models.Profile{}, &models.ProfileImage{}, &models.Education{}, &models.User{}, &models.Institution{}, &models.Department{}, &models.Degree{})
	Database = dbInstance{DB: db}
}
