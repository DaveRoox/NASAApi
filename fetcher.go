package nasaapi

import (
	"encoding/json"
	"fmt"
	"nasaapi/types"
	donki "nasaapi/types/donki"
	"net/http"
	"net/url"
	"time"
)

// Fetcher represent an interface that exposes methods to access NASA's public APIs
type Fetcher interface {
	Apod(date time.Time, hd bool) (*types.Apod, error)
	NeoFeed(startDate, endDate time.Time) (*types.NeoFeed, error)
	NeoLookup(asteroidID int64) (*types.NeoLookup, error)
	NeoBrowse() (*types.NeoBrowse, error)
	CoronalMassEjections(startDate, endDate time.Time) (*donki.CoronalMassEjections, error)
	CoronalMassEjectionsAnalyses(
		startDate, endDate time.Time,
		mostAccurateOnly, completeEntryOnly bool,
		lowerSpeed, lowerHalfAngle int64,
		catalog donki.Catalog,
		keyword string) (*donki.CoronalMassEjectionsAnalyses, error)
}

// New takes an api key and returns a newly created Fetcher object binded to the key
func New(apiKey string) Fetcher {
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
	fmt.Printf("Request at %v\n", u)
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
