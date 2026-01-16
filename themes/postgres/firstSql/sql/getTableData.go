package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetTableData(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	// select
	sqlString := `
		SELECT id, title, description, completed, completed_at FROM tasks
		WHERE completed = false
		ORDER BY id
	`

	// error
	rows, err := conn.Query(ctx, sqlString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// writing
	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.Completed_at); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
