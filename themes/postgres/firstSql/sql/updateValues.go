package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateValues(
	ctx context.Context,
	conn *pgx.Conn,
	task TaskModel,
) error {
	sqlString := `
		UPDATE tasks
		SET title = $1, description = $2, completed_at = $3
		WHERE id = $4
	`

	_, err := conn.Exec(ctx, sqlString, task.Title, task.Description, task.Completed_at, task.ID)

	return err
}
