package v1

// TYPES
type LuckyPoint struct {
	Point int    `json:"point"`
	Desc  string `json:"desc"`
}

type LuckyPointGroup struct {
	Group  string `json:"group"`
	Points []int  `json:"points"`
	Desc   string `json:"desc"`
}

type PlateCalculationResult struct {
	FirstPart struct {
		Value string `json:"value"`
		Sum   int    `json:"sum"`
	} `json:"firstPart"`
	SecondPart struct {
		Value      string     `json:"value"`
		Sum        int        `json:"sum"`
		LuckyPoint LuckyPoint `json:"luckyPoint"`
	} `json:"secondPart"`
	Total struct {
		Sum        int              `json:"sum"`
		LuckyGroup *LuckyPointGroup `json:"luckyGroup"`
	} `json:"total"`
}
