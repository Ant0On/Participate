package main

import (
	"log"

	"Participate/controllers"
	"Participate/middlewares"
	"Participate/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/customer", controllers.CurrentCustomer)

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

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	} // listen and serve on 0.0.0.0:8080
}
