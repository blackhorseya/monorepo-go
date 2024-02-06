//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package notify

import (
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// Notifier is an interface for sending notifications.
type Notifier interface {
	SendText(ctx contextx.Contextx, message string) error
}
