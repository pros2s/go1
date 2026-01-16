package first_sql

import (
	"context"
	"fmt"
	"time"

	"go1/themes/postgres/firstSql/connection"
	"go1/themes/postgres/firstSql/sql"

	"github.com/k0kubun/pp"
)

func CheckConnection(ctx context.Context) {
	connect, err := connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := sql.CreateTable(ctx, connect); err != nil {
		panic(err)
	}

	// newTask := sql.NewTask("Model title1", "It is model title1")
	// if err := sql.InsertValues(ctx, connect, newTask); err != nil {
	// 	panic(err)
	// }
	if err := sql.DeleteValues(ctx, connect, []int{9, 10}); err != nil {
		panic(err)
	}

	tasks, err := sql.GetTableData(ctx, connect)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 8 {
			task.Title = "Eight title"
			task.Description = "Eight description"
			timeNow := time.Now()
			task.Completed_at = &timeNow

			if err := sql.UpdateValues(ctx, connect, task); err != nil {
				panic(err)
			}

			break
		}
	}

	pp.Println(sql.GetTableData(ctx, connect))
	fmt.Println("Connection success")
}
