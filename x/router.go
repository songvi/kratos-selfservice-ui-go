package x

import (
	//"fmt"
	//"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	*httprouter.Router
}

func NewRouter() *Router {
	return &Router{
		Router: httprouter.New(),
	}
}


