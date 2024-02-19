package main

import (
	"errors"
	"fmt"
)

func customError() {
	var password string

	fmt.Scanln(&password)

	// valid, err := validPassword(password)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(valid)
	// }

	if valid, err := validPassword(password); err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

func validPassword(password string) (string, error) {

	lengthPassword := len(password)

	if lengthPassword < 5 {
		return "", errors.New("Password has to have more than 4 characters")
	}

	return "Valid password", nil
}
