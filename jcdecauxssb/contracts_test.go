package jcdecauxssb

import "testing"

func TestContracts(t *testing.T) {

	client, err := getTestClient()
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Contracts()
	if err != nil {
		t.Fatal(err)
	}
}
