package customer

import (
	"bufio"
	"fmt"
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
		panic(err)
	}

	log.Println("[INFO] Saving on DB")
	if err := customerService.SaveOnDB(customers); err != nil {
		panic(err)
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println("error reading the file:", err)
	}
}
