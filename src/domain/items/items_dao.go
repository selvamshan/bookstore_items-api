package items

import (
	"fmt"
	"errors"
	"strings"
	"encoding/json"
	//"github.com/olivere/elastic"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
	"github.com/selvamshan/bookstore_items-api/src/domain/queries"
	"github.com/selvamshan/bookstore_items-api/src/domain/documents"
	"github.com/selvamshan/bookstore_items-api/src/clients/elasticsearch"
)

const (
	indexItems = "items"
	typeItem = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItem,  i.Id)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found wiht id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying ot get id %s", i.Id), errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to parse database response"),  errors.New("database error"))
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to parse database response"),  errors.New("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	
	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err = json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("databaase error"))
		}
		//fmt.Println(hit.Id)
		//item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found for given creteria")
	}	

	return items, nil
}

func (i *Item) Update(index string,  doc documents.EsDoc)(*UpdateResponse, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Update(indexItems, typeItem, i.Id, doc.Build())
	if err != nil {
		//fmt.Println(err.Error())
		if strings.Contains(err.Error(), "404") {
			return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return nil, rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to update id %s", i.Id), errors.New("database error"))
	}
	isUpdated := result.Result == "updated"
	updateRes := &UpdateResponse{
		DocId : result.Id, 
		IsUpdated: isUpdated,
	}

	return updateRes, nil

}


func (i *Item) Delete(index string)(*DeleteResponse,  rest_errors.RestErr) {
	result, err := elasticsearch.Client.Delete(indexItems, typeItem, i.Id)
	if err != nil {
		//fmt.Println(err.Error())
		if strings.Contains(err.Error(), "404") {
			return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return nil, rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to delete id %s", i.Id), errors.New("database error"))
	}

	isDeleted := (result.Result == "deleted")
	
	deleteRes := &DeleteResponse{
		DocId : result.Id, 
		IsDeleted: isDeleted,
	}

	return deleteRes, nil

}