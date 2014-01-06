// Holds the networking aspects of the library

package jcdecauxssb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Error struct {
	Response *http.Response
	Message  string `json:"error,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Request failed for %v, with code %v and message %v \n", e.Response.Request.URL.String(), e.Response.StatusCode, e.Message)
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified
func (c *Client) NewRequest(method, urlStr string, body string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	q := u.Query()
	if c.Key != "" && q.Get("apiKey") == "" {
		q.Set("apiKey", c.Key)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

	// Doing the http request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// TODO: Checks rate limit

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

// CheckResponse checks the API response for error, and returns it
// if present. A response is considered an error if it has non StatusOK
// code.
func CheckResponse(r *http.Response) error {
	if r.StatusCode == http.StatusOK {
		return nil
	}
	resp := &Error{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, resp)
	}
	return resp
}
