package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document
type Link struct{ Href, Text string }

// Parse will take in an HTML document and will return a slice of links parsed
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var (
		nodes = linkNodes(doc)
		links []Link
	)
	for _, n := range nodes {
		links = append(links, buildLink(n))
	}
	return links, nil
}

func buildLink(n *html.Node) (res Link) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			res.Href = attr.Val
			break
		}
	}
	res.Text = strings.Join(strings.Fields(text(n)), " ")
	return
}

func text(n *html.Node) (res string) {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res += text(c) + " "
	}
	return
}

func linkNodes(n *html.Node) (res []*html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		return append(res, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, linkNodes(c)...)
	}
	return
}

// func traverse(n *html.Node, padding string) {
// 	msg := n.Data
// 	if n.Type == html.ElementNode {
// 		msg = "<" + msg + ">"
// 	}
// 	fmt.Println(padding, msg)

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		traverse(c, padding+"  ")
// 	}
// }
