package scraper

import (
	"strconv"
	"strings"
	"testing"
)

func TestAppSucess(t *testing.T) {
	options := Options{
		Language: "de",
		Country:  "de",
	}
	got, err := App("1400408720", options)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if got.TrackName != "hvv hop" {
		t.Errorf("Got trackName %s, but want %s", got.TrackName, "hvv Hop")
	}
	if got.Version == "" {
		t.Errorf("Got Version %s, but want %s", got.Version, "hvv Hop")
	} else {
		version := strings.Split(got.Version, ".")
		major := version[0]
		majorNum, err := strconv.Atoi(major)
		if err != nil || majorNum < 3 {
			t.Errorf("Got Version %s, but wanted at least major version %d", got.Version, 3)
		}
	}
	var minimumWantedRating float64 = 4.7
	if got.AverageUserRating < minimumWantedRating {
		t.Errorf("Got Rating of %f, but wanted %f or higher", got.AverageUserRating, minimumWantedRating)
	}
	if got.TrackViewURL != "https://apps.apple.com/de/app/hvv-hop/id1400408720?uo=4" {
		t.Errorf("Got trackViewUrl %s, but want %s", got.TrackViewURL, "https://apps.apple.com/us/app/hvv-hop/id1400408720?uo=4")
	}
	if got.ArtworkURL512 != "https://is2-ssl.mzstatic.com/image/thumb/Purple116/v4/93/bc/62/93bc6267-e6be-d901-ae71-d137678e5902/AppIcon-0-1x_U007ephone-0-85-220.png/512x512bb.jpg" {
		t.Errorf("Got trackName %s, but want %s", got.ArtworkURL512, "https://is2-ssl.mzstatic.com/image/thumb/Purple116/v4/93/bc/62/93bc6267-e6be-d901-ae71-d137678e5902/AppIcon-0-1x_U007ephone-0-85-220.png/512x512bb.jpg")
	}
}
