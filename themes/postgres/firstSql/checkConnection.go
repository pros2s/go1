package first_sql

import (
	"context"
	"fmt"

	"go1/themes/postgres/firstSql/connection"
	"go1/themes/postgres/firstSql/sql"
)

func CheckConnection(ctx context.Context) {
	connect, err := connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := sql.CreateTable(ctx, connect); err != nil {
		panic(err)
	}

	newIvs := sql.NewInsertValues("Ivs1", "ivs description")
	if err := sql.InsertValues(ctx, connect, newIvs); err != nil {
		panic(err)
	}

	fmt.Println("Connection success")
}
