# Thai plate prophecy GO

This project is designed to help customers predict or "prophecy" potential license plate numbers. By leveraging algorithms and user input, it provides insights and suggestions for plate numbers that may hold significance or align with user preferences.  

Library สำหรับใช้ตรวจสอบทะเบียนที่เหมาะสมกับตัวเรา และยังสามารถช่วยแนะนำทะเบียนที่เหมาะสมกับตัวเราผ่านวันเดือนปีเกิดได้อีกด้วย ที่จะช่วยส่งเสริมในด้านต่างๆตามความเชื่อ 

---
## How to install
`$ go get github.com/naruebaet/thai-plate-prophecy-go`

## Example
```Go
package main

import (
	"encoding/json"
	"fmt"

	tpp "github.com/naruebaet/thai-plate-prophecy-go/v1"
)

func main() {
	raw, _ := tpp.CalculatePlateData("1กข", "1234")
	prettyJson, _ := json.MarshalIndent(raw, "", "  ")
	fmt.Printf("%s\n", prettyJson)

	advice, _ := tpp.AdvicePlateData("22", "02", "1993")
	prettyJsonAdvice, _ := json.MarshalIndent(advice, "", "  ")
	fmt.Printf("%s\n", prettyJsonAdvice)
}


```

## Result
```json
{
  "firstPart": {
    "value": "1กข",
    "sum": 4
  },
  "secondPart": {
    "value": "1234",
    "sum": 1,
    "luckyPoint": {
      "point": 1,
      "desc": "เลขทะเบียนรถที่มีผลรวมเท่ากับ 1 เหมาะสำหรับผู้ที่มีตำแหน่งหน้าที่การงานสูง เช่น ข้าราชการ หัวหน้างาน หรือเจ้าของกิจการ เลขนี้ช่วยเสริมสร้างความเป็นผู้นำ อำนาจวาสนา และเพิ่มความน่าเชื่อถือในสายตาผู้อื่น"
    }
  },
  "total": {
    "sum": 14,
    "luckyGroup": {
      "group": "best",
      "points": [
        4,
        5,
        6,
        9,
        14,
        15,
        19,
        23,
        24,
        32,
        36,
        40,
        41,
        42,
        44,
        45,
        46,
        50,
        51,
        54
      ],
      "desc": "เลขมงคลระดับดีเยี่ยม ช่วยเสริมดวงให้เจ้าของรถประสบความสำเร็จในชีวิต ส่งเสริมความก้าวหน้าในหน้าที่การงาน เพิ่มโชคลาภด้านการเงิน นำมาซึ่งความมั่งคั่งร่ำรวย และเป็นที่เคารพนับถือของคนรอบข้าง ช่วยเสริมสิริมงคลและความเจริญรุ่งเรืองให้กับชีวิต"
    }
  }
}
```