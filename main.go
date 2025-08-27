package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	// var name string
	// // fmt.Println(name)
	// name = "Chirag"
	// fmt.Println(name)

	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	//

	defer func() {
		sum := 0
		for i := 0; i < 100; i++ {
			sum += i

		}
		fmt.Println(sum)
	}()

	sum := 0
	for i := 0; i < 100000000000; i++ {
		sum += i
	}
	fmt.Println(sum)

}
