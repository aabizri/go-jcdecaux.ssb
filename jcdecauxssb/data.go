package jcdecauxssb

import "net/url"

// A contract is the service provided by JcDecaux to the city(es)
type Contract struct {
	Name           string   `json:"name,omitempty"`            //is the identifier of the contract
	CommercialName string   `json:"commercial_name,omitempty"` //is the commercial name of the contract, the one users usually know
	CountryCode    string   `json:"country_code,omitempty"`    //is the code (ISO 3166) of the country
	Cities         []string `json:"cities,omitempty"`          //the cities that are concerned by this contract
}

// A station hosts bike stands
type Station struct {

	// Static data
	Number   uint     `json:"number,omitempty"`   //number of the station. This is NOT an id, thus it is unique only inside a contract.
	Name     string   `json:"name,omitempty"`     //name of the station
	Address  string   `json:"address,omitempty"`  //address of the station. As it is raw data, sometimes it will be more of a comment than an address.
	Position position `json:"position,omitempty"` //position of the station in WGS84 format
	Banking  bool     `json:"banking,omitempty"`  //indicates whether this station has a payment terminal
	Bonus    bool     `json:"bonus,omitempty"`    //indicates whether this is a bonus station

	// Dynamic data
	Status          string `json:"status,omitempty"`                //indicates whether this station is CLOSEDor OPEN
	BikeStands      uint   `json:"bike_stands,omitempty"`           //the number of operational bike stands at this station
	AvailableStands uint   `json:"available_bike_stands,omitempty"` //the number of available bike stands at this station
	AvailableBikes  uint   `json:"available_bikes,omitempty"`       //the number of available and operational bikes at this station
	LastUpdate      int64  `json:"last_update,omitempty"`           //timestamp indicating the last update time in milliseconds since Epoch (UNIX timestamp*1000)
}

//Position of the station in WGS84 format
type position struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

//Returns all the contracts of the JcDecaux self-serving bikes service
func (c *Client) Contracts() ([]Contract, error) {
	req, err := c.NewRequest("GET", "contracts", "")
	if err != nil {
		return nil, err
	}

	contracts := new([]Contract)

	_, err = c.Do(req, contracts)
	return *contracts, err
}

// If parameter == "" :
//	Retrieves all stations of all contracts
// Else:
//	Retrieves all stations of the specified contract
func (c *Client) Stations(contract string) ([]Station, error) {

	u := "stations"

	params := url.Values{}
	if contract != "" {
		params.Add("contract", contract)
		u += "?" + params.Encode()
	}

	req, err := c.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	stations := new([]Station)

	_, err = c.Do(req, stations)
	return *stations, err
}

func (c *Client) StationData(contract string, number uint) (*Station, error) {

	params := url.Values{}
	params.Add("contract", contract)
	u := "stations?" + params.Encode()

	req, err := c.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	station := new(Station)

	_, err = c.Do(req, station)
	return station, err
}
