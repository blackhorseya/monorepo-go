package mariadb

import (
	"time"

	"github.com/blackhorseya/monorepo-go/pkg/configx"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	defaultConns = 100
)

// NewClient create a new mariadb client.
func NewClient(config *configx.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.Storage.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "open mariadb failed")
	}

	conns := defaultConns
	if config.Storage.Conns > 0 {
		conns = config.Storage.Conns
	}

	db.SetConnMaxLifetime(15 * time.Minute)
	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns)

	return db, nil
}
