package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type DNSRecordService struct {
	dnsRecords map[string]DNSRecord
	sync.RWMutex
}

func NewService(dnsRecords map[string]DNSRecord) *DNSRecordService {
	return &DNSRecordService{
		dnsRecords: dnsRecords,
	}
}

func (s *DNSRecordService) ListenAndServe() error {
	r := mux.NewRouter()
	r.HandleFunc("/dnsrecords/{uuid}", s.GetDNSRecord).Methods("GET")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}
	return nil
}
