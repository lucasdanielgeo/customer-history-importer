package validation

import "testing"

func TestValidateCPF(t *testing.T) {
	testCases := []struct {
		name, input string
		expected    bool
	}{
		{"valid CPF", "75220186094", true},
		{"valid formated CPF", "752.201.860-94", true},
		{"valid CPF with extra characters", "123.456.789--@00", false},
		{"valid CPF with random digits", "75220186094", true},
		{"invalid CPF with wrong both validation code", "443.496.800-88", false},
		{"invalid CPF with equal numbers", "00000000000", false},
		{"invalid empty input", "", false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validationErr := ValidateCPF(tc.input)
			got := validationErr == nil
			if got != tc.expected {
				t.Errorf("%v for input %s", validationErr, tc.input)
			}
		})
	}
}

func TestValidateFirstVerificationCode(t *testing.T) {
	testCases := []struct {
		name, input string
		expected    bool
	}{
		{"valid first verication code random numbers", "44349680098", true},
		{"valid first verication code sequence of 9 numbers", "12345678909", true},
		{"invalid first verication code", "44349680008", false},
		{"invalid input string with non-numeric characters", "abcdefghijkl", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validationErr := validateFirstVerificationCodeCPF(tc.input)
			got := validationErr == nil
			if got != tc.expected {
				t.Errorf("%v for input %s", validationErr, tc.input)
			}
		})
	}
}

func TestValidateSecondVerificationCode(t *testing.T) {
	testCases := []struct {
		name, input string
		expected    bool
	}{
		{"valid second verication code diferent then 0", "44349680098", true},
		{"valid second verication code diferent equals 0", "40312508000140", true},
		{"invalid second verication code", "44349680008", false},
		{"invalid input string with non-numeric characters", "abcdefghijkl", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validationErr := validateSecondVerificationCode(tc.input)
			got := validationErr == nil
			if got != tc.expected {
				t.Error(validationErr)
			}
		})
	}
}
