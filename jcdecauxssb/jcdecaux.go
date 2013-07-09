/*
Package jcdecauxssb provides access to the JcDecaux self-serving bikes API.

Create a new client this way:

	client := JcDecaux.NewClient(APIKEY,nil)

With the client set, you can query the api this way:

	// Get the list of contracts:
	contracts := client.Contracts()
	// Get all the station of a specific contract (in this case Paris)
	parisStations := client.Stations("Paris")
	// Retrieve my specific station
	myStation := parisStations[42]
	// Update it
	myStation = client.Station("Paris",myStation.Number)

The full JcDecaux API is documented at https://developer.jcdecaux.com/#/opendata/vls?page=dynamic.
*/
package jcdecauxssb

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	// LibraryVersion represents this library version
	LibraryVersion = "2013.07.09"

	// BaseURL represents JcDecaux API base URL
	BaseURL = "https://api.jcdecaux.com/vls/v1/"

	// UserAgent represents this client User-Agent
	UserAgent = "go-jcdecaux/" + LibraryVersion
)

// A Client manages communication with the JcDecaux API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// UserAgent agent used when communicating with JcDecaux API.
	UserAgent string

	// Application Key
	Key string
}

// New returns a new JcDecaux API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func New(key string, httpClient *http.Client) (*Client, error) {
	if key == "" {
		return nil, errors.New("Key field must be populated")
	}

	// As a key is 40 characters long, check for conformance
	if len(key) != 40 {
		return nil, errors.New("The key format is invalid, it must be 40 characters long")
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: UserAgent,
		Key:       key,
	}
	return c, err
}
