package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateValues(
	ctx context.Context,
	conn *pgx.Conn,
	updateId int,
	newValue bool,
) error {
	sqlString := `
		UPDATE tasks
		SET completed = $2
		WHERE id = $1
	`

	_, err := conn.Exec(ctx, sqlString, updateId, newValue)
	return err
}
