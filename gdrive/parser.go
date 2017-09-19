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

// Pull the style attribute text out from a node.
func getStyle(attrs []html.Attribute) string {
	var style string
	for _, a := range attrs {
		if a.Key == "style" {
			style = a.Val
			break
		}
	}
	return style
}

// Format the text based on the CSS string from the style attribute.
func formatStyle(styleAttr string, content string) string {
	if strings.Contains(styleAttr, "font-weight:700") {
		content = fmt.Sprintf("**%s**", content)
	}
	if strings.Contains(styleAttr, "font-style:italic") {
		content = fmt.Sprintf("*%s*", content)
	}
	if strings.Contains(styleAttr, "text-decoration:line-through") {
		content = fmt.Sprintf("~~%s~~", content)
	}
	return content
}

// Count the number of children for a particular node.
func numberOfChildren(n *html.Node) int {
	if n == nil {
		return -1
	}

	count := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count++
	}

	return count
}

func walk(b *bytes.Buffer, n *html.Node) {
	if n.Type == html.TextNode {

		primaryTag := n.Parent.Parent.DataAtom.String()
		style := getStyle(n.Parent.Attr)
		iContent := formatStyle(style, strings.Trim(n.Data, " "))

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
				// i := strconv.Itoa(numberOfChildren(n.Parent.Parent))
				// fmt.Println("number of kids " + i)
				b.WriteString(fmt.Sprintf("%s\n\n", iContent))
			case "li":
				superTag := n.Parent.Parent.Parent.DataAtom.String()
				if superTag == "ol" {
					b.WriteString(fmt.Sprintf("1. %s\n", iContent))
				} else {
					b.WriteString(fmt.Sprintf("* %s\n", iContent))
				}
			}
		}
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
