package controllers

import (
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	public := r.Group("/api")
	public.POST("/login", Login)
	public.GET("/offers", GetOffers)

	register := r.Group("api/register")
	register.POST("/customer", RegisterCustomer)
	register.POST("/host", RegisterHost)

	host := r.Group("api/host")
	host.Use(middlewares.JwtAuthMiddleware("host"))
	host.POST("/create", CreateOffer)
	host.DELETE("/delete/:id", DeleteOffer)
	host.PUT("/update/:id", UpdateOffer)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware("admin"))
	protected.GET("/user", CurrentUser)

	country := r.Group("/api/country")
	country.GET("/get/all", GetAllCountry)
}
