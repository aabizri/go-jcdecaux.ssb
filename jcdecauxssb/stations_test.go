package jcdecauxssb

import "testing"

func TestAllStations(t *testing.T) {


	client, err := getTestClient()
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Stations("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestContractStations(t *testing.T) {
	client, err := getTestClient()
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Stations("Paris")
	if err != nil {
		t.Fatal(err)
	}
}
