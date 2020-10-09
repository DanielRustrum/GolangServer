package handler

// TODO: CORS Support
// TODO: File Handler > Inprogress
// TODO: REST Handler
// TODO: Debug Handler

import (
	"math/rand"
	"net/http"
)

//* Handler Function

var idMap map[int]bool

func genHandlerID() int {
	for {
		number := rand.Int()
		if !idMap[number] {
			idMap[number] = true
			return number
		}
	}
}

//* Public

// Setup is ...
func Setup() {
	idMap = make(map[int]bool)
}

// Handler is ..
type Handler struct {
	ID      int
	Handler func(http.ResponseWriter, *http.Request)
}
