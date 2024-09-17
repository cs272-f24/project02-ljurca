package main

import(
"net/http"
"log"
"io"
)

func Download(url string)([]byte, error){

	// Get the Body using http.Get 
	rsp, err := http.Get(url)

	if err != nil{
		log.Fatalf("http.Get returned %v\n", err)
		return nil, err
	}

	// When body goes out of scope, defer and close 
	defer rsp.Body.Close() 
	
	// Read the response body into a byte slice
	body, err := io.ReadAll(rsp.Body)

	if err != nil{
		log.Fatalf("io.ReadAll returned %v\n", err)
		return nil, err
	}

	return body, nil
}