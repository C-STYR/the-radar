package main

import (
	"reflect"
	"regexp"
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

		reviewLink := regexp.MustCompile(`https://[a-z\.]*/reviews/[a-z\-_]*`)
		got := reviewLink.FindString(entry.raw)
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

	// got := ReadIndex("./example-html/test-index.html")
	// want := []string{"https://www.undertheradarmag.com/reviews/fantasy_black_channel", "https://www.undertheradarmag.com/reviews/the_finally_lp"}

}

func TestParseReview(t *testing.T) {

	t.Run("correctly returns a nil value when ratings threshholds are unmet", func(t *testing.T) {

		// in this case, want is the nil value of a pointer to Ratings struct
		var want *Ratings
		got := ParseReview("./example-html/test-review1.html")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("returns pointer to Ratings struct when ratings thresholds are met", func(t *testing.T) {

		got := ParseReview("./example-html/test-review2.html")
		want := &Ratings{group: "Gorillaz", albumTitle: "Cracker Island", readerRating: 4, authorRating: 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
