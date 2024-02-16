package validation

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateCNPJ(cnpj *string) (bool, error) {
	if cnpj == nil {
		return false, nil
	}

	sanitizeIdentifier := SanitizeIdentifier(*cnpj)

	if err := validateLength(14, sanitizeIdentifier); err != nil {
		return false, err
	}
	if err := validateAllDigitsNotEqual(sanitizeIdentifier); err != nil {
		return false, err
	}
	if err := validateFirstVerificationCodeCNPJ(sanitizeIdentifier); err != nil {
		return false, err
	}
	if err := validateSecondVerificationCodeCNPJ(sanitizeIdentifier); err != nil {
		return false, err
	}

	return true, nil
}

func validateFirstVerificationCodeCNPJ(cnpj string) error {
	cnpjDigits := strings.Split(cnpj, "")

	firstvalidationCode, err := strconv.Atoi(cnpjDigits[12])
	if err != nil {
		return fmt.Errorf("could not parse first validation code, err: %w", err)
	}
	multiplyer := 5
	sum := 0
	for _, digit := range cnpjDigits[:12] {
		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			return err
		}
		sum += digitInt * multiplyer
		multiplyer--
		if multiplyer == 1 {
			multiplyer = 9
		}
	}

	calculation := sum % 11
	if calculation < 2 {
		calculation = 0
	}
	if calculation >= 2 {
		calculation = 11 - calculation
	}

	if calculation != firstvalidationCode {
		return fmt.Errorf("first validation code is not valid, it should be %d, but got %d", calculation, firstvalidationCode)
	}

	return nil
}

func validateSecondVerificationCodeCNPJ(cnpj string) error {
	cnpjDigits := strings.Split(cnpj, "")
	secondValidatiorCode, err := strconv.Atoi(cnpjDigits[13])
	if err != nil {
		return fmt.Errorf("could not parse second validation code, err: %w", err)
	}

	multiplyer := 6
	sum := 0
	for _, digit := range cnpjDigits[:13] {
		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			return err
		}
		sum += digitInt * multiplyer
		multiplyer--
		if multiplyer == 1 {
			multiplyer = 9
		}
	}

	calculation := sum % 11
	if calculation < 2 {
		calculation = 0
	}
	if calculation >= 2 {
		calculation = 11 - calculation
	}

	if calculation != secondValidatiorCode {
		return fmt.Errorf("second validation code is not valid, it should be %d, but got %d", calculation, secondValidatiorCode)
	}

	return nil
}
