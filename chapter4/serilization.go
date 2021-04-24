package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

const IssueURL = "https://api.github.com/search/issues"

func main() {
	testIssueSearch()
}

// ------------------------------------------------------------------

func testIssueSearch() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Total counts: %d.\n", result.TotalCount)

	var lastMonth []Issue
	var lastYear []Issue
	var overOneYear []Issue

	now := time.Now()

	for _, item := range result.Items {
		elapsedHours := now.Sub(item.CreatedAt).Hours()
		if elapsedHours < 24*30 {
			lastMonth = append(lastMonth, *item)
		} else if elapsedHours < 24*30*365 {
			lastYear = append(lastYear, *item)
		} else {
			overOneYear = append(overOneYear, *item)
		}
	}
	fmt.Printf("last month: \n")
	printIssues(lastMonth)
	fmt.Printf("last year: \n")
	printIssues(lastYear)
	fmt.Printf("over a year: \n")
	printIssues(overOneYear)
}

func printIssues(issues []Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, err
}

type IssuesSearchResult struct {
	TotalCount int      `json:"total_count,omitempty"`
	Items      []*Issue `json:"items,omitempty"`
}

type Issue struct {
	Number    int       `json:"number,omitempty"`
	HTMLURL   string    `json:"html_url,omitempty"`
	Title     string    `json:"title,omitempty"`
	State     string    `json:"state,omitempty"`
	User      *User     `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Body      string    `json:"body,omitempty"`
}

type User struct {
	Login   string `json:"login,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
}

// ------------------------------------------------------------------

type Movie struct {
	Title  string   `json:"title,omitempty"`
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors,omitempty"`
}

func testMarshal() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var s []Movie
	err = json.Unmarshal([]byte(data), &s)
	fmt.Printf("%#v\n", s)
}
