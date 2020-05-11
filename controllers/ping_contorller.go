package controllers


import (
	
	"net/http"
)

const (
	pong = "pong"
)

type PingHandler interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingHandler struct {

}

func NewPingHandler() PingHandler {
	return &pingHandler{}
}



func (handler *pingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}