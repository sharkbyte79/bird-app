package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	BaseUrl = "https://api.ebird.org/v2" // Root for eBird API 2.0
)

type EBirdClient interface {
	RecentObsByRegion(p RegionSearchParams) ([]BirdObservation, error)
	NotableObsByRegion(p RegionSearchParams) ([]BirdObservation, error)
}

type Client struct {
	apiKey     string       // API key for authenticating requests
	baseUrl    string       // URL base for eBird API
	httpClient *http.Client // Pointer to the HttpClient attached to this wrapper
}

type BirdObservation struct {
	SpeciesCode string  `json:"speciesCode"` // designates particular bird species
	ComName     string  `json:"comName"`     // natural language name for bird species in question
	SciName     string  `json:"sciName"`     // scientific name for bird species in question
	LocID       string  `json:"locId"`       // ID for location of observation
	LocName     string  `json:"locName"`     // natural language name for bird species in question
	ObsDt       string  `json:"obsDt"`       // date-time string of observation
	HowMany     int     `json:"howMany"`     // quantity of this bird species observed
	Lat         float64 `json:"lat"`         // latitudinal coordinates of this observation
	Lng         float64 `json:"lng"`         // longitudinal coordinates of this observation
	ObsValid    bool    `json:"obsValid"`
	ObsReviewed bool    `json:"obsReviewed"` // has this observation report been reviewed
	LocPrivate  bool    `json:"locationPrivate"`
	SubID       string  `json:"subId"`
}

type RegionSearchParams struct {
	RegionCode string
	Back       int
	MaxResults int
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
		baseUrl:    BaseUrl,
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

	// Attach eBird API key/token to request headers
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

func (c *Client) RecentObsByRegion(p RegionSearchParams) ([]BirdObservation, error) {
	rc, back, max := p.RegionCode, p.Back, p.MaxResults

	// Reject attempt for region code without correct shape
	ok := verifyRegionCode(rc)
	if !ok {
		return nil, fmt.Errorf("string %s does not match region code format", rc)
	}

	// Convert URL string to type URL to safely add query parameters, then
	// convert it back to string form.
	u, err := url.Parse(fmt.Sprintf("/data/obs/%s/recent", rc))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("back", strconv.Itoa(back))
	q.Add("maxResults", strconv.Itoa(max))
	u.RawQuery = q.Encode()

	endpoint := u.String()
	res, err := c.eBirdFetch(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch recent observations, %w", err)
	}

	return res, nil
}

func (c *Client) NotableObsByRegion(p RegionSearchParams) ([]BirdObservation, error) {
	// Unpack search parameters to append them to handle appending to query
	rc, back, max := p.RegionCode, p.Back, p.MaxResults

	// Reject attempt for region code without correct shape
	ok := verifyRegionCode(rc)
	if !ok {
		return nil, fmt.Errorf("string %s does not match region code format", rc)
	}

	// Convert URL string to type URL to safely add query parameters, then
	// convert it back to string form.
	u, err := url.Parse(fmt.Sprintf("/data/obs/%s/notable", rc))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("back", strconv.Itoa(back))
	q.Add("maxResults", strconv.Itoa(max))
	u.RawQuery = q.Encode()

	endpoint := u.String()
	res, err := c.eBirdFetch(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch notable observations, %w", err)
	}

	return res, nil
}
