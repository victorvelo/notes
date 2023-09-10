package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/victorvelo/notes/internal/models"
)

func getFilePath() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	path := os.Getenv("path")
	return path, nil
}

func InitConfig() (*models.Config, error) {
	filePath, err := getFilePath()
	if err != nil {
		return nil, fmt.Errorf("couldn't load path : %w", err)
	}
	configuration, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("couldn't load configuration file : %w", err)
	}

	var config models.Config
	if err := json.Unmarshal(configuration, &config); err != nil {
		return nil, fmt.Errorf("could't unmarshal configuration file : %w", err)
	}

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error reading env: %w", err)
	}
	config.Postgres.Password = os.Getenv("DB_PASSWORD")

	return &config, nil
}
