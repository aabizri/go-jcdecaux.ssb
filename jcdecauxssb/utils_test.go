package jcdecauxssb

import (
	"flag"
)

var testKey = flag.String("key","","The api key (MANDATORY)")

func getTestClient() (*Client,error) {
	flag.Parse()
	client, err := New(*testKey, nil)
	return client,err
}
