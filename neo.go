package nasaapi

import (
	"nasaapi/types"
	"net/url"
	"strconv"
	"time"
)

func (f *fetcher) NeoFeed(startDate, endDate time.Time) (*types.NeoFeed, error) {

	u := f.buildURL(
		"neo/rest/v1/feed",
		map[string]string{
			"start_date": startDate.Format(dateFormat),
			"end_date":   endDate.Format(dateFormat),
		},
	)
	return getNeoFeed(u)
}

func getNeoFeed(u *url.URL) (*types.NeoFeed, error) {
	n := &types.NeoFeed{}
	if err := getAndParse(u.String(), n); err != nil {
		return nil, err
	}
	return n, nil
}

func (f *fetcher) NeoLookup(asteroidID int64) (*types.NeoLookup, error) {
	u := f.buildURL(
		"neo/rest/v1/neo/"+strconv.FormatInt(asteroidID, 10),
		nil,
	)
	return getNeoLookup(u)
}

func getNeoLookup(u *url.URL) (*types.NeoLookup, error) {
	n := &types.NeoLookup{}
	if err := getAndParse(u.String(), n); err != nil {
		return nil, err
	}
	return n, nil
}

func (f *fetcher) NeoBrowse() (*types.NeoBrowse, error) {
	u := f.buildURL(
		"neo/rest/v1/neo/browse",
		nil,
	)
	return getNeoBrowse(u)
}

func getNeoBrowse(u *url.URL) (*types.NeoBrowse, error) {
	n := &types.NeoBrowse{}
	if err := getAndParse(u.String(), n); err != nil {
		return nil, err
	}
	return n, nil
}
