package nasaapi

import (
	donki "nasaapi/types/donki"
	"net/url"
	"strconv"
	"time"
)

// Catalog is the type for the CMEAnalyses' `catalog` parameter
type catalog string

// Enum values for CMEAnalyses' `catalog` parameters
const (
	All             = catalog("ALL")
	SWRCCatalog     = catalog("SWRC_CATALOG")
	JangEtAlCatalog = catalog("JANG_ET_AL_CATALOG")
)

func (f *fetcher) CoronalMassEjections(startDate, endDate time.Time) (*donki.CoronalMassEjections, error) {

	u := f.buildURL(
		"DONKI/CME",
		map[string]string{
			"startDate": startDate.Format(dateFormat),
			"endDate":   endDate.Format(dateFormat),
		},
	)

	return getCME(u)
}

func getCME(u *url.URL) (*donki.CoronalMassEjections, error) {
	a := &donki.CoronalMassEjections{}
	if err := getAndParse(u.String(), a); err != nil {
		return nil, err
	}
	return a, nil
}

func (f *fetcher) CoronalMassEjectionsAnalyses(
	startDate, endDate time.Time,
	mostAccurateOnly, completeEntryOnly bool,
	lowerSpeed, lowerHalfAngle int64,
	catalog catalog,
	keyword string) (*donki.CoronalMassEjectionsAnalyses, error) {

	u := f.buildURL(
		"DONKI/CMEAnalysis",
		map[string]string{
			"startDate":         startDate.Format(dateFormat),
			"endDate":           endDate.Format(dateFormat),
			"mostAccurateOnly":  strconv.FormatBool(mostAccurateOnly),
			"completeEntryOnly": strconv.FormatBool(completeEntryOnly),
			"speed":             strconv.FormatInt(lowerSpeed, 10),
			"halfAngle":         strconv.FormatInt(lowerHalfAngle, 10),
			"catalog":           string(catalog),
			"keyword":           keyword,
		},
	)

	return getCMEAnalyses(u)
}

func getCMEAnalyses(u *url.URL) (*donki.CoronalMassEjectionsAnalyses, error) {
	a := &donki.CoronalMassEjectionsAnalyses{}
	if err := getAndParse(u.String(), a); err != nil {
		return nil, err
	}
	return a, nil
}
