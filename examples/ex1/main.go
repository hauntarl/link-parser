package main

import (
	"fmt"
	"strings"

	"github.com/hauntarl/link-parser"
)

func main() {
	var (
		reader     = strings.NewReader(ex1HTML)
		links, err = link.Parse(reader)
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

const ex1HTML = `<html>

<body>
    <h1>Hello!</h1>
    <a href="/other-page">
		A link to another page 
		<span> some span </span>
	</a>
	<a href="/page-two">A link to second page</a>
</body>

</html>`
