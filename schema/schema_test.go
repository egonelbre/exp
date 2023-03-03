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

func BenchmarkSchema(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		WithDatabase(ctx, b, func(b *testing.B, db *pgx.Conn) {
			_, err := db.Exec(ctx, "CREATE TABLE accounts ( user_id serial PRIMARY KEY )")
			if err != nil {
				b.Fatal(err)
			}
		})
	}
}

var pgaddr = flag.String("database", os.Getenv("DATABASE_URL"), "database address")

func WithDatabase[TB testing.TB](ctx context.Context, t TB, fn func(t TB, db *pgx.Conn)) {
	if *pgaddr == "" {
		t.Skip("-database flag not defined")
	}

	var schema [8]byte
	rand.Read(schema[:])
	schemaname := t.Name() + "/" + hex.EncodeToString(schema[:])

	db, err := pgx.Connect(ctx, connstrWithSchema(*pgaddr, schemaname))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		t.Fatal(err)
	}
	defer func() { _ = db.Close(ctx) }()

	if err := createSchema(ctx, db, schemaname); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := dropSchema(ctx, db, schemaname); err != nil {
			t.Fatal(err)
		}
	}()

	t.Helper()
	fn(t, db)
}

func sanitizeSchema(schema string) string {
	return pgx.Identifier{schema}.Sanitize()
}

func connstrWithSchema(connstr, schema string) string {
	return connstr + "&search_path=" + url.QueryEscape(sanitizeSchema(schema))
}

func createSchema(ctx context.Context, db *pgx.Conn, schema string) error {
	_, err := db.Exec(ctx, `create schema if not exists `+sanitizeSchema(schema)+`;`)
	return err
}

func dropSchema(ctx context.Context, db *pgx.Conn, schema string) error {
	_, err := db.Exec(ctx, `drop schema `+sanitizeSchema(schema)+` cascade;`)
	return err
}
