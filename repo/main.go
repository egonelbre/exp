package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/slog"
)

type Entity interface {
	TableName() string
}

type Repo[T Entity] struct {
	db *pgxpool.Pool
}

func NewRepo[T Entity](db *pgxpool.Pool) *Repo[T] {
	return &Repo[T]{
		db: db,
	}
}

func (r *Repo[T]) ByID(ctx context.Context, id uint64) (T, error) {
	var z T
	rows, err := r.db.Query(ctx, `SELECT * FROM `+z.TableName()+` WHERE id = $1`, id)
	if err != nil {
		return z, err
	}
	z, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[T])
	return z, err
}

func (r *Repo[T]) ListAll(ctx context.Context) ([]T, error) {
	var z T
	rows, err := r.db.Query(ctx, `SELECT * FROM `+z.TableName())
	if err != nil {
		return nil, err
	}
	zs, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[T])
	return zs, err
}

func (r *Repo[T]) DeleteByID(ctx context.Context, id uint64) error {
	var z T
	_, err := r.db.Exec(ctx, `DELETE FROM `+z.TableName()+` WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo[T]) Insert(ctx context.Context, v T) error {
	return nil
}

func (r *Repo[T]) Update(ctx context.Context, v T) error {
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		slog.Error("run", err)
		os.Exit(1)
	}
}

type User struct {
	ID   uint64
	Name string
}

func (User) TableName() string { return "users" }

func run(ctx context.Context) error {
	conn, err := pgxpool.New(ctx, `host=localhost port=5432 dbname=storj user=storj password=storj`)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	if err := conn.Ping(ctx); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	_, err = conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT)`)
	if err != nil {
		return fmt.Errorf("database creation failed: %w", err)
	}

	var z [8]byte
	rand.Read(z[:])

	_, err = conn.Exec(ctx, `INSERT INTO users (name) values ($1)`, hex.EncodeToString(z[:]))
	if err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}

	repo := NewRepo[User](conn)
	items, err := repo.ListAll(ctx)
	if err != nil {
		return fmt.Errorf("ListAll failed: %w", err)
	}
	for _, item := range items {
		fmt.Println(item.ID, item.Name)
		if err := repo.DeleteByID(ctx, item.ID); err != nil {
			slog.Error("deletion failed", err)
		}
	}

	return nil
}
