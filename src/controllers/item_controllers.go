package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/selvamshan/bookstore_items-api/src/domain/items"
	"github.com/selvamshan/bookstore_items-api/src/domain/queries"
	"github.com/selvamshan/bookstore_items-api/src/domain/documents"
	"github.com/selvamshan/bookstore_items-api/src/services"
	"github.com/selvamshan/bookstore_items-api/src/utils/http_utils"
	"github.com/selvamshan/bookstore_oauth-go/oauth"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
)

type itemHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)	
}

type itemHandler struct {
	service services.ItemService
}

func NewItemHandler(service services.ItemService) itemHandlerInterface {
	return &itemHandler{
		service: service,
	}
}

func buildUpdateDoc(itemId string) documents.EsDoc{
	doc := documents.EsDoc{
		Docs: []documents.FieldValues{
			documents.FieldValues{
				Field: "id",
				Value: itemId,
			},
		},
	}
	//fmt.Println(doc)
	return doc
}

func (handler *itemHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.Authenticate(r); err != nil {
		http_utils.RespondJson(w, err.Status(), err)
		//http_utils.RespondError(w, *err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("unauthorized")
		fmt.Println(respErr.Status(), respErr.Message())
		http_utils.RespondJson(w, respErr.Status(), respErr.Message())
		// http_utils.RespondError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondJson(w, respErr.Status(), respErr)
		// http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondJson(w, respErr.Status(), respErr)
		// http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := handler.service.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondJson(w, createErr.Status(), createErr)
		//http_utils.RespondError(w, createErr)
		return
	}

	//fmt.Println(result.Id)	
	doc := buildUpdateDoc(result.Id)
	_, updateErr := handler.service.Update(result.Id, doc)
	if updateErr != nil {
		fmt.Printf("Created item doc id not updated %s", result.Id)
	}
	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (handler *itemHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := handler.service.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, item)
}

func (handler *itemHandler) Search(w http.ResponseWriter, r *http.Request) {
	
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()
	
	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	
	items, searchErr := handler.service.Search(query)
	if searchErr != nil {
		http_utils.RespondError(w, searchErr)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, items)
}


func (handler *itemHandler) Update(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("in item controller update line 123")
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var doc documents.EsDoc
	if err := json.Unmarshal(bytes, &doc); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}

	resp, updateErr := handler.service.Update(itemId, doc)
	if updateErr != nil {
		http_utils.RespondError(w, updateErr)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, resp)
}


func (handler *itemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	//mt.Println("in item controller delete line 154")
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	resp, deleteErr := handler.service.Delete(itemId)
	if deleteErr != nil {
		http_utils.RespondError(w, deleteErr)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, resp)
}