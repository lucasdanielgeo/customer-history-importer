package validation

import (
	"fmt"
	"regexp"
	"strings"
)

var nonNumericRegex = regexp.MustCompile(`\D`)

func removeStringSpaces(value string) string {
	return strings.ReplaceAll(value, " ", "")
}

func removeNonNumericCharacters(value string) string {
	return nonNumericRegex.ReplaceAllString(value, "")
}

func sanitizeIdentifier(identifier string) string {
	identifier = removeStringSpaces(identifier)
	identifier = removeNonNumericCharacters(identifier)

	return identifier
}

func validateLength(validLenght int, identifier string) error {
	if len(identifier) != validLenght {
		return fmt.Errorf("CPF should have 11 characters, but got %d", len(identifier))
	}

	return nil
}

func validateAllDigitsNotEqual(identifier string) error {
	for i := 1; i < len(identifier); i++ {
		if identifier[i] != identifier[i-1] {
			return nil
		}
	}
	return fmt.Errorf("invallid identifier, all digits are equal: %v", identifier)
}
