package utilities

import "testing"

var randomCases = []struct {
	input int
	err   error
}{
	{-1, errRandomZero},
	{0, errRandomZero},
	{1, nil},
}

// Quite hard to test the randoms. But lets test the errors at least
func TestRandomInt(t *testing.T) {
	for _, test := range randomCases {
		_, err := RandomInt(test.input)
		if err != nil {
			if err != test.err {
				t.Errorf("Input : %d. Expected error: %s. Observed: %s", test.input, test.err, err)
			}
		}
		if test.input >= 1 && err != nil {
			t.Errorf("Input : %d. Expected error to be nil. Observed: %s", test.input, err)
		}
	}
}
