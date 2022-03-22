package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	_, b, _, _ = runtime.Caller(0)

	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

/**
This function is to create a configuration
file with the file name .env
*/
func Config(key string) string {
	// Load .env file
	err := godotenv.Load(".env")

	// Check if failed to load file
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	// Take env file parameters
	return os.Getenv(key)
}
