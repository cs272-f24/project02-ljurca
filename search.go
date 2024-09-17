package main

import (
	"log"
	"strings"
	"github.com/kljensen/snowball"
)

func Search(index indexOuter, word string) frequency {

	// Make word lowercase, get stemmedWord
	word = strings.ToLower(word)
	stemmedWord, err := snowball.Stem(word, "english", true)
	if err != nil{
		log.Fatalf("Error stemming word: %v", err)
	}

	// Return frequency from map
	if frequency, ok := index[stemmedWord]; ok{
		return frequency
	}

	log.Fatalf("Error stemming word: %v", err)
	return nil
}