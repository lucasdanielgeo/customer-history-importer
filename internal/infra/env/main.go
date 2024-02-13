package env

import (
	"fmt"
	"os"
)

func GetPostgresHost() string {
	return os.Getenv("POSTGRES_HOST")
}

func GetPostgresDatabase() string {
	return os.Getenv("POSTGRES_DATABASE")
}

func GetPostgresUser() string {
	return os.Getenv("POSTGRES_USER")
}

func GetPostgresPassword() string {
	return os.Getenv("POSTGRES_PASSWORD")
}

func GetPostgresConnectionString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		GetPostgresUser(),
		GetPostgresPassword(),
		GetPostgresHost(),
		GetPostgresDatabase(),
	)
}
