// Holds the general aspects of the library

package jcdecauxssb

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	// BaseURL represents JcDecaux API base URL
	BaseURL = "https://api.jcdecaux.com/vls/v1/"

	// The default User-Agent used in the library
	// TODO: Add a library version
	UserAgent = "go-jcdecaux/"
)

// A Client manages communication with the JcDecaux API.
// It is the only way to "hack" the base URL
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

// Checks the key
func CheckKey(key string) error {
	// If there is no key, return an error
	if key == "" {
		return errors.New("Key field must be populated")
	}

	// As a key is 40 characters long, check for conformance
	if len(key) != 40 {
		return errors.New("The key format is invalid, it must be 40 characters long")
	}
	return nil
}

// New returns a new JcDecaux API client. If a nil httpClient is provided, http.DefaultClient will be used.
func New(key string, httpClient *http.Client) (*Client, error) {

	err := CheckKey(key) 
	if err != nil {
		return nil, err
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
