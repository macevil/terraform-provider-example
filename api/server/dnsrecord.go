package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type DNSRecord struct {
	Name string `json:"name"`
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

	err := json.NewEncoder(w).Encode(s.dnsRecords[uuid])
	if err != nil {
		log.Println(err)
		return
	}
}