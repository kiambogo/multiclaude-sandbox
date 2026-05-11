// Package stringutil provides simple string utility functions.
package stringutil

import (
	"strings"
	"unicode"
)

// Slugify converts s to a URL-friendly slug: lowercase, with runs of
// non-alphanumeric characters collapsed to a single hyphen, and leading/
// trailing hyphens trimmed.
func Slugify(s string) string {
	var b strings.Builder
	inSep := false
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			inSep = false
		} else if !inSep && b.Len() > 0 {
			b.WriteByte('-')
			inSep = true
		}
	}
	result := b.String()
	return strings.TrimRight(result, "-")
}

// Reverse returns the string s with its characters in reverse order.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome reports whether s reads the same forward and backward,
// ignoring case and non-letter characters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) {
			letters = append(letters, r)
		}
	}
	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		if letters[i] != letters[j] {
			return false
		}
	}
	return true
}

// WordCount returns a map of each unique word in s to the number of times it appears.
// Words are compared case-insensitively.
func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(s) {
		word = strings.ToLower(strings.Trim(word, ".,!?;:\"'"))
		if word != "" {
			counts[word]++
		}
	}
	return counts
}
