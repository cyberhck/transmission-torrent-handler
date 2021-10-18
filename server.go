package main

import (
	"fmt"
	"github.com/cyberhck/torrent-handler/config"
	"github.com/cyberhck/torrent-handler/internal/handlers"
	"github.com/cyberhck/torrent-handler/internal/transmission"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	cfg := config.MustLoadConfig()
	client := transmission.New(cfg.TransmissionConfig.Endpoint)
	router.Handle("/", handlers.Index(cfg)).Methods("GET")
	router.Handle("/open", handlers.Open(cfg)).Methods("GET")
	router.Handle("/start-torrent", handlers.Start(client)).Methods("POST")
	fmt.Printf("starting on %d", cfg.SelfPort)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", cfg.SelfPort), router))
}
