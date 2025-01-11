package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  SqUirTle CHARMANDER bulbasaur",
			expected: []string{"squirtle", "charmander", "bulbasaur"},
		},
		{
			input:    "             lots        of        white            space             ",
			expected: []string{"lots", "of", "white", "space"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("num words did not match. actual: %d, expected: %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words did not match. actual: %s, expected: %s", word, expectedWord)
			}
		}
	}
}
