//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/agg"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// ListOptions is the options for list issues.
type ListOptions struct {
	Offset  int
	Limit   int
	OwnerID string
}

// IIssueRepo is the interface for issue repository.
type IIssueRepo interface {
	List(ctx contextx.Contextx, opts ListOptions) (items []agg.Issue, err error)
	GetByID(ctx contextx.Contextx, id string) (issue agg.Issue, err error)
	Create(ctx contextx.Contextx, item agg.Issue) (id string, err error)
}
