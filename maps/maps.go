package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myMaps = make([]map[string]string, 0)

	for i := 0; i < 5; i++ {
		var myMap = make(map[string]string)
		myMap["firstName"] = "Vinh" + strconv.Itoa(i)
		myMap["lastName"] = "Bui"
		myMap["email"] = "vinhbh@ptit.com"

		myMaps = append(myMaps, myMap)
	}

	for _, item := range myMaps {
		firstName := item["firstName"]
		fmt.Println(firstName)

	}
}
