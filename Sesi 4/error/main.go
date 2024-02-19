package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var err error

	number, err = strconv.Atoi("12345H")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
}
