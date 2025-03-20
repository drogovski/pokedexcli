package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Buongiorno ",
			expected: []string{"buongiorno"},
		},
		{
			input:    "Welcome to  the jungle  ",
			expected: []string{"welcome", "to", "the", "jungle"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("Actual and expected slices have different lengths. %v != %v", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words are not the same. Actual: %s Expected: %s", word, expectedWord)
			}
		}
	}
}
