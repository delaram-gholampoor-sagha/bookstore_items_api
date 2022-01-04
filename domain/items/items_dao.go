package items

import (
	"errors"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_items_api/clients/elastic_search"
	"github.com/Delaram-Gholampoor-Sagha/bookstore_utils-go/rest_errors"
)



const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result , err := elastic_search.Client.Index(indexItems , i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save items" , errors.New("database error"))
	}
	i.Id = result.Id
	return nil 
}