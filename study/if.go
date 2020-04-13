package main

import "fmt"

func main() {

	var a int = 20
	b := 30

	if a >= 15 {
		fmt.Println("15 이상")
	}
	if b >= 25 {
		fmt.Println("25 이상")
	}
	if c := 40; c >= 35 {
		fmt.Println("35 이상")
	}

	if a >= 65 {
		fmt.Println("65 이상")
	} else {
		fmt.Println("65 미만")
	}
	if c := 90; c >= 100 {
		fmt.Println("100 이상")
	} else {
		fmt.Println("100 미만")
	}

	i := 100
	if i >= 120 {
		fmt.Println("120 이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100 이상 120 미만")
	} else {
		fmt.Println("100 미만")
	}
}
