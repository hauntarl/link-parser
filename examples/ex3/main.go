package main

import (
	"fmt"
	"os"

	"github.com/hauntarl/link-parser"
)

func main() {
	file, err := os.Open("files/ex3.html")
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(file)
	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Println(link.String())
	}
}
