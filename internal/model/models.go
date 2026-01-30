package model

import (
	"time"
)

// User models a single Birdup account.
type User struct {
	FirebaseId string    `json:"firebaseID"` // User ID assigned by Firebase
	Email      string    `json:"email"`      // Email provided at sign up
	CreatedAt  time.Time `json:"createdAt"`  // Timestamp for account creation
}

// BirdFollow models the relation between a single User account and a bird species whose
// observations that User is following.
type BirdFollow struct {
	FirebaseId  string    // Identifier for User associated with this follow relation
	SpeciesCode string    // Id assigned by eBird service for a bird species
	CreatedAt   time.Time // Timestamp for follow creation
}
