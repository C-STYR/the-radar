package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://duckduckgo.com"
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", html)
}
