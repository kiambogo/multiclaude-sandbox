# multiclaude-sandbox

A Go string utility library used as a sandbox for experimenting with [multiclaude](https://github.com/dlorenc/multiclaude).

## Package

`stringutil` — simple, pure string manipulation functions with no external dependencies.

### Existing functions

- `Reverse(s string) string` — reverses a string
- `IsPalindrome(s string) bool` — checks if a string reads the same forward and backward (ignores case and non-letters)
- `WordCount(s string) map[string]int` — returns a word frequency map

## Running tests

```bash
go test ./...
```

## Suggested agent tasks

Good candidates for parallel agent work — each is independent and fully testable:

- Add `Truncate(s string, maxLen int, suffix string) string` — truncate to maxLen chars, appending suffix (e.g. "...") if truncated
- Add `Slugify(s string) string` — convert to lowercase, replace spaces with hyphens, strip non-alphanumeric characters
- Add `CamelToSnake(s string) string` — convert camelCase or PascalCase to snake_case
- Add `SnakeToCamel(s string) string` — convert snake_case to camelCase
- Add `CountVowels(s string) int` — count vowel characters (a, e, i, o, u, case-insensitive)
- Add `TitleCase(s string) string` — capitalize the first letter of each word
- Add `RemoveDuplicateWords(s string) string` — remove duplicate words, preserving order and case

Each task should include the function and a `_test.go` file with table-driven tests.
