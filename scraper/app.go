package scraper

import (
	"fmt"
	"net/url"
	"strconv"
)

func App(appId string, options Options) (Result, error) {
	values := url.Values{}
	values.Add("id", appId)
	if options.Country != "" {
		values.Add("country", options.Country)
	}
	if options.Language != "" {
		values.Add("lang", options.Language)
	}
	values.Add("limit", strconv.Itoa(options.Limit))
	results, err := execute(values.Encode())
	if err != nil {
		return Result{}, err
	}
	for _, result := range results {
		if result.WrapperType == "software" {
			return result, nil
		}
	}
	return Result{}, fmt.Errorf("Coun't find app with id '%s'", appId)
}
