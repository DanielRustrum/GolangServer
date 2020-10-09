package handler

import (
	"net/http"
)

//RESTServer is ...
func RESTServer() Handler {
	id := genHandlerID()
	return Handler{
		ID: id,
		Handler: func(response http.ResponseWriter, request *http.Request) {

		},
	}
}
