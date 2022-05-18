package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

type DNSRecordService struct {
	dnsRecords map[string]DNSRecord
	sync.RWMutex
}

func NewService(dnsRecords map[string]DNSRecord) *DNSRecordService{
	return &DNSRecordService{
		dnsRecords: dnsRecords
	}
}

func (s *DNSRecordService) ListenAndServe() error {
	r := mux.NewRouter()
	//r.HandleFunc("/dnsrecords", ).Methods("POST")
	//r.HandleFunc("/dnsrecords", ).Methods("GET")
	r.HandleFunc("/dnsrecords/{uuid}", s.GetDNSRecord).Methods("GET")
	//r.HandleFunc("/dnsrecords/{uuid}", handler).Methods("PUT")
	//r.HandleFunc("/dnsrecords/{uuid}", handler).Methods("DELETE")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}
	return nil
}