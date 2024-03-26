package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello")
	//singleVariableTests()
	//multipleVariableTests()
	//constants()
	//arrays()
	//slicing()
	//rangeFunction()
	//a := function_with_named_return(10, 15)
	//fmt.Println(a)

	//For a Function with Multiple Return Types
	a, b := function_with_multiple_return(10, "Hello")
	fmt.Println(a, b)
	//If one of the return types are required,use underscore
	_, c := function_with_multiple_return(10, "Hello")
	fmt.Println(c)

	result := recursiveFunction(10)
	fmt.Println(result)
}
func singleVariableTests() {
	//Variable Definition with initial value.
	//Type is inferred
	firstName := "Tom"
	//With var keyword.Type is inferred
	var lastName = "Onyango"
	//With var keyword and type specified
	var middleName string = "George"
	fmt.Println("Hello " + firstName + " " + middleName + " " + lastName)

	//Variable Definition without Initial Value.Default type value is assigned
	var isAdult bool          //assigned false
	var age int               //assigned zero
	var jobDescription string //assigned ""
	fmt.Println(isAdult)
	fmt.Println(age)
	fmt.Println(jobDescription)

	//Assign Value after Declaration
	isAdult = true
	age = 35
	jobDescription = "Engineer"
	fmt.Println(isAdult)
	fmt.Println(age)
	fmt.Println(jobDescription)

}
func multipleVariableTests() {
	//With var keyword
	var firstName, lastName, age, isAdult = "Tom", "Onyango", 35, true
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(isAdult)
	//With :=
	jobdescription, location := "Engineer", "Kampala"
	fmt.Println(jobdescription)
	fmt.Println(location)
	//In a block
	var (
		maritalStatus string = "Single"
		noOfChildren  int
	)
	fmt.Println(maritalStatus)
	fmt.Println(noOfChildren)

}
func constants() {
	const PI = 3.142
	fmt.Println(PI)
}
func arrays() {
	//Known Lengths
	var myArray = [6]int{0, 1, 2, 3, 4, 5}
	myOtherArray := [2]string{"Tom", "Onyango"}
	fmt.Println(myArray)
	fmt.Println(myOtherArray)

	//Inferred Lengths
	var myArray2 = [...]int{0, 1, 2, 3}
	myOtherArray2 := [...]string{"Tom", "Onyango"}
	//Access Values
	fmt.Println(myArray2[2])
	fmt.Println(myOtherArray2[1])
	//Change Values
	myArray2[2] = 10
	fmt.Println(myArray2[2])
}
func slicing() {
	my_slice := []int{}
	fmt.Println(my_slice)

	//Slice from Array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]
	fmt.Printf("myslice = %v\n", myslice)

	//Make Function
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	//Append Values to Slice
	myslice1 = append(myslice1, 4, 5)
	fmt.Printf("myslice1 appended = %v\n", myslice1)
	//Append Slices
	myslice_1 := []int{1, 2, 3}
	myslice_2 := []int{4, 5, 6}
	myslice_3 := append(myslice_1, myslice_2...)
	fmt.Printf("myslice_3 = %v\n", myslice_3)
	//Change length by re-slicing the array
	myslice1 = arr1[1:3]
	//Change length by appending items
	myslice1 = append(myslice1, 20, 21, 22, 23)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

}
func if_else_condition() {
	var a int = 10
	var b int = 220
	if a > b {
		fmt.Println("a is greater")
	} else {
		fmt.Println("b is greater")
	}
}
func else_if_Condition() {
	var a int = 10
	var b int = 220
	if a > b {
		fmt.Println("a is greater")
	} else if b > a { //Else must be on the same line as the closing bracket of the if condition } else
		fmt.Println("b is greater")
	}
}
func switch_single_condition() {
	//Similar to JAVA
	var a int = 10
	var b int = 13
	switch true {
	case a > b:
		fmt.Println("a is greater")
	case a < b:
		fmt.Println("b is greater")
	default:
		fmt.Println("a is equal to b")
	}
}
func switch_multiple_condition() {
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}
func for_loop() {
	//for is the only loop in Go
	//continue is used to skip an iteration
	//break terminates the iteration and exits the loop
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
func range_function() {
	//range is used to iterate over an array,slice or map
	//returns both index and value
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits { //idx holds the index,val holds the value
		fmt.Printf("%v\t%v\n", idx, val)
	}
	//If only value is required
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
	//If only index is required
	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}

}

// for functions with return types, specify the return type before the opening {
func function_with_params(a int, b int) int {

	return a + b
}

// If the return type is named, you can use only return
func function_with_named_return(a int, b int) (result int) {
	result = a + b
	//return
	//or
	return result
}

// multiple return types
func function_with_multiple_return(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

// recursive function
func recursiveFunction(a int) int {
	if a >= 11 {
		return 0
	}
	fmt.Println(a)
	return recursiveFunction(a + 1)

}
