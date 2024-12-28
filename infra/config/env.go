package config

import (
	"log"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func LoadEnv() {
	once.Do(func() {
		_, filename, _, _ := runtime.Caller(0)
		projectRoot := filepath.Join(filepath.Dir(filename), "../../")

		envPath := filepath.Join(projectRoot, ".env")

		err := godotenv.Load(envPath)
		if err != nil {
			log.Printf("No .env file found at %s. Proceeding with system environment variables.\n", envPath)
		}
	})
}
