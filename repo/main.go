package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgxutil"
	"golang.org/x/exp/slog"
)

type Entity interface {
	TableName() string
	Data() map[string]any
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
	v, err := pgxutil.SelectRow(ctx, r.db, `SELECT * FROM `+z.TableName()+` WHERE id = $1`, []any{id}, pgx.RowToStructByNameLax[T])
	if err != nil {
		return z, fmt.Errorf("select by id from %v failed: %w", z.TableName(), err)
	}
	return v, nil
}

func (r *Repo[T]) ListAll(ctx context.Context) ([]T, error) {
	var z T
	vs, err := pgxutil.Select(ctx, r.db, `SELECT * FROM `+z.TableName(), nil, pgx.RowToStructByNameLax[T])
	if err != nil {
		return nil, fmt.Errorf("select all from %v failed: %w", z.TableName(), err)
	}
	return vs, nil
}

func (r *Repo[T]) DeleteByID(ctx context.Context, id uint64) error {
	var z T
	_, err := r.db.Exec(ctx, `DELETE FROM `+z.TableName()+` WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete from %v failed: %w", z.TableName(), err)
	}
	return nil
}

func (r *Repo[T]) Insert(ctx context.Context, v T, vs ...T) error {
	xs := []map[string]any{v.Data()}
	for _, v := range vs {
		xs = append(xs, v.Data())
	}
	for _, x := range xs {
		delete(x, "id")
	}

	_, err := pgxutil.Insert(ctx, r.db, v.TableName(), xs)
	if err != nil {
		return fmt.Errorf("insert into %v failed: %w", v.TableName(), err)
	}
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

func (u User) Data() map[string]any {
	return map[string]any{
		"id":   u.ID,
		"name": u.Name,
	}
}

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

	repo := NewRepo[User](conn)

	for i := 0; i < 10; i++ {
		var name [8]byte
		rand.Read(name[:])
		err := repo.Insert(ctx, User{Name: hex.EncodeToString(name[:])})
		if err != nil {
			slog.Error("random user", err)
		}
	}

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
