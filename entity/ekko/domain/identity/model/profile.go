package model

// Profile is a value object that represents a user's profile
type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
