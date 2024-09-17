package main

import (
	"log"
	"strings"
	"golang.org/x/net/html"
	"unicode"
)


func Extract(body []byte) ([]string, []string){

	// Initialize words and hrefs as empty slices, returns 
	// empty slice instead of nil if blank 
	words := []string{}
	hrefs := []string{}

	// Error
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
		return words, hrefs
	}

	// Define recursive function, tree structure 
	var f func(*html.Node)
	f = func(n *html.Node) {
		var styleTag bool

		// Set the flag to true and skip over if style tag is found 
		if n.Type == html.ElementNode && n.Data == "style" {
			styleTag = true
			return 
		}
	
		// Reset the flag
		if styleTag && n.Type == html.ElementNode && n.Data != "style" {
			styleTag = false
		}
	
		// Append hrefs to href's slice if found 
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					hrefs = append(hrefs, a.Val)
					break
				}
			}
		}

		// Append words to words's slice if found 
		if n.Type == html.TextNode {
			text := strings.FieldsFunc(n.Data, unicode.IsSpace) // split text on whitespace char
			if len(text) > 0 {
				words = append(words, text...)
			}
		}

		// Traverse through HTML tree structure
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
}
	f(doc)

	return words, hrefs
}

