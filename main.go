package main

import (
	"fmt"

	. "github.com/ahmetb/go-linq/v3"
)

type Car struct {
	id    int
	year  int
	owner string
}

func main() {
	//##LinQ
	linQ()

	//##Validate
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
