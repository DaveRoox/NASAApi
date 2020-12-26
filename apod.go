package nasaapi

import (
	"net/url"
	"strconv"
	"time"
)

const dateFormat = "2006-01-02"

// Apod represents an object returned from a valid APOD request
type Apod struct {
	Copyright      string
	Date           time.Time //YYYY-MM-DD
	Explanation    string
	HDURL          *url.URL
	MediaType      string
	ServiceVersion string
	Title          string
	URL            *url.URL
}

func (f *fetcher) ApodDefault() (*Apod, error) {
	return f.Apod(time.Now(), false)
}

func (f *fetcher) Apod(date time.Time, hd bool) (*Apod, error) {

	u := f.buildURL(
		"planetary/apod",
		map[string]string{
			"date": date.Format(dateFormat),
			"hd":   strconv.FormatBool(hd),
		},
	)

	a := &struct {
		Copyright      string `json:"copyright"`
		Date           string `json:"date"`
		Explanation    string `json:"explanation"`
		HDURL          string `json:"hdurl"`
		MediaType      string `json:"media_type"`
		ServiceVersion string `json:"service_version"`
		Title          string `json:"title"`
		URL            string `json:"url"`
	}{}

	if err := getAndParse(u, a); err != nil {
		return nil, err
	}

	d, err := time.Parse(dateFormat, a.Date)
	if err != nil {
		return nil, err
	}

	hdurl, err := url.Parse(a.HDURL)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(a.URL)
	if err != nil {
		return nil, err
	}

	return &Apod{
		Copyright:      a.Copyright,
		Date:           d,
		Explanation:    a.Explanation,
		HDURL:          hdurl,
		MediaType:      a.MediaType,
		ServiceVersion: a.ServiceVersion,
		Title:          a.Title,
		URL:            url,
	}, nil
}
