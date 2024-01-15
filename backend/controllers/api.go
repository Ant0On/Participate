package controllers

import (
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	public := r.Group("/api")
	public.POST("/login", Login)
	public.GET("/offers", GetOffers)
	public.GET("/offers/:id", GetOfferByID)

	register := r.Group("api/register")
	register.POST("/customer", RegisterCustomer)
	register.POST("/host", RegisterHost)

	host := r.Group("api/host")
	host.Use(middlewares.JwtAuthMiddleware("host"))
	host.POST("/create", CreateOffer)
	host.DELETE("/delete/:id", DeleteOffer)
	host.PUT("/update/:id", UpdateOffer)

	country := r.Group("/api/country")
	country.GET("/get/all", GetAllCountries)

	animal := r.Group("/api/animal")
	animal.POST("/add", AddAnimal)

	discount := r.Group("/api/discount")
	discount.POST("/add", AddDiscount)

	payment := r.Group("/api/payment")
	payment.POST("/get", GetAllPayments)

	town := r.Group("/api/town")
	town.POST("/add", AddTown)
	town.GET("/get/:id", GetTownByID)

	townType := r.Group("/api/town_type")
	townType.GET("/get", GetAllTownTypes)

	grade := r.Group("/api/grade")
	grade.GET("/get", GetAllGrades)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware("admin"))
	protected.GET("/user", CurrentUser)
}
