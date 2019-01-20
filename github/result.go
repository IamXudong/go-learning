// Package github 提供了 github issue 跟踪接口的 GO API
package github

import "time"

const IssueUrl = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // Markdown 格式
}

type User struct {
	Login   string
	HTMLURL string `jsong: "html_url"`
}
