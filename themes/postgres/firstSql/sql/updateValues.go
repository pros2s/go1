package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateValues(
	ctx context.Context,
	conn *pgx.Conn,
	updateCompleted bool,
	newValue string,
) error {
	sqlString := `
		UPDATE tasks
		SET description = $2
		WHERE completed = $1
	`

	_, err := conn.Exec(ctx, sqlString, updateCompleted, newValue)
	return err
}
