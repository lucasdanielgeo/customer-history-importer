package customer

import (
	"testing"
)

func TestParseNullString(t *testing.T) {
	testCases := []struct {
		name, input string
		expect      bool
	}{
		{"parse NULL string", "NULL", true},
		{"parse empty string", "", true},
		{"parse not NULL or empty string", "Lorem ipsum", false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsedString := ParseNullString(tc.input)

			got := parsedString == nil
			if got != tc.expect {
				t.Errorf("could not parse string to nil or return pointer")
			}
		})
	}
}
func TestParseBool(t *testing.T) {
	testCases := []struct {
		name, input string
		expect      bool
	}{
		{"empty string", "", false},
		{"invalid string", "invalid", false},
		{"NULL string", "NULL", true},
		{"non-empty string", "true", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseBool(tc.input)
			got := err == nil
			if got != tc.expect {
				t.Error(err)
			}
		})
	}
}
func TestParseFloat64(t *testing.T) {
	testCases := []struct {
		name, input string
		expect      bool
	}{
		{"empty string", "", false},
		{"invalid string", "invalid", false},
		{"NULL string", "NULL", true},
		{"non-empty string float with a coma as decimal separator", "2,6", true},
		{"non-empty string float with a dot as decimal separator", "2.6", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseFloat64(tc.input)
			got := err == nil
			if got != tc.expect {
				t.Error(err)
			}
		})
	}
}
