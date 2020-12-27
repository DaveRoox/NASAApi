package nasaapi

import (
	"net/url"
	"strconv"
	"time"
)

// NeoFeed represents an object returned by a valid NeoFeed request
type NeoFeed struct {
	Links            links                        `json:"links"`
	ElementCount     int                          `json:"element_count"`
	NearEarthObjects map[string][]nearEarthObject `json:"near_earth_objects"`
}

// NeoBrowse represents an object returned by a valid NeoBrowse request
type NeoBrowse struct {
	Links            links             `json:"links"`
	Page             page              `json:"page"`
	NearEarthObjects []nearEarthObject `json:"near_earth_objects"`
}

// NeoLookup represents an object returned by a valid NeoLookup request
type NeoLookup struct {
	Links                          links               `json:"links"`
	ID                             string              `json:"id"`
	NeoReferenceID                 string              `json:"neo_reference_id"`
	Name                           string              `json:"name"`
	Designation                    string              `json:"designation"`
	NasaJplURL                     string              `json:"nasa_jpl_url"`
	AbsoluteMagnitudeH             float64             `json:"absolute_magnitude_h"`
	EstimatedDiameter              estimatedDiameter   `json:"estimated_diameter"`
	IsPotentiallyHazardousAsteroid bool                `json:"is_potentially_hazardous_asteroid"`
	CloseApproachData              []closeApproachData `json:"close_approach_data"`
	OrbitalData                    orbitalData         `json:"orbital_data"`
	IsSentryObject                 bool                `json:"is_sentry_object"`
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

type orbitalData struct {
	OrbitID                   string     `json:"orbit_id"`
	OrbitDeterminationDate    string     `json:"orbit_determination_date"`
	FirstObservationDate      string     `json:"first_observation_date"`
	LastObservationDate       string     `json:"last_observation_date"`
	DataArcInDays             int        `json:"data_arc_in_days"`
	ObservationsUsed          int        `json:"observations_used"`
	OrbitUncertainty          string     `json:"orbit_uncertainty"`
	MinimumOrbitIntersection  string     `json:"minimum_orbit_intersection"`
	JupiterTisserandInvariant string     `json:"jupiter_tisserand_invariant"`
	EpochOsculation           string     `json:"epoch_osculation"`
	Eccentricity              string     `json:"eccentricity"`
	SemiMajorAxis             string     `json:"semi_major_axis"`
	Inclination               string     `json:"inclination"`
	AscendingNodeLongitude    string     `json:"ascending_node_longitude"`
	OrbitalPeriod             string     `json:"orbital_period"`
	PerihelionDistance        string     `json:"perihelion_distance"`
	PerihelionArgument        string     `json:"perihelion_argument"`
	AphelionDistance          string     `json:"aphelion_distance"`
	PerihelionTime            string     `json:"perihelion_time"`
	MeanAnomaly               string     `json:"mean_anomaly"`
	MeanMotion                string     `json:"mean_motion"`
	Equinox                   string     `json:"equinox"`
	OrbitClass                orbitClass `json:"orbit_class"`
}

type orbitClass struct {
	OrbitClassType        string `json:"orbit_class_type"`
	OrbitClassDescription string `json:"orbit_class_description"`
	OrbitClassRange       string `json:"orbit_class_range"`
}

type page struct {
	Size          int `json:"size"`
	TotalElements int `json:"total_elements"`
	TotalPages    int `json:"total_pages"`
	Number        int `json:"number"`
}

func (f *fetcher) NeoFeed(startDate, endDate time.Time) (*NeoFeed, error) {

	u := f.buildURL(
		"neo/rest/v1/feed",
		map[string]string{
			"start_date": startDate.Format(dateFormat),
			"end_date":   endDate.Format(dateFormat),
		},
	)
	return getNeoFeed(u)
}

func (f *fetcher) NeoLookup(asteroidID int64) (*NeoLookup, error) {
	u := f.buildURL(
		"neo/rest/v1/neo/"+strconv.FormatInt(asteroidID, 10),
		nil,
	)
	return getNeoLookup(u)
}

func (f *fetcher) NeoBrowse() (*NeoBrowse, error) {
	u := f.buildURL(
		"neo/rest/v1/neo/browse",
		nil,
	)
	return getNeoBrowse(u)
}

func getNeoFeed(u *url.URL) (*NeoFeed, error) {
	n := &NeoFeed{}
	if err := getAndParse(u.String(), n); err != nil {
		return nil, err
	}
	return n, nil
}

func getNeoLookup(u *url.URL) (*NeoLookup, error) {
	n := &NeoLookup{}
	if err := getAndParse(u.String(), n); err != nil {
		return nil, err
	}
	return n, nil
}

func getNeoBrowse(u *url.URL) (*NeoBrowse, error) {
	n := &NeoBrowse{}
	if err := getAndParse(u.String(), n); err != nil {
		return nil, err
	}
	return n, nil
}
