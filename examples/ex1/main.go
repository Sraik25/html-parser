package main

import (
	"fmt"
	"github.com/Sraik25/html-parser"
	"strings"
)

var exampleHtml = `
<html>
	<body>
		<h1>Hello!</h1>
		<a href="/other-page">A link to another page</a>
	</body>
,/html>
`

func main() {
	r := strings.NewReader(exampleHtml)

	links, err := html_parser.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
