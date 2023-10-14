package main

import (
	"Participate/controllers"
	"Participate/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	// TODO To be deleted, left as an example
	//r.GET("/users/:id", func(c *gin.Context) {
	//	var user Osoba
	//
	//	if err := db.Find(&user).Error; err != nil {
	//		err.Error()
	//		c.JSON(500, gin.H{"error": "Failed to fetch users"})
	//		return
	//	}
	//
	//	c.JSON(200, user)
	//})

	err := r.Run(":8080")
	if err != nil {
		err.Error()
		return
	} // listen and serve on 0.0.0.0:8080
}
