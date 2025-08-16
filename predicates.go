package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Predicate struct {
	Name string
	Func func(word string) bool
}

// Returns a function that checks if a word starts with the given prefix.
func wordStartsWith(prefix string) func(word string) bool {
	return func(word string) bool {
		return strings.HasPrefix(word, prefix)
	}
}

// Returns a function that checks if a word ends with the given suffix.
func wordEndsWith(suffix string) func(word string) bool {
	return func(word string) bool {
		return strings.HasSuffix(word, suffix)
	}
}

// Returns a function that checks if a word starts with the given prefix and ends with the given suffix.
func wordStartsAndEndsWith(prefix, suffix string) func(word string) bool {
	return func(word string) bool {
		return strings.HasPrefix(word, prefix) && strings.HasSuffix(word, suffix)
	}
}

// Returns a function that checks if a word contains all given substrings.
func wordContains(substrings []string) func(word string) bool {
	return func(word string) bool {
		for _, substring := range substrings {
			if !strings.Contains(word, substring) {
				return false
			}
		}
		return true
	}
}

// Returns a function that checks if a word does not contain all given substring.
func wordDoesNotContain(substrings []string) func(word string) bool {
	return func(word string) bool {
		for _, substring := range substrings {
			if strings.Contains(word, substring) {
				return false
			}
		}
		return true
	}
}

// Returns a function that checks if a word has any double letters.
func wordHasDoubleLetter() func(word string) bool {
	return func(word string) bool {
		for i := range len(word) - 1 {
			if word[i] == word[i+1] {
				return true
			}
		}
		return false
	}
}

// Returns a function that checks if a word's length is greater than the given length.
func wordLengthGreaterThan(length int) func(word string) bool {
	return func(word string) bool {
		return len(word) > length
	}
}

// Returns a function that checks if a word's length is less than the given length.
func wordLengthLessThan(length int) func(word string) bool {
	return func(word string) bool {
		return len(word) < length
	}
}

// Returns a function that checks if a word's length is equal to the given length.
func wordLengthEqualsTo(length int) func(word string) bool {
	return func(word string) bool {
		return len(word) == length
	}
}

func wordLengthBetween(low, high int) func(word string) bool {
	return func(word string) bool {
		wordLen := len(word)
		return wordLen >= low && wordLen <= high
	}
}

// Returns a function that checks if a word contains more than one occurrence of a given letter.
func wordContainsMoreThanOne(letter string) func(word string) bool {
	return func(word string) bool {
		return strings.Count(word, letter) > 1
	}
}

// Parses a string like "Starts with mi" and returns a corresponding predicate.
func parsePredicate(predicate string) func(word string) bool {
	predicate = strings.ToLower(predicate)

	switch {
	// Starts with X - The word must start with X.
	case strings.HasPrefix(predicate, "starts with"):
		return wordStartsWith(strings.TrimPrefix(predicate, "starts with "))
	// Ends with X - The word must end with X.
	case strings.HasPrefix(predicate, "ends with"):
		return wordEndsWith(strings.TrimPrefix(predicate, "ends with "))
	// Contains the letter X
	case strings.HasPrefix(predicate, "contains the letter"):
		letter := strings.TrimSpace(strings.TrimPrefix(predicate, "contains the letter "))
		return wordContains([]string{letter})
	// Contains X, Y, Z - Must include each letter anywhere in the word.
	// Contains XY - Must contain the exact sequence.
	case strings.HasPrefix(predicate, "contains"):
		chars := strings.Split(strings.TrimPrefix(predicate, "contains "), ",")
		for i, c := range chars {
			chars[i] = strings.TrimSpace(c)
		}
		// fmt.Println("chars:", chars)
		return wordContains(chars)
	// Does not contain X, Y, Z
	case strings.HasPrefix(predicate, "does not contain"):
		chars := strings.Split(strings.TrimPrefix(predicate, "does not contain "), ",")
		for i, c := range chars {
			chars[i] = strings.TrimSpace(c)
		}
		return wordDoesNotContain(chars)
	// Between X and Y letters
	case strings.HasPrefix(predicate, "between"):
		low := 0
		high := 0
		fmt.Sscanf(predicate, "between %d and %d letters", &low, &high)
		return wordLengthBetween(low, high)
	// Multiple letter X’s - More than one occurrence of X.
	case strings.HasPrefix(predicate, "multiple letter"):
		rest := strings.TrimPrefix(predicate, "multiple letter ")
		letter := strings.TrimSuffix(rest, "'s")
		return wordContainsMoreThanOne(letter)
	// Multiple X’s - More than one occurrence of X.
	case strings.HasPrefix(predicate, "multiple"):
		rest := strings.TrimPrefix(predicate, "multiple ")
		letter := strings.TrimSuffix(rest, "'s")
		return wordContainsMoreThanOne(letter)
	// Double letter - Includes two identical letters in a row.
	case predicate == "double letter":
		return wordHasDoubleLetter()
	// Starts & ends with X
	case strings.HasPrefix(predicate, "starts & ends with"):
		s := strings.TrimPrefix(predicate, "starts & ends with ")
		return wordStartsAndEndsWith(s, s)
	// X letters or fewer
	case strings.HasSuffix(predicate, "letters or fewer"):
		num, err := strconv.Atoi(strings.TrimSuffix(predicate, " letters or fewer"))
		check(err)
		return wordLengthLessThan(num + 1)
	// X letters or more
	case strings.HasSuffix(predicate, "letters or more"):
		num, err := strconv.Atoi(strings.TrimSuffix(predicate, " letters or more"))
		check(err)
		return wordLengthGreaterThan(num - 1)
	// X letter word - The word must have that many letters.
	case strings.HasSuffix(predicate, "letter word"):
		num := strings.TrimSuffix(predicate, " letter word")
		switch num {
		case "two":
			return wordLengthEqualsTo(2)
		case "three":
			return wordLengthEqualsTo(3)
		case "four":
			return wordLengthEqualsTo(4)
		case "five":
			return wordLengthEqualsTo(5)
		case "six":
			return wordLengthEqualsTo(6)
		case "seven":
			return wordLengthEqualsTo(7)
		case "eight":
			return wordLengthEqualsTo(8)
		case "nine":
			return wordLengthEqualsTo(9)
		case "ten":
			return wordLengthEqualsTo(10)
		default:
			panic("Cannot parse number: " + num)
		}
	default:
		panic("Cannot parse predicate: " + predicate)
	}
}
