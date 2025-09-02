package main

import (
	"testing"
)

func TestWordStartsWith(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		word     string
		expected bool
	}{
		{
			name:     "empty prefix matches any word",
			prefix:   "",
			word:     "hello",
			expected: true,
		},
		{
			name:     "prefix matches word start",
			prefix:   "hel",
			word:     "hello",
			expected: true,
		},
		{
			name:     "prefix does not match word start",
			prefix:   "abc",
			word:     "hello",
			expected: false,
		},
		{
			name:     "prefix equals word",
			prefix:   "hello",
			word:     "hello",
			expected: true,
		},
		{
			name:     "prefix longer than word",
			prefix:   "helloworld",
			word:     "hello",
			expected: false,
		},
		{
			name:     "case sensitive match",
			prefix:   "Hel",
			word:     "hello",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordStartsWith(tt.prefix)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordStartsWith(%q)(%q) = %v, want %v", tt.prefix, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordEndsWith(t *testing.T) {
	tests := []struct {
		name     string
		suffix   string
		word     string
		expected bool
	}{
		{
			name:     "empty suffix matches any word",
			suffix:   "",
			word:     "hello",
			expected: true,
		},
		{
			name:     "suffix matches word end",
			suffix:   "llo",
			word:     "hello",
			expected: true,
		},
		{
			name:     "suffix does not match word end",
			suffix:   "abc",
			word:     "hello",
			expected: false,
		},
		{
			name:     "suffix equals word",
			suffix:   "hello",
			word:     "hello",
			expected: true,
		},
		{
			name:     "suffix longer than word",
			suffix:   "worldhello",
			word:     "hello",
			expected: false,
		},
		{
			name:     "case sensitive match",
			suffix:   "LLO",
			word:     "hello",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordEndsWith(tt.suffix)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordEndsWith(%q)(%q) = %v, want %v", tt.suffix, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordStartsAndEndsWith(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		suffix   string
		word     string
		expected bool
	}{
		{
			name:     "empty prefix and suffix match any word",
			prefix:   "",
			suffix:   "",
			word:     "hello",
			expected: true,
		},
		{
			name:     "prefix and suffix match word",
			prefix:   "he",
			suffix:   "lo",
			word:     "hello",
			expected: true,
		},
		{
			name:     "prefix matches but suffix doesn't",
			prefix:   "he",
			suffix:   "xx",
			word:     "hello",
			expected: false,
		},
		{
			name:     "suffix matches but prefix doesn't",
			prefix:   "xx",
			suffix:   "lo",
			word:     "hello",
			expected: false,
		},
		{
			name:     "neither prefix nor suffix match",
			prefix:   "xx",
			suffix:   "yy",
			word:     "hello",
			expected: false,
		},
		{
			name:     "same character for prefix and suffix",
			prefix:   "a",
			suffix:   "a",
			word:     "abba",
			expected: true,
		},
		{
			name:     "prefix and suffix overlap for short word",
			prefix:   "ab",
			suffix:   "bc",
			word:     "abc",
			expected: true,
		},
		{
			name:     "case sensitive matches",
			prefix:   "He",
			suffix:   "lo",
			word:     "hello",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordStartsAndEndsWith(tt.prefix, tt.suffix)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordStartsAndEndsWith(%q, %q)(%q) = %v, want %v",
					tt.prefix, tt.suffix, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordContains(t *testing.T) {
	tests := []struct {
		name       string
		substrings []string
		word       string
		expected   bool
	}{
		{
			name:       "empty substring slice matches any word",
			substrings: []string{},
			word:       "hello",
			expected:   true,
		},
		{
			name:       "single substring that matches",
			substrings: []string{"ell"},
			word:       "hello",
			expected:   true,
		},
		{
			name:       "single substring that doesn't match",
			substrings: []string{"xyz"},
			word:       "hello",
			expected:   false,
		},
		{
			name:       "multiple substrings that all match",
			substrings: []string{"h", "e", "l", "o"},
			word:       "hello",
			expected:   true,
		},
		{
			name:       "multiple substrings with one that doesn't match",
			substrings: []string{"h", "e", "z"},
			word:       "hello",
			expected:   false,
		},
		{
			name:       "case sensitive matching",
			substrings: []string{"H"},
			word:       "hello",
			expected:   false,
		},
		{
			name:       "substring equals word",
			substrings: []string{"hello"},
			word:       "hello",
			expected:   true,
		},
		{
			name:       "substring longer than word",
			substrings: []string{"helloworld"},
			word:       "hello",
			expected:   false,
		},
		{
			name:       "overlapping substrings",
			substrings: []string{"el", "ll"},
			word:       "hello",
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordContains(tt.substrings)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordContains(%v)(%q) = %v, want %v", tt.substrings, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordDoesNotContain(t *testing.T) {
	tests := []struct {
		name       string
		substrings []string
		word       string
		expected   bool
	}{
		{
			name:       "empty substring slice matches any word",
			substrings: []string{},
			word:       "hello",
			expected:   true,
		},
		{
			name:       "single substring that doesn't match",
			substrings: []string{"xyz"},
			word:       "hello",
			expected:   true,
		},
		{
			name:       "single substring that matches",
			substrings: []string{"ell"},
			word:       "hello",
			expected:   false,
		},
		{
			name:       "multiple substrings that all don't match",
			substrings: []string{"x", "y", "z"},
			word:       "hello",
			expected:   true,
		},
		{
			name:       "multiple substrings with one that matches",
			substrings: []string{"x", "y", "e"},
			word:       "hello",
			expected:   false,
		},
		{
			name:       "case sensitive matching",
			substrings: []string{"H"},
			word:       "hello",
			expected:   true,
		},
		{
			name:       "substring equals word",
			substrings: []string{"hello"},
			word:       "hello",
			expected:   false,
		},
		{
			name:       "substring longer than word",
			substrings: []string{"helloworld"},
			word:       "hello",
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordDoesNotContain(tt.substrings)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordDoesNotContain(%v)(%q) = %v, want %v", tt.substrings, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordHasDoubleLetter(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected bool
	}{
		{
			name:     "word with double letter",
			word:     "hello",
			expected: true,
		},
		{
			name:     "word with multiple double letters",
			word:     "bookkeeper",
			expected: true,
		},
		{
			name:     "word without double letter",
			word:     "world",
			expected: false,
		},
		{
			name:     "empty word",
			word:     "",
			expected: false,
		},
		{
			name:     "single letter word",
			word:     "a",
			expected: false,
		},
		{
			name:     "repeated but not adjacent letters",
			word:     "banana",
			expected: false,
		},
		{
			name:     "double letter at beginning",
			word:     "sspeak",
			expected: true,
		},
		{
			name:     "double letter at end",
			word:     "pass",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordHasDoubleLetter()
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordHasDoubleLetter()(%q) = %v, want %v", tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordLengthGreaterThan(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		word     string
		expected bool
	}{
		{
			name:     "word length greater than specified length",
			length:   3,
			word:     "hello",
			expected: true,
		},
		{
			name:     "word length equal to specified length",
			length:   5,
			word:     "hello",
			expected: false,
		},
		{
			name:     "word length less than specified length",
			length:   6,
			word:     "hello",
			expected: false,
		},
		{
			name:     "empty word with zero length",
			length:   0,
			word:     "",
			expected: false,
		},
		{
			name:     "empty word with negative length",
			length:   -1,
			word:     "",
			expected: true,
		},
		{
			name:     "single character word",
			length:   0,
			word:     "a",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordLengthGreaterThan(tt.length)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordLengthGreaterThan(%d)(%q) = %v, want %v", tt.length, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordLengthLessThan(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		word     string
		expected bool
	}{
		{
			name:     "word length less than specified length",
			length:   6,
			word:     "hello",
			expected: true,
		},
		{
			name:     "word length equal to specified length",
			length:   5,
			word:     "hello",
			expected: false,
		},
		{
			name:     "word length greater than specified length",
			length:   4,
			word:     "hello",
			expected: false,
		},
		{
			name:     "empty word with non-zero length",
			length:   1,
			word:     "",
			expected: true,
		},
		{
			name:     "empty word with zero length",
			length:   0,
			word:     "",
			expected: false,
		},
		{
			name:     "single character word",
			length:   2,
			word:     "a",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordLengthLessThan(tt.length)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordLengthLessThan(%d)(%q) = %v, want %v", tt.length, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordLengthEqualsTo(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		word     string
		expected bool
	}{
		{
			name:     "word length equals specified length",
			length:   5,
			word:     "hello",
			expected: true,
		},
		{
			name:     "word length less than specified length",
			length:   6,
			word:     "hello",
			expected: false,
		},
		{
			name:     "word length greater than specified length",
			length:   4,
			word:     "hello",
			expected: false,
		},
		{
			name:     "empty word with zero length",
			length:   0,
			word:     "",
			expected: true,
		},
		{
			name:     "empty word with non-zero length",
			length:   1,
			word:     "",
			expected: false,
		},
		{
			name:     "single character word",
			length:   1,
			word:     "a",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordLengthEqualsTo(tt.length)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordLengthEqualsTo(%d)(%q) = %v, want %v", tt.length, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordLengthBetween(t *testing.T) {
	tests := []struct {
		name     string
		low      int
		high     int
		word     string
		expected bool
	}{
		{
			name:     "word length within range",
			low:      4,
			high:     6,
			word:     "hello",
			expected: true,
		},
		{
			name:     "word length equal to low bound",
			low:      5,
			high:     10,
			word:     "hello",
			expected: true,
		},
		{
			name:     "word length equal to high bound",
			low:      1,
			high:     5,
			word:     "hello",
			expected: true,
		},
		{
			name:     "word length below range",
			low:      6,
			high:     10,
			word:     "hello",
			expected: false,
		},
		{
			name:     "word length above range",
			low:      1,
			high:     4,
			word:     "hello",
			expected: false,
		},
		{
			name:     "empty word with zero low bound",
			low:      0,
			high:     5,
			word:     "",
			expected: true,
		},
		{
			name:     "empty word with non-zero low bound",
			low:      1,
			high:     5,
			word:     "",
			expected: false,
		},
		{
			name:     "equal low and high bounds matching word length",
			low:      5,
			high:     5,
			word:     "hello",
			expected: true,
		},
		{
			name:     "equal low and high bounds not matching word length",
			low:      6,
			high:     6,
			word:     "hello",
			expected: false,
		},
		{
			name:     "negative low bound",
			low:      -1,
			high:     5,
			word:     "hello",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordLengthBetween(tt.low, tt.high)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordLengthBetween(%d, %d)(%q) = %v, want %v", tt.low, tt.high, tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordContainsMoreThanOne(t *testing.T) {
	tests := []struct {
		name     string
		letter   string
		word     string
		expected bool
	}{
		{
			name:     "word contains more than one of the letter",
			letter:   "l",
			word:     "hello",
			expected: true,
		},
		{
			name:     "word contains exactly one of the letter",
			letter:   "h",
			word:     "hello",
			expected: false,
		},
		{
			name:     "word contains none of the letter",
			letter:   "z",
			word:     "hello",
			expected: false,
		},
		{
			name:     "word contains many of the letter",
			letter:   "s",
			word:     "mississippi",
			expected: true,
		},
		{
			name:     "case sensitivity - lowercase letter in word",
			letter:   "E",
			word:     "hello",
			expected: false,
		},
		{
			name:     "case sensitivity - uppercase letter in word",
			letter:   "e",
			word:     "HELLO",
			expected: false,
		},
		{
			name:     "multi-character sequence",
			letter:   "ll",
			word:     "hello",
			expected: false,
		},
		{
			name:     "multi-character sequence appearing multiple times",
			letter:   "ab",
			word:     "abababa",
			expected: true,
		},
		{
			name:     "empty word",
			letter:   "a",
			word:     "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := wordContainsMoreThanOne(tt.letter)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("wordContainsMoreThanOne(%q)(%q) = %v, want %v", tt.letter, tt.word, result, tt.expected)
			}
		})
	}
}

func TestParsePredicate(t *testing.T) {
	tests := []struct {
		name      string
		predicate string
		word      string
		expected  bool
	}{
		// Starts with
		{
			name:      "starts with prefix",
			predicate: "Starts with he",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "starts with prefix - mismatch",
			predicate: "Starts with xy",
			word:      "hello",
			expected:  false,
		},
		// Ends with
		{
			name:      "ends with suffix",
			predicate: "Ends with lo",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "ends with suffix - mismatch",
			predicate: "Ends with xy",
			word:      "hello",
			expected:  false,
		},
		// Contains the letter
		{
			name:      "contains the letter",
			predicate: "Contains the letter e",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "contains the letter - mismatch",
			predicate: "Contains the letter z",
			word:      "hello",
			expected:  false,
		},
		// Contains multiple
		{
			name:      "contains multiple letters",
			predicate: "Contains e, l",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "contains multiple letters - one missing",
			predicate: "Contains e, z",
			word:      "hello",
			expected:  false,
		},
		{
			name:      "contains sequence",
			predicate: "Contains el",
			word:      "hello",
			expected:  true,
		},
		// Does not contain
		{
			name:      "does not contain letters",
			predicate: "Does not contain z, y",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "does not contain letters - one present",
			predicate: "Does not contain z, e",
			word:      "hello",
			expected:  false,
		},
		// Between X and Y letters
		{
			name:      "between length range - within range",
			predicate: "Between 4 and 6 letters",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "between length range - below range",
			predicate: "Between 6 and 10 letters",
			word:      "hello",
			expected:  false,
		},
		{
			name:      "between length range - above range",
			predicate: "Between 2 and 4 letters",
			word:      "hello",
			expected:  false,
		},
		// Multiple letter X's
		{
			name:      "multiple letter",
			predicate: "Multiple letter l's",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "multiple letter - only one occurrence",
			predicate: "Multiple letter h's",
			word:      "hello",
			expected:  false,
		},
		// Multiple X's
		{
			name:      "multiple shorthand",
			predicate: "Multiple l's",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "multiple shorthand - only one occurrence",
			predicate: "Multiple h's",
			word:      "hello",
			expected:  false,
		},
		// Double letter
		{
			name:      "double letter - present",
			predicate: "Double letter",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "double letter - absent",
			predicate: "Double letter",
			word:      "world",
			expected:  false,
		},
		// Starts & ends with
		{
			name:      "starts and ends with same letter",
			predicate: "Starts & ends with a",
			word:      "area",
			expected:  true,
		},
		{
			name:      "starts and ends with same letter - mismatch",
			predicate: "Starts & ends with h",
			word:      "hello",
			expected:  false,
		},
		// X letters or fewer
		{
			name:      "letters or fewer - within limit",
			predicate: "5 letters or fewer",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "letters or fewer - exceeds limit",
			predicate: "4 letters or fewer",
			word:      "hello",
			expected:  false,
		},
		// X letters or more
		{
			name:      "letters or more - meets minimum",
			predicate: "5 letters or more",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "letters or more - below minimum",
			predicate: "6 letters or more",
			word:      "hello",
			expected:  false,
		},
		// X letter word
		{
			name:      "letter word - five",
			predicate: "five letter word",
			word:      "hello",
			expected:  true,
		},
		{
			name:      "letter word - four",
			predicate: "four letter word",
			word:      "hello",
			expected:  false,
		},
		// Infinity
		{
			name:      "infinity predicate",
			predicate: "infinity",
			word:      "anyword",
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := parsePredicate(tt.predicate)
			result := predicate(tt.word)
			if result != tt.expected {
				t.Errorf("parsePredicate(%q)(%q) = %v, want %v", tt.predicate, tt.word, result, tt.expected)
			}
		})
	}
}
