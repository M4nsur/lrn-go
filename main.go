package main

import (
	"fmt"
	"math"
)

func main() {
	isRepeat := true
	for isRepeat  {
		userHeight, userWeight := getUserValues()
		getResult(userHeight, userWeight)

		isRepeat = handleRepeat()
	}
}

func getResult( userHeight, userWeight float64 ) {
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
		fmt.Println("Overweigh")
	default:
		fmt.Println("degree obesity")
	}
 }

func getUserValues () (float64, float64) {
	var userHeight float64
	var userWeight float64
	
	fmt.Print("Enter your height (in sm) ")
	fmt.Scan(&userHeight)

	fmt.Print("Enter your weight (in kg) ")
	fmt.Scan(&userWeight)

	return userHeight, userWeight
}


func handleRepeat () (bool) {
	var answer string 
	for answer != "y" && answer != "n" {
		fmt.Println("Again ? y/n")
		fmt.Scan(&answer)
	}
	return answer == "y" 
} 