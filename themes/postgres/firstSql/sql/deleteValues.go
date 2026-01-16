package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteValues(ctx context.Context, conn *pgx.Conn, deleteIds []int) error {
	sqlString := `
		DELETE FROM tasks
		WHERE id = ANY($1)
	`

	_, err := conn.Exec(ctx, sqlString, deleteIds)
	return err
}
