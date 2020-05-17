package app

import (	
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/selvamshan/bookstore_items-api/src/clients/elasticsearch"
)


var (
	router = mux.NewRouter()
)


func StartApplication() {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
        Handler:      router,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}