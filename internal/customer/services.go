package customer

import (
	"bufio"
	"fmt"
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
	ignoreFileHeader(s.scanner)
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
	cpf := fields[0]

	IsValidCPF := true
	if err := validation.ValidateCPF(cpf); err != nil {
		IsValidCPF = false
	}

	lastPurchaseDate := ParseNullString(fields[3])

	mostFrequentStore := ParseNullString(fields[6])
	isMostFrequentStoreValid := true
	if mostFrequentStore == nil || validation.ValidateCNPJ(*mostFrequentStore) != nil {
		isMostFrequentStoreValid = false

	}

	lastPurchaseStore := ParseNullString(fields[7])
	islastPurchaseStoreValid := true
	if lastPurchaseStore == nil || validation.ValidateCNPJ(*lastPurchaseStore) != nil {
		islastPurchaseStoreValid = false
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
		lastPurchaseStore, isMostFrequentStoreValid, islastPurchaseStoreValid,
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
