package error

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
)

type (
	ErrorHandlerProvider interface {
		ErrorHandler() *Handler
	}

	Handler struct {
		c configuration.ConfigProvider
	}
)

func (h *Handler) RegisterRouter(r *httprouter.Router) {
	r.GET(h.c.ErrorUrl(), func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// check request
		fmt.Println("Error request")
		requestId := r.URL.Query()["error"]

		if len(requestId) == 0 {
			// redirect to http://127.0.0.1:4455/.ory/kratos/public/self-service/browser/flows/login
			http.Redirect(w, r, h.c.KratosAdminUrl() +"/errors?error=", http.StatusFound)
			return
		}	

		//fmt.Printf("Request ID: %s", requestId[0])
		// Got Login Request JSON Payload

		errorUrl := h.c.KratosAdminUrl() + "/errors?error=" + requestId[0]

		//fmt.Println("LOGGGGGING :" + loginUrl)
		resp, err := http.Get(errorUrl)
		if err != nil {
			fmt.Printf("Error %v", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error %v", err)
		}

		fmt.Println(string(body))
	})
}

func NewErrorHandler(cfg configuration.ConfigProvider) *Handler {
	return &Handler{c: cfg}
}
