package main

import (
	"log"

	"github.com/macevil/terraform-provider-example/api/server"
)

func main() {
	dnsRecord1 := server.DNSRecord{"test", "1234"}
	dnsRecords := map[string]server.DNSRecord{"1234": dnsRecord1}
	dnsRecordService := server.NewService(dnsRecords)
	err := dnsRecordService.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
