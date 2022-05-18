package server

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Greeting struct {
	Language string `json:"language"`
	Content  string `json:"content"`
	//dnsRecords map[string]DNSRecord
}

func NewService() *Greeting{
	return &Greeting{}
}

func templateGreeting(s string) string {
	if s == "" {
		s = "World"
	}
	return fmt.Sprintf("Hello, %s!", s)
}

func handler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		log.Printf("%s %s", method, path)
		handlerFunc(w, r)
		return
	}
}

func (g *Greeting) ListenAndServe() error {
	r := mux.NewRouter()
	r.HandleFunc("/dnsrecords", handler).Methods("POST")
	r.HandleFunc("/dnsrecords", handler(g.GetDNSRecord)).Methods("GET")
	r.HandleFunc("/dnsrecords/{uuid}", g.GetDNSRecord).Methods("GET")
	r.HandleFunc("/dnsrecords/{uuid}", handler).Methods("PUT")
	r.HandleFunc("/dnsrecords/{uuid}", handler).Methods("DELETE")
	//http.Handle("/", r)
	log.Fatal("Starting server on %s", s.connectionString)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}
	return nil
}