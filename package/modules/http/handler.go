package http

// Reader is ...
type Reader struct {
}

// Writer is ...
type Writer struct {
}

// HandlerFunc is ...
type HandlerFunc func(Reader) Writer

// Handler is ...
type Handler struct {
	ID      int
	Handler HandlerFunc
}
