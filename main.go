package main

import (
	"fmt"
	"strings"

	// "time"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/gookit/validate"
)

type Car struct {
	id    int
	year  int
	owner string
}

func main() {

	//## Contrain string in array
	// str := []string{"api/login", "api/logOut"}
	// has := Contains(str, "Out")
	// fmt.Println("Has is: ", has)

	//##LinQ
	// linQ()

	//##Validate
	Valid()
}

//message:"กรุณากรอกค่า {field}"
// UserForm struct
type UserForm struct {
	Name string `validate:"required|minLen:4|maxLen:10" message:"required: ห้ามว่าง| minLen:ค่าต้องไม่น้อยกว่า 7 |maxLen:ค่าต้องไม่เกิน 10"`
	// Email    string    `validate:"email" message:"email is invalid"`
	// Age int `validate:"required|int|min:1|max:99" message:"int:age must int| min: age min value is 1"`
	// CreateAt int       `validate:"min:1"`
	// Safe     int       `validate:"-"`
	// UpdateAt time.Time `validate:"required"`
	// Code     string    `validate:"customValidator"`
}

// CustomValidator custom validator in the source struct.
func (f UserForm) CustomValidator(val string) bool {
	return len(val) == 4
}

func Valid() {
	u := &UserForm{
		Name: "12345678901",
	}

	v := validate.Struct(u)
	// v := validate.New(u)

	if v.Validate() { // validate ok
		// do something ...
	} else {
		fmt.Println(v.Errors) // all error messages
		// fmt.Println(v.Errors.One())         // returns a random error message text
		// fmt.Println(v.Errors.Field("Name")) // returns error messages of the field
	}
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
