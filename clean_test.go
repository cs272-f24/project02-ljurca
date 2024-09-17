package main

import (
	"reflect"
	"testing"
)

func TestCleanHref(t *testing.T){

	// Test case struct
	tests := []struct {
		name string
		base string
		hrefs []string
		output []string
	}{ 
		{
			name: "* relative & absolute * ",
			base: "https://cs272-f24.github.io/",
			hrefs : []string{"/", "/help/", "/syllabus/", "https://gobyexample.com/"},
			output: []string{
				"https://cs272-f24.github.io/",
				"https://cs272-f24.github.io/help/",
				"https://cs272-f24.github.io/syllabus/",
				"https://gobyexample.com/",
				},
		},
		{ 
			name: "* absolute *",
			base: "https://cs272-f24.github.io/",
			hrefs: []string{"https://example.com/", "https://other.com/path", "https://gobyexample.com/"},
			output: []string{ 
				"https://example.com/",
				"https://other.com/path",
				"https://gobyexample.com/",
				},
		},
		{
			name: "* clean *",
			base: "https://example.com/",
			hrefs: []string{"https://example.com/about", "https://example.com/contact"},
			output: []string{ 
				"https://example.com/about", 
				"https://example.com/contact",
				},
		},
		{
			name: "* empty *",
			base: "https://example.com/",
			hrefs: []string{""},
			output: []string{"https://example.com/"},
		},
}

	// Loop over test cases, use DeepEqual to compare structs
	for _, test := range tests {
		got := Clean(test.base, test.hrefs)
		if (!reflect.DeepEqual(got, test.output)) {
			t.Errorf("\n %v TEST FAILED. \n WANTED: %v \n GOT: %v\n", test.name, test.output, got)
		}
	}
}