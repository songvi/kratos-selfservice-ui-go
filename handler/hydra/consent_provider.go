package hydra

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
	//"github.com/songvi/kratos-selfservice-ui-go/render"
)

type (
	HydraConsentHandlerProvider interface {
		HydraConsentHandler() *ConsentHandler
	}
	ConsentHandler struct {
		c configuration.ConfigProvider
	}
)

func (k *ConsentHandler) RegisterRouter(r *httprouter.Router) {
	r.GET("/oauth2/consent", k.hydraConsentHandler())
}

func NewHydraConsentHandler(cfg configuration.ConfigProvider) *ConsentHandler {
	return &ConsentHandler{c: cfg}
}

func (k *ConsentHandler) hydraConsentHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("Hydra login handler request")
	}
}
