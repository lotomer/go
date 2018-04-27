package router

import (
	"github.com/julienschmidt/httprouter"
)

// DefaultRouter is static
var DefaultRouter = &router{*httprouter.New()}

type router struct {
	httprouter.Router
}
