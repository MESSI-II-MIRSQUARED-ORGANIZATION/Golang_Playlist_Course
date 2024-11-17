package main

import (
	"fmt"

)

// Student Grade Tracker
func addGrade (grades []int, grade int) []int {
	return append(grades, grade)
}


func updateInventory (inventory map[string]int, item string, quantity int) {
	inventory[item] = quantity
}
func main (){
	// var grades [5]int


	// grades[0] = 95
	// grades[1] = 85
	// grades[2] = 75
	// grades[3] = 65


	// fmt.Println("Grades:", grades)

	// fmt.Println("First grade: ", grades[0])

	// -----SLICES----------
	// grades := []int {90, 85, 60}

	// grades = append(grades, 88)

	// fmt.Println("Grades: ", grades)
	// fmt.Println("Length of grades", len(grades))
	// fmt.Println("Capacity", cap(grades))

// -----------STUDENT GRADE TRACKER--------
	// grades := []int{}


	// grades = addGrade(grades, 95)
	// grades = addGrade(grades, 85)
	// grades = addGrade(grades, 75)


	// fmt.Println("Grades: ", grades)


// -------------MAPS-----------

//  inventory := map[string]int{
// 	"apples" : 30,
// 	"bananas" : 15,
//  }


//  fmt.Println("Apples in stock: ", inventory["apples"])

// inventory["oranges"] = 40

// inventory["bananas"] = 50

// fmt.Println("Inventory: ", inventory)


// -------------------
inventory := map[string]int{}

updateInventory(inventory, "apples", 30)
updateInventory(inventory, "bananas", 15)

fmt.Println("Inventory: ", inventory)


updateInventory(inventory, "bananas", 30)

fmt.Println("Inventory: ", inventory)

}