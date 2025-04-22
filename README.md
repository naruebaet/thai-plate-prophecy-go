# Thai plate prophecy GO

This project is designed to help customers predict or "prophecy" potential license plate numbers. By leveraging algorithms and user input, it provides insights and suggestions for plate numbers that may hold significance or align with user preferences.  

---
Example

```Go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/naruebaet/thai-plate-prophecy-go/v1"
)

func main() {
    result, err := CalculatePlateData("1กข", "1234")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    jsonResult, _ := json.MarshalIndent(result, "", "  ")
    fmt.Println(string(jsonResult))
}
```