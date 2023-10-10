//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

// IUserBiz is a user biz interface
type IUserBiz interface {
}
