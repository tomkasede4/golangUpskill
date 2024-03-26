package models

import (
	"fmt"
)

func createMap() map[string]string {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	return a
}

func createMapWithMake() map[string]string {
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	return a
}

func createEmptyMap() map[string]string {
	var a map[string]string // The map is nil
	return a
}
func mapFunctions() {
	//Maps are reference types
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	//Delete with KeyName
	delete(a, "brand")
	//Check if Key Exists
	//Checking for existing key and its value
	val1, ok1 := a["brand"] //Result:  Ford true
	//Checking for non-existing key and its value
	val2, ok2 := a["color"] //Result:   false

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)

	//Iterate over Map with range
	b := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range b {
		fmt.Printf("%v : %v, ", k, v)
	}

}
