package parser

import (
	"strconv"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

type Test struct {
	tag    string
	result string
}

func TestGetText(t *testing.T) {
	test := Test{
		tag:    `<h1>Some text<h2>more text</h2></h1>`,
		result: `Some textmore text`,
	}

	r := strings.NewReader(test.tag)
	doc, err := html.Parse(r)
	if err != nil {
		t.Fatalf("Can't parse test html. Err - %s", err)
	}
	buf := strings.Builder{}
	getTagText(doc, &buf)
	if buf.String() != test.result {
		t.Errorf("Expected %s to equal %s", buf.String(), test.result)
	}

}

func TestGetNodes(t *testing.T) {
	test := Test{
		tag:    `<a href="#">Text<a href="1">more<span>text</span></a></a>`,
		result: "1",
	}
	r := strings.NewReader(test.tag)
	doc, err := html.Parse(r)
	if err != nil {
		t.Fatalf("Can't parse test html. Err - %s", err)
	}
	nodes := findNodes(doc)
	if strconv.Itoa(len(nodes)) != test.result {
		t.Errorf("Expected %v to have length of %s", nodes, test.result)
	}
}
