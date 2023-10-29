package mariadb

import (
	"time"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	defaultConns = 100
)

// NewClient create a new mariadb client.
func NewClient(config *configx.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.Persistence.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "open mariadb failed")
	}

	conns := defaultConns
	if config.Persistence.Conns > 0 {
		conns = config.Persistence.Conns
	}

	db.SetConnMaxLifetime(15 * time.Minute)
	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns)

	return db, nil
}
