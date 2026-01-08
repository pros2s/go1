package sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type InsertValuesStruct struct {
	title      string
	desc       string
	created_at time.Time
	completed  bool
}

func NewInsertValues(title string, desc string) InsertValuesStruct {
	return InsertValuesStruct{
		title:      title,
		desc:       desc,
		created_at: time.Now(),
		completed:  false,
	}
}

func InsertValues(ctx context.Context, conn *pgx.Conn, ivs InsertValuesStruct) error {
	sqlString := `
		INSERT INTO tasks(title, description, created_at, completed)
		VALUES ($1, $2, $3, $4)
	`

	_, err := conn.Exec(ctx, sqlString, ivs.title, ivs.desc, ivs.created_at, ivs.completed)
	return err
}
