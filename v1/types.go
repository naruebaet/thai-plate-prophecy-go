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
	FirstPart  FirstPartData  `json:"firstPart"`
	SecondPart SecondPartData `json:"secondPart"`
	Total      TotalData      `json:"total"`
}

type FirstPartData struct {
	Value string `json:"value"`
	Sum   int    `json:"sum"`
}

type SecondPartData struct {
	Value      string     `json:"value"`
	Sum        int        `json:"sum"`
	LuckyPoint LuckyPoint `json:"luckyPoint"`
}

type TotalData struct {
	Sum        int              `json:"sum"`
	LuckyGroup *LuckyPointGroup `json:"luckyGroup"`
}
