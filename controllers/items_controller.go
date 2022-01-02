package controllers

import (
	"fmt"
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_items_api/domain/items"
	"github.com/Delaram-Gholampoor-Sagha/bookstore_items_api/services"
	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth-go/oauth"
)

var ItemsController itemsControllerInterface = &itemsController{}

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		return
	}
	fmt.Println(result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
