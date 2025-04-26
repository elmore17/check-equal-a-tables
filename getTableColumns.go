package main

import "database/sql"

func getTableColumns(db *sql.DB, tablename string) ([]string, error) {
	query := `SELECT COLUMN_NAME 
		FROM INFORMATION_SCHEMA.COLUMNS 
		WHERE TABLE_NAME = @p1
		ORDER BY ORDINAL_POSITION`

	rows, err := db.Query(query, tablename)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	return columns, nil
}
