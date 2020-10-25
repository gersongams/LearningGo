package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var length float64 = 1.2
	var width int = 2

	//This will give us an error
	fmt.Println("Area is", length *  float64(width))
	fmt.Println("length > width", length > float64(width))

	var now = time.Now()
	var year = now.Year()
	fmt.Println(year)

	sentence := "Python is the best language"
	replacer := strings.NewReplacer("Python", "Go")
	fixedSentence := replacer.Replace(sentence)
	fmt.Println(fixedSentence)

}
