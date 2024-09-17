package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"reflect"

)

func TestDownload(t *testing.T) {

	// Test case struct
	tests := []struct {
		name string
		url string
		want_body string
	}{
		{
			name: "* simple * ",
			url: "", 			// overridden by the mock server URL
			want_body:`<html><body>Hello CS 272, there are no links here. </body></html>`, 
		},
		{
			name: "* href * ",
			url: "",
			want_body: `<html><body>For a simple example, 
			see <a href="/tests/project01/simple.html">simple.html</a></body></html>`, 
		},
		{
			name: "* style * ",
			url: "",
			want_body: `<html><head><title>Style</title><style>a.blue {color: blue;}
			a.red {color: red; }</style><body><p> Here is a blue link to <a 
			class="blue" href="/tests/project01/href.html">href.html</a></p><p>And a red 
			link to <a class="red" href="/tests/project01/simple.html">simple.html</a></p>
			</body></html>`, 
		},
		{
			name: "* index * ",
			url: "",
			want_body: `<html><body><ul><li><a href="/tests/project01/simple.html">
			simple.html</a></li><li><a href="/tests/project01/href.html">href.html</a></li><li>
			<a href="/tests/project01/style.html">style.html</a></ul></body></html>`, 
		},
	}

	// Loop through test cases, use DeepEqual to compare structs
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Mock handler 
			handler := func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(test.want_body)) // Creates content
			}
			// New mock server
			server := httptest.NewServer(http.HandlerFunc(handler))
			defer server.Close()

			// Download function is called with the mock server URL
			got_body, err := Download(server.URL + test.url)

			if err != nil {
				t.Errorf("Failed to download: %v", err)
			}

			// Compare the actual body with the expected body
			if !reflect.DeepEqual(string(got_body), test.want_body) {
				t.Errorf("\n %v TEST FAILED. \n WANTED: %v \n GOT: %v\n", test.name, test.want_body, got_body)
			}
		})
	}
}