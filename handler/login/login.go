package login

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"

	//"github.com/songvi/kratos-selfservice-ui-go/render"
	"github.com/songvi/kratos-selfservice-ui-go/render/flows"
)

type (
	LoginHandlerProvider interface {
		LoginHandler() *Handler
	}

	Handler struct {
		c configuration.ConfigProvider
		l logrus.FieldLogger
	}
)

func (k *Handler) RegisterRouter(r *httprouter.Router) {
	fmt.Printf("Register handler %s\n", k.c.LoginUrl())
	r.GET(k.c.LoginUrl(), k.loginHandler)
}

func NewLoginHandler(cfg configuration.ConfigProvider, log logrus.FieldLogger) *Handler {
	return &Handler{c: cfg, l: log}
}

func (k *Handler) loginHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// TODO
	// GET /auth/login?request=abcde
	//fmt.Printf("call login handler %s\n", r.URL.String())
	k.l.Infof("Resquest is called %s", k.c.LoginUrl())

	// check request
	requestId := r.URL.Query()["request"]

	if len(requestId) == 0 {
		// redirect to http://127.0.0.1:4455/.ory/kratos/public/self-service/browser/flows/login
		baseUrl, _ := url.Parse("http://127.0.0.1:8084/oauth/")
		params := url.Values{}
		params.Add("login_challenge", "1234")
		baseUrl.RawQuery = params.Encode()
		//k.l.Infof(k.c.KratosPublicFlowsUrl() + "/login" + "?return_to=" + baseUrl.String())
		k.l.Info("Redirect to " + k.c.KratosPublicFlowsUrl() + "/login" + "?return_to=" + baseUrl.String())
		http.Redirect(w, r, k.c.KratosPublicFlowsUrl()+"/login"+"?return_to="+baseUrl.String(), http.StatusFound)
		return
	}

	//fmt.Printf("Request ID: %s", requestId[0])
	// Got Login Request JSON Payload

	loginUrl := k.c.KratosLoginFlowUrl() + "?request=" + requestId[0]

	//fmt.Println("LOGGGGGING :" + loginUrl)
	resp, err := http.Get(loginUrl)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	var loginStruct flows.PageLogin
	loginStruct.Title = "Login"

	//fmt.Println(string(body))

	err = json.Unmarshal(body, &loginStruct)
	if err != nil {
		// TODO
		fmt.Errorf("Error during unmarshal json payload %v", err)
	}

	//fmt.Printf("htmlForm %v", loginStruct)

	// Render
	//render := render.NewHtmlRender(k.c)
	//render.RenderTemplate(w, r, loginStruct, "login")

	t := template.New("Login Page")
	t = template.Must(t.ParseFiles(k.c.TemplateDir()+"/login.html", k.c.TemplateDir()+"/layout.html"))
	err = t.ExecuteTemplate(w, "layout", loginStruct)

	if err != nil {
		fmt.Errorf("Template execution: %s", err)
	}
}
