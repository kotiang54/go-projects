package main

import "fmt"

func main() {
	// var mapvariable map[keyType]valueType
	// mapvariable := make(map[keyType]valueType)
	// using a Map literal
	// mapvariable := map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)

	// Access the map values using keys
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["code"])

	// modify a map value
	myMap["code"] = 21
	fmt.Println(myMap)

	// delete a map element
	// deleting a non-existent key does not cause a panic error
	delete(myMap, "key")
	fmt.Println(myMap)

	if _, ok := myMap["code"]; ok {
		delete(myMap, "code")
	}
	fmt.Println(myMap)

	// delete the entire mare
	myMap["key1"] = 11
	myMap["code1"] = 18
	myMap["key2"] = 12
	myMap["code2"] = 19
	fmt.Println(myMap)

	clear(myMap)
	fmt.Println(myMap)

}
