package app

import (	
	"net/http"
	//"github.com/gorilla/mux"
	"github.com/selvamshan/bookstore_items-api/src/controllers"	
	"github.com/selvamshan/bookstore_items-api/src/services"
)


func mapUrls() {
	itemHandler := controllers.NewItemHandler(services.NewItemService())
	pingHandler := controllers.NewPingHandler()

	router.HandleFunc("/ping", pingHandler.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", itemHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", itemHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/itmes/search", itemHandler.Search).Methods(http.MethodPost)
	router.HandleFunc("/itmes/update/{id}", itemHandler.Update).Methods(http.MethodPost)
	router.HandleFunc("/itmes/delete/{id}", itemHandler.Delete).Methods(http.MethodDelete)
}