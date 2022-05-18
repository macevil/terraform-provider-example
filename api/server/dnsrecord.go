package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type DNSRecord struct {
	content string `json:"content"`
	name string `json:"name"`
	uuid string `json:"uuid"`
}

func (g *Greeting) GetDNSRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	if uuid == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
}