package items

import (
	"fmt"
	"errors"
	"strings"
	"encoding/json"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
	"github.com/selvamshan/bookstore_items-api/clients/elasticsearch"
)

const (
	indexItems = "items"
	typeItem = "item"
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
