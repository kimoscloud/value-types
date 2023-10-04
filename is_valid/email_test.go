package is_valid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"firstname.lastname@example.com", true},
		{"email@subdomain.example.com", true},
		{"firstname+lastname@example.com", true},
		{"email@123.123.123.123", false},
		{"email@[123.123.123.123]", false},
		{"\"email\"@example.com", false},

		// Invalid emails
		{"plainaddress", false},
		{"@missingusername.com", false},
		{"email@.com", false},
		{"email@123.123.123.123.123", false},
		{"email@[123.123.123.123].123", false},
		{"_@example.com", false},
		{"example.com", false},
		{"A@b@c@example.com", false},
	}

	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			assert.Equal(t, test.expected, IsValidEmail(test.email))
		})
	}
}
