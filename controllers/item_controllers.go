package controllers

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/selvamshan/bookstore_oauth-go/oauth"
	"github.com/selvamshan/bookstore_items-api/domain/items"
	"github.com/selvamshan/bookstore_items-api/services"
	"github.com/selvamshan/bookstore_items-api/utils/http_utils"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
)

type itemHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemHandler struct {
   service services.ItemService
}

func NewItemHandler(service services.ItemService) itemHandlerInterface {
	return &itemHandler{
		service: service,
	}
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
		fmt.Println(respErr.Status(), respErr.Message(), )
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

	//fmt.Println(result)
	http_utils.RespondJson(w, http.StatusCreated, result)

}


func (handler *itemHandler) Get(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := handler.service.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
	}

	http_utils.RespondJson(w, http.StatusOK, item)
}