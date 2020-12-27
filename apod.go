package nasaapi

import (
	"net/url"
	"strconv"
	"time"
)

// Apod represents an object returned by a valid APOD request
type Apod struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HDURL          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

func (f *fetcher) Apod(date time.Time, hd bool) (*Apod, error) {

	u := f.buildURL(
		"planetary/apod",
		map[string]string{
			"date": date.Format(dateFormat),
			"hd":   strconv.FormatBool(hd),
		},
	)

	return getApod(u)
}

func getApod(u *url.URL) (*Apod, error) {
	a := &Apod{}
	if err := getAndParse(u.String(), a); err != nil {
		return nil, err
	}
	return a, nil
}
