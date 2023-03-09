package models

import (
	"gorm.io/gorm"
	"github.com/go-ozzo/ozzo-validation"
)

type Room struct {
	gorm.Model
	Name  string
	Posts []Post `gorm:"foreignKey:RoomID"`
}

type CreateRoomInput struct {
	Name string `json:"name" binding:"required"`
}

func FindAllRooms(db *gorm.DB) ([]Room, error) {
	var rooms []Room
	result := db.Find(&rooms)

	return rooms, result.Error
}

func (room *Room) Create(db *gorm.DB) (Room, error) {
	newRoom := Room{
		Name: room.Name,
	}
	result := db.Create(&newRoom)

	return newRoom, result.Error
}

func (room *Room) Validate() error {
	return validation.ValidateStruct(room,
		validation.Field(&room.Name,
			validation.Required.Error("Name is required"),
			validation.Length(1, 255).Error("Name must be between 1 and 255 characters"),
		),
	)
}