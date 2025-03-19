package build

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	maxConnTime        = time.Minute
	maxOpenConnections = 5
	maxIdleConnections = 10
)

func (b *Builder) postgresClient() *bun.DB {
	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(b.conf.PostgresDSN)))

	conn.SetConnMaxIdleTime(maxConnTime)
	conn.SetConnMaxLifetime(maxConnTime)
	conn.SetMaxOpenConns(maxOpenConnections)
	conn.SetMaxIdleConns(maxIdleConnections)

	db := bun.NewDB(conn, pgdialect.New())

	return db
}
