package main

import (
	"fmt"
	"strconv"
)

func add(a int, b int) int {
	return a + b;

}

func parseNumber(s string)(int, error){
	num, err := strconv.Atoi(s);
	return num, err;
}


///BMI CALCULATOR
func calculateBMI(weight, height float64) float64 {
	return weight / (height * height);
}

func main(){
	// result := add(5, 7)
    // fmt.Println("The sum of 5 and 7 is:", result)


	// result, err := parseNumber("43");

	// if err!= nil {
    //     fmt.Println("Error parsing the number:", err)
    // } else {
	// 	fmt.Println("The number is: ", result);
	// }


	// --------------------------------
	var weight, height float64

	fmt.Print("Enter Weight (kg): ")
	fmt.Scanln(&weight)
	fmt.Print("Enter Height (m): ")
	fmt.Scanln(&height);

	bmi := calculateBMI(weight, height);
	fmt.Printf("Your BMI is: %.2f\n", bmi);


	if bmi < 18.5 {
		fmt.Println("UnderWeight");
	} else if bmi >= 18.5 && bmi < 24.9 {
		fmt.Println("Normal Weight");
	} else if bmi >= 25 && bmi <= 29.9 {
		fmt.Println("Overweight");
	} else {
		fmt.Println("Obesity")
	}
}