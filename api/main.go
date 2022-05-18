package main

import (
	"github.com/macevil/terraform-provider-example/api/server"
	//"log"
	"fmt"
)

func main() {
	dnsRecordService := server.NewService()
	//err := dnsRecordService.ListenAndServe()
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Printf("Hello, World\n")
}