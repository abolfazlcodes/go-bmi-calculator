package main

import (
	"errors"
	"fmt"
	"os"
)

// Mark and John are trying to compare their BMI
// bmi formula: mass / height ** 2 or mass / height * height

// Tasks:
// 1. Store Mark's and John's mass and height in variable / a file
// 2. calculate both their BMIs
// 3. store the values in a file and before calculating again, show the previous comparison before calculating again

// Test data:
// § Data 1: Marks weights 78 kg and is 1.69 m tall. John weights 92 kg and is 1.95
// m tall.
// § Data 2: Marks weights 95 kg and is 1.88 m tall. John weights 85 kg and is 1.76
// m tall.

func getUserInput(label string) (float64, error) {
	var userValue float64
	fmt.Printf("Enter %v: ", label)
	fmt.Scan(&userValue)

	if userValue <= 0 {
		errorMessage := errors.New("Invalid input")
		return 0, errorMessage
	}

	return userValue, nil
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func storeUserBMI(userName string, bmiValue float64) {
	// sanitize username for a valid file name
	fileName := fmt.Sprintf("%s_bmi.txt", userName)

	// convert it to a string
	result := fmt.Sprintf("%v's BMI is %.2f", userName, bmiValue)

	// store it in a file
	error := os.WriteFile(fileName, []byte(result), 0644)

	if error != nil {
		fmt.Printf("Error writing to file: %v\n", error)
		return
	}

	fmt.Printf("BMI was successfully stored in file: %s\n ", fileName)
}

func main() {
	fmt.Println("------ Welcome to BMI Calculator ------")

	var userName string
	fmt.Println("Please enter the user's name:")
	fmt.Scan(&userName)

	userWeight, weightErrorMsg := getUserInput("Weight")
	userHeight, heightErrorMsg := getUserInput("Height")

	if weightErrorMsg != nil {
		fmt.Println(weightErrorMsg)
		return
	}

	if heightErrorMsg != nil {
		fmt.Println(heightErrorMsg)
		return
	}

	var bmi = calculateBMI(userWeight, userHeight)
	storeUserBMI(userName, bmi)

	fmt.Println(userName, userHeight, userWeight, bmi)
}
