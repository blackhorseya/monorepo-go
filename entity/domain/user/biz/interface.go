//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	userM "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IUserBiz is a user biz interface
type IUserBiz interface {
	// Login serve caller to given username and password to login system
	Login(ctx contextx.Contextx, username, password string) (account *userM.UserAccount, err error)

	// Signup serve caller to given username and password to signup system
	Signup(ctx contextx.Contextx, username, password string) (account *userM.UserAccount, err error)

	// Logout serve caller to logout system
	Logout(ctx contextx.Contextx, token string) error
}
