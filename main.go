package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func word_starts_with(prefix string) func(word string) bool {
	// Returns a function that checks if a word starts with the given prefix.
	return func(word string) bool {
		return strings.HasPrefix(word, prefix)
	}
}

func word_ends_with(suffix string) func(word string) bool {
	// Returns a function that checks if a word ends with the given suffix.
	return func(word string) bool {
		return strings.HasSuffix(word, suffix)
	}
}

func word_contains(substring string) func(word string) bool {
	// Returns a function that checks if a word contains the given substring.
	return func(word string) bool {
		return strings.Contains(word, substring)
	}
}

func word_does_not_contain(substring string) func(word string) bool {
	// Returns a function that checks if a word does not contain the given substring.
	return func(word string) bool {
		return !strings.Contains(word, substring)
	}
}

func word_has_double_letter() func(word string) bool {
	// Returns a function that checks if a word has any double letters.
	return func(word string) bool {
		for i := range len(word) - 1 {
			if word[i] == word[i+1] {
				return true
			}
		}
		return false
	}
}

func word_length_greater_than(length int) func(word string) bool {
	// Returns a function that checks if a word's length is greater than the given length.
	return func(word string) bool {
		return len(word) > length
	}
}

func word_length_less_than(length int) func(word string) bool {
	// Returns a function that checks if a word's length is less than the given length.
	return func(word string) bool {
		return len(word) < length
	}
}

func word_length_equal_to(length int) func(word string) bool {
	// Returns a function that checks if a word's length is equal to the given length.
	return func(word string) bool {
		return len(word) == length
	}
}

func word_contains_more_than_one(letter string) func(word string) bool {
	// Returns a function that checks if a word contains more than one occurrence of a given letter.
	return func(word string) bool {
		return strings.Count(word, letter) > 1
	}
}

func pretty_print(words []string, columns int) string {
	// Returns the array of words in a three-column formatted string.
	if len(words) == 0 {
		fmt.Println("No words found.")
		return ""
	}

	max_length := 0
	for _, word := range words {
		if len(word) > max_length {
			max_length = len(word)
		}
	}

	rows := (len(words) + columns - 1) / columns // Calculate number of rows needed
	var sb strings.Builder

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			index := i + j*rows
			if index < len(words) {
				fmt.Fprintf(&sb, "%-*s ", max_length, words[index])
			} else {
				fmt.Fprintf(&sb, "%-*s ", max_length, "")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func main() {
	raw_words, err := os.ReadFile("words.txt")
	check(err)
	words := strings.Split(string(raw_words), "\n")

	predicates := []func(string) bool{
		word_starts_with("DE"),
		word_contains("F"),
	}

	filtered := slices.DeleteFunc(words, func(w string) bool {
		// Delete words that do not match any of the predicates.
		for _, pred := range predicates {
			if !pred(w) {
				return true
			}
		}
		// All predicates match, so we keep the word.
		return false
	})

	fmt.Println(pretty_print(filtered, 3))
	fmt.Println("Total words found:", len(filtered))
}
