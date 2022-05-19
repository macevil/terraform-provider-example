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

func (s *DNSRecordService) GetDNSRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	if uuid == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	s.RLock()
	defer s.RUnlock()

	dnsRecord := &DNSRecord{"test2", "5678"}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(dnsRecord)
	//err := json.NewEncoder(w).Encode(s.dnsRecords[uuid])
	if err != nil {
		log.Println(err)
		return
	}
}
