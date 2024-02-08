package model

// Ticket is a value object that represents a ticket.
type Ticket struct {
	ID        string `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}
