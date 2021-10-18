package handlers

import (
	"github.com/cyberhck/torrent-handler/config"
	"net/http"
	"text/template"
)

type OpenTemplateData struct {
	TransmissionPublicURL string
}

func Open(cfg *config.Config) http.Handler {
	tmpl := template.Must(template.ParseFiles("./internal/templates/open.html"))
	data := &OpenTemplateData{
		TransmissionPublicURL: cfg.TransmissionConfig.PublicURL,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, data)
	})
}
