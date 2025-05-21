package utils

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hieu2304/order-food-be/config"
	"github.com/joho/godotenv"
)

func SetupTest(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	envPath := filepath.Join(basepath, "..", ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		t.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}
	_, err = config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
}
