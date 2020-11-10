package passwords

import (
	"day-4/internal/passwords"
	"testing"
)

var testCases = []struct {
	password string
	valid    bool
}{
	{"122345", true},
	{"112233", true},
	{"111144", true},
	{"223450", false},
	{"123789", false},
	{"122234", false},
}

func TestValidPasswords(t *testing.T) {
	for _, test := range testCases {
		if passwords.IsValid(test.password) != test.valid {
			t.Fatalf("password %v, expected %t", test.password, test.valid)
		}
	}
}
