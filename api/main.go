package main

import (
	"github.com/macevil/terraform-provider-example/api/server"
	"log"
)

func main() {
	dnsRecords := map[string]server.DNSRecord{}
	dnsRecordService := server.NewService()
	err := dnsRecordService.ListenAndServe(dnsRecords)
	if err != nil {
		log.Fatal(err)
	}
}