package dashboard

import (
	"fmt"
	"net/http"

	"html/template"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
	//"github.com/songvi/kratos-selfservice-ui-go/render"
)

type (
	DashBoardHandlerProvider interface {
		DashBoardHandler() *Handler
	}

	Handler struct {
		c configuration.ConfigProvider
		l logrus.FieldLogger
	}
)

func (k *Handler) RegisterRouter(r *httprouter.Router) {
	fmt.Printf("Register handler %s\n", k.c.DashboardUrl())
	r.GET(k.c.DashboardUrl(), k.dashBoardHandler)
}

func NewDashBoardHandler(cfg configuration.ConfigProvider, log logrus.FieldLogger) *Handler {
	return &Handler{c: cfg, l: log}
}

func (k *Handler) dashBoardHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	t := template.New("Dashboard Page")
	t = template.Must(t.ParseFiles(k.c.TemplateDir()+"/dashboard.html", k.c.TemplateDir()+"/layout.html"))
	err := t.ExecuteTemplate(w, "layout", nil)

	if err != nil {
		k.l.Error(err)
	}
}
