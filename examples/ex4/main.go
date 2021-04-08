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

<body>
    <a href="/dog-cat">dog cat
        <!-- commented text SHOULD NOT be included! -->
    </a>
</body>

</html>`
