package hydra

import (
	"net/http"
	"fmt"

	"github.com/julienschmidt/httprouter"

	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
	//"github.com/songvi/kratos-selfservice-ui-go/render"
)

type (
	HydraLoginHandlerProvider interface {
		HydraLoginHandler() *LoginHandler
	}

	LoginHandler struct {
		c configuration.ConfigProvider
	}
)

func (k *LoginHandler) RegisterRouter(r *httprouter.Router) {
	r.GET("/oauth2/login", k.hydraLoginHandler())
}

func NewHydraLoginHandler(cfg configuration.ConfigProvider) *LoginHandler {
	return &LoginHandler{c: cfg}
}

func (k *LoginHandler) hydraLoginHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("Hydra login handler request")
	}
}
