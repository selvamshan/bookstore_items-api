package services

import (
	//"net/http"
	"github.com/selvamshan/bookstore_items-api/src/domain/items"
	"github.com/selvamshan/bookstore_items-api/src/domain/queries"
	"github.com/selvamshan/bookstore_items-api/src/domain/documents"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
)

// var (
// 	ItemService itemsServiceInterface = &itemService{}
// )

type ItemService interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)	
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)	
	Update(string, documents.EsDoc) (*items.UpdateResponse, rest_errors.RestErr)
	Delete(string) (*items.DeleteResponse, rest_errors.RestErr)
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

func (s *itemService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	
	return dao.Search(query)
	
}

func (s *itemService) Update(id string, doc documents.EsDoc) (*items.UpdateResponse, rest_errors.RestErr){
	dao := items.Item{Id: id}

	return dao.Update(id, doc)

}


func (s *itemService) Delete(id string) (*items.DeleteResponse, rest_errors.RestErr){
	dao := items.Item{Id: id}

	return dao.Delete(id)

}

