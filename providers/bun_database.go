package providers

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
	"time"
)

type BunDatabase struct {
	Db *bun.DB
}

func NewBunDatabase() *BunDatabase {
	return &BunDatabase{}
}

func (b *BunDatabase) Setup() *BunDatabase {
	pgConn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(os.Getenv("DATABASE_ADDRESS")),
		pgdriver.WithTLSConfig(nil),
		pgdriver.WithUser(os.Getenv("DATABASE_USER")),
		pgdriver.WithPassword(os.Getenv("DATABASE_PASSWORD")),
		pgdriver.WithDatabase(os.Getenv("DATABASE_NAME")),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)
	sqldb := sql.OpenDB(pgConn)
	b.Db = bun.NewDB(sqldb, pgdialect.New())
	return b
}
