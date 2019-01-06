package main

import (
	"local/nearby/places"
	"local/nearby/server"
	"log"
	"net/http"
	"os"
)

var (
	defaultURL = "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
	port       = ":8080"
)

func main() {
	l := log.New(os.Stdout, "", 0)
	c := places.NewClient(defaultURL, &places.HTTPClient{HTTPClient: http.DefaultClient})
	s := server.NewServer(l, &c)
	s.Logger.Printf("listening on port %s", port)
	http.HandleFunc("/nearby", s.Nearby)
	http.ListenAndServe(":8080", nil)
}
