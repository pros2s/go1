package first_sql

import (
	"context"
	"fmt"
)

func CheckConnection(ctx context.Context) {
	connect, err := CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := CreateTable(ctx, connect); err != nil {
		panic(err)
	}

	fmt.Println("Connection success")
}
