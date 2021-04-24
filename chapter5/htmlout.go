package htmlout

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

func ForEachElement(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachElement(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++

		} else {
			fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func Output(r io.Reader) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	ForEachElement(doc, startElement, endElement)
}
