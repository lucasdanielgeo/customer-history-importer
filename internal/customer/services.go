package customer

import (
	"bufio"
	"fmt"
	"log"
	"regexp"

	"github.com/lucasdanielgeo/customer-history-importer/internal/customer/validation"
)

type CustomerHistoryService interface {
	SaveOnDB([]CustomerHistory) error
	ReadLines() ([]CustomerHistory, error)
}

type customerHistoryServiceImpl struct {
	db      CustomerHistoryRepository
	scanner *bufio.Scanner
}

func NewCustomerService(db CustomerHistoryRepository, scanner *bufio.Scanner) CustomerHistoryService {
	return &customerHistoryServiceImpl{
		db:      db,
		scanner: scanner,
	}
}

func (s customerHistoryServiceImpl) SaveOnDB(customers []CustomerHistory) error {
	if err := s.db.SaveBatch(customers); err != nil {
		return fmt.Errorf("could not persist customers on DB. Error: %w", err)
	}

	return nil
}

func (s *customerHistoryServiceImpl) ReadLines() ([]CustomerHistory, error) {
	s.ignoreFileHeader()
	var customers []CustomerHistory
	for s.scanner.Scan() {
		line := s.scanner.Text()
		fields, err := s.parseFileLine(line)
		if err != nil {
			return []CustomerHistory{}, err
		}

		customerHistory := s.parseCustomerFromFileFields(fields)

		if err != nil {
			return []CustomerHistory{}, err
		}

		customers = append(customers, customerHistory)
	}

	return customers, nil
}

func (s *customerHistoryServiceImpl) parseCustomerFromFileFields(fields []string) CustomerHistory {
	cpf := validation.SanitizeIdentifier(fields[0])

	IsValidCPF, err := validation.ValidateCPF(cpf)
	if err != nil {
		log.Println(err)
	}

	lastPurchaseDate := ParseNullString(fields[3])

	mostFrequentStore := ParseNullString(fields[6])
	if mostFrequentStore != nil {
		mostFrequentStore = validation.SanitizeNullableIdentifier(mostFrequentStore)
	}

	isMostFrequentStoreValid, err := validation.ValidateCNPJ(mostFrequentStore)
	if err != nil {
		log.Println(err)
	}

	lastPurchaseStore := ParseNullString(fields[7])

	isLastPurchaseStoreValid, err := validation.ValidateCNPJ(lastPurchaseStore)
	if err != nil {
		log.Println(err)
	}
	if lastPurchaseStore != nil {
		lastPurchaseStore = validation.SanitizeNullableIdentifier(lastPurchaseStore)
	}

	private, err := ParseBool(fields[1])
	if err != nil {
		private = nil
	}

	incomplete, err := ParseBool(fields[2])
	if err != nil {
		incomplete = nil
	}

	averageTicket, err := ParseFloat64(fields[4])
	if err != nil {
		averageTicket = nil
	}

	lastPurchaseTicket, err := ParseFloat64(fields[5])
	if err != nil {
		lastPurchaseTicket = nil
	}

	customer := NewCustomerHistory(
		cpf, IsValidCPF, private, incomplete, lastPurchaseDate,
		averageTicket, lastPurchaseTicket, mostFrequentStore,
		lastPurchaseStore, isMostFrequentStoreValid, isLastPurchaseStoreValid,
	)

	return customer
}

var splitExpression = regexp.MustCompile(`\s+|\s+\s`)

func (s customerHistoryServiceImpl) parseFileLine(line string) ([]string, error) {
	fields := splitExpression.Split(line, -1)

	if err := ValidateSliceLength(8, fields); err != nil {
		return nil, err
	}

	return fields, nil
}

func (s customerHistoryServiceImpl) ignoreFileHeader() {
	s.scanner.Scan()
}
