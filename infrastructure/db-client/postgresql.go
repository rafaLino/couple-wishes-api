package dbclient

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rafaLino/couple-wishes-api/infrastructure/db"
)

type DbContext struct {
	client  *db.Queries
	context context.Context
}

func NewDBContext() *DbContext {
	return &DbContext{}
}

func (c *DbContext) Connect() error {
	connectionString := os.Getenv("CONNECTION_STRING")
	ctx := context.Background()
	connection, err := pgx.Connect(ctx, connectionString)

	if err != nil {
		return err
	}

	c.client = db.New(connection)
	c.context = ctx

	Ping(connection, ctx)
	return err
}

func (c *DbContext) GetClient() (*db.Queries, error) {
	return c.client, nil
}

func (c *DbContext) GetContext() context.Context {
	return c.context
}

func Ping(connection *pgx.Conn, ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := connection.Ping(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
