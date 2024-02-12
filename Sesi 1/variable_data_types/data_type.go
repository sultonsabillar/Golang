package main

import "fmt"

func main() {

	// ----------------------- INTEGER ----------------------- //
	first := 22
	second := -23

	var third uint8 = 98
	var fourth int8 = -9
	fmt.Printf("First data type : %T \n", first)
	fmt.Printf("Second data type : %T \n", second)

	fmt.Printf("Third data type : %T \n", third)
	fmt.Printf("Fourth data type : %T \n", fourth)
	// ------------------- END OF INTEGER -------------------- //

	// ------------------------ FLOAT ------------------------ //
	decimalNumber := 3.45

	fmt.Printf("Decimal number: %f \n", decimalNumber)
	fmt.Printf("Decimal number: %.3f \n", decimalNumber)
	// ------------------- END OF INTEGER ------------------- //

	// ----------------------- BOOLEAN ---------------------- //
	condition := false
	fmt.Printf("Is it golang course for men? %t \n", condition)
	// ------------------- END OF BOOLEAN ------------------- //

	// ----------------------- STRING ----------------------- //
	var message string = "Annyeong!"
	fmt.Printf("%s", message)
	// -------------------- END OF STRING ------------------- //

}
