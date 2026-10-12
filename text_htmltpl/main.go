package main

import (
	"encoding/json"
	"fmt"
	htmltpl "html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	texttpl "text/template"
	"time"
)

func main() {
	t1()
	t2()
}

func t1() {
	const issueURL = "https://api.github.com/search/issues"

	type user struct {
		Login   string
		HTMLURL string `json:"html_url"`
	}
	type issue struct {
		Number    int
		HTMLURL   string `json:"html_url"`
		Title     string
		State     string
		User      *user
		CreatedAt time.Time `json:"created_at"`
		Body      string    // in Markdown format
	}
	type issuesSearchResult struct {
		TotalCount int `json:"total_count"`
		Items      []*issue
	}

	const templ = `
{{.TotalCount}} issues:
{{range limit .Items 3 }}
------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
	`

	daysAgo := func(t time.Time) int {
		return int(time.Since(t).Hours() / 24)
	}

	limit := func(items []*issue, n int) []*issue {
		if len(items) > n {
			return items[:n]
		}
		return items
	}

	report := texttpl.Must(texttpl.New("issuelist").Funcs(texttpl.FuncMap{"daysAgo": daysAgo, "limit": limit}).Parse(templ))

	searchIssues := func(terms []string) (*issuesSearchResult, error) {
		q := url.QueryEscape(strings.Join(terms, " "))
		resp, err := http.Get(issueURL + "?q=" + q)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("search query failed: %s", resp.Status)
		}
		var result issuesSearchResult
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	}

	result, err := searchIssues([]string{"is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, *result); err != nil {
		log.Fatal(err)
	}
}

func t2() {
	const templ = `<p>A:{{.A}}</p> <p>B:{{.B}}</p>`
	t := htmltpl.Must(htmltpl.New("escape").Parse(templ))
	var data struct {
		A string       // untrusted plain text
		B htmltpl.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>" // < > 不是直接嵌入了，先转义成lt; gt;
	data.B = "<b>Hello!</b>"
	f, _ := os.Create("index2.html")
	if err := t.Execute(f, data); err != nil {
		log.Fatal(err)
	}
}
