package main

import "fmt"

func GetPrice(age int) string {
	s := ""
	if age < 0 || age > 120 {
		s = "Wrong age. Try again"
	} else if age >= 0 && age <= 9 {
		s = "Free"
	} else if age >= 10 && age <= 19 {
		s = "30000"
	} else if age >= 20 && age <= 59 {
		s = "50000"
	} else {
		s = "30000"
	}
	return s
}

func main() {
	var age int

	fmt.Print("Input: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Output:", GetPrice(age))
}
