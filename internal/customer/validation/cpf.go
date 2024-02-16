package validation

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateCPF(cpf string) (bool, error) {
	sanitizedValue := SanitizeIdentifier(cpf)

	if err := validateLength(11, sanitizedValue); err != nil {
		return false, err
	}
	if err := validateAllDigitsNotEqual(sanitizedValue); err != nil {
		return false, err
	}
	if err := validateFirstVerificationCodeCPF(sanitizedValue); err != nil {
		return false, err
	}
	if err := validateSecondVerificationCode(sanitizedValue); err != nil {
		return false, err
	}

	return true, nil
}

func validateFirstVerificationCodeCPF(cpf string) error {
	cpfDigits := strings.Split(cpf, "")

	firstvalidationCode, err := strconv.Atoi(cpfDigits[9])
	if err != nil {
		return fmt.Errorf("could not parse first validation code, err: %w", err)
	}

	multiplyer := 10
	sum := 0
	for _, digit := range cpfDigits[:9] {
		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			return err
		}
		sum += digitInt * multiplyer
		multiplyer--
	}

	calculation := (sum * 10) % 11
	if calculation == 10 {
		calculation = 0
	}

	if calculation != firstvalidationCode {
		return fmt.Errorf("first validation code is not valid, it should be %d, but got %d", calculation, firstvalidationCode)
	}

	return nil
}

func validateSecondVerificationCode(cpf string) error {
	cpfDigits := strings.Split(cpf, "")
	secondValidatiorCode, err := strconv.Atoi(cpfDigits[10])
	if err != nil {
		return fmt.Errorf("could not parse first validation code, err: %w", err)
	}

	multiplyer := 11
	sum := 0
	for _, digit := range cpfDigits[:10] {
		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			return err
		}
		sum += digitInt * multiplyer
		multiplyer--
	}

	calculation := (sum * 10) % 11
	if calculation == 10 {
		calculation = 0
	}

	if calculation != secondValidatiorCode {
		return fmt.Errorf("second validation code is not valid, it should be %d, but got %d", calculation, secondValidatiorCode)
	}

	return nil
}
