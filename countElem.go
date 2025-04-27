package main

import (
	"context"
	"database/sql"
	"log"
)

func countElem(ctx context.Context, db *sql.DB, tablename string) (int, error) {
	query := `SELECT COUNT(*) FROM ` + tablename + `;`
	var count int
	err := db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса:", err)
	}
	return count, nil
}
