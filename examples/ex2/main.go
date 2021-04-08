package main

import (
	"fmt"
	"strings"

	"github.com/hauntarl/link-parser"
)

func main() {
	var (
		reader     = strings.NewReader(html)
		links, err = link.Parse(reader)
	)
	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Println(link.String())
	}
}

const html = `<html>

<head>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>

<body>
    <h1>Social stuffs</h1>
    <div>
        <a href="https://www.twitter.com/joncalhoun">
            Check me out on twitter
            <i class="fa fa-twitter" aria-hidden="true"></i>
        </a>
        <a href="https://github.com/gophercises">
            Gophercises is on <strong>Github</strong>!
        </a>
    </div>
</body>

</html>`
