package logout

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
)

type (
	LogoutHandlerProvider interface {
		LogoutHandler() *Handler
	}

	Handler struct {
		c configuration.ConfigProvider
	}
)

func NewLogoutHandler(cfg configuration.ConfigProvider) *Handler {
	return &Handler{c: cfg}
}

func (h *Handler) RegisterRouter(r *httprouter.Router) {
	r.GET(h.c.LogoutUrl(), func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("Logout request")
		http.Redirect(w, r, h.c.KratosLogoutFlowsUrl(), http.StatusFound)
	})
}
