package main

import (
	"reflect"
	"testing"
)

func TestExtract(t *testing.T) {

	// Test case struct
	tests := []struct {
		name string
		input []byte
		want_words[]string
		want_hrefs[]string
	}{
		{
			name: " * simple * ",
			input: []byte("<p>Hello World!</p>"),
			want_words: []string{"Hello", "World!"},
			want_hrefs: []string{},
		},
		{
			name: " * link *",
			input: []byte(`<a href="http://example.com">Example</a>`),
			want_words: []string{"Example"},
			want_hrefs: []string{"http://example.com"},
		},
		{
			name:"* empty body * ",
			input: []byte(""), 
			want_words: []string{},
			want_hrefs: []string{},
		},
		{
			name: "* whitespace only *",
			input: []byte(`<body>   </body>`), 
			want_words: []string{}, 
			want_hrefs: []string{}, 
		},
		{
			name: "* empty text node *",
			input: []byte(`<body><p></p></body>`),
			want_words: []string{},
			want_hrefs: []string{}, 
		},
		{
			name: "* nested tags *",
			input: []byte(`<div><p>Hello<a href="http://example.com">World!</a></p></div>`), 
			want_words: []string{"Hello", "World!"},
			want_hrefs: []string{"http://example.com"},
		},
		{
			name: "* multiple *",
			input: []byte(`<p>Visit <a href="http://example.com">Example</a> or <a href="http://test.com">Test.</a></p>`),
			want_words: []string{"Visit", "Example", "or", "Test."},
			want_hrefs: []string{"http://example.com", "http://test.com"},
		},
		{
			name: "* special char *",
			input: 	[]byte(`<p>C++ and Python.</p>`),
			want_words:  []string{"C++", "and", "Python."},
			want_hrefs: []string{},
		},
		{
			name: "* broken *",
			input: []byte(`<div><p>Broken <a href="http://example.com">Link</p></div>`), 
			want_words: []string{"Broken", "Link"},
			want_hrefs: []string{"http://example.com"}, 
		},
		{
			name: "* mixed *",
			input: []byte(`<A HREF="http://example.com">Link</A>`),
			want_words: []string{"Link"},
			want_hrefs: []string{"http://example.com"},
		},
	}

	// Loop over test cases, use DeepEqual to compare structs
	for _, test := range tests {
		got_words, got_hrefs := Extract(test.input)
		if (!reflect.DeepEqual(got_words, test.want_words)) {
			t.Errorf("\n %v WORD TEST FAILED. \n WANTED: %v \n GOT: %v\n", test.name, test.want_words, got_words)
		}
		if (!reflect.DeepEqual(got_hrefs, test.want_hrefs)) {
			t.Errorf("\n %v HREF TEST FAILED. \n WANTED: %v \n GOT: %v\n", test.name, test.want_hrefs, got_hrefs)
		}
	}
}
