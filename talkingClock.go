package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	hours()
	//minutes10()
}

func hours() {

	var input string
	fmt.Print("Enter the time : ")
	_, err := fmt.Scan(&input)
	inputInt, _ := strconv.Atoi(input)
	suffix := "am"

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if inputInt > 12 {
		inputInt = inputInt - 12
		suffix = "pm"
	}

	switch inputInt {
	case 1:
		fmt.Println("one", suffix)
	case 2:
		fmt.Println("two", suffix)
	case 3:
		fmt.Println("three", suffix)
	case 4:
		fmt.Println("four", suffix)
	case 5:
		fmt.Println("five", suffix)
	case 6:
		fmt.Println("six", suffix)
	case 7:
		fmt.Println("seven", suffix)
	case 8:
		fmt.Println("eight", suffix)
	case 9:
		fmt.Println("nine", suffix)
	case 10:
		fmt.Println("ten", suffix)
	case 11:
		fmt.Println("eleven", suffix)
	case 12:
		fmt.Println("twelve", suffix)
	}

	// func minutes10() {

	// 	m1 := 4
	// 	switch m1 {
	// 	case 0:
	// 		fmt.Println("oh ")
	// 	case 1:
	// 		fmt.Println("oh")
	// 	case 2:
	// 		fmt.Println("twenty")
	// 	case 3:
	// 		fmt.Println("thirty")
	// 	case 4:
	// 		fmt.Println("forty")
	// 	case 5:
	// 		fmt.Println("fifty")
	// 	}

}
