//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IStringBiz string biz interface.
type IStringBiz interface {
	// Uppercase serve caller to given string value to uppercase.
	Uppercase(ctx contextx.Contextx, value string) (upper string, err error)

	// Count serve caller to given string value to count.
	Count(ctx contextx.Contextx, value string) int
}
