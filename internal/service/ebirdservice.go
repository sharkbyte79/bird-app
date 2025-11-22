// Package ebirdservice provides a thin service layer over the eBird API client.
package ebirdservice

import (
	"fmt"
	"net/http"
	"regexp"

	ac "github.com/sharkbyte79/bird-app/internal/client"
)

type EBirdService struct {
	client ac.EBirdClient
}

// validRegionCode returns true if rc matches the format of a value
// region code as accepted by the eBird API, and false otherwise.
func validRegionCode(rc string) bool {
	pattern := "^[A-Za-z][A-Za-z](-[0-9][0-9])?"
	ok, err := regexp.Match(pattern, []byte(rc))
	if err != nil {
		// assume false or some malformity for error
		return false
	}

	return ok

}
func NewEBirdService(tok string, hc *http.Client) (*EBirdService, error) {
	ebc, err := ac.NewClient(tok, hc)
	if err != nil {
		return nil, fmt.Errorf("failed to create eBird service instance: %w", err)
	}

	return &EBirdService{client: ebc}, nil
}

func (s *EBirdService) RecentObsByRegion(rc string, back, max int) ([]ac.BirdObservation, error) {
	// Reject region code that doesn't match expected format
	if !validRegionCode(rc) {
		return nil, fmt.Errorf("invalid region code: %s", rc)
	}

	params := ac.RegionSearchParams{RegionCode: rc, Back: back, MaxResults: max}
	return s.client.RecentObsByRegion(params)
}

func (s *EBirdService) NotableObsByRegion(rc string, back, max int) ([]ac.BirdObservation, error) {
	// Reject region code that doesn't match expected format
	if !validRegionCode(rc) {
		return nil, fmt.Errorf("invalid region code: %s", rc)
	}

	params := ac.RegionSearchParams{RegionCode: rc, Back: back, MaxResults: max}
	return s.client.NotableObsByRegion(params)
}
