package src

import (
	"os"

	"golang.org/x/net/html"
)

// Parser parses through code and finds tokens.
type Parser interface {
	Parse(*os.File, chan html.Token)
}

// HTMLParser implements Parser.
type HTMLParser struct {
}

// NewHTMLParser returns and new instance of HTMLParser.
func NewHTMLParser() Parser {
	return &HTMLParser{}
}

// Parse starts a go routine to ParseHTML for tokens.
func (h HTMLParser) Parse(data *os.File, elements chan html.Token) {
	go ParseHTML(data, elements)
}

// ParseHTML creates a tokenizer and begins generating elements.
func ParseHTML(data *os.File, elements chan html.Token) {
	tokenizer := html.NewTokenizer(data)

	generateElements(elements, tokenizer)
	close(elements)
}

// generateElements finds html elements or tokens and sends them on the elements channel.
func generateElements(elements chan html.Token, tokenizer *html.Tokenizer) {
	ended := false
	for !ended {
		t := tokenizer.Next()

		if t == html.ErrorToken {
			// The html document has ended
			ended = true
		} else if t == html.StartTagToken {
			t2 := tokenizer.Token()
			elements <- t2 // Add found elements to channel.
		}
	}
}
