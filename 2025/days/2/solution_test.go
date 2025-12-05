package main

import "testing"

func TestContainsRepeats(t *testing.T) {
	tests := []struct {
		testName       string
		input          string
		expectedOutput bool
	}{
		{
			"same digit - odd length",
			"11111",
			true,
		},
		{
			"same digit - even length",
			"1111",
			true,
		},
		{
			"two digit seq - positive case",
			"121212",
			true,
		},
		{
			"two digit seq - negative case",
			"1212121",
			false,
		},
		{
			"three digit seq - positive case",
			"123123",
			true,
		},
		{
			"three digit seq - negative case",
			"12312345",
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			acutalOuput := containsRepeats(test.input)
			if test.expectedOutput != acutalOuput {
				t.Errorf("expected %v but got %v\n", test.expectedOutput, acutalOuput)
			}
		})
	}
}
