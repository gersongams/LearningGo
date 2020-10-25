package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter your age:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	age, _ := strconv.Atoi(input)
	fmt.Println(age)

	if age < 20 {
		fmt.Println("you are underage")
	}else {
		fmt.Println("you are good")
	}

	//fileInfo, errFile := os.Stat("1_firstapp.go")
	//if errFile != nil{
	//	log.Fatal(errFile)
	//}
	//fmt.Println(fileInfo.Size())


}
