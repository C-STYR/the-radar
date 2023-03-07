package main

type UrlList struct {
	Urls chan ReviewUrl
}

type ReviewUrl string

// CreateIndexUrl returns a url string typed as IndexUrl
func CreateReviewUrl(link string) ReviewUrl {
	return ReviewUrl(link)
}

// Enqueue sends a ReviewUrl down the Urls channel
func (u *UrlList) Enqueue(link ReviewUrl) {
	u.Urls <- link
}

// CreateURLChannel creates a UrlList struct with a Urls channel with a 5000 cap
func CreateURLChannel() UrlList {
	return UrlList{make(chan ReviewUrl, 5000)}
}
