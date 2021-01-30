package parse

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

// type Link struct {
// 	Href string
// 	text string
// }

func Parse(r io.Reader) {
	node, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	dfs(node)
}

func dfs(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		fmt.Println(node)
		return
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		dfs(node)
	}
}
