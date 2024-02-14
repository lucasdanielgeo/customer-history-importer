package validation

import "testing"

func TestValidateCNPJ(t *testing.T) {
	testCases := []struct {
		name, input string
		expected    bool
	}{
		{"valid CNPJ", "64252700000156", true},
		{"valid formated CNPJ", "64.252.700/0001-56", true},
		{"valid CNPJ with sequence numbers with dots and slashes", "12.345.678/0001-95", true},
		{"invalid CNPJ with wrong first validation code", "65.374.415/0001-10", false},
		{"invalid CNPJ with wrong second validation code", "65.374.415/0001-71", false},
		{"invalid CNPJ with all nines", "99999999999999", false},
		{"invalid random input with 14 characters", "Loremipsum12345", false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validationErr := ValidateCNPJ(tc.input)
			got := validationErr == nil
			if got != tc.expected {
				t.Errorf("%v for input %s", validationErr, tc.input)
			}
		})
	}
}

func TestValidateFirstVerificationCodeCNPJ(t *testing.T) {
	testCases := []struct {
		name, input string
		expected    bool
	}{
		{"valid first verification code random numbers", "12345678900188", true},
		{"valid first with 0 as first validation code", "59118609000102", true},
		{"invalid input string with non-numeric characters", "abcdefghijkl", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validationErr := validateFirstVerificationCodeCNPJ(tc.input)
			got := validationErr == nil
			if got != tc.expected {
				t.Errorf("%v for input %s", validationErr, tc.input)
			}
		})
	}
}

func TestValidateSecondVerificationCodeCNPJ(t *testing.T) {
	testCases := []struct {
		name, input string
		expected    bool
	}{
		{"valid second verification code sequence of 9 numbers", "12345678900188", true},
		{"invalid second verification code", "44349680000108", false},
		{"invalid input string with non-numeric characters", "abcdefghijklj", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validationErr := validateSecondVerificationCodeCNPJ(tc.input)
			got := validationErr == nil
			if got != tc.expected {
				t.Errorf("%v for input %s", validationErr, tc.input)
			}
		})
	}
}
