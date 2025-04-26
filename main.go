package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func main() {
	config, err := readConfig("config_db.json")
	if err != nil {
		log.Fatal("Ошибка чтения конфигурации: %v", err)
	}

	db, err := connectToDB(config.Config)
	if err != nil {
		log.Fatal("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	fmt.Println("\nДанные из таблицы:", config.Params.FirstTable)
	date_first_table, err := getTableColumns(db, config.Params.FirstTable)
	if err != nil {
		log.Printf("Ошибка при выводе данных из таблицы %s: %v", config.Params.FirstTable, err)
	}
	fmt.Println(date_first_table)

	fmt.Println("\nДанные из таблицы:", config.Params.SecondTable)
	date_second_table, err := getTableColumns(db, config.Params.SecondTable)
	if err != nil {
		log.Printf("Ошибка при выводе данных из таблицы %s: %v", config.Params.SecondTable, err)
	}
	fmt.Println(date_second_table)
}
