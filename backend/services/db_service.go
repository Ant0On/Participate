package services

import (
	"errors"
	"fmt"
	"log"
	"os"

	"backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DbHost, DbUser, DbPassword, DbName, DbPort)

	DB, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("Connection error:", err)
	}
	fmt.Println("We are connected to the database")

	if err = DB.AutoMigrate(&models.Customer{}); err != nil {
		log.Fatal("Migration error:", err)
		return
	}
}

func GetCustomerByID(uid uint) (*models.Customer, error) {
	var c models.Customer

	if err := DB.First(&c, uid).Error; err != nil {
		return &c, errors.New("user not found")
	}

	return &c, nil
}

func CheckCustomerEmailExists(email string, c *models.Customer) error {
	if err := DB.Model(models.Customer{}).Where("email = ?", email).Take(&c).Error; err != nil {
		return fmt.Errorf("DB.Model.Where.Take: %w", err)
	}
	return nil
}

func CheckHostEmailExists(email string, h *models.Host) error {
	if err := DB.Model(models.Host{}).Where("email = ?", email).Take(&h).Error; err != nil {
		return fmt.Errorf("DB.Model.Where.Take: %w", err)
	}
	return nil
}

func SaveCustomer(c *models.Customer) (*models.Customer, error) {
	if err := DB.Create(&c).Error; err != nil {
		return &models.Customer{}, fmt.Errorf("DB.Create: %w", err)
	}
	return c, nil
}

func SaveHost(h *models.Host) (*models.Host, error) {
	if err := DB.Create(&h).Error; err != nil {
		return &models.Host{}, fmt.Errorf("DB.Create: %w", err)
	}
	return h, nil
}
