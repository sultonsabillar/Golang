package main

import "fmt"

func main() {
	var dataa []string
	dataa = []string{"aaa", "bbb"}
	for key, data := range dataa {
		fmt.Println("No", key+1, ":", data)
	}

	datas := []string{"a", "b", "c"}
	for _, data := range datas {
		fmt.Println(data)
	}
}
