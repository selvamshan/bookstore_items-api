package queries

import ()

type EsQuery struct {
	Equals []FieldValues `json:"equals"`
}

type FieldValues struct {
	Field string 		`json:"field"`
	Value interface {}  `json:"value"`
}