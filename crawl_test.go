package main

import(
	"testing"
	"reflect"
)

func TestCrawl(t *testing.T){
	// Testing struct
	tests := []struct{
		name string
		url string
		want []string
	}{
		{
			name: "* simple *",
			url: "https://cs272-f24.github.io/tests/project01/simple.html",
			want: []string{"https://cs272-f24.github.io/tests/project01/simple.html"},
		},
		{
			name:"* href *",
			url:"https://cs272-f24.github.io/tests/project01/href.html",
			want: []string{"https://cs272-f24.github.io/tests/project01/href.html", "https://cs272-f24.github.io/tests/project01/simple.html"},
		},
		{ 
			name:"* style *",
			url:"https://cs272-f24.github.io/tests/project01/style.html",
			want: []string{"https://cs272-f24.github.io/tests/project01/style.html", "https://cs272-f24.github.io/tests/project01/href.html", "https://cs272-f24.github.io/tests/project01/simple.html"},
		},
		{
			name:"* index *",
			url:"https://cs272-f24.github.io/tests/project01/index.html",
			want: []string{"https://cs272-f24.github.io/tests/project01/index.html", "https://cs272-f24.github.io/tests/project01/simple.html", "https://cs272-f24.github.io/tests/project01/href.html", "https://cs272-f24.github.io/tests/project01/style.html"},
		},
	}

	// Loop over test cases, use DeepEqual to compare structs
	for _, test := range tests {
		_, got := Crawl(test.url)
		if (!reflect.DeepEqual(got, test.want)) {
			t.Errorf("\n %v TEST FAILED. \n WANTED: %v \n GOT: %v\n", test.name, test.want, got)
		}
	}
}