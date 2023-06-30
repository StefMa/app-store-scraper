# app-store-scraper
scrape data from the itunes app store

Go scraper to get data from the iTunes/Mac App Store via the [iTunes Search API](https://performance-partners.apple.com/search-api.
The goal is to provide an interface as close as possible to the [app-store-scraper Node.js module](https://github.com/facundoolano/app-store-scraper).

## Installation

```
go get -u github.com/StefMa/app-store-scraper/
```

## Usage

### App 
Retrieves the full details of an application.

```go
import "github.com/StefMa/app-store-scraper/scraper"

func main() {
	options := scraper.Options{
		Country:  "[COUNTRY_CODE]",
		Language: "[LANGUAGE]",
	}
	result, err := scraper.App("[APP_ID]", options)
	if err != nil {
		// Handle error
	}
	// Do something with the result
}
```

[Full working example at the Go Playground](https://go.dev/play/p/-v1uigDt_fu).

### Developer 
Retrieves a list of apps by the given developer id.

```go
import "github.com/StefMa/app-store-scraper/scraper"

func main() {
	options := scraper.Options{
		Country:  "[COUNTRY_CODE]",
		Language: "[LANGUAGE]",
		Limit:    [LIMIT],
	}
	results, err := scraper.Developer("[DEVELOPER_ID]", options)
	if err != nil {
		// Handle error
	}
	// Do something with the result
}
```

[Full working example at the Go Playground](https://go.dev/play/p/HNtxOYEq6N8).
