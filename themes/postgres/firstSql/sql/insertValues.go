package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertValues(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlString := `
		INSERT INTO tasks(title, description, created_at, completed)
		VALUES ($1, $2, $3, $4)
	`

	_, err := conn.Exec(ctx, sqlString, task.Title, task.Description, task.Created_at, task.Completed)
	return err
}
