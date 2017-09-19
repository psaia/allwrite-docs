package gdrive

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

var allowedTags = map[string]bool{
	"a":    true,
	"p":    true,
	"li":   true,
	"img":  true,
	"span": true,
	"h1":   true,
	"h2":   true,
	"h3":   true,
	"h4":   true,
	"h5":   true,
	"h6":   true,
}

func getStyle(attrs []html.Attribute) string {
	style := ""
	for _, a := range attrs {
		if a.Key == "style" {
			style = a.Val
			break
		}
	}
	return style
}

func walk(b *bytes.Buffer, n *html.Node) {
	if n.Type == html.TextNode {

		primaryTag := n.Parent.Parent.DataAtom.String()
		// style := getStyle(n.Parent.Attr)
		iContent := strings.Trim(n.Data, " ")

		if allowedTags[primaryTag] {
			switch primaryTag {
			case "h1":
				b.WriteString(fmt.Sprintf("# %s\n\n", iContent))
			case "h2":
				b.WriteString(fmt.Sprintf("## %s\n\n", iContent))
			case "h3":
				b.WriteString(fmt.Sprintf("### %s\n\n", iContent))
			case "h4":
				b.WriteString(fmt.Sprintf("#### %s\n\n", iContent))
			case "h5":
				b.WriteString(fmt.Sprintf("##### %s\n\n", iContent))
			case "h6":
				b.WriteString(fmt.Sprintf("###### %s\n\n", iContent))
			case "p":
				b.WriteString(fmt.Sprintf("%s\n\n", iContent))
			case "li":
				superTag := n.Parent.Parent.Parent.DataAtom.String()
				if superTag == "ol" {
					b.WriteString(fmt.Sprintf("1. %s\n", iContent))
				} else {
					b.WriteString(fmt.Sprintf("* %s\n", iContent))
				}
			}
			// b.WriteString("-----\n")
			// b.WriteString("Tag: " + primaryTag + "\n")
			// b.WriteString("Content: " + n.Data + "\n")
			// b.WriteString("Style: " + style + "\n")
			// b.WriteString("-----\n")
		}
		// 	// b.WriteString(n.DataAtom.String())
		// 	// b.WriteString("\n")
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walk(b, c)
	}
}

// MarshalMarkdownFromHTML transforms a nasty google doc to markdown.
func MarshalMarkdownFromHTML(r io.Reader) (string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	walk(&buffer, doc)
	return buffer.String(), nil
}
