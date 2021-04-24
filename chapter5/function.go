package htmlout

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func test() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	// for _, link := range visit(nil, doc) {
	// 	fmt.Println(link)
	// }
	// outline(nil, doc)
	counter := map[string]int{}
	countTags(counter, doc)
	fmt.Println(counter)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func countTags(counter map[string]int, n *html.Node) {
	// count := map[string]int{}
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countTags(counter, c)
	}
}
