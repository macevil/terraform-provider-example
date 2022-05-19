package main

import (
	"log"

	"github.com/macevil/terraform-provider-example/api/server"
)

func main() {
	dnsRecords := map[string]server.DNSRecord{}
	dnsRecordService := server.NewService()
	err := dnsRecordService.ListenAndServe(dnsRecords)
	if err != nil {
		log.Fatal(err)
	}
}
