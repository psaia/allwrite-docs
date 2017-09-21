package parsers

import (
	"bytes"

	"golang.org/x/net/html"
)

// P parses paragraph tags.
func P(b *bytes.Buffer, n *html.Node) {
	InlineWalker(b, n, GetAttr(n.Attr, "style"))
	b.WriteString("\n\n")
}
