package parsers

import (
	"bytes"

	"golang.org/x/net/html"
)

// Ul parses UL types of tags.
func Ul(b *bytes.Buffer, n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			b.WriteString("* ")
			InlineWalker(b, c, GetAttr(n.Attr, "style"))
			b.WriteString("\n")
		}
	}
	b.WriteString("\n\n")
}
