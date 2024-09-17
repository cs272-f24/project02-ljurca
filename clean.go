package main

import (
	"log"
	"net/url"
)

func Clean(host string, hrefs []string) []string {

	// Declare cleaned URL
	var cleaned []string

	// Error checking
	u, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Loop over hrefs, parse with host 
	for _, href := range hrefs {
		rel, err := u.Parse(href)
		if err != nil {
			log.Fatal(err)
			continue
		}
		cleaned = append(cleaned, rel.String())
	}

	return cleaned
}
