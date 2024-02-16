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

func SanitizeIdentifier(identifier string) string {
	identifier = removeStringSpaces(identifier)
	identifier = removeNonNumericCharacters(identifier)

	return identifier
}

func SanitizeNullableIdentifier(nullableIdentifier *string) *string {
	identifier := *nullableIdentifier
	identifier = removeStringSpaces(identifier)
	identifier = removeNonNumericCharacters(identifier)

	return &identifier
}

func validateLength(validLength int, identifier string) error {
	if len(identifier) != validLength {
		return fmt.Errorf("identifier should have %d characters, but got %d", validLength, len(identifier))
	}

	return nil
}

func validateAllDigitsNotEqual(identifier string) error {
	for i := 1; i < len(identifier); i++ {
		if identifier[i] != identifier[i-1] {
			return nil
		}
	}
	return fmt.Errorf("invalid identifier, all digits are equal: %v", identifier)
}
