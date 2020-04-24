package whoami

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
	//"github.com/songvi/kratos-selfservice-ui-go/render"
)

type (
	WhoamiHandlerProvider interface {
		WhoamiHandler() *Handler
	}

	Handler struct {
		c configuration.ConfigProvider
	}
)

func (k *Handler) RegisterRouter(r *httprouter.Router) {
	r.GET("/", k.whoamiHandler())
}

func NewWhoamiHandler(cfg configuration.ConfigProvider) *Handler {
	return &Handler{c: cfg}
}

func (k *Handler) whoamiHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// TODO
		fmt.Println("Whoami requested")
		headers := map[string][]string{
			"Accept": []string{"application/json"},
		}
		var body []byte
		req, err := http.NewRequest("GET", "http://127.0.0.1:4455/.ory/kratos/public/sessions/whoami", bytes.NewBuffer(body))
		req.Header = headers
		cookie, _ := r.Cookie("ory_kratos_session")
		req.AddCookie(cookie)

		client := &http.Client{}
		resp, err := client.Do(req)

		//resp, err := http.Get(profileUrl)
		if err != nil {
			fmt.Printf("Error %v", err)
			return
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error %v", err)
			return
		}
		fmt.Printf("WHO AM I: %v", string(body))
		w.Write(body)
	}
}
