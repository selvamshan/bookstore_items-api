package elasticsearch


import (
	"context"
	//"encoding/json"
	"fmt"
	//"reflect"
	"time"

	"github.com/olivere/elastic/v7"
	//"github.com/selvamshan/bookstore_items-api/logger"
	"github.com/selvamshan/bookstore_utils-go/logger"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{})(*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)	
	Search(string, elastic.Query) (*elastic.SearchResult, error)
	Update(string, string, string, map[string]interface{}) (*elastic.UpdateResponse, error)
	Delete(string, string, string) (*elastic.DeleteResponse, error) 
}

type esClient struct {
	client *elastic.Client
}

func Init() {	
	log := logger.GetLogger()
	client, err := elastic.NewClient(		
		elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9200"),
		// elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetRetrier(NewCustomRetrier()),
		// elastic.SetGzip(true),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
		// elastic.SetHeaders(http.Header{
		//   "X-Caller-Id": []string{"..."},
		// }),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
					Index(index).
					Type(docType).
					BodyJson(doc).
					Do(ctx)
	if err != nil {
		logger.Error(
			fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil

}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
				Index(index).
				//Type(docType).
				Id(id).
				Do(ctx)
	if err != nil {
		logger.Error(
			fmt.Sprintf("error when trying to get id %s", id), 
			err,
		)
		return nil, err
	}
	
	return result, nil
}


func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	//fmt.Println("in elastic client", query)
	
	result, err := c.client.Search(index).
					Query(query).
					RestTotalHitsAsInt(true).
					Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, err
	}

	return result, nil

}

func (c *esClient) Update(index string, docType string, id string, doc map[string]interface{}) (*elastic.UpdateResponse, error) {
	//fmt.Println(doc)
	ctx := context.Background()
	result, err := c.client.Update().
			Index(index).
			//Type(docType).
			Id(id).
			Doc(doc).			
			Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to update document id %s", id), err)
		return nil, err
	}

	return result, nil
}


func (c *esClient) Delete(index string, docType string, id string) (*elastic.DeleteResponse, error) {

	ctx := context.Background()
	result, err := c.client.Delete().Index(index).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to delete document id %s", id), err)
		return nil, err
	}

	return result, nil

}