//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

// IStringBiz string biz interface.
type IStringBiz interface {
}
