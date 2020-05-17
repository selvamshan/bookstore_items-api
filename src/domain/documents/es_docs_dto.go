package documents


type EsDoc struct {
	Docs []FieldValues `json:"docs"`
}

type FieldValues struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}