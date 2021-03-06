package items

import ()


type Item struct {
	Id 				  string 		 `json:"id"`
	Seller 			  int64 		 `json:"seller"`
	Title 			  string 		 `json:"title"`
	Description 	  Desctription 	 `json:"description"`
	Picutres    	  []Picture    	 `json:"pictures"`
	Video       	  string       	 `json:"video"`
	Price       	  float32      	 `json:"price"`
	AvailableQuantity int    		 `json:"available_quantity"`
	SoldQuantity      int    		 `json:"sold_quantity"`
	Status 			  string 		 `json:"status"`
} 		  

type Desctription struct {
	PlainText  string `json:"plain_text"`
	Html       string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}

type UpdateResponse struct {
	DocId string  `json:"id"`
	IsUpdated bool `json:isupdated"`
}


type DeleteResponse struct {
	DocId string  `json:"id"`
	IsDeleted bool `json:isdeleted"`
}

