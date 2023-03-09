package schema_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	dockerPool *dockertest.Pool
)

func TestMain(m *testing.M) {
	os.Exit(mainWithDocker(m))
}

func mainWithDocker(m *testing.M) int {
	var err error
	dockerPool, err = dockertest.NewPool("")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	dockerPool.MaxWait = 120 * time.Second

	return m.Run()
}

func BenchmarkDocker(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		WithDocker(ctx, b, func(b *testing.B, db *pgx.Conn) {
			_, err := db.Exec(ctx, DatabaseSchema)
			if err != nil {
				b.Fatal(err)
			}
		})
	}
}

func WithDocker[TB testing.TB](ctx context.Context, tb TB, fn func(tb TB, db *pgx.Conn)) {
	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15",
		Cmd: []string{"-c", "fsync=off"},
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user",
			"POSTGRES_DB=main",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		tb.Fatalf("Could not start resource: %s", err)
	}
	defer func() {
		if err := dockerPool.Purge(resource); err != nil {
			tb.Errorf("failed to stop: %v", err)
		}
	}()

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseConnstr := fmt.Sprintf("postgres://user:secret@%s/main?sslmode=disable", hostAndPort)

	err = resource.Expire(2 * 60) // hard kill the container after 2 minutes, just in case.
	if err != nil {
		tb.Fatalf("Unable to set container expiration: %v", err)
	}

	var db *pgx.Conn
	err = dockerPool.Retry(func() error {
		db, err = pgx.Connect(ctx, databaseConnstr)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		tb.Fatal("unable to connect to Postgres", err)
	}

	defer func() { _ = db.Close(ctx) }()

	fn(tb, db)
}
