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
	Name       string   `json:"name"`
	Condition1 string   `json:"condition_1"`
	Condition2 string   `json:"condition_2"`
	Words      []string `json:"words"`
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
	diffDays := int(diff.Hours() / 24)

	return START_GAME_NUMBER + diffDays
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

	// Add timestamp to results data
	type ResultsData struct {
		GameNumber int       `json:"game_number"`
		Timestamp  time.Time `json:"timestamp"`
		Results    []Result  `json:"results"`
	}

	resultsData := ResultsData{
		GameNumber: getGameNumber(),
		Timestamp:  time.Now().UTC(),
		Results:    results,
	}

	// Write results to a JSON file
	jsonData, err := json.MarshalIndent(resultsData, "", "  ")
	check(err)
	err = os.WriteFile("./web/src/results.json", jsonData, 0644)
	check(err)
	fmt.Println("Results written to ./web/src/results.json")
}
