package nasaapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Fetcher represent an interface that exposes methods to access NASA's public APIs
type Fetcher interface {
	Apod(date time.Time, hd bool) (*Apod, error)
	ApodDefault() (*Apod, error)
	NeoFeed(startDate, endDate time.Time) (*NeoFeed, error)
	NeoLookup(asteroidID int64) (*NeoLookup, error)
	NeoBrowse() (*NeoBrowse, error)
}

// GetFetcher takes an api key and returns a Fetcher object
func GetFetcher(apiKey string) Fetcher {
	return &fetcher{
		key: apiKey,
	}
}

// fetcher is the underlying implementation of the Fetcher interface
type fetcher struct {
	key string
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

func getAndParse(u string, a interface{}) error {
	fmt.Printf("Requesting at %v\n", u)
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http GET status: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(a)
}

const dateFormat = "2006-01-02"
