package profile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
	"github.com/songvi/kratos-selfservice-ui-go/render/flows"
)

type (
	ProfileHandlerProvider interface {
		ProfileHandler() *Handler
	}

	Handler struct {
		c configuration.ConfigProvider
	}
)

func (k *Handler) RegisterRouter(r *httprouter.Router) {
	r.GET(k.c.ProfileUrl(), k.profileHandler())
}

func NewProfileHandler(cfg configuration.ConfigProvider) *Handler {
	return &Handler{c: cfg}
}

func (k *Handler) profileHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// check request
		requestId := r.URL.Query()["request"]

		if len(requestId) == 0 {
			// redirect to http://127.0.0.1:4455/.ory/kratos/public/self-service/browser/flows/login
			http.Redirect(w, r, k.c.KratosPublicFlowsUrl()+"/profile", http.StatusFound)
			return
		}

		//fmt.Printf("Request ID: %s", requestId[0])
		// Got Login Request JSON Payload

		profileUrl := k.c.KratosProfileFlowUrl() + "?request=" + requestId[0]

		headers := map[string][]string{
			"Accept": []string{"application/json"},
		}
		var body []byte
		req, err := http.NewRequest("GET", profileUrl, bytes.NewBuffer(body))
		req.Header = headers

		client := &http.Client{}
		resp, err := client.Do(req)

		//resp, err := http.Get(profileUrl)
		if err != nil {
			// TODO
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			// TODO
		}

		var profile flows.PageProfile
		profile.Title = "Profile"

		//fmt.Println(string(body))

		err = json.Unmarshal(body, &profile)
		if err != nil {
			// TODO
			fmt.Errorf("Error during unmarshal json payload %v", err)
		}

		// Render
		t := template.New("Profile Page")
		t = template.Must(t.ParseFiles(k.c.TemplateDir()+"/profile.html", k.c.TemplateDir()+"/layout.html"))
		err = t.ExecuteTemplate(w, "layout", profile)

		if err != nil {
			fmt.Errorf("Template execution: %s", err)
		}
	}
}
