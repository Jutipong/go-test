package main

import (
	"fmt"
	"strings"

	. "github.com/ahmetb/go-linq/v3"
)

type Car struct {
	id    int
	year  int
	owner string
}

func main() {

	//## Contrain string in array
	str := []string{"api/login", "api/logOut"}
	has := Contains(str, "Out")
	fmt.Println("Has is: ", has)

	//##LinQ
	// linQ()

	//##Validate
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}
	return false
}

func linQ() {
	cars := []Car{}
	cars = append(cars,
		Car{
			id:    111111,
			year:  1,
			owner: "11",
		}, Car{
			id:    222222,
			year:  2,
			owner: "22",
		}, Car{
			id:    333333,
			year:  3,
			owner: "33",
		}, Car{
			id:    444444,
			year:  4,
			owner: "44",
		})

	//## Linq
	var result []Car
	From(cars).WhereT(func(c Car) bool {
		return c.id == 222222
	}).SelectT(func(c Car) Car {
		return c
	}).ToSlice(&result)

	fmt.Printf("%+v", result)
	fmt.Println()
}
