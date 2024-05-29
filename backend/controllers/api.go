package controllers

import (
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	public := r.Group("/api")
	public.POST("/login", Login)
	public.POST("/register", Register)
	public.GET("/host/:id", GetHostByID)
	public.GET("/offers/accommodations", GetAccommodations)
	public.GET("/offers/accommodation/:id", GetAccommodationByID)
	public.GET("/offers/activities", GetActivities)
	public.GET("/offers/activity/:id", GetActivityByID)
	public.GET("/offers/events", GetEvents)
	public.GET("/offers/event/:id", GetEventByID)

	host := r.Group("/api/host")
	host.PUT("/:id/change/description", ChangeDescription)
	host.PUT("/:id/change/phone_number", ChangePhoneNumber)
	host.PUT("/:id/change/bank_account", ChangeBankAccount)

	activity := r.Group("api/host/activity")
	activity.Use(middlewares.JwtAuthMiddleware("host"))
	activity.POST("/create", CreateActivityOffer)
	activity.DELETE("/delete/:id", DeleteActivity)
	activity.PUT("/update/:id", UpdateActivity)
	activity.PUT("/discount/:offerID", DiscountActivity)
	activity.GET("/:id/reservations", GetCurrentActivityReservations)
	activity.POST("/:id/equipment/add", AddEquipment)
	activity.GET("/:id/offers", GetActivitiesForHost)
	activity.PUT("/price/:id", ChangeActivityPrice)

	accommodation := r.Group("api/host/accommodation")
	accommodation.Use(middlewares.JwtAuthMiddleware("host"))
	accommodation.POST("/create", CreateAccommodationOffer)
	accommodation.DELETE("/delete/:id", DeleteAccommodation)
	accommodation.PUT("/update/:id", UpdateAccommodation)
	accommodation.PUT("/discount/:offerID", DiscountAccommodation)
	accommodation.GET("/:id/reservations", GetCurrentAccommodationReservations)
	accommodation.GET("/:id/offers", GetAccommodationsForHost)
	accommodation.PUT("/price/:id", ChangeAccommodationPrice)
	accommodation.POST("/:id/facilities/add", AddGeneralFacilities)

	event := r.Group("api/host/event")
	event.Use(middlewares.JwtAuthMiddleware("host"))
	event.POST("/create", CreateEventOffer)
	event.DELETE("/delete/:id", DeleteEvent)
	event.PUT("/update/:id", UpdateEvent)
	event.PUT("/discount/:offerID", DiscountEvent)
	event.GET("/:id/reservations", GetCurrentEventReservations)
	event.GET("/:id/offers", GetEventsForHost)
	event.PUT("/price/:id", ChangeEventPrice)

	customer := r.Group("api/customer")
	customer.Use(middlewares.JwtAuthMiddleware("customer"))
	customer.GET(":id/reservations/accommodation/history", GetReservationsAccommodationHistory)
	customer.GET(":id/reservations/activity/history", GetReservationsActivityHistory)
	customer.GET(":id/reservations/event/history", GetReservationsEventHistory)
	customer.GET(":id/reservations/room/history", GetReservationsRoomHistory)
	customer.PUT(":id/change/first_name", ChangeFirstName)
	customer.PUT(":id/change/last_name", ChangeLastName)
	customer.PUT(":id/change/email", ChangeEmail)
	customer.PUT(":id/change/picture", ChangeImage)
	customer.PUT(":id/change/password", ChangePassword)
	customer.POST("/offer/accommodation/:id/rate", GradeAccommodationReservation)
	customer.POST("/offer/activity/:id/rate", GradeActivityReservation)
	customer.POST(":id/promote", PromoteToHost)

	country := r.Group("/api/country")
	country.GET("/get/all", GetAllCountries)

	generalFacility := r.Group("/api/general_facility")
	generalFacility.POST("/get", GetAllGeneralFacilities)

	roomFacility := r.Group("/api/room_facility")
	roomFacility.POST("/get", GetAllRoomFacilities)

	equipment := r.Group("/api/equipment")
	equipment.POST("/get", GetAllEquipment)

	payment := r.Group("/api/payment")
	payment.GET("/get", GetAllPayments)

	town := r.Group("/api/town")
	town.Use(middlewares.JwtAuthMiddleware("host"))
	town.POST("/add", AddTown)

	reservationActivity := r.Group("/api/reservation/activity")
	reservationActivity.GET("/get/:id", GetActivityReservationById)
	reservationActivity.GET("/:state", GetActivityReservationsByState)
	reservationActivity.POST("/add", middlewares.JwtAuthMiddleware("customer"), AddActivityReservation)
	reservationActivity.POST("/:id/:state", middlewares.JwtAuthMiddleware("host"), ChangeActivityReservationState)

	reservationEvent := r.Group("/api/reservation/event")
	reservationEvent.GET("/get/:id", GetEventReservationById)
	reservationEvent.GET("/:state", GetEventReservationsByState)
	reservationEvent.POST("/add", middlewares.JwtAuthMiddleware("customer"), AddEventReservation)
	reservationEvent.POST("/:id/:state", middlewares.JwtAuthMiddleware("host"), ChangeEventReservationState)

	reservationAccommodation := r.Group("/api/reservation/accommodation")
	reservationAccommodation.GET("/get/:id", GetAccommodationReservationById)
	reservationAccommodation.GET("/:state", GetAccommodationReservationsByState)
	reservationAccommodation.POST("/add", middlewares.JwtAuthMiddleware("customer"), AddAccommodationReservation)
	reservationAccommodation.POST("/:id/:state", middlewares.JwtAuthMiddleware("host"), ChangeAccommodationReservationState)

	reservationRoom := r.Group("/api/reservation/room")
	reservationRoom.GET("/get/:id", GetRoomReservationById)
	reservationRoom.GET("/:state", GetRoomReservationsByState)
	reservationRoom.POST("/add", middlewares.JwtAuthMiddleware("customer"), AddRoomReservation)
	reservationRoom.POST("/:id/:state", middlewares.JwtAuthMiddleware("host"), ChangeRoomReservationState)

	room := r.Group("/api/host/room")
	room.Use(middlewares.JwtAuthMiddleware("host"))
	room.POST("/create", CreateRooms)
	room.GET("/:id/get", GetRoomsForAccommodation)
	room.GET("/:id/reservations", GetCurrentRoomReservations)
	room.POST("/:id/facilities/add", AddRoomFacilities)
}
