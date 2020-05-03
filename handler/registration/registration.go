package registration

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
	//"github.com/songvi/kratos-selfservice-ui-go/render"
	"github.com/songvi/kratos-selfservice-ui-go/render/flows"
)

type (
	RegistrationHandlerProvider interface {
		RegistrationHandler() *Handler
	}

	Handler struct {
		c configuration.ConfigProvider
	}
)

func (k *Handler) RegisterRouter(r *httprouter.Router) {
	r.GET(k.c.RegistrerUrl(), k.registrationHandler())
}

func NewRegistrationHandler(cfg configuration.ConfigProvider) *Handler {
	return &Handler{c: cfg}
}

func (k *Handler) registrationHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// TODO
		// GET /auth/login?request=abcde
		// check request
		requestId := r.URL.Query()["request"]

		if len(requestId) == 0 {
			// redirect to http://127.0.0.1:4455/.ory/kratos/public/self-service/browser/flows/login
			http.Redirect(w, r, k.c.KratosPublicUrl() + k.c.KratosBrowserInitPath() + "/registration", http.StatusFound)
			return
		}

		// Got Login Request JSON Payload
		loginUrl := k.c.KratosAdminUrl() + k.c.KratosBrowserRequestPath() + "/registration" + "?request=" + requestId[0]

		resp, err := http.Get(loginUrl)
		if err != nil {
			// TODO
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// TODO
		}

		var registration flows.PageRegistration
		err = json.Unmarshal(body, &registration)
		if err != nil {
			// TODO
			fmt.Errorf("Error during unmarshal json payload %v", err)
		}
		registration.Title = "Registration"
	
		// Render
		t := template.New("Registration Page")
		t = template.Must(t.ParseFiles(k.c.TemplateDir()+"/registration.html", k.c.TemplateDir()+"/layout.html"))
		err = t.ExecuteTemplate(w, "layout", registration)

		if err != nil {
			fmt.Errorf("Template execution: %s", err)
		}
	}
}
