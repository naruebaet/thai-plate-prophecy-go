# Thai plate prophecy GO(TH-Version)

This project is designed to help customers predict or "prophecy" potential license plate numbers. By leveraging algorithms and user input, it provides insights and suggestions for plate numbers that may hold significance or align with user preferences.  

โครงการนี้ถูกออกแบบมาเพื่อช่วยให้ผู้ใช้งานสามารถทำนายหรือ "พยากรณ์" หมายเลขทะเบียนรถที่เป็นไปได้ โดยใช้การประมวลผลผ่านอัลกอริทึมและข้อมูลที่ผู้ใช้งานให้มา เพื่อให้คำแนะนำและข้อมูลเชิงลึกเกี่ยวกับหมายเลขทะเบียนที่อาจมีความหมายหรือสอดคล้องกับความชอบของผู้ใช้งาน


## How to install
```shell
$ go get github.com/naruebaet/thai-plate-prophecy-go
 ```

## Feature
| Feature Name       | Description                                                                 |
|--------------------|-----------------------------------------------------------------------------|
| Plate Prediction   | Predicts potential license plate numbers based on user input and algorithms. |
| User Preferences   | Allows customization of predictions to align with user preferences.         |
| Algorithm Insights | Provides transparency into how predictions are generated.                  |

## Public Functions and Usage

### 1. `AdviceByPlateData`
**Description**: Calculates the numerological value and meaning of a Thai license plate.

**Parameters**:
- `firstData` (string): The first part of the license plate (e.g., Thai characters).
- `secondData` (string): The numeric part of the license plate.

**Returns**:
- `PlateCalculationResult`: Contains the calculated values and meanings.
- `error`: If validation fails or an error occurs.

**Example**:
```go
result, err := AdviceByPlateData("กข", "1234")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Result: %+v\n", result)
}
```

---

### 2. `AdviceByDMY`
**Description**: Provides lucky number advice based on a given date, month, and year.

**Parameters**:
- `date` (string): The day of the month (e.g., "01").
- `month` (string): The month (e.g., "01" for January).
- `year` (string): The year (e.g., "2023").

**Returns**:
- `LuckyNumberAdvice`: Contains the lucky number advice.
- `error`: If the date format is invalid or no advice is found.

**Example**:
```go
advice, err := AdviceByDMY("01", "01", "2023")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Advice: %+v\n", advice)
}
```

---

### 3. `AdviceByWeekDay`
**Description**: Provides lucky number advice based on the day of the week.

**Parameters**:
- `day` (WeekDay): The day of the week (0 for Sunday, 1 for Monday, etc.).

**Returns**:
- `LuckyNumberAdvice`: Contains the lucky number advice.
- `error`: If no advice is found for the given day.

**Example**:
```go
advice, err := AdviceByWeekDay(1) // 1 for Monday
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Advice: %+v\n", advice)
}
```

### Support me coffee
[Buy me a coffee](https://www.buymeacoffee.com/alpakalab)

Crafted by : [Alpaka LAB](https://alpakalab.com)