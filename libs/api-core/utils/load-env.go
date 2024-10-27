package utils

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file.
//
// Parameters:
// - level: int, the number of levels of directories to search for the .env file.
// No return type.
func LoadEnv(level int) Env {
	envs := findEnvFiles(level)

	if len(envs) == 0 {
		log.Fatal("Could not resolve .env file")
	}

	err := godotenv.Overload(envs...)
	if err != nil {
		log.Println("Warning: error loading .env file", err)
	}

	return List()
}

// findEnvFiles finds .env files recursively up to the specified level.
func findEnvFiles(level int) []string {
	var envs []string

	// Iterate through directories up to the specified level
	for i := level; i >= 0; i-- {
		// Construct the path to the .env file
		envPath := filepath.Join(repeat("../", i), ".env")
		// Check if the .env file exists
		if _, err := os.Stat(envPath); err == nil {
			envs = append(envs, envPath)
		}
	}

	return envs
}

// repeat repeats a string 's' 'count' times.
func repeat(s string, count int) string {
	returnString := ""
	for i := 0; i < count; i++ {
		returnString += s
	}
	return returnString
}

type Env struct {
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_DRIVER   string
	APP_NAME    string
	APP_PORT    int
	APP_ENV     string
}

func List() Env {
	driver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	appName := os.Getenv("APP_NAME")
	appPort := os.Getenv("APP_PORT")
	appEnv := os.Getenv("APP_ENV")

	requiredEnvVars := map[string]string{
		"DB_NAME":     dbName,
		"DB_USER":     dbUser,
		"DB_PASSWORD": dbPassword,
		"DB_HOST":     dbHost,
		"DB_PORT":     dbPort,
		"DB_DRIVER":   driver,
		"APP_NAME":    appName,
		"APP_PORT":    appPort,
		"APP_ENV":     appEnv,
	}

	for envVar, value := range requiredEnvVars {
		if len(value) == 0 {
			panic("ENV " + envVar + " IS NOT SET")
		}
	}

	dbPortInt, _ := strconv.Atoi(dbPort)
	appPortInt, _ := strconv.Atoi(appPort)

	return Env{
		DB_NAME:     dbName,
		DB_USER:     dbUser,
		DB_PASSWORD: dbPassword,
		DB_HOST:     dbHost,
		DB_PORT:     dbPortInt,
		DB_DRIVER:   driver,
		APP_NAME:    appName,
		APP_PORT:    appPortInt,
		APP_ENV:     appEnv,
	}
}
