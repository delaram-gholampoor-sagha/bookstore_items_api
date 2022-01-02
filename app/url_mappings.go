package app

import (
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_items_api/controllers"
)

func MapUrls() {
	router.HandleFunc("/ping", controllers.PongController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
