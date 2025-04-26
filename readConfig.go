package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readConfig(fileName string) (*AppConfig, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Не удалось прочитать файл конфигурации: %w", err)
	}

	var config AppConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("Ошибка парсинга JSON: %w", err)
	}
	return &config, nil
}
