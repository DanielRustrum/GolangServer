package http

import "crypto/rand"

// HTTPHandler is ...
type HTTPHandler struct {
	ID      int
	Handler func(http.ResponseWriter, *http.Request)
}

var idMap map[int]bool

// GetHandlerID is ...
func GetHandlerID() int {
	for {
		number := rand.Int()
		if !idMap[number] {
			idMap[number] = true
			return number
		}
	}
}

