package customer

import (
	"bufio"
	"io"
	"log"
)

type CustomerHistoryImporter struct {
	Repository CustomerHistoryRepository
	Reader     io.Reader
}

func (c CustomerHistoryImporter) Execute() {
	log.Println("[INFO] Connected to PostgreSQL database")

	fileScanner := bufio.NewScanner(c.Reader)

	customerService := NewCustomerService(c.Repository, fileScanner)

	log.Println("[INFO] Reading file")
	customers, err := customerService.ReadLines()
	if err != nil {
		log.Fatalf("[ERROR] Could not read file content: %v", err)
	}

	log.Println("[INFO] Saving on DB")
	if err := customerService.SaveOnDB(customers); err != nil {
		log.Fatalf("[ERROR] Could not persist data on db: %v", err)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("[INFO] error reading the file: %v", err)
	}
}
