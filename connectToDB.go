package main

import (
	"context"
	"database/sql"
	"fmt"
)

func connectToDB(cfg Config) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		cfg.ServerName, cfg.User, cfg.Password, cfg.Port, cfg.Database)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия соединения: %w", err)
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("Ошибка проверки соединения: %w", err)
	}

	return db, nil
}
