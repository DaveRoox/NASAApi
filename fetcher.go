package nasaapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type fetcher struct {
	key string
}

const dateFormat = "2006-01-02"

func Fetcher(apiKey string) *fetcher {
	return &fetcher{
		key: apiKey,
	}
}

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
		Date           string `json:"date" time_format:"2006-01-02"`
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

func (f *fetcher) buildURL(path string, params map[string]string) *url.URL {
	u := &url.URL{
		Scheme: "https",
		Host:   "api.nasa.gov",
	}
	u.Path = path
	q := u.Query()
	q.Set("api_key", f.key)
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u
}

func getAndParse(u *url.URL, a interface{}) error {
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http GET status: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(a)
}
