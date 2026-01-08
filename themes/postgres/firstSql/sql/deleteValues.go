package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteValues(ctx context.Context, conn *pgx.Conn, deleteId int) error {
	sqlString := `
		DELETE FROM tasks
		WHERE id = $1
	`

	_, err := conn.Exec(ctx, sqlString, deleteId)
	return err
}
