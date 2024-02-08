package model

// User is a domain model that represents a user
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`

	Profile Profile `json:"profile"`
}
