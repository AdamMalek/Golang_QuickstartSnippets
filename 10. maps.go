package main

import "fmt"

func main() {
	mapObj := make(map[string]string)
	fmt.Println("empty map", mapObj)

	inlineMap := map[string]string{}
	fmt.Println("inline map", inlineMap)

	mapObj["key1"] = "value1"
	inlineMap["inline-key1"] = "inline-value1"

	fmt.Println("both works the same", mapObj, inlineMap)

	withItems := map[string]string{"init1": "valueInit1"}
	fmt.Println("initialized inline", withItems)

	fmt.Println("Getting key/value")
	value, keyExists := mapObj["key1"]
	fmt.Println("value/keyExists:", value, "/", keyExists)
	value, keyExists = mapObj["key324"]
	fmt.Println("value/keyExists:", value, "/", keyExists)

	fmt.Println("usage")
	PrintIfExists := func(mapToCheck map[string]string, key string) {
		val, exists := mapToCheck[key]
		if exists {
			fmt.Println(key, ": exists and it's value is:", val)
		} else {
			fmt.Println(key, "does not exist")
		}
	}

	PrintIfExists(mapObj, "key1")
	PrintIfExists(mapObj, "NonExistent key")

	fmt.Println("We can delete keys")
	fmt.Println("Before deleting key1", mapObj)
	delete(mapObj, "key1")
	fmt.Println("After", mapObj)

	fmt.Println("Before deleting invalid key", inlineMap)
	delete(inlineMap, "invalid key")
	fmt.Println("After", inlineMap)

	fmt.Println("Before adding element", inlineMap)
	inlineMap["newItem"] = "newValue"
	fmt.Println("After", inlineMap)

	fmt.Println("Before updating element", inlineMap)
	inlineMap["newItem"] = "newValue34143124"
	fmt.Println("After", inlineMap)
}
