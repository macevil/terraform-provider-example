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
	dnsRecords map[string]DNSRecord
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

func handler(w http.ResponseWriter, r *http.Request) {
	g := &Greeting{Language: "go", Content: templateGreeting(r.URL.Query().Get("name"))}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func (g *Greeting) ListenAndServe() error {
	r := mux.NewRouter()
	r.HandleFunc("/dnsrecords", handler).Methods("POST")
	r.HandleFunc("/dnsrecords", handler).Methods("GET")
	r.HandleFunc("/dnsrecords/{uuid}", g.GetDNSRecord).Methods("GET")
	r.HandleFunc("/dnsrecords/{uuid}", handler).Methods("PUT")
	r.HandleFunc("/dnsrecords/{uuid}", handler).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}