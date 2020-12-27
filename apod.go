package nasaapi

import (
	"nasaapi/types"
	"net/url"
	"strconv"
	"time"
)

func (f *fetcher) Apod(date time.Time, hd bool) (*types.Apod, error) {

	u := f.buildURL(
		"planetary/apod",
		map[string]string{
			"date": date.Format(dateFormat),
			"hd":   strconv.FormatBool(hd),
		},
	)

	return getApod(u)
}

func getApod(u *url.URL) (*types.Apod, error) {
	a := &types.Apod{}
	if err := getAndParse(u.String(), a); err != nil {
		return nil, err
	}
	return a, nil
}
