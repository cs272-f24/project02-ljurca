package main

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T){

	// Test case struct
	tests := []struct {
		name string
		url string
		word string
		output frequency
	}{ 
		{
			name: "* sceneI_30.0 Verona * ",
			url: "https://cs272-f24.github.io/tests/rnj/sceneI_30.0.html",
			word : "Verona",
			output: frequency{"https://cs272-f24.github.io/tests/rnj/sceneI_30.0.html": 1},
		},
		{ 
			name: "* sceneI_30.1 Benvolio *",
			url: "https://cs272-f24.github.io/tests/rnj/sceneI_30.1.html",
			word: "Benvolio",
			output: frequency{"https://cs272-f24.github.io/tests/rnj/sceneI_30.1.html": 26},
		},
		{ 
			name: "* all Romeo *",
			url: "https://cs272-f24.github.io/tests/rnj/",
			word: "Romeo",
			output: frequency{"https://cs272-f24.github.io/tests/rnj/sceneI_30.0.html":  2,
				"https://cs272-f24.github.io/tests/rnj/sceneI_30.1.html":  22,
				"https://cs272-f24.github.io/tests/rnj/sceneI_30.3.html":  2,
				"https://cs272-f24.github.io/tests/rnj/sceneI_30.4.html":  17,
				"https://cs272-f24.github.io/tests/rnj/sceneI_30.5.html":  15,
				"https://cs272-f24.github.io/tests/rnj/sceneII_30.2.html": 42,
				"https://cs272-f24.github.io/tests/rnj/":                  199,
				"https://cs272-f24.github.io/tests/rnj/sceneI_30.2.html":  15,
				"https://cs272-f24.github.io/tests/rnj/sceneII_30.0.html": 3,
				"https://cs272-f24.github.io/tests/rnj/sceneII_30.1.html": 10,
				"https://cs272-f24.github.io/tests/rnj/sceneII_30.3.html": 13},
		},
	}

	// Loop over test cases, use DeepEqual to compare structs
	for _, test := range tests {

		// Call Crawl to create inverted index, then find hits through Search 
		index, _ := Crawl(test.url)
		hits := Search(index, test.word)
		if (!reflect.DeepEqual(hits, test.output)) {
			t.Errorf("\n %v TEST FAILED. \n WANTED: %v \n GOT: %v\n", test.name, test.output, hits)
		}
	}
}