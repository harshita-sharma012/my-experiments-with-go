package parse

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) {
	node, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	dfs(node)
}

func fetchHref(node *html.Node, link *Link) {
	for _, value := range node.Attr {
		if value.Key == "href" {
			link.Href = value.Val
		}
	}
	//fetchText(node)
}

func fetchText(node *html.Node, link *Link) {
	link.Text = node.Data
}

func dfs(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		var link Link
		fetchHref(node, &link)
		next := node.FirstChild
		if next.Type == html.TextNode {
			fetchText(next, &link)
		}
		fmt.Println(link)
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		dfs(node)
	}
}
