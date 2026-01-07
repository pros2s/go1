package first_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlString := `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(100) NOT NULL,
			description VARCHAR(1000) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			completed BOOLEAN NOT NULL,
			completed_at TIMESTAMP
		)
	`

	_, err := conn.Exec(ctx, sqlString)
	if err != nil {
		return err
	}

	return nil
}
