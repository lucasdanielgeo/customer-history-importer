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
	"github.com/lucasdanielgeo/customer-history-importer/internal/infra/localfs"
)

const (
	file = "./data/base_teste.txt"
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

	done := make(chan struct{})
	go func() {
		defer close(done)
		Run()
	}()

	select {
	case <-ctx.Done():
		log.Fatalf("[ERROR] importing process too long, cancelled after %vs", timeout)

	case <-done:
		elapsedTime := time.Since(start)
		log.Printf("[SUCCESS] Importing took: %vs\n", elapsedTime.Seconds())
	}
}

func Run() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("[INFO] Connected to PostgreSQL database")

	postgresCustomerHistoryRepository := customer.NewPostgresCustomerHistoryRepository(db)

	scanner, file, err := localfs.NewFileScanner(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	customerService := customer.NewCustomerService(postgresCustomerHistoryRepository, scanner)

	log.Println("[INFO] Reading file")
	customers, err := customerService.ReadLines()
	if err != nil {
		panic(err)
	}

	log.Println("[INFO] Saving on DB")
	if err := customerService.SaveOnDB(customers); err != nil {
		panic(err)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading the file:", err)
	}
}
