package handlers

import (
	"github.com/cyberhck/torrent-handler/config"
	"net/http"
	"strconv"
	"text/template"
)

type IndexTemplateData struct {
	SelfHost string
	SelfPort string
}

func Index(cfg *config.Config) http.Handler {
	tmpl := template.Must(template.ParseFiles("./internal/templates/index.html"))
	templateData := &IndexTemplateData{
		SelfHost: cfg.SelfURL,
		SelfPort: strconv.Itoa(cfg.SelfPort),
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, templateData)
	})
}
