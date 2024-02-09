package agg

import (
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/model"
)

// Issue is an aggregate root that represents an issue.
type Issue struct {
	ticket *model.Ticket
}

// NewIssue creates a new issue.
func NewIssue(title string) (Issue, error) {
	ticket, err := model.NewTicket(title)
	if err != nil {
		return Issue{}, err
	}

	return Issue{
		ticket: ticket,
	}, nil
}

// GetID returns the issue ID.
func (x *Issue) GetID() string {
	return x.ticket.ID
}

// SetID sets the issue ID.
func (x *Issue) SetID(id string) {
	x.ticket.ID = id
}

// GetTitle returns the issue title.
func (x *Issue) GetTitle() string {
	return x.ticket.Title
}

// GetCompleted returns the issue completed status.
func (x *Issue) GetCompleted() bool {
	return x.ticket.Completed
}

// SetCompleted sets the issue completed status.
func (x *Issue) SetCompleted(completed bool) {
	x.ticket.Completed = completed
}
