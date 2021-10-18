package handlers

import (
	"encoding/json"
	"github.com/cyberhck/torrent-handler/internal/transmission"
	"net/http"
)

type Request struct {
	URL string `json:"url"`
}

func Start(client *transmission.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := &Request{}
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request: " + err.Error()))
			return
		}
		err = client.AddTorrent(request.URL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal error: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
