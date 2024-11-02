package main
//1. var keyword

var age int = 25

//2. Type inference with var

var name = "Alice"

//3. Short declaration ( := )
score := 10


// Data Types
//1 Numeric Types
// Integer Type int, int8 32  (signed integers)
//uint8, uint32 (unsigned integers)

// Float Tye float 32 & 64
var temperature float64 = 36.6

//2 Boolean Type
var isActive bool = false

//3. String Type
var greeting string = "Alice"

//4. Complex Types
//Arrays 
var numbers [5]int = [5]int{1,2,3,4,5}

var fruits = []string{"apple", "banana"}

//Maps
var cities = map[string]string{"USA" : "Washington D.C"}

//Structs
type Person struct{
	Name string
	Age  int
}

var john = Person{"John", 30}

//constants
const pi = 3.141