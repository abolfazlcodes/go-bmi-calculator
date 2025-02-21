package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func storeUserBMI(userName string, bmiValue float64) error {
	// sanitize username for a valid file name
	fileName := fmt.Sprintf("%s_bmi.txt", userName)

	// convert it to a string
	result := fmt.Sprintf("%v's BMI is %.2f", userName, bmiValue)

	// store it in a file
	error := os.WriteFile(fileName, []byte(result), 0644)

	if error != nil {
		fmt.Printf("Error writing to file: %v\n", error)
		return errors.New("Something went wrong during writing to file. Please try again")
	}

	fmt.Printf("BMI was successfully stored in file: %s\n ", fileName)
	return nil
}

func validateUserHeight(height float64) float64 {
	var result float64
	if height > 3 {
		// it means the user had entered his height in centimeters ==> we need to convert it to meters
		result = height / 100
		return result
	}

	// otherwise it is already in meters ==> return it
	return height
}

func readBMIFromFile(userName string) (float64, error) {
	fileName := fmt.Sprintf("%s_bmi.txt", userName)

	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return 0, errors.New("Something went wrong during reading from file. Please try again")
	}

	// Convert bytes to string
	dataString := string(data)

	// Split at " is " and get the second part
	splittedValue := strings.Split(dataString, " is ")
	if len(splittedValue) < 2 {
		return 0, errors.New("BMI value is missing or not properly formatted")
	}

	// Trim any newline or spaces
	bmiStr := strings.TrimSpace(splittedValue[1])

	// Convert to float64
	bmiValue, parseErr := strconv.ParseFloat(bmiStr, 64)
	if parseErr != nil {
		return 0, errors.New("BMI value is not a valid number")
	}

	return bmiValue, nil
}

func main() {
	fmt.Println("------ Welcome to BMI Calculator ------")

	var userName string
	fmt.Println("Please enter the user's name:")
	fmt.Scan(&userName)

	// 1. check if the user already has calculated its BMI
	userPrevBMI, errorPrev := readBMIFromFile(userName)

	if errorPrev != nil {
		panic(errorPrev)
	}

	fmt.Printf("Your previous BMI is: %.2f", userPrevBMI)

	var recalculateBMIAnswer string
	fmt.Println("Do you want to recalculate your BMI? (y/n): ")
	fmt.Scan(&recalculateBMIAnswer)

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

	var checkedHeight = validateUserHeight(userHeight)

	var bmi = calculateBMI(userWeight, checkedHeight)
	error := storeUserBMI(userName, bmi)

	if error != nil {
		fmt.Println(error)
		return
	}
}
