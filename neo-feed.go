package nasaapi

import (
	"time"
)

// NeoFeed represents an object returned by a valid NeoFeed request
type NeoFeed struct {
	Links            links                        `json:"links"`
	ElementCount     int                          `json:"element_count"`
	NearEarthObjects map[string][]nearEarthObject `json:"near_earth_objects"`
}

type links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
	Self string `json:"self"`
}

type nearEarthObject struct {
	Links                          links               `json:"links"`
	ID                             string              `json:"id"`
	NeoReferenceID                 string              `json:"neo_reference_id"`
	Name                           string              `json:"name"`
	NasaJplURL                     string              `json:"nasa_jpl_url"`
	AbsoluteMagnitudeH             float64             `json:"absolute_magnitude_h"`
	EstimatedDiameter              estimatedDiameter   `json:"estimated_diameter"`
	IsPotentiallyHazardousAsteroid bool                `json:"is_potentially_hazardous_asteroid"`
	CloseApproachData              []closeApproachData `json:"close_approach_data"`
	IsSentryObject                 bool                `json:"is_sentry_object"`
}

type estimatedDiameter struct {
	Kilometers diameterEstimation `json:"kilometers"`
	Meters     diameterEstimation `json:"meters"`
	Miles      diameterEstimation `json:"miles"`
	Feet       diameterEstimation `json:"feet"`
}

type diameterEstimation struct {
	EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
	EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
}

type closeApproachData struct {
	CloseApproachDate      string           `json:"close_approach_date"`
	CloseApproachDateFull  string           `json:"close_approach_date_full"`
	EpochDateCloseApproach int64            `json:"epoch_date_close_approach"`
	RelativeVelocity       relativeVelocity `json:"relative_velocity"`
	MissDistance           missDistance     `json:"miss_distance"`
	OrbitingBody           string           `json:"orbiting_body"`
}

type relativeVelocity struct {
	KilometersPerSecond string `json:"kilometers_per_second"`
	KilometersPerHour   string `json:"kilometers_per_hour"`
	MilesPerHour        string `json:"miles_per_hour"`
}

type missDistance struct {
	Astronomical string `json:"astronomical"`
	Lunar        string `json:"lunar"`
	Kilometers   string `json:"kilometers"`
	Miles        string `json:"miles"`
}

func (f *fetcher) NeoFeed(startDate, endDate time.Time) (*NeoFeed, error) {

	u := f.buildURL(
		"neo/rest/v1/feed",
		map[string]string{
			"start_date": startDate.Format(dateFormat),
			"end_date":   endDate.Format(dateFormat),
		},
	)

	a := &NeoFeed{}
	if err := getAndParse(u, a); err != nil {
		return nil, err
	}

	return a, nil
}
