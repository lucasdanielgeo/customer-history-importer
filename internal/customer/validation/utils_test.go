package validation

import (
	"testing"
)

const (
	defaultErrorExpectErrorMessage = "got error: %v, expected error: %v"
	defaultErrorMessage            = "got: %v, expected: %v"
)

func TestValidateLength(t *testing.T) {
	cpfValidLength := 11
	cnpjValidLength := 14
	testCases := []struct {
		name   string
		cpf    string
		want   int
		expect bool
	}{
		{"more than 11 characters", "104936409080", cpfValidLength, false},
		{"exactly 11 characters", "10493640908", cpfValidLength, true},
		{"less than 11 characters", "1049364090", cpfValidLength, false},
		{"more than 14 characters", "793794910008506", cnpjValidLength, false},
		{"exactly 14 characters", "79379491000850", cnpjValidLength, true},
		{"less than 14 characters", "7937949100085", cnpjValidLength, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := validateLength(tc.want, tc.cpf)
			if (got == nil) != tc.expect {
				t.Errorf(defaultErrorExpectErrorMessage, got, tc.expect)
			}
		})
	}
}

func TestRemoveStringSpaces(t *testing.T) {
	testCases := []struct {
		name, input, expected string
	}{
		{"String with spaces", "1 2 3", "123"},
		{"Empty input", "", ""},
		{"Alphanumeric string with spaces", "a b c", "abc"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := removeStringSpaces(tc.input)
			if got != tc.expected {
				t.Errorf(defaultErrorMessage, got, tc.expected)
			}
		})
	}

}

func TestRemoveNonNumericCharacters(t *testing.T) {
	testCases := []struct {
		name, input, expected string
	}{
		{"Alphanumeric string with numbers", "12a34b56", "123456"},
		{"Alphanumeric string", "abc123def", "123"},
		{"String with special characters", "12#45@78", "124578"},
		{"String with spaces", "1 2 3 4 5", "12345"},
		{"String with Unicode characters", "٤٥٦", ""},
		{"Chinese numbers", "一二三四五六", ""},
		{"Japanese numbers", "一二三四五六", ""},
		{"Hindi numbers", "१२३४५६", ""},
		{"Thai numbers", "๑๒๓๔๕๖", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := removeNonNumericCharacters(tc.input)
			if got != tc.expected {
				t.Errorf(defaultErrorMessage, got, tc.expected)
			}

		})
	}
}

func TestSanitizeIdentifier(t *testing.T) {
	testeCases := []struct {
		name, input, expected string
	}{
		{"Sanitized CPF with dots and hyphen", "123.456.789-00", "12345678900"},
		{"CPF with alphanumeric characters", "abc123.456.789-00def", "12345678900"},
		{"Sanitized CNPJ with dots, hyphen and slash", "12.345.678/0001-00", "12345678000100"},
		{"CNPJ with alphanumeric characters", "abc12.345.678/0001-00def", "12345678000100"},
	}

	for _, tc := range testeCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SanitizeIdentifier(tc.input)
			if got != tc.expected {
				t.Errorf(defaultErrorMessage, got, tc.expected)
			}
		})
	}
}

func TestValidateAllDigitsNotEqual(t *testing.T) {
	testCases := []struct {
		name, input string
		expected    bool
	}{
		{"all digits not equal", "12345678900", true},
		{"invalid CPF with all digits equal", "00000000000", false},
		{"invalid CPF with all digits equal", "11111111111", false},
		{"invalid CPF with all digits equal", "99999999999", false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validationErr := validateAllDigitsNotEqual(tc.input)
			got := validationErr == nil
			if got != tc.expected {
				t.Error(validationErr)
			}
		})
	}
}
