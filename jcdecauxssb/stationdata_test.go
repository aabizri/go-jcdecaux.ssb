package jcdecauxssb

import "testing"

func TestStationData(t *testing.T) {

	client, err := getTestClient()
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.StationData("Paris", 10023)
	if err != nil {
		t.Fatal(err)
	}
}
