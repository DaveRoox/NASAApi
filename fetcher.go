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
	Apod(time.Time, bool) (*Apod, error)
	ApodDefault() (*Apod, error)
}

// GetFetcher takes an api key and returns a Fetcher object that stores is
func GetFetcher(apiKey string) Fetcher {
	return &fetcher{
		key: apiKey,
	}
}

// fetcher is the underlying implementation of the Fetcher interface
// The users don't use this directly
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
