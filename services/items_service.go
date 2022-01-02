package services

import (
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_items_api/domain/items"
	"github.com/Delaram-Gholampoor-Sagha/bookstore_utils-go/rest_errors"
)

var ItemsService itemsServiceInterface = &itemsService{}

// what is the interface that you want to provide for this service
type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(items.Item) (*items.Item, rest_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *itemsService) Get(items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}
