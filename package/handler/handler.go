package handler

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

// Handler is ..
type Handler struct {
	ID      int
	Handler func(http.ResponseWriter, *http.Request)
}
