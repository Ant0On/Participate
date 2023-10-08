package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Osoba struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Imie string `json:"imie"`
}

func main() {
	dsn := "user=postgres dbname=participate port=5432"

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		var user Osoba

		if err := db.Find(&user).Error; err != nil {
			err.Error()
			c.JSON(500, gin.H{"error": "Failed to fetch users"})
			return
		}

		c.JSON(200, user)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
