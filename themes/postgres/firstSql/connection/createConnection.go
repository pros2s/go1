package connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	conn_path := os.Getenv("CONNECTION_PATH")
	return pgx.Connect(ctx, conn_path)
}
