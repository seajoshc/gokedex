package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "hello, world",
			expected: []string{
				"hello,",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("lengths are not equal: actual %v and expected %v", len(actual), len(cs.expected))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]

			if actualWord != expectedWord {
				t.Errorf("words do not match: actual %v and expected %v", actualWord, expectedWord)
			}
		}
	}
}
