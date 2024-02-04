//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package finmindx

import (
	"time"

	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// Dialer is used to dial the finmindx service.
type Dialer interface {
	// Do is used to do the request.
	Do(ctx contextx.Contextx, dataset string, params map[string]string, v any) error

	TaiwanStockPrice(
		ctx contextx.Contextx,
		symbol string,
		start, end time.Time,
	) (res *TaiwanStockPriceResponse, err error)
}
