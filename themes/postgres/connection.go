package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()

	connect, err := pgx.Connect(ctx, "postgres://postgres:7355608@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := connect.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Connection success")
}
