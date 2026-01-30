package dto

type CreateUserRequest struct {
	FirebaseID string `json:"firebaseId"`
	Email      string `json:"email"`
}

type GetUserRequest struct {
	FirebaseID string `json:"firebaseId"`
}
