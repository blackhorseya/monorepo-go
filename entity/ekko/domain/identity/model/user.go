package model

// User is a domain model that represents a user
type User struct {
	ID     string `json:"id"`
	Active bool   `json:"active,omitempty"`

	Profile Profile `json:"profile,omitempty"`
}
