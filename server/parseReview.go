package main

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type Ratings struct {
	group      string
	albumTitle string
	// reviewLink   string will need to be added during final phase
	readerRating int
	authorRating int
}

// parseReview parses an html file and returns a Ratings struct of collected info
func ParseReview(url ReviewUrl) *Ratings {
	var ratings Ratings

	resp, err := http.Get(string(url))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	fileScanner := bufio.NewScanner(resp.Body)
	fileScanner.Split(bufio.ScanLines)

	findGroup := regexp.MustCompile(`<h3>([\w\s]+)</h3>`)
	findAlbumTitle := regexp.MustCompile(`<h4><i>([\w\s]+)</i></h4>`)
	findAuthorRating := regexp.MustCompile(`Author rating: <b>([0-9])</b>`)
	findReaderRating := regexp.MustCompile(`reader rating: <b>([0-9])</b>`)

	for fileScanner.Scan() {
		Group := findGroup.FindStringSubmatch(fileScanner.Text())
		aTitle := findAlbumTitle.FindStringSubmatch(fileScanner.Text())
		aRating := findAuthorRating.FindStringSubmatch(fileScanner.Text())
		rRating := findReaderRating.FindStringSubmatch(fileScanner.Text())

		// this regexp has several possible matches per page - target the first one only
		if ratings.group == "" && len(Group) > 1 {
			group := Group[1]
			ratings.group = group
		}

		if len(aTitle) > 1 {
			title := aTitle[1]
			ratings.albumTitle = title
		}

		if len(aRating) > 1 {
			num, err := strconv.Atoi(aRating[1])
			if err != nil {
				fmt.Println(err)
			}
			ratings.authorRating = num
		}

		if len(rRating) > 1 {
			num, err := strconv.Atoi(rRating[1])
			if err != nil {
				fmt.Println(err)
			}
			ratings.readerRating = num
		}
	}
	if ratings.authorRating >= 9 || ratings.readerRating >= 9 {
		return &ratings
	}
	return nil
}
