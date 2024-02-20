package main

import "fmt"

func maiin() {
	var randomValues interface{}

	// _ = randomValues

	randomValues = "Jalan Sudirman"

	randomValues = 20

	randomValues = true

	randomValues = []string{"Pizza", "Udon"}

	fmt.Println(randomValues)

	randomMap := map[string]interface{}{
		"Name": "Fitri",
		"Age":  20,
	}

	fmt.Println(randomMap)

}
