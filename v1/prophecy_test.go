package v1

import (
	"reflect"
	"testing"
)

func TestCalculatePlateData(t *testing.T) {
	// Define a helper function to create expected PlateCalculationResult
	createExpectedResult := func(
		firstValue string, firstSum int,
		secondValue string, secondSum int, secondPoint LuckyPoint,
		totalSum int, luckyGroup *LuckyPointGroup,
	) PlateCalculationResult {
		return PlateCalculationResult{
			FirstPart: FirstPartData{
				Value: firstValue,
				Sum:   firstSum,
			},
			SecondPart: SecondPartData{
				Value:      secondValue,
				Sum:        secondSum,
				LuckyPoint: secondPoint,
			},
			Total: TotalData{
				Sum:        totalSum,
				LuckyGroup: luckyGroup,
			},
		}
	}

	// Setup some test LuckyPoint and LuckyPointGroup for expected results
	// You might need to adapt this based on your actual data structures
	point1 := LuckyPoint{Point: 1}
	point5 := LuckyPoint{Point: 5}
	point9 := LuckyPoint{Point: 9}

	// Mock the luckyPointStore for testing
	// This assumes luckyPointStore is exported or you have accessor functions
	originalLuckyPointStore := luckyPointStore
	luckyPointStore = []LuckyPoint{point1, point5, point9}
	defer func() { luckyPointStore = originalLuckyPointStore }()

	// Mock luckyPointGroups
	group1 := LuckyPointGroup{Points: []int{10, 19, 28}}
	group2 := LuckyPointGroup{Points: []int{13, 20, 29}}
	originalLuckyPointGroups := luckyPointGroups
	luckyPointGroups = []LuckyPointGroup{group1, group2}
	defer func() { luckyPointGroups = originalLuckyPointGroups }()

	tests := []struct {
		name        string
		firstData   string
		secondData  string
		want        PlateCalculationResult
		expectError bool
	}{
		{
			name:       "Valid Thai characters and numbers",
			firstData:  "กข",
			secondData: "1234",
			want: createExpectedResult(
				"กข", 3, // firstValue, firstSum
				"1234", 1, point1, // secondValue, secondSum, secondPoint (assuming 1+2+3+4=10, 1+0=1)
				13, &group2, // totalSum, luckyGroup
			),
			expectError: false,
		},
		{
			name:        "Invalid first part - too long",
			firstData:   "กขคง",
			secondData:  "1234",
			want:        PlateCalculationResult{},
			expectError: true,
		},
		{
			name:        "Invalid first part - invalid character",
			firstData:   "A",
			secondData:  "1234",
			want:        PlateCalculationResult{},
			expectError: true,
		},
		{
			name:        "Invalid second part - non-numeric",
			firstData:   "กข",
			secondData:  "123A",
			want:        PlateCalculationResult{},
			expectError: true,
		},
		{
			name:        "Invalid second part - too long",
			firstData:   "กข",
			secondData:  "12345",
			want:        PlateCalculationResult{},
			expectError: true,
		},
		{
			name:       "Numeric first part",
			firstData:  "1",
			secondData: "234",
			want: createExpectedResult(
				"1", 1, // firstValue, firstSum
				"234", 9, point9, // secondValue, secondSum, secondPoint (2+3+4=9)
				1+2+3+4, &group1, // totalSum, luckyGroup
			),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculatePlateData(tt.firstData, tt.secondData)

			// Check error expectations
			if (err != nil) != tt.expectError {
				t.Errorf("CalculatePlateData() error = %v, expectError %v", err, tt.expectError)
				return
			}

			// If expecting an error, don't compare results
			if tt.expectError {
				return
			}

			// Compare results
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePlateData() = %v, want %v", got, tt.want)
			}
		})
	}
}
