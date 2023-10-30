package controllers

import (
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	public := r.Group("/api")
	public.POST("/login", Login)

	register := r.Group("api/register")
	register.POST("/customer", RegisterCustomer)
	register.POST("/host", RegisterHost)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", CurrentUser)
}
