package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetTableData(ctx context.Context, conn *pgx.Conn) error {
	sqlString := `
		SELECT id, title, completed, completed_at FROM tasks
		WHERE completed = false
		ORDER BY id
	`

	rows, err := conn.Query(ctx, sqlString)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var completed bool
		var completed_at *time.Time

		if err := rows.Scan(&id, &title, &completed, &completed_at); err != nil {
			return err
		}

		fmt.Printf("ID: %d, Title: %s, Completed: %t, Time %v\n", id, title, completed, completed_at)
	}

	return nil
}
