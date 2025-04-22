package main

import (
	"encoding/json"
	"fmt"

	tpp "github.com/naruebaet/thai-plate-prophecy-go/v1"
)

func main() {
	// Example 1: Using AdviceByPlateData
	firstData := "กข"
	secondData := "1234"
	plateResult, err := tpp.AdviceByPlateData(firstData, secondData)
	if err != nil {
		fmt.Println("Error in AdviceByPlateData:", err)
	} else {
		plateResultJSON, _ := json.MarshalIndent(plateResult, "", "  ")
		fmt.Println("AdviceByPlateData Result:")
		fmt.Println(string(plateResultJSON))
	}

	// Example 2: Using AdviceByDMY
	date := "01"
	month := "01"
	year := "2023"
	dmyResult, err := tpp.AdviceByDMY(date, month, year)
	if err != nil {
		fmt.Println("Error in AdviceByDMY:", err)
	} else {
		dmyResultJSON, _ := json.MarshalIndent(dmyResult, "", "  ")
		fmt.Println("AdviceByDMY Result:")
		fmt.Println(string(dmyResultJSON))
	}

	// Example 3: Using AdviceByWeekDay
	day := tpp.WeekDay(1) // 1 for Monday
	weekdayResult, err := tpp.AdviceByWeekDay(day)
	if err != nil {
		fmt.Println("Error in AdviceByWeekDay:", err)
	} else {
		weekdayResultJSON, _ := json.MarshalIndent(weekdayResult, "", "  ")
		fmt.Println("AdviceByWeekDay Result:")
		fmt.Println(string(weekdayResultJSON))
	}
}
