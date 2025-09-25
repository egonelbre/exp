package gormdb

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"site.test/user"
)

type DB struct {
	underlying *sql.DB
	db         *gorm.DB
}

func OpenSqlite(ctx context.Context, path string) (*DB, error) {
	underlying, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	dialect := sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        path,
		Conn:       underlying,
	}

	db, err := gorm.Open(dialect)
	if err != nil {
		underlying.Close()
		return nil, err
	}

	return &DB{
		underlying: underlying,
		db:         db,
	}, nil
}

func (db *DB) Migrate(ctx context.Context) error {
	err := db.db.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Close() error {
	return db.underlying.Close()
}
