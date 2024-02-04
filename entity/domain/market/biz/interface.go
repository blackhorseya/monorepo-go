//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

// IMarketBiz is the interface for market biz.
type IMarketBiz interface {

}
