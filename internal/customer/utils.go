package customer

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseNullString(value string) *string {
	if value == "NULL" || value == "" {
		return nil
	}
	return &value
}

func ParseBool(value string) (*bool, error) {
	if value == "NULL" {
		return nil, nil
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return nil, err
	}

	return &boolValue, nil
}

func ParseFloat64(value string) (*float64, error) {
	if value == "NULL" {
		return nil, nil
	}
	value = strings.ReplaceAll(value, ",", ".")
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}

	return &floatValue, nil
}

func ValidateSliceLength[T any](validLength int, fields []T) error {
	if len(fields) != validLength {
		return fmt.Errorf("unexpected length: got %d, expected 8", len(fields))
	}
	return nil
}
