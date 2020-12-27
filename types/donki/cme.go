package donki

// CoronalMassEjections represent an object returned by a valid CME request
type CoronalMassEjections []coronalMassEjection

type coronalMassEjection struct {
	ActivityID      string        `json:"activityID"`
	Catalog         string        `json:"catalog"`
	StartTime       string        `json:"startTime"`
	SourceLocation  string        `json:"sourceLocation"`
	ActiveRegionNum *int          `json:"activeRegionNum"`
	Link            string        `json:"link"`
	Note            string        `json:"note"`
	Instruments     []instrument  `json:"instruments"`
	CMEAnalyses     []cmeAnalysis `json:"cmeAnalyses"`
	LinkedEvents    []event       `json:"linkedEvents"`
}

type instrument struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
}

type cmeAnalysis struct {
	Time21_5       string  `json:"time21_5"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	HalfAngle      float64 `json:"halfAngle"`
	Speed          float64 `json:"speed"`
	Type           string  `json:"type"`
	IsMostAccurate bool    `json:"isMostAccurate"`
	Note           string  `json:"note"`
	LevelOfData    int     `json:"levelOfData"`
	Link           string  `json:"link"`
	EnlilList      []enlil `json:"enlilList"`
}

type enlil struct {
	ModelCompletionTime       string   `json:"modelCompletionTime"`
	Au                        float64  `json:"au"`
	EstimatedShockArrivalTime *string  `json:"estimatedShockArrivalTime"`
	EstimatedDuration         *float64 `json:"estimatedDuration"`
	RminRe                    *float64 `json:"rmin_re"`
	KP18                      *int     `json:"kp_18"`
	KP90                      *int     `json:"kp_90"`
	KP135                     *int     `json:"kp_135"`
	KP180                     *int     `json:"kp_180"`
	IsEarthGB                 bool     `json:"isEarthGB"`
	Link                      string   `json:"link"`
	ImpactList                []impact `json:"impactList"`
	CMEIDs                    []string `json:"cmeIDs"`
}

type impact struct {
	IsGlancingBlow bool   `json:"isGlancingBlow"`
	Location       string `json:"location"`
	ArrivalTime    string `json:"arrivalTime"`
}

type event struct {
	ActivityID string `json:"activityID"`
}
