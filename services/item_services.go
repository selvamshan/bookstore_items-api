package services

import (
	//"net/http"
	"github.com/selvamshan/bookstore_items-api/domain/items"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
)

// var (
// 	ItemService itemsServiceInterface = &itemService{}
// )

type ItemService interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)	
}

type itemService struct{}

func NewItemService() ItemService {
	return &itemService{}
}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}


func (s *itemService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}