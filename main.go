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
}
