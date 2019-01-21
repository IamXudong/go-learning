package main

import (
	"github.com/stevzhang01/go-learning/github"
	"html/template"
	"log"
	"os"
)

const templ = `
<head>
	<title>Github Issues</title>
	<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
</head>
<body>
<div class="container">
<h1>{{.TotalCount}} Issues</h1>
<table class="table table-bordered">
	<tr>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
		<td>{{.Number}}</td>
		<td>{{.State}}</td>
		<td><a href="{{.User.HTMLURL}}">{{.User.Login}}</a></td>
		<td><a href="{{.HTMLURL}}">{{.Title}}</a></td>
	</tr>
	{{end}}
</table>
</div>
<body>`

func main() {
	var report = template.Must(template.New("issuelist").Parse(templ))
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
