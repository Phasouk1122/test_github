package main

import (
	"fmt"

	"rsc.io/quote"
)

type struct1 struct {
	firstName string
	lastName  string
	a         int
	b         int
}

func byval(q *int) {
	fmt.Printf("3. byval -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
	*q = 4143
	fmt.Printf("4. byval -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
	q = nil
}

func mergeName(firstName string, lastName string) string {
	f := fmt.Sprintf("%v %v", firstName, lastName)
	return f
}
func add(x int, y int) int {
	return x + y
}

func mergeNameAndAddNumber(firstName string, lastName string, x int, y int) (int, string) {
	f := mergeName(firstName, lastName)
	a := add(x, y)
	return a, f
}

func (s struct1) addForStruct1() int {
	return s.a + s.b
}

func main() {
	//defer & print
	fmt.Println("connect database success!!!")
	defer fmt.Println("close database!!!")
	defer fmt.Println("log path add numbers!!!")
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	//variable
	var param1 float64 = 20.0
	param2 := 20.0
	fmt.Printf("param1 : %v param2 : %v\n", param1, param2)

	//array & loop
	arrays := [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}
	arrays2 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	for i := 0; i < len(arrays); i++ {
		fmt.Println(i, arrays[i])
	}
	for i, array := range arrays2 {
		fmt.Println(i, array)
	}
	//function & return value
	fullname := mergeName("thamrong", "chaiwong")
	addNumber := add(1, 2)
	a, f := mergeNameAndAddNumber("thamrong", "chaiwong", 1, 2)
	fmt.Printf("fullname is : %v & add numbers is : %v\n", fullname, addNumber)
	fmt.Printf("fullname is : %v & add numbers is : %v\n", f, a)
	//function & return value & struct
	variable := struct1{
		firstName: "thamrong",
		lastName:  "chaiwong",
		a:         1,
		b:         2,
	}
	fullnameX := mergeName(variable.firstName, variable.lastName)
	addNumberX := add(variable.a, variable.b)
	aX, fX := mergeNameAndAddNumber(variable.firstName, variable.lastName, variable.a, variable.b)
	fmt.Printf("fullname is : %v & add numbers is : %v\n", fullnameX, addNumberX)
	fmt.Printf("fullname is : %v & add numbers is : %v\n", fX, aX)
	//function specifically for struct
	specialFunc := variable.addForStruct1()
	fmt.Println(specialFunc)
	//pointer
	i := 42
	fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
	p := &i
	fmt.Printf("2. main  -- p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
	byval(p)
	fmt.Printf("5. main  -- p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
	fmt.Printf("6. main  -- i  %T: &i=%p i=%v\n", i, &i, i)

}