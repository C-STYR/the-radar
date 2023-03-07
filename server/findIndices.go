package main

import (
	"fmt"
	"net/http"
)

/*
FindIndices starts at "https://www.undertheradarmag.com/reviews/category/music"
then proceeds to add "P10", "P20", etc in a loop, visiting each site and returning the html
*/
func FindIndices(urlChan *UrlList) {

	urls := []string{
		"https://www.undertheradarmag.com/reviews/category/music",
		"https://www.undertheradarmag.com/reviews/category/music/P10",
	}

	for _, page := range urls {

		resp, err := http.Get(page)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		links := ReadIndex(resp.Body)

		for url := range links {
			urlChan.Enqueue(CreateReviewUrl(url))
		}
	}
}
