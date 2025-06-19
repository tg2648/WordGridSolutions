package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Predicate struct {
	Name string
	Func func(word string) bool
}

type Result struct {
	Name  string
	Words []string
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

// Returns the array of words in a three-column formatted string.
func prettyPrint(words []string, columns int) string {
	if len(words) == 0 {
		fmt.Println("No words found.")
		return ""
	}

	maxLength := 0
	for _, word := range words {
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}

	rows := (len(words) + columns - 1) / columns // Calculate number of rows needed
	var sb strings.Builder

	for i := range rows {
		for j := range columns {
			index := i + j*rows
			if index < len(words) {
				fmt.Fprintf(&sb, "%-*s  ", maxLength, words[index])
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func getText(n *html.Node) string {
	var sb strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			sb.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)

	// Collapse any whitespace in the text
	return strings.Join(strings.Fields(sb.String()), " ")
}

// Parses a string like "Starts with mi" and returns a corresponding predicate.
func parsePredicate(predicate string) func(word string) bool {
	predicate = strings.ToLower(predicate)

	switch {
	case strings.HasPrefix(predicate, "starts with"):
		return wordStartsWith(strings.TrimPrefix(predicate, "starts with "))
	case strings.HasPrefix(predicate, "contains"):
		chars := strings.Split(strings.TrimPrefix(predicate, "contains "), ",")
		for i, c := range chars {
			chars[i] = strings.TrimSpace(c)
		}
		// fmt.Println("chars:", chars)
		return wordContains(chars)
	case strings.HasPrefix(predicate, "does not contain"):
		chars := strings.Split(strings.TrimPrefix(predicate, "does not contain "), ",")
		for i, c := range chars {
			chars[i] = strings.TrimSpace(c)
		}
		return wordDoesNotContain(chars)
	case strings.HasPrefix(predicate, "between"):
		low := 0
		high := 0
		fmt.Sscanf(predicate, "between %d and %d letters", &low, &high)
		return wordLengthBetween(low, high)
	default:
		panic("Cannot parse predicate: " + predicate)
	}
}

func main() {

	// page := rod.New().MustConnect().NoDefaultDevice().MustPage("https://wordgrid.clevergoat.com/")
	// html := page.MustWaitStable().MustHTML()

	// err := os.WriteFile("wordgrid.html", []byte(html), 0644)

	raw_html, err := os.ReadFile("wordgrid.html")
	check(err)
	html_str := string(raw_html)

	// z := html.NewTokenizer(strings.NewReader(html_str))
	doc, err := html.Parse(strings.NewReader(html_str))
	check(err)

	var spans []string
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == atom.Span.String() {
			spans = append(spans, getText(n))
		}
	}

	columns := make([]Predicate, 3)
	columns[0] = Predicate{Name: spans[0], Func: parsePredicate(spans[0])}
	columns[1] = Predicate{Name: spans[1], Func: parsePredicate(spans[1])}
	columns[2] = Predicate{Name: spans[2], Func: parsePredicate(spans[2])}

	rows := make([]Predicate, 3)
	rows[0] = Predicate{Name: spans[3], Func: parsePredicate(spans[3])}
	rows[1] = Predicate{Name: spans[4], Func: parsePredicate(spans[4])}
	rows[2] = Predicate{Name: spans[5], Func: parsePredicate(spans[5])}

	raw_words, err := os.ReadFile("words.txt")
	check(err)
	words := strings.Split(string(raw_words), "\n")

	results := make([]Result, 0, len(columns)*len(rows))
	for _, col := range columns {
		for _, row := range rows {
			var filtered []string
			for _, w := range words {
				if col.Func(w) && row.Func(w) {
					filtered = append(filtered, w)
				}
			}
			resultName := fmt.Sprintf("%s & %s", col.Name, row.Name)
			results = append(results, Result{Name: resultName, Words: filtered})
		}
	}

	for {
		for i, res := range results {
			fmt.Printf("%d) %s; words found: %d\n", i+1, res.Name, len(res.Words))
		}
		fmt.Println()

		var selection int
		fmt.Printf("Select a result (1-%d) or 0 to exit: ", len(results))
		fmt.Scanf("%d", &selection)

		if selection < 0 || selection > len(results) {
			fmt.Println("Invalid selection. Exiting.")
			break
		}
		if selection == 0 {
			fmt.Println("Exiting.")
			break
		}
		selectedResult := results[selection-1]
		fmt.Println(prettyPrint(selectedResult.Words, 5))
		fmt.Println(selectedResult.Name)
		fmt.Println("Press Enter to continue...")
		fmt.Scanln() // Wait for user to press Enter
	}

}
