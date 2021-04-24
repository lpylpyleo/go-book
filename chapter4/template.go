package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const textTempl = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(textTempl))

func main() {
	testIssueSearch()
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

const IssueURL = "https://api.github.com/search/issues"

func testIssueSearch() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err = issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
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
