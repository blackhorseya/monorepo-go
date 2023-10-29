package mariadb

import (
	"github.com/jmoiron/sqlx"
)

const (
	defaultConns = 100
)

// NewClient create a new mariadb client.
func NewClient() (*sqlx.DB, error) {
	// todo: 2023/10/30|sean|impl me
	panic("implement me")
}
