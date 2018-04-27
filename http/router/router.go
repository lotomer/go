package router

import (
	"github.com/julienschmidt/httprouter"
)

// DefaultRouter is static
var DefaultRouter = new()

type router struct {
	httprouter.Router
}

func new() *router {
	return &router{*httprouter.New()}
}
