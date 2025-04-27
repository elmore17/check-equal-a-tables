package main

import (
	"context"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
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

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Не удалось подключиться к серверу:", err)
	}

	fmt.Println("\nДанные из таблицы:", config.Params.FHD)
	date_FHD, err := getTableColumns(db, config.Params.FHD)
	if err != nil {
		log.Printf("Ошибка при выводе данных из таблицы %s: %v", config.Params.FHD, err)
	}
	fmt.Println(date_FHD)

	fmt.Println("\nДанные из таблицы:", config.Params.KHD)
	date_KHD, err := getTableColumns(db, config.Params.KHD)
	if err != nil {
		log.Printf("Ошибка при выводе данных из таблицы %s: %v", config.Params.KHD, err)
	}
	fmt.Println(date_KHD)

	fmt.Println("\nКоличество данных из таблицы:", config.Params.FHD)
	count_FHD, err := countElem(ctx, db, config.Params.FHD)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса: %v", err)
	}

	fmt.Println("\nКоличество данных из таблицы:", config.Params.KHD)
	count_KHD, err := countElem(ctx, db, config.Params.KHD)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса: %v", err)
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	if err := f.DeleteSheet("Sheet1"); err != nil {
		log.Fatal(err)
	}
	columnName := []string{"ФХД", "КХД"}
	createSheetXlsx(f, "Кол-во записей", columnName, strconv.Itoa(count_FHD), strconv.Itoa(count_KHD))

	if err := f.SaveAs("NameFile.xlsx"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Отчет успешно создан: NameFile.xlsx")
}
