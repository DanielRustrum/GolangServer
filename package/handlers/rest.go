package presets

import (
	"net/http"

	serverHttp "github.com/DanielRustrum/Https-Go-Server/package/servers/http"
)

//RESTServer is ...
func RESTServer() serverHttp.HTTPHandler {
	id := serverHttp.GetHandlerID()
	return serverHttp.HTTPHandler{
		ID: id,
		Handler: func(response http.ResponseWriter, request *http.Request) {

		},
	}
}
