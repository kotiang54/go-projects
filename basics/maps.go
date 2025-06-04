package basics

import (
	"fmt"
	"maps"
)

func main() {
	/*	var mapVariable map[keyType]valueType
		mapVariable := make(map[keyType]valueType)

		using a Map literal
		mapVariable := map[keyType]valueType{
		key1: value1,
		key2: value2,
		}
	*/

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

	// delete the entire map
	myMap["key1"] = 11
	myMap["code1"] = 18
	myMap["key2"] = 12
	myMap["code2"] = 19
	fmt.Println(myMap)

	// clear(myMap)
	fmt.Println(myMap)

	value, unknownValue := myMap["key1"]
	fmt.Println("value: ", value)
	fmt.Println("Is a value associated with key1?: ", unknownValue)

	// the unknown component in the map call above is important
	// like demonstrated in the -if statement- below:

	if _, ok := myMap["code"]; ok {
		delete(myMap, "code")
	}
	fmt.Println(myMap)

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	myMap3 := map[string]int{"a": 1, "b": 2, "c": 3}

	if maps.Equal(myMap2, myMap3) {
		fmt.Println("myMap2 and myMap3 are equal!")
	}

	// iterating over a map
	for key, value := range myMap2 {
		fmt.Println(key, value)
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value.")
	} else {
		fmt.Println("The map is not initialized to nil value.")
	}

	val := myMap4["key"]
	fmt.Println(val)

	// if you initialize a map as above i.e., var myMap4 map[string]string
	// you cannot populate the map using the map["key"] = value statement as shown below
	// myMap4["key"] = "10"
	// fmt.Println(myMap4)

	// Correct way:
	myMap4 = make(map[string]string)
	myMap4["key"] = "10"
	fmt.Println(myMap4)

	fmt.Println("myMap4 length is", len(myMap4))

	// nested maps
	myMap5 := make(map[string]map[string]string)
	fmt.Println(myMap5)

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
