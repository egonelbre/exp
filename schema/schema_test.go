package schema_test

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var pgaddr = flag.String("database", os.Getenv("DATABASE_URL"), "database address")

func BenchmarkSchema(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		WithSchema(ctx, b, func(b *testing.B, db *pgx.Conn) {
			_, err := db.Exec(ctx, DatabaseSchema)
			if err != nil {
				b.Fatal(err)
			}
		})
	}
}

func WithSchema[TB testing.TB](ctx context.Context, tb TB, fn func(t TB, db *pgx.Conn)) {
	if *pgaddr == "" {
		tb.Skip("-database flag not defined")
	}
	dbaddr := *pgaddr

	var id [8]byte
	rand.Read(id[:])
	uniqueName := tb.Name() + "/" + hex.EncodeToString(id[:])

	connstr, err := connstrWithSchema(dbaddr, uniqueName)
	if err != nil {
		tb.Fatal(err)
	}

	db, err := pgx.Connect(ctx, connstr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		tb.Fatal(err)
	}
	defer func() { _ = db.Close(ctx) }()

	if err := createSchema(ctx, db, uniqueName); err != nil {
		tb.Fatal(err)
	}
	defer func() {
		if err := dropSchema(ctx, db, uniqueName); err != nil {
			tb.Fatal(err)
		}
	}()

	fn(tb, db)
}

func sanitizeSchema(schema string) string {
	return pgx.Identifier{schema}.Sanitize()
}

func connstrWithSchema(connstr, schema string) (string, error) {
	u, err := url.Parse(connstr)
	if err != nil {
		return "", fmt.Errorf("invalid connstr: %q", connstr)
	}
	q := u.Query()
	q.Set("search_path", sanitizeSchema(schema))
	u.RawQuery  = q.Encode()
	return u.String(), nil
}

func createSchema(ctx context.Context, db *pgx.Conn, schema string) error {
	_, err := db.Exec(ctx, `create schema if not exists `+sanitizeSchema(schema)+`;`)
	return err
}

func dropSchema(ctx context.Context, db *pgx.Conn, schema string) error {
	_, err := db.Exec(ctx, `drop schema `+sanitizeSchema(schema)+` cascade;`)
	return err
}
