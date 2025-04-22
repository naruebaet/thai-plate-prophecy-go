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

type LuckyNumberAdvice struct {
	Day           int      `json:"day"`
	LuckyNumDesc  string   `json:"lucky_num_desc"`
	LuckyNum      []int    `json:"lucky_num"`
	AvoidNumDesc  string   `json:"avoid_num_desc"`
	AvoidNum      []int    `json:"avoid_num"`
	AvoidCharDesc string   `json:"avoid_char_desc"`
	AvoidChar     []string `json:"avoid_char"`
}
