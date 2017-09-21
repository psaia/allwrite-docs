package gdrive

import (
	"bytes"
	"io"

	"github.com/LevInteractive/allwrite-docs/gdrive/parsers"

	"golang.org/x/net/html"
)

// MarshalMarkdownFromHTML transforms a nasty google doc to markdown.
func MarshalMarkdownFromHTML(r io.Reader) (string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	parsers.Walk(&buffer, doc)
	return buffer.String(), nil
}
