package main

import (
	"reflect"
	"testing"
)

var RegexpData = []struct {
	raw   string
	match string
}{
	{`<a href="https://www.undertheradarmag.com/reviews/tonight_franz_ferdinand"><img class="dominant" src="https://undertheradarmag.com/uploads/review_images/FranzFerdinandTonight.jpg" width="150" alt="" /></a>`, `https://www.undertheradarmag.com/reviews/tonight_franz_ferdinand`},
	{`<a href="https://www.undertheradarmag.com/reviews/intimacy/"><img src="https://undertheradarmag.com/uploads/review_images/BlocPartyIntimacy.jpg" width="428"/></a>`, `https://www.undertheradarmag.com/reviews/intimacy`},
}

func TestMatchRegex(t *testing.T) {
	for _, entry := range RegexpData {
		got := findLink(entry.raw)
		want := entry.match

		if len(got) == 0 {
			got = "no match found"
		}

		if got != want {
			t.Errorf("for %s, expected %s but got %s", entry.raw, want, got)
		}
	}
}

func TestReadIndex(t *testing.T) {
	got := readIndex("./example-html/test-index.html")
	// want := []string{"https://www.undertheradarmag.com/reviews/fantasy_black_channel", "https://www.undertheradarmag.com/reviews/the_finally_lp"}
	want := map[string]bool{
		"https://www.undertheradarmag.com/reviews/fantasy_black_channel": true,
		"https://www.undertheradarmag.com/reviews/the_finally_lp":        true,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReadRating(t *testing.T) {
	got := parseReview("./example-html/test-review.html")
	want := Ratings{group: "Gorillaz", albumTitle: "Cracker Island", readerRating: 4, authorRating: 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
