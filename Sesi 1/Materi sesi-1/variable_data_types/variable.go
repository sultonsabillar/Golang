package main

// WITH DATA TYPE
/* func main() {
	// var title string = "Golang for women"
	// var totalMember int = 18

	var title string
	var totalMember int
	title = "Golang for women"
	totalMember = 18

	fmt.Println("Course title is: ", title)
	fmt.Println("The total member is: ", totalMember)
} */

// WITHOUT DATA TYPE
/* func main() {
	// var title = "Golang for women"
	// var totalMember = 18

	title := "Golang for women" //short declaration technique
	totalMember := 18

	fmt.Printf("%T, %T \n", title, totalMember)
	fmt.Println("Course title is: ", title)
	fmt.Println("The total member is: ", totalMember)
} */

// MULTIPLE DECLARATION VARIABLE
/* func main() {
	var firstName, surName, lastName string = "Fitri", "Ayu", "Anggraini"

	// var age1, age2, age3 int
	// age1, age2, age3 = 20, 21, 22
	age1, age2, age3 := 20, 21, 22

	fmt.Println(firstName, surName, lastName)
	fmt.Println(age1, age2, age3)

	fmt.Printf("Hello, my name is %s and my age is %d years old \n", firstName, age2)
} */
