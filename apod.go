package nasaapi

import (
	"strconv"
	"time"
)

// Apod represents an object returned from a valid APOD request
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

func (f *fetcher) ApodDefault() (*Apod, error) {
	return f.Apod(time.Now(), false)
}

func (f *fetcher) Apod(day time.Time, hd bool) (*Apod, error) {

	u := f.buildURL(
		"planetary/apod",
		map[string]string{
			"date": day.Format(dateFormat),
			"hd":   strconv.FormatBool(hd),
		},
	)

	a := &Apod{}
	if err := getAndParse(u, a); err != nil {
		return nil, err
	}

	return a, nil
}
