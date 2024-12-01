package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nastya231213/go-hotel-management/pkg/models"
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
