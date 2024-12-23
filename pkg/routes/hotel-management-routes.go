package routes

import (
	"github.com/Nastya231213/go-hotel-management/pkg/controllers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var RegisterHotelManagementRoutes = func(router *mux.Router) {
	router.HandleFunc("/hotel/", controllers.CreateHotel).Methods("POST")
	router.HandleFunc("/hotel/", controllers.GetHotel).Methods("GET")
	router.HandleFunc("/hotel/{hotelId}", controllers.GetHotelById).Methods("GET")
	router.HandleFunc("/hotel/{hotelId}", controllers.UpdateHotel).Methods("PUT")
	router.HandleFunc("/hotel/{hotelId}", controllers.DeleteHotel).Methods("DELETE")

}
