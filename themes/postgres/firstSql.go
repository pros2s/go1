package postgres

import (
	"context"

	first_sql "go1/themes/postgres/firstSql"
)

func TestFirstSQL() {
	ctx := context.Background()

	first_sql.CheckConnection(ctx)
}
