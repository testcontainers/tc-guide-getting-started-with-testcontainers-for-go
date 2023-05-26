package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	Container        testcontainers.Container
	CloseFn          func()
	ConnectionString string
}

// SetupPostgres creates an instance of the postgres container type
func SetupPostgres(ctx context.Context) (*PostgresContainer, error) {
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.2-alpine"),
		postgres.WithInitScripts(filepath.Join("testdata", "init-user-db.sh")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &PostgresContainer{
		Container: container,
		CloseFn: func() {
			if err := container.Terminate(ctx); err != nil {
				log.Fatalf("error terminating postgres container: %s", err)
			}
		},
		ConnectionString: connStr,
	}, nil
}

var customerRepository *CustomerRepository

func TestMain(m *testing.M) {
	ctx := context.Background()
	pgContainer, err := SetupPostgres(ctx)
	if err != nil {
		log.Fatalf("failed to setup Postgres container")
	}
	customerRepository = NewCustomerRepository(ctx, pgContainer.ConnectionString)
	defer pgContainer.CloseFn()

	os.Exit(m.Run())
}
