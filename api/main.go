package main

import (
	"github.com/macevil/terraform-provider-example/api/server"
	"log"
)

func main() {
	dnsRecordService := server.NewService()
	err := dnsRecordService.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}