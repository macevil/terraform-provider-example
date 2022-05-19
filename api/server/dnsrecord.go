package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type DNSRecord struct {
	name string `json:"name"`
	uuid string `json:"uuid"`
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

	dnsRecord := &DNSRecord{name: "test", uuid: "test"}
	err := json.NewEncoder(w).Encode(dnsRecord)
	//err := json.NewEncoder(w).Encode(s.dnsRecords[uuid])
	if err != nil {
		log.Println(err)
		return
	}
}
