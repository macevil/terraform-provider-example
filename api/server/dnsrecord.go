package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type DNSRecord struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

func (s *DNSRecordService) PostDNSRecord(w http.ResponseWriter, r *http.Request) {
	var dnsRecord DNSRecord
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&dnsRecord)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	s.Lock()
	defer s.Unlock()

	if s.dnsRecordExists(dnsRecord.Uuid) {
		http.Error(w, fmt.Sprintf("dnsRecord %s already exists", dnsRecord.Name), http.StatusBadRequest)
		return
	}

	s.dnsRecords[dnsRecord.Uuid] = dnsRecord
	log.Printf("added dnsRecord: %s", dnsRecord.Name)
	err = json.NewEncoder(w).Encode(dnsRecord)
	if err != nil {
		log.Printf("error sending response - %s", err)
	}
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

	if !s.dnsRecordExists(uuid) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(s.dnsRecords[uuid])
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *DNSRecordService) dnsRecordExists(uuid string) bool {
	if _, ok := s.dnsRecords[uuid]; ok {
		return true
	}
	return false
}
