package validation

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateCNPJ(cnpj string) error {
	sanitizedValue := sanitizeIdentifier(cnpj)
	if err := validateLength(14, sanitizedValue); err != nil {
		return err
	}
	if err := validateAllDigitsNotEqual(sanitizedValue); err != nil {
		return err
	}
	if err := validateFirstVerificationCodeCNPJ(sanitizedValue); err != nil {
		return err
	}
	if err := validateSecondVerificationCodeCNPJ(sanitizedValue); err != nil {
		return err
	}

	return nil
}

func validateFirstVerificationCodeCNPJ(cnpj string) error {
	cnpjDigits := strings.Split(cnpj, "")
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

	firstvalidationCode, err := strconv.Atoi(cnpjDigits[12])
	if err != nil {
		return fmt.Errorf("could not parse first validation code, err: %w", err)
	}

	if calculation != firstvalidationCode {
		return fmt.Errorf("first validation code is not valid, it should be %d, but got %d", calculation, firstvalidationCode)
	}

	return nil
}

func validateSecondVerificationCodeCNPJ(cnpj string) error {
	cnpjDigits := strings.Split(cnpj, "")

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

	secondValidatiorCode, err := strconv.Atoi(cnpjDigits[13])

	if err != nil {
		return fmt.Errorf("could not parse second validation code, err: %w", err)
	}

	if calculation != secondValidatiorCode {
		return fmt.Errorf("second validation code is not valid, it should be %d, but got %d", calculation, secondValidatiorCode)
	}

	return nil
}
