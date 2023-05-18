package main

import "fmt"

type cal struct {
	number1, number2 int
}

func (c *cal) Add() {
	Add := c.number1 + c.number2
	fmt.Println("The result is:", Add)

}
func (c *cal) Sub() {
	Sub := c.number1 - c.number2
	fmt.Println("The result is:", Sub)

}
func (c *cal) Multi() {
	Multi := c.number1 * c.number2
	fmt.Println("The result is:", Multi)

}
func (c *cal) Div() {
	Div := c.number1 / c.number2
	fmt.Println("The result is:", Div)

}

func main() {
	var number1, number2 int
	fmt.Println("enter number1:")
	fmt.Scanf("%d", &number1)
	fmt.Println("enter number2:")
	fmt.Scanf("%d", &number2)
	s := cal{
		number1: number1,
		number2: number2,
	}
	var choice int
	fmt.Println("1. For Add")
	fmt.Println("2. For Sub")
	fmt.Println("3. For Multi")
	fmt.Println("4. For Div")
	fmt.Println("Enter your choice:")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		s.Add()

	case 2:
		s.Sub()

	case 3:
		s.Multi()

	case 4:
		s.Div()

	}

}
