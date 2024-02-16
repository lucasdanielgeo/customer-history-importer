package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lucasdanielgeo/customer-history-importer/internal/customer"
	"github.com/lucasdanielgeo/customer-history-importer/internal/infra/db"
)

const (
	filePath = "./data/base_teste.txt"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)
	log.SetPrefix(fmt.Sprintf("[%d] [customer-history-importer] ", os.Getpid()))
	godotenv.Load()
}

func main() {
	start := time.Now()
	timeout := 60
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	postgresCustomerHistoryRepository := customer.NewPostgresCustomerHistoryRepository(db)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	importer := customer.CustomerHistoryImporter{
		Repository: postgresCustomerHistoryRepository,
		Reader:     file,
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		importer.Execute()

	}()

	select {
	case <-ctx.Done():
		log.Fatalf("[ERROR] importing process too long, cancelled after %vs", timeout)

	case <-done:
		elapsedTime := time.Since(start)
		log.Printf("[SUCCESS] Importing took: %vs\n", elapsedTime.Seconds())
	}
}
