// File name: ...\s03\exercises\set-1\bmi\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// File name: ...\s03\exercises\set-1\bmi\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

/*
Exercise: Write a program to calculate body mass index (BMI).

Read 'weight' (in pounds) and 'height' (in inches) from console.

bmi = weight  / (height  * height )
	  (in kg)	(in m)	   (in m)

Display bmi and weight status from the following table.

bmi < 18.5	--> Underweight
bmi < 25	--> Normal
bmi < 30	--> Overweight
otherwise	-->	Obese

Examples:
weight	height	bmi			status
140		90		39.371896	Underweight
140		70		12.151820	Normal
140		60		27.341595	Overweight
140		50		39.371896	Obese
*/

func main() {

	const kilogramsPerPound = 0.45359237
	const metersPerInch = 0.0254

	var weight float64
	fmt.Print("Enter weight in pounds:")
	fmt.Scanf("%f ", &weight)

	var height float64
	fmt.Print("Enter height in inch:")
	fmt.Scanf("%f ", &height)

	// Covert Pounds to Kilograms & Inches to Meters
	weightInKilograms := weight * kilogramsPerPound
	heightInMeters := height * metersPerInch

	// Calculate BMI
	bmi := weightInKilograms / (heightInMeters * heightInMeters)

	weightStatus := ""
	if bmi < 18.5 {
		weightStatus = "Underweight"
	} else if bmi < 25 {
		weightStatus = "Normal"
	} else if bmi < 30 {
		weightStatus = "Overweight"
	} else {
		weightStatus = "Obese"
	}

	fmt.Println()
	fmt.Printf("%-10s %-10s %-12s %-10s \n", "Weight", "Height", "BMI", "Status")
	fmt.Printf("%-10.2f %-10.2f %-12.6f %-10s \n", weight, height, bmi, weightStatus)

}
