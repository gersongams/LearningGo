package main

import "fmt"

func main() {

	//start := 2
	//end := 10
	//
	//for x := start; x < end; x++ {
	//	fmt.Println(x)
	//}

	//fmt.Println(x)

	//end := 2
	//start := 10
	//
	//for x := start; x > end; x-- {
	//	fmt.Println(x)
	//}

	//for x := 0; x <= 15; x+=5 {
	//	fmt.Println(x)
	//}

	//var x int
	//
	//for x = 5; x < 10; x++ {
	//	fmt.Println(x)
	//}
	//
	//fmt.Println("Out of the scope", x)

	var y int

	for y = 0; y < 10; y++ {
		fmt.Println(y)
		if y == 5 {
			fmt.Println("Breaking")
			break
		}
	}

}
