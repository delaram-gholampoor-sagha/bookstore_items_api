package app

import (
	"net/http"
	"time"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_items_api/src/clients/elastic_search"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	elastic_search.Init()
	MapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8081",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
