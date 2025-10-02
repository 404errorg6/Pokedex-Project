package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " a b  c   d     e      f",
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			input:    " HELLo wORlD",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected Length: %d, Actual Length: %d\n", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected: %v, Actual: %v", expectedWord, word)
			}
		}
	}
}
