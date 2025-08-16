package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Result struct {
	Name       string
	Condition1 string
	Condition2 string
	Words      []string
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

func getSolutions(row_predicates, col_predicates []Predicate) []Result {
	fmt.Println("Loading dictionary...")
	raw_words, err := os.ReadFile("words.txt")
	check(err)
	words := strings.Split(strings.ReplaceAll(string(raw_words), "\r\n", "\n"), "\n")

	fmt.Println("Calculating results...")
	results := make([]Result, 0, len(col_predicates)*len(row_predicates))
	for _, col := range col_predicates {
		for _, row := range row_predicates {
			var filtered []string
			for _, w := range words {
				if col.Func(w) && row.Func(w) {
					filtered = append(filtered, w)
				}
			}
			resultName := fmt.Sprintf("%s & %s", col.Name, row.Name)
			results = append(results, Result{
				Name:       resultName,
				Condition1: col.Name,
				Condition2: row.Name,
				Words:      filtered,
			})
		}
	}

	return results
}

const START_GAME_NUMBER = 442

var START_GAME_NUMBER_DATE = time.Date(2025, 8, 16, 0, 0, 0, 0, time.UTC)

func getGameNumber() int {
	now := time.Now().UTC()
	diff := now.Sub(START_GAME_NUMBER_DATE)

	return START_GAME_NUMBER + int(diff.Hours()/24)
}

func getPredicates() ([]Predicate, []Predicate) {
	resp, err := http.Get("https://api.clevergoat.com/wordgrid/game/" + fmt.Sprint(getGameNumber()))
	check(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Failed to get predicates")
	}

	var data struct {
		Rows []struct {
			Text string `json:"text"`
			Code string `json:"code"`
		} `json:"rows"`
		Columns []struct {
			Text string `json:"text"`
			Code string `json:"code"`
		} `json:"columns"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	check(err)

	rowPredicates := make([]Predicate, len(data.Rows))
	for i, row := range data.Rows {
		rowPredicates[i] = Predicate{
			Name: row.Text,
			Func: parsePredicate(row.Text),
		}
	}

	colPredicates := make([]Predicate, len(data.Columns))
	for i, col := range data.Columns {
		colPredicates[i] = Predicate{
			Name: col.Text,
			Func: parsePredicate(col.Text),
		}
	}

	return rowPredicates, colPredicates
}

func main() {
	rowPredicates, colPredicates := getPredicates()
	results := getSolutions(rowPredicates, colPredicates)

	var outputBuilder strings.Builder
	for _, res := range results {
		outputBuilder.WriteString("<p>")
		outputBuilder.WriteString("<b>")
		outputBuilder.WriteString(fmt.Sprintf("%s -- words found: %d\n", res.Name, len(res.Words)))
		outputBuilder.WriteString("</b>")
		outputBuilder.WriteString(prettyPrint(res.Words, 5))
		outputBuilder.WriteString("</p>")
	}
	fmt.Println()

	// Replace "{# placeholder #}" in ./web/index.html with solutions
	rawHtml, err := os.ReadFile("./web/index.html")
	check(err)
	err = os.WriteFile("./web/index.html", []byte(strings.ReplaceAll(string(rawHtml), "{# placeholder #}", outputBuilder.String())), 0644)
	check(err)
}
