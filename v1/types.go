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

type WeekDay int

func (w WeekDay) String() string {
	switch w {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	default:
		return "Unknown"
	}
}

func (w WeekDay) Int() int {
	return int(w)
}
