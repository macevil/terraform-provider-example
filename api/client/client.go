package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/macevil/terraform-provider-example/api/server"
	"io"
	"net/http"
)

type Client struct {
	hostname   string
	port       int
	httpClient *http.Client
}

func NewClient(hostname string, port int) *Client {
	return &Client{
		hostname:   hostname,
		port:       port,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetDNSRecord(name string) (*server.DNSRecord, error) {
	body, err := c.httpRequest(fmt.Sprintf("dnsrecords/%v", name), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	dnsRecord := &server.DNSRecord{"",""}
	err = json.NewDecoder(body).Decode(dnsRecord)
	if err != nil {
		return nil, err
	}
	return dnsRecord, nil
}