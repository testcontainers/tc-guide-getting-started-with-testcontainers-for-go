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
	*postgres.PostgresContainer
	ConnectionString string
}

// SetupPostgres creates an instance of the postgres container type
func SetupPostgres(ctx context.Context) (*PostgresContainer, error) {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.2-alpine"),
		postgres.WithInitScripts(filepath.Join("testdata", "init-db.sh")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}, nil
}

var customerRepository *CustomerRepository

func TestMain(m *testing.M) {
	ctx := context.Background()
	pgContainer, err := SetupPostgres(ctx)
	if err != nil {
		log.Fatalf("failed to setup Postgres container")
	}

	customerRepository, err = NewCustomerRepository(ctx, pgContainer.ConnectionString)
	if err != nil {
		log.Fatalf("failed to initialize customerRepository")
	}
	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Fatalf("error terminating postgres container: %s", err)
		}
	}()

	os.Exit(m.Run())
}
