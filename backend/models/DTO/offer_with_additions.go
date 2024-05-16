package DTO

type AccommodationRoom struct {
	Title             string  `gorm:"size:100;not null" json:"title" binding:"required,min=2,max=100"`
	PricePerDay       float64 `gorm:"not null" json:"price_per_day" binding:"required,min=1"`
	RoomName          string  `gorm:"not null" json:"room_name" binding:"required,min=3"`
	RoomDescription   string  `gorm:"not null" json:"room_description" binding:"required,min=10"`
	RoomFacilities    string  `gorm:"not null" json:"room_facilities" binding:"required"`
	Capacity          int     `gorm:"not null" json:"room_capacity" binding:"required,gt=0"`
	GeneralFacilities string  `gorm:"not null" json:"general_facilities" binding:"required"`
}

type ActivityEquipment struct {
	Title       string  `gorm:"size:100;not null" json:"title" binding:"required,min=2,max=100"`
	Description string  `gorm:"size:300;not null" json:"description" binding:"required,min=30,max=300"`
	Price       float64 `gorm:"not null" json:"price" binding:"required,gt=0"`
	Equipment   string  `gorm:"not null" json:"equipment" binding:"required" type:"text"`
}
