package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type DNSRecord struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

func (s *DNSRecordService) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	if uuid == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	s.RLock()
	defer s.RUnlock()
	if !s.itemExists(uuid) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(s.dnsRecords[uuid])
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *DNSRecordService) itemExists(uuid string) bool {
	if _, ok := s.dnsRecords[uuid]; ok {
		return true
	}
	return false
}
