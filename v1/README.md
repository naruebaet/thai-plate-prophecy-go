# Thai Plate Prophecy Go

This Go package provides functionality to calculate numerological values and meanings for Thai license plates and dates. It includes validation, utility functions, and main calculation methods.

## Features

- **Validation**: Ensures the correctness of Thai license plate data.
- **Numerology Calculations**: Computes sums and lucky points for license plates.
- **Date Advice**: Provides lucky number advice based on a given date.

## Functions

### Validation Helpers
- `isValidThaiChar(char rune) bool`: Checks if a character is a valid Thai character.
- `isValidFirstChar(char rune) bool`: Validates the first character of a license plate.
- `isValidNumber(char rune) bool`: Checks if a character is a valid number.

### Utility Functions
- `getStringCount(s string) int`: Returns the number of characters in a string, handling UTF-8 properly.
- `getThaiCharValue(char rune) (int, error)`: Converts a Thai character to its numerical value.
- `getLuckyPoint(point int) (LuckyPoint, error)`: Retrieves the lucky point description for a given number.
- `getLuckyPointGroup(sum int) *LuckyPointGroup`: Determines the lucky group for a sum.

### Validation Functions
- `validateFirstData(firstData string) error`: Validates the first part of a Thai license plate.
- `validateSecondData(secondData string) error`: Validates the second part of a Thai license plate.

### Main Calculation Functions
- `AdviceByPlateData(firstData, secondData string) (PlateCalculationResult, error)`: Calculates the numerological value and meaning of a Thai license plate.
- `AdviceByDMY(date, month, year string) (LuckyNumberAdvice, error)`: Provides lucky number advice based on a given date.
- `AdviceByWeekDay(day int) (LuckyNumberAdvice, error)`: Retrieves advice for a specific day of the week.

## Usage

### License Plate Calculation
```go
result, err := AdviceByPlateData("1ก", "234")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Result: %+v\n", result)
}
```

### Date Advice
```go
advice, err := AdviceByDMY("15", "08", "2023")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Advice: %+v\n", advice)
}
```

## Data Structures

### PlateCalculationResult
Represents the result of a license plate calculation.

### LuckyPoint
Represents a lucky point description.

### LuckyPointGroup
Represents a group of lucky points.

### LuckyNumberAdvice
Represents advice based on a day of the week.

## Error Handling
All functions return detailed error messages for invalid inputs or missing data.

## License
This project is licensed under the MIT License.
