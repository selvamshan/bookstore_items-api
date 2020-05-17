package documents

func(d EsDoc) Build() map[string]interface{} {
	
	docs := make(map[string]interface{})
	for _, dc := range d.Docs {
		docs[dc.Field] = dc.Value
	}
	//fmt.Println(docs)
	return docs
}


// type EsDoc struct {
// 	Docs []FieldValues `json:"docs"`
// }

// type FieldValues struct {
// 	Field string      `json:"field"`
// 	Value interface{} `json:"value"`
// }

// func (d EsDoc) Build() map[string]interface{} {

// 	docs := make(map[string]interface{})
// 	for _, dc := range d.Docs {
// 		docs[dc.Field] = dc.Value
// 	}
// 	//fmt.Println(docs)
// 	return docs
// }

// func main() {
// 	esJson := `{"docs":[{"field":"id", "value":"4"},{"field":"status", "value":"active"}]}`
// 	var d EsDoc
// 	json.Unmarshal([]byte(esJson), &d)

// 	//docs := make(map[string]interface{})
// 	//for _, d := range q.Docs {
// 	//	docs[d.Field] = d.Value
// 	//}
// 	fmt.Println(d.Build())
// }