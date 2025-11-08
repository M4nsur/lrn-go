package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	isRepeat := true
	for isRepeat {
		userHeight, userWeight := getUserValues()
		getResult(userHeight, userWeight)
		isRepeat = handleRepeat()
	}
}

func getResult(userHeight, userWeight float64) {
	const IMTPower = 2
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)

	switch {
	case IMT < 16:
		fmt.Println("Severe body mass deficit")
	case IMT < 18.5:
		fmt.Println("Body mass deficit")
	case IMT < 25:
		fmt.Println("Normal")
	case IMT < 30:
		fmt.Println("Overweight")
	default:
		fmt.Println("Obesity")
	}
}

func getUserValues() (float64, float64) {
	userHeight := validateValue("height (in cm)")
	userWeight := validateValue("weight (in kg)")
	return userHeight, userWeight
}

func handleRepeat() bool {
	var answer string
	for answer != "y" && answer != "n" {
		fmt.Println("Again? y/n")
		fmt.Scan(&answer)
	}
	return answer == "y"
}

func validateValue(valueName string) float64 {
    var input string
    for {
        fmt.Printf("Enter your %s: ", valueName)
        if _, err := fmt.Scanln(&input); err != nil {
            fmt.Println("Error: please enter a valid positive number")
            continue
        }
        trimmed := strings.TrimSpace(input)
        value, err := strconv.ParseFloat(trimmed, 64)
        if err != nil || value <= 0 {
            fmt.Println("Error: please enter a valid positive number")
            continue
        }
        return value
    }
}