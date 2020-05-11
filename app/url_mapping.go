package app

import (	
	"net/http"
	//"github.com/gorilla/mux"
	"github.com/selvamshan/bookstore_items-api/controllers"	
	"github.com/selvamshan/bookstore_items-api/services"
)


func mapUrls() {
	itemHandler := controllers.NewItemHandler(services.NewItemService())
	pingHandler := controllers.NewPingHandler()

	router.HandleFunc("/ping", pingHandler.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", itemHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", itemHandler.Get).Methods(http.MethodGet)
}