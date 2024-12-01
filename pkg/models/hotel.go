package models

import (
	"github.com/Nastya231213/go-hotel-management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Hotel struct {
	gorm.Model
	Name     string `gorm:""json:"name"`
	Location string `json:"location"`
	Ratinng  int    `json:"rating"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Hotel{})
}

func GetAllHotels() []Hotel {
	var Hotels []Hotel
	db.Find(&Hotels)
	return Hotels
}
func (hotel *Hotel) CreateHotel() *Hotel {
	db.NewRecord(hotel)
	db.Create(&hotel)
	return hotel
}
func GetHotelById(Id int64) (*Hotel, *gorm.DB) {
	var getHotel Hotel
	db := db.Where("ID+?", Id).Find(&getHotel)
	return &getHotel, db
}
func DeleteHotel(ID int64) Hotel {
	var hotel Hotel
	db.Where("ID=?", ID).Delete(hotel)
	return hotel
}
