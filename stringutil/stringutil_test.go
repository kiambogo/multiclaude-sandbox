package stringutil_test

import (
	"testing"

	"github.com/kiambogo/multiclaude-sandbox/stringutil"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"racecar", "racecar"},
		{"Go!", "!oG"},
	}
	for _, tt := range tests {
		got := stringutil.Reverse(tt.input)
		if got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"A man a plan a canal Panama", true},
		{"hello", false},
		{"", true},
		{"Was it a car or a cat I saw", true},
		{"not a palindrome", false},
	}
	for _, tt := range tests {
		got := stringutil.IsPalindrome(tt.input)
		if got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		input string
		want  map[string]int
	}{
		{
			"the cat sat on the mat",
			map[string]int{"the": 2, "cat": 1, "sat": 1, "on": 1, "mat": 1},
		},
		{
			"Hello, hello! HELLO.",
			map[string]int{"hello": 3},
		},
		{
			"",
			map[string]int{},
		},
	}
	for _, tt := range tests {
		got := stringutil.WordCount(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("WordCount(%q) returned %d entries, want %d", tt.input, len(got), len(tt.want))
			continue
		}
		for word, count := range tt.want {
			if got[word] != count {
				t.Errorf("WordCount(%q)[%q] = %d, want %d", tt.input, word, got[word], count)
			}
		}
	}
}
