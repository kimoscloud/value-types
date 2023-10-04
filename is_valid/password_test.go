package is_valid

import (
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{"Password1!", true},
		{"password1!", false},                         // Missing uppercase
		{"PASSWORD1!", false},                         // Missing lowercase
		{"Password!", false},                          // Missing digit
		{"Password1", false},                          // Missing special character
		{"Passw1!", false},                            // Too short
		{"PasswordPasswordPasswordPassword1!", false}, // Too long
		{"", false},                                   // Empty
	}

	for _, test := range tests {
		if IsValidPassword(test.password) != test.valid {
			t.Errorf("Expected IsValidPassword(%q) to be %v", test.password, test.valid)
		}
	}
}

func TestIsSpecialCharacter(t *testing.T) {
	tests := []struct {
		char  rune
		valid bool
	}{
		{'!', true},
		{'a', false},
		{'A', false},
		{'1', false},
		{'{', true},
	}

	for _, test := range tests {
		if isSpecialCharacter(test.char) != test.valid {
			t.Errorf("Expected isSpecialCharacter(%q) to be %v", test.char, test.valid)
		}
	}
}
