package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/*
- start at https://www.undertheradarmag.com/reviews/category/music
- next page is https://www.undertheradarmag.com/reviews/category/music/P10
- visit each page in a goroutine and search the source code for regex matches to
	"https://www.undertheradarmag.com/reviews/*" where * excludes "category"
- each will be inside a div cn="headline", inside an h3
- there will be repeats
- build a map of those links, which will then be traversed to look for ratings
- then filter for author ratings or reader ratings > 8.5
- for each review page with a rating greater than 8.5, make an object with the two ratings, the name of the album, the name of the band, a link to the album review page
- create a slice of objects, serialize and send to the front end

*/

func main() {
	readIndex("./example-html/example-index.html")
}


// readIndex parses an html file and returns a slice of urls that match a given regexp
func readIndex(filePath string) []string {
	// could be optimized if we know the number of lines in the file?
	var list []string

	// get the file
	readFile, err := os.Open(filePath)
	
	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// scan each line and pull out matching urls 
		link := findLink(fileScanner.Text())

		/* 
		we need to avoid the pattern "...undertheradarmag.com/reviews/category"
		but grab the pattern "...undertheradarmap.com/reviews/(anything else here)"
		golang regexp has no negative lookahead support, so we are manually checking for "category"
		TODO: refine regexp or look for another solution
		*/
		if strings.Contains(link, "category") {
			link = ""
		}
		if len(link) != 0 {
			list = append(list, link)
		}
	}
	return list
}

// findLink returns the first url in a string that matches the target pattern
func findLink(line string) string {
	reviewLink := regexp.MustCompile(`https://[a-z\.]*/reviews/[a-z\-_]*`)
	return reviewLink.FindString(line)
}

func getIndexAndPrint(url string) {
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
