package model

// Ticket is a value object that represents a ticket.
type Ticket struct {
	ID        string `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

// NewTicket creates a new ticket.
func NewTicket(id, title string) (*Ticket, error) {
	return &Ticket{
		ID:    id,
		Title: title,
	}, nil
}
