package server

import (
	"github.com/gorilla/mux"
	"github.com/koind/calendar/app/delivery/http/service"
	"net/http"
)

// HttpServer
type HttpServer struct {
	domain string
	router http.Handler
	s      *service.EventService
}

// Start fires up the http server
func (s *HttpServer) Start() error {
	return http.ListenAndServe(s.domain, s.router)
}

// NewHTTPServer returns http server that wraps event business logic
func NewHTTPServer(handleService *service.EventService, domain string) *HttpServer {

	r := mux.NewRouter()
	hs := HttpServer{router: r, domain: domain, s: handleService}

	r.HandleFunc("/event", handleService.CreateHandle).Methods("POST")
	r.HandleFunc("/event/{uuid}", handleService.UpdateHandle).Methods("PUT")
	r.HandleFunc("/event/{uuid}", handleService.DeleteHandle).Methods("DELETE")

	http.Handle("/", r)

	return &hs
}
