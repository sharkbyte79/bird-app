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
	ObsByRegion(p RegionSearchParams) ([]BirdObservation, error)
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
	SubID       string  `json:"subId"` // unique ID for this observation/checklist
}

type RegionSearchParams struct {
	RegionCode string // alphanumeric code for region to filter observations by
	Back       int    // number of days to search backwards to
	MaxResults int    // maximum number of observations to get back
	Notable    bool   // retrieve observations marked as "notable"
}

// NewClient returns an API client for the eBird API configured with
// the given API key 'key' or an error.
func NewClient(tok string, hc *http.Client) (*Client, error) {
	// WARN This does not by any means validate that a non-empty key string
	// is a valid eBird API key, just that any string is provided.
	if strings.TrimSpace(tok) == "" {
		return nil, errors.New("no eBird API key was provided for eBird service")
	}

	c := &Client{
		apiKey:     tok,
		baseUrl:    BaseUrl,
		httpClient: hc,
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

func (c *Client) ObsByRegion(p RegionSearchParams) ([]BirdObservation, error) {
	rc, back, max, notable := p.RegionCode, p.Back, p.MaxResults, p.Notable

	endpointPattern := "/data/obs/%s/recent"
	if notable {
		endpointPattern += "/notable"
	}

	// Convert URL string to type URL to safely add query parameters, then
	// convert it back to string form.
	u, err := url.Parse(fmt.Sprintf(endpointPattern, rc))
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
