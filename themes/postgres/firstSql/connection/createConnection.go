package connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var PGX_CONNECT_STRING = "postgres://postgres:7355608@localhost:5432/postgres"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, PGX_CONNECT_STRING)
}
