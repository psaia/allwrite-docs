package parsers

import (
	"bytes"

	"golang.org/x/net/html"
)

// Header parses h1, h2, h3, etc, tags.
func Header(b *bytes.Buffer, n *html.Node) {
	switch n.DataAtom.String() {
	case "h1":
		b.WriteString("# ")
	case "h2":
		b.WriteString("## ")
	case "h3":
		b.WriteString("### ")
	case "h4":
		b.WriteString("#### ")
	case "h5":
		b.WriteString("##### ")
	case "h6":
		b.WriteString("###### ")
	}

	InlineWalker(b, n, GetAttr(n.Attr, "style"))
	b.WriteString("\n\n")
}
