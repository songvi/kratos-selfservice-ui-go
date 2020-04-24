package render

import (
	"html/template"
	"net/http"

	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
)

type HtmlRender struct {
	template string
}

func NewHtmlRender(c configuration.ConfigProvider) *HtmlRender {
	return &HtmlRender{template: c.TemplateDir()}
}

func (h *HtmlRender) RenderTemplate(w http.ResponseWriter, r *http.Request, data interface{}, page string) {
	t, _ := template.ParseFiles(h.template + "/" + page + ".html")
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
