package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nastya231213/go-hotel-management/pkg/models"
	"github.com/Nastya231213/go-hotel-management/pkg/utils"
	"github.com/gorilla/mux"
)

var NewHotel models.Hotel

func GetHotel(w http.ResponseWriter, r *http.Request) {
	newHotels := models.GetAllHotels()
	res, _ := json.Marshal(newHotels)
	w.Header().Set("Content-Type", "pkgplication/json")
	w.WriteHeader((http.StatusOK))
	w.Write(res)
}
func GetHotelById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hotelId := vars["hotelId"]
	ID, err := strconv.ParseInt(hotelId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	hotelDetails, _ := models.GetHotelById(ID)
	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("Content-Type", "pkgplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	CreateHotel := &models.Hotel{}
	utils.ParseBody(r, CreateHotel)
	hotel := CreateHotel.CreateHotel()
	res, _ := json.Marshal(hotel)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hotelId := vars["hotelId"]
	ID, err := strconv.ParseInt(hotelId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	hotel := models.DeleteHotel(ID)
	res, _ := json.Marshal(hotel)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	var updateHotel = &models.Hotel{}
	utils.ParseBody(r, updateHotel)
	vars := mux.Vars(r)
	hotelId := vars["hotelID"]
	ID, err := strconv.ParseInt(hotelId, 0, 0)

	if err != nil {
		fmt.Println("error")
	}
	hotelDetails, db := models.GetHotelById(ID)
	if updateHotel.Name != "" {
		hotelDetails.Name = updateHotel.Name
	}
	if updateHotel.Location != "" {
		hotelDetails.Location = updateHotel.Location
	}
	if updateHotel.Ratinng != 0 {
		hotelDetails.Ratinng = updateHotel.Ratinng
	}
	db.Save(&hotelDetails)
	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
