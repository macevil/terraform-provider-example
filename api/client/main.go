package main

import (
	"log"

	"github.com/macevil/terraform-provider-example/api/client"
)

func main() {
	client := client.NewClient("localhost", 8080)
	log.Fatal(client.GetDNSRecord("test"))
}