package gdrive

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func getFixture(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(file)
}

func TestMarshalMarkdownFromHTML(t *testing.T) {
	mdDoc := getFixture("./fixtures/sample-doc.md")
	htmlDoc := getFixture("./fixtures/sample-doc.html")

	r := strings.NewReader(htmlDoc)
	transformedMd, err := MarshalMarkdownFromHTML(r)

	if err != nil {
		t.Error(err)
	}

	if transformedMd != mdDoc {
		t.Error("HTML did not translate to markdown properly.")
	}
}

func TestMarshalMarkdownFromHTMLImages(t *testing.T) {
	mdDoc := getFixture("./fixtures/image-doc.md")
	htmlDoc := getFixture("./fixtures/image-doc.html")

	r := strings.NewReader(htmlDoc)
	transformedMd, err := MarshalMarkdownFromHTML(r)
	fmt.Println(transformedMd)

	if err != nil {
		t.Error(err)
	}

	if transformedMd != mdDoc {
		t.Error("HTML did not translate to image markdown properly.")
	}
}
