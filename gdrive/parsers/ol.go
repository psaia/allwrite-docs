package parsers

import (
	"bytes"
	"strconv"

	"golang.org/x/net/html"
)

// Ol is a lot like a Ul but ordered.
func Ol(b *bytes.Buffer, n *html.Node) {
	index := 1
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			b.WriteString(strconv.Itoa(index) + " ")
			InlineWalker(b, c, GetAttr(n.Attr, "style"))
			b.WriteString("\n")
			index++
		}
	}
	b.WriteString("\n\n")
}
