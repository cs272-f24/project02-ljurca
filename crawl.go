package main

import (
	"log"
)

// Crawl function to crawl URLs and build the inverted index
func Crawl(seedURL string) (indexOuter, []string) {
	visitedURLs := make(map[string]bool) // Keep track of visited URLs
	queue := []string{seedURL}           // Initialize the queue with the seed URL
	var crawled []string                 // Store crawled URLs
	index := make(indexOuter)			// create inverted index

	for len(queue) > 0 {
		// Pop from the queue
		currentURL := queue[0]
		queue = queue[1:]

		// Skip if already visited
		if visitedURLs[currentURL] {
			continue
		}

		visitedURLs[currentURL] = true
		crawled = append(crawled, currentURL)

		// Download the body
		body, err := Download(currentURL)
		if err != nil {
			log.Printf("Failed to download URL: %s, error: %v", currentURL, err)
			continue
		}

		// Extract hrefs from body
		_, hrefs := Extract([]byte(body))

		// Update the index with words and frequency per URL 
		index = Index([]string{currentURL}, currentURL, index)

		// Clean and validate URLs
		cleanedURLs := Clean(currentURL, hrefs)
		for _, cleanedURL := range cleanedURLs {
			// Skip specific unwanted URLs, in this instance Gutenberg links 
			if cleanedURL == "https://www.gutenberg.org/" || cleanedURL == "https://www.gutenberg.org/donate/" {
				continue
			}

			// Add cleaned URLs to the queue if not visited and valid
			if !visitedURLs[cleanedURL] {
					queue = append(queue, cleanedURL)
					visitedURLs[cleanedURL] = false // URL is in queue, but not yet visited
			}
		}
	}
	return index, crawled
}