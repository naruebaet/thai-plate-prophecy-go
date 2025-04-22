package v1

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// VALIDATION HELPERS
func isValidThaiChar(char rune) bool {
	re := regexp.MustCompile(`^[ก-ฮ]$`)
	return re.MatchString(string(char))
}

func isValidFirstChar(char rune) bool {
	re := regexp.MustCompile(`^[0-9ก-ฮ]$`)
	return re.MatchString(string(char))
}

func isValidNumber(char rune) bool {
	re := regexp.MustCompile(`^[0-9]$`)
	return re.MatchString(string(char))
}

// UTILITY FUNCTIONS
// GetStringCount returns the number of characters in a string
// This properly handles UTF-8 characters including Thai script
func getStringCount(s string) int {
	return len([]rune(s))
}

// GetThaiCharValue converts a Thai character to its numerical value
func getThaiCharValue(char rune) (int, error) {
	if value, exists := thaiCharToNumberMap[char]; exists {
		return value, nil
	}

	if isValidNumber(char) {
		num, _ := strconv.Atoi(string(char))
		return num, nil
	}

	return 0, errors.New("invalid Thai character: " + string(char))
}

// GetLuckyPoint retrieves the lucky point description for a given number
func getLuckyPoint(point int) (LuckyPoint, error) {
	moduloPoint := point % 9
	if moduloPoint == 0 {
		moduloPoint = 9
	}

	for _, luckyPoint := range luckyPointStore {
		if luckyPoint.Point == moduloPoint {
			return luckyPoint, nil
		}
	}

	return LuckyPoint{}, errors.New("invalid point: " + strconv.Itoa(point))
}

// GetLuckyPointGroup determines which lucky group a sum belongs to
func getLuckyPointGroup(sum int) *LuckyPointGroup {
	for i, group := range luckyPointGroups {
		for _, p := range group.Points {
			if p == sum {
				return &luckyPointGroups[i]
			}
		}
	}
	return nil
}

// VALIDATION FUNCTIONS
// ValidateFirstData validates the first part of a Thai license plate
func validateFirstData(firstData string) error {

	if getStringCount(firstData) > 3 {
		return errors.New("invalid first part of license plate: too long")
	}

	runes := []rune(firstData)
	if !isValidFirstChar(runes[0]) {
		return errors.New("invalid first character in license plate")
	}

	for i := 1; i < len(runes); i++ {
		if !isValidThaiChar(runes[i]) {
			return errors.New("invalid Thai character in license plate")
		}
	}

	return nil
}

// ValidateSecondData validates the second part of a Thai license plate
func validateSecondData(secondData string) error {
	matched, _ := regexp.MatchString(`^\d{1,4}$`, secondData)
	if !matched {
		return errors.New("invalid second part of license plate")
	}
	return nil
}

// MAIN CALCULATION FUNCTION
// AdviceByPlateData calculates the numerological value and meaning of a Thai license plate.
// It takes two string inputs: `firstData` (the first part of the license plate) and `secondData` (the numeric part).
// Returns a PlateCalculationResult containing the calculated values and meanings, or an error if validation fails.
func AdviceByPlateData(firstData, secondData string) (PlateCalculationResult, error) {
	var result PlateCalculationResult

	// Validate inputs
	if err := validateFirstData(firstData); err != nil {
		return result, err
	}

	if err := validateSecondData(secondData); err != nil {
		return result, err
	}

	// Calculate sums
	firstDataSum := 0
	for _, char := range firstData {
		value, err := getThaiCharValue(char)
		if err != nil {
			return result, err
		}
		firstDataSum += value
	}

	secondDataSum := 0
	for _, digit := range secondData {
		num, _ := strconv.Atoi(string(digit))
		secondDataSum += num
	}

	totalSum := firstDataSum + secondDataSum

	// Handle double-digit secondDataSum
	realSecondDataSum := secondDataSum
	if realSecondDataSum > 9 {
		tempSum := 0
		for _, digit := range strconv.Itoa(secondDataSum) {
			num, _ := strconv.Atoi(string(digit))
			tempSum += num
		}
		realSecondDataSum = tempSum
	}

	// Get meanings
	secondDataLuckyPoint, err := getLuckyPoint(realSecondDataSum)
	if err != nil {
		return result, err
	}

	totalLuckyGroup := getLuckyPointGroup(totalSum)

	// Build result
	result.FirstPart.Value = firstData
	result.FirstPart.Sum = firstDataSum

	result.SecondPart.Value = secondData
	result.SecondPart.Sum = realSecondDataSum
	result.SecondPart.LuckyPoint = secondDataLuckyPoint

	result.Total.Sum = totalSum
	result.Total.LuckyGroup = totalLuckyGroup

	return result, nil
}

// AdviceByDMY provides lucky number advice based on a given date, month, and year.
// It takes three string inputs: `date`, `month`, and `year`.
// Returns a LuckyNumberAdvice containing the advice or an error if the date format is invalid or no advice is found.
func AdviceByDMY(date, month, year string) (LuckyNumberAdvice, error) {

	inputFormat := fmt.Sprintf("%s-%s-%s", year, month, date)

	// format date from input
	t, err := time.Parse("2006-01-02", inputFormat)
	if err != nil {
		return LuckyNumberAdvice{}, errors.New("invalid date format")
	}

	// get day from date
	dayInt := WeekDay(int(t.Weekday()))

	// Find luckyNumberAdvice by Day
	data, err := AdviceByWeekDay(dayInt)
	if err != nil {
		return LuckyNumberAdvice{}, errors.New("no advice found for the given day")
	}

	return data, nil
}

// AdviceByWeekDay provides lucky number advice based on the day of the week.
// It takes an integer input `day` representing the day of the week (0 for Sunday, 1 for Monday, etc.).
// Returns a LuckyNumberAdvice containing the advice or an error if no advice is found for the given day.
func AdviceByWeekDay(day WeekDay) (LuckyNumberAdvice, error) {
	for _, advice := range luckyNumberAdvice {
		if advice.Day == day.Int() {
			return advice, nil
		}
	}
	return LuckyNumberAdvice{}, errors.New("no advice found for the given day")
}
