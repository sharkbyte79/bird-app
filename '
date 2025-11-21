package client

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const baseUrl = "https://api.ebird.org/v2" // Root for eBird API 2.0

type Client struct {
	apiKey     string       // API key for authenticating requests
	baseUrl    string       // URL base for eBird API
	httpClient *http.Client // Actual HttpClient attached to this wrapper
}

type BirdObservation struct {
	SpeciesCode string    `json:"speciesCode"`
	ComName     string    `json:"comName"`
	SciName     string    `json:"sciName"`
	LocID       string    `json:"locId"`
	LocName     string    `json:"locName"`
	ObsDt       time.Time `json:"obsDt"`
	HowMany     int       `json:"howMany"`
	Lat         float64   `json:"lat"`
	Lng         float64   `json:"lng"`
	ObsValid    bool      `json:"obsValid"`
	ObsReviewed bool      `json:"obsReviewed"`
	LocPrivate  bool      `json:"locationPrivate"`
}

// NewClient returns an API client for the eBird API configured with
// the given API key 'key' or an error.
func NewClient(key string, httpClient *http.Client) (*Client, error) {
	// WARN: This does not by any means validate that a non-empty key string
	// is a valid eBird API key, just that any string is provided.
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("No eBird API key was provided for eBird service")
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
		return nil, errors.New("Client requires a valid endpoint for query")
	}

	// Construct full url to query
	url := fmt.Sprintf("%s/%s", c.baseUrl, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create eBird request: %w", err)
	}

	// TODO: configure request correctly and submit http req
	bearer := "Bearer " + c.apiKey
	req.Header.Add("Authorization", bearer)

	return nil, nil
}
