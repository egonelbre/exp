package schema_test

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/url"
	"testing"

	"github.com/jackc/pgx/v5"
)

func BenchmarkDatabase_FsyncOff(b *testing.B) {
	ctx := context.Background()
	WithDocker(ctx, b, false, func(b *testing.B, db *pgx.Conn, connstr string) {
		b.ResetTimer()
		defer b.StopTimer()

		for i := 0; i < b.N; i++ {
			withDatabase(ctx, connstr, b, func(b *testing.B, db *pgx.Conn) {
				_, err := db.Exec(ctx, DatabaseSchema)
				if err != nil {
					b.Fatal(err)
				}
			})
		}
	})
}

func BenchmarkDatabase_FsyncOn(b *testing.B) {
	ctx := context.Background()
	WithDocker(ctx, b, true, func(b *testing.B, db *pgx.Conn, connstr string) {
		b.ResetTimer()
		defer b.StopTimer()

		for i := 0; i < b.N; i++ {
			withDatabase(ctx, connstr, b, func(b *testing.B, db *pgx.Conn) {
				_, err := db.Exec(ctx, DatabaseSchema)
				if err != nil {
					b.Fatal(err)
				}
			})
		}
	})
}

func WithDatabase[TB testing.TB](ctx context.Context, tb TB, fn func(t TB, db *pgx.Conn)) {
	if *pgaddr == "" {
		tb.Skip("-database flag not defined")
	}
	withDatabase(ctx, *pgaddr, tb, fn)
}

func withDatabase[TB testing.TB](ctx context.Context, connstr string, tb TB, fn func(t TB, db *pgx.Conn)) {
	if connstr == "" {
		tb.Skip("connstr not defined")
	}

	var id [8]byte
	rand.Read(id[:])
	uniqueName := tb.Name() + "/" + hex.EncodeToString(id[:])

	// Create the first connection that we use to create the database.
	maindb, err := pgx.Connect(ctx, connstr)
	if err != nil {
		tb.Fatalf("Unable to connect to database: %v", err)
	}

	// Run the database creation query and defer the database cleanup query.
	if err := createDatabase(ctx, maindb, uniqueName); err != nil {
		tb.Fatalf("unable to create database: %v", err)
	}
	defer func() {
		if err := dropDatabase(ctx, maindb, uniqueName); err != nil {
			tb.Fatalf("unable to drop database: %v", err)
		}
	}()

	// Make a connection the new database.
	connstr, err = connstrWithDatabase(connstr, uniqueName)
	if err != nil {
		tb.Fatal(err)
	}

	db, err := pgx.Connect(ctx, connstr)
	if err != nil {
		tb.Fatalf("Unable to connect to database: %v", err)
	}
	defer func() { _ = db.Close(ctx) }()

	// Run our test code.
	fn(tb, db)
}

func sanitizeDatabaseName(schema string) string {
	return pgx.Identifier{schema}.Sanitize()
}

func connstrWithDatabase(connstr, database string) (string, error) {
	u, err := url.Parse(connstr)
	if err != nil {
		return "", fmt.Errorf("invalid connstr: %q", connstr)
	}
	u.Path = database
	return u.String(), nil
}

func createDatabase(ctx context.Context, db *pgx.Conn, name string) error {
	_, err := db.Exec(ctx, `CREATE DATABASE `+sanitizeDatabaseName(name)+`;`)
	return err
}

func dropDatabase(ctx context.Context, db *pgx.Conn, name string) error {
	_, err := db.Exec(ctx, `DROP DATABASE `+sanitizeDatabaseName(name)+`;`)
	return err
}
