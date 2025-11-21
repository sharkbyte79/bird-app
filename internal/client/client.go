package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const baseUrl = "https://api.ebird.org/v2" // Root for eBird API 2.0

type Client struct {
	apiKey     string       // API key for authenticating requests
	baseUrl    string       // URL base for eBird API
	httpClient *http.Client // Pointer to the HttpClient attached to this wrapper
}

type BirdObservation struct {
	SpeciesCode string  `json:"speciesCode"` // designates particular bird species
	ComName     string  `json:"comName"`     // natural language name for bird species in question
	SciName     string  `json:"sciName"`
	LocID       string  `json:"locId"`
	LocName     string  `json:"locName"`
	ObsDt       string  `json:"obsDt"`
	HowMany     int     `json:"howMany"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	ObsValid    bool    `json:"obsValid"`
	ObsReviewed bool    `json:"obsReviewed"`
	LocPrivate  bool    `json:"locationPrivate"`
	SubID       string  `json:"subId"`
}

// NewClient returns an API client for the eBird API configured with
// the given API key 'key' or an error.
func NewClient(key string, httpClient *http.Client) (*Client, error) {
	// WARN: This does not by any means validate that a non-empty key string
	// is a valid eBird API key, just that any string is provided.
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("no eBird API key was provided for eBird service")
	}

	c := &Client{
		apiKey:     key,
		baseUrl:    baseUrl,
		httpClient: httpClient,
	}

	return c, nil
}

func (c *Client) eBirdFetch(endpoint string) ([]BirdObservation, error) {
	if strings.TrimSpace(endpoint) == "" {
		return nil, errors.New("client requires a valid endpoint for query")
	}

	req, err := http.NewRequest("GET", c.baseUrl+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create eBird request: %w", err)
	}

	req.Header.Add("X-eBirdApiToken", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to receive eBird response: %w", err)
	}

	switch res.StatusCode {
	case http.StatusOK:
		var data []BirdObservation

		// Attempt to unmarshal body for successful response
		err := json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			return nil, fmt.Errorf("failed to decode eBird response: %w", err)
		}

		return data, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("eBird API key is invalid")
	default:
		return nil, fmt.Errorf("failed to get observation data, %w", err)
	}
}

// verifyRegionCode returns true if the given string rc represents
// a valid region code of the form "XX-00-...", and false otherwise.
func verifyRegionCode(rc string) bool {
	// TODO: check if region code string matches expected format
	// by regex
	// b, err := regexp.Match("", []byte(rc))
	return true
}

func (c *Client) RecentObs(rc string) ([]BirdObservation, error) {
	// Reject attempt with region code without correct shape
	ok := verifyRegionCode(rc)
	if !ok {
		return nil, fmt.Errorf("string %s does not match region code format", rc)
	}

	endpoint := fmt.Sprintf("/data/obs/%s/recent", rc)
	res, err := c.eBirdFetch(endpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch recent observations, %w", err)
	}

	return res, nil
}
