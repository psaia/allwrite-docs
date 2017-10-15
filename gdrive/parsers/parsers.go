package parsers

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// TagParsers is a map for handling each type of block element.
var TagParsers = map[string]func(*bytes.Buffer, *html.Node){
	"p":  P,
	"h1": Header,
	"h2": Header,
	"h3": Header,
	"h4": Header,
	"h5": Header,
	"h6": Header,
	"ul": Ul,
	"ol": Ol,
}

// Pulls a link out of our custom format.
var linkRegexp *regexp.Regexp = regexp.MustCompile(
	"___LINKΔΔΔ([^Δ]*)ΔΔΔ",
)

// Pulls an image out of our custom format.
var imgRegexp *regexp.Regexp = regexp.MustCompile(
	"___IMGΔΔΔ([^Δ]*)ΔΔΔ",
)

// FormatStyle an inline style. All text gets procssed through here so we can
// also run some general clean up. Weird characters and things like that.
func FormatStyle(css string, content string) string {

	// Strongs, italics, and strike throughs.
	if strings.Contains(css, "font-weight:700") {
		content = fmt.Sprintf("**%s**", content)
	}
	if strings.Contains(css, "font-style:italic") {
		content = fmt.Sprintf("_%s_", content)
	}
	if strings.Contains(css, "text-decoration:line-through") {
		content = fmt.Sprintf("~~%s~~", content)
	}

	// This is a special case for google's "title" header type. It doesn't
	// actually make it an <h1>, but rather adds font-size: 26pt to it.
	if strings.Contains(css, "font-size:26pt") {
		content = fmt.Sprintf("# %s", content)
	}

	// Handle bad quotes.
	content = strings.Replace(content, "“", "\"", -1)
	content = strings.Replace(content, "”", "\"", -1)

	// Handle images.
	if imgRegexp.Match([]byte(css)) {
		result := imgRegexp.FindStringSubmatch(css)
		if len(result) == 2 {
			parts := strings.Split(result[1], "∏")
			content = fmt.Sprintf("![%s](%s)", parts[0], parts[1])
		} else {
			log.Println("Could not find alt and src from image! " + css)
		}
	}

	// Handle links.
	if linkRegexp.Match([]byte(css)) {
		result := linkRegexp.FindStringSubmatch(css)
		if len(result) == 2 {
			content = fmt.Sprintf("[%s](%s)", content, result[1])
		} else {
			log.Println("Could not find link! " + css)
		}
	}

	return content
}

// GetAttr will get a attribute out of the html.Attribute slice.
func GetAttr(attrs []html.Attribute, attr string) string {
	var style string
	for _, a := range attrs {
		if a.Key == attr {
			style = a.Val
			break
		}
	}
	return style
}

// InlineWalker walks inline.
func InlineWalker(b *bytes.Buffer, n *html.Node, parentCSS string) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			styles := parentCSS + GetAttr(c.Attr, "style")

			// If it's a link, inject the link href. Hackylovely AF.
			if c.DataAtom.String() == "a" {
				styles += "___LINKΔΔΔ" + GetAttr(c.Attr, "href") + "ΔΔΔ___"
			}

			// If at an image, there is no reason to go deeper in since nothing can go
			// in a image tag. Just write it out.
			if c.DataAtom.String() == "img" {
				alt := GetAttr(c.Attr, "alt")
				src := GetAttr(c.Attr, "src")
				styles += "___IMGΔΔΔ" + alt + "∏" + src + "ΔΔΔ___"
				b.WriteString(FormatStyle(styles, c.Data))
				return
			}

			InlineWalker(b, c, styles)
		} else if c.Type == html.TextNode {
			b.WriteString(FormatStyle(parentCSS, c.Data))
		} else {
			InlineWalker(b, c, parentCSS)
		}
	}
}

// Walk .. rather, tippy toe through the DOM.
// This is where you should initiate the parsing.
func Walk(b *bytes.Buffer, n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if fn := TagParsers[c.DataAtom.String()]; fn != nil {
			fn(b, c)
		} else {
			Walk(b, c)
		}
	}
}
