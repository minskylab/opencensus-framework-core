package client

import (
	"opencensus/core/ent"
	"os"

	"github.com/joho/godotenv"
)

type pgConfig struct {
	hostname string
	port     int
	dbName   string
	user     string
	password string
}

func getPGConfig() pgConfig {
	_ = godotenv.Load() // load envs from .env file (create your own).

	hostname := os.Getenv("PG_HOSTNAME")
	// TODO: add more postgres configuration fields.

	return pgConfig{
		hostname: hostname,
		// TODO: Fully fill it.
	}
}

func NewClient() (*ent.Client, error) {
	// TODO: Implement entgo auto migration and return ready client
	return nil, nil
}
