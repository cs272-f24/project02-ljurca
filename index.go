package main

import (
	"log"
	"strings"
	"unicode"
	"github.com/kljensen/snowball"
)

// Intializing for inverted index
type frequency map[string]int
type indexOuter map[string]frequency

func Check(word string) string{
	word = strings.ToLower(word)

	// Check to see if the word needs to be cleaned
	var cleaned strings.Builder
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned.WriteRune(r)
		}
	}
	return cleaned.String()
}

func Index(body []string, url string, index indexOuter) indexOuter {
	for _, url := range body {

		// Download URls and Extract words  
		body, err := Download(url)
		if err != nil {
			log.Printf("Failed to download URL: %s", url)
			continue
		}
		words, _ := Extract([]byte(body))

		// Process each word
		for _, word := range words {
			cleaned := Check(word)

			// Stem the word
			stemmedWord, err := snowball.Stem(cleaned, "english", true)
			if err != nil {
				log.Printf("Error stemming word: %v", err)
				continue
			}
			
			// Create new entry if needed
			if _, ok := index[stemmedWord]; !ok{
				index[stemmedWord] = make(frequency)
			}
			index[stemmedWord][url]++
		}
	}
	return index
}