package parser

import (
	"fmt"
	"io"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

//Link type represent <a> tag
type Link struct {
	Href string
	Text string
}

func findNodes(n *html.Node) []*html.Node {
	var nodes []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	for tag := n.FirstChild; tag != nil; tag = n.NextSibling {
		nodes = append(nodes, findNodes(tag)...)
	}
	return nodes
}

func getTagText(n *html.Node, buf *strings.Builder) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	} else if n.Type != html.ElementNode && n.Type != html.DocumentNode {
		return
	} else {
		for tag := n.FirstChild; tag != nil; tag = html.NextSibling {
			getTagText(tag, buf)
		}
	}
}

func readTag(n *html.Node) Link {
	var link Link
	text := strings.Builder{}
	for _, a := range n.Attr {
		if a.Key == "href" {
			link.Href = a.Val
		}
	}
	getTagText(n, &text)
	tmp := strings.FieldsFunc(
		text.String(), func(r rune) bool { return unicode.IsSpace(r) },
	)
	link.Text = strings.Join(tmp, " ")
	return link
}

//Parse function return slice of links which are in HTML document
func Parse(r io.Reader) ([]Link, error) {
	var links []Link
	doc, err := html.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("Error parsing html decument: %s", err)
	}
	aTags := findNodes(doc)
	for _, aTag := range aTags {
		links = append(links, readTag(aTag))
	}
	return links, nil
}

func DFS(n *html.Node, padding string) {
	fmt.Println(padding, n.Data)
	for tag := n.FirstChild; tag != nil; tag = n.NextSibling {
		DFS(tag, padding+" ")
	}
}
