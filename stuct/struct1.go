package main

import "fmt"

func menu() {
	var choice int
	fmt.Println("1. For Add")
	fmt.Println("2. For Sub")
	fmt.Println("3. For Multi")
	fmt.Println("2. For Div")
	fmt.Println("Enter your choice:")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		Add()
		menu()
	case 2:
		Sub()
		menu()
	case 3:
		Multi()
		menu()
	case 4:
		Div()
		menu()

	}
}
func Add() {
	var a, b int
	fmt.Println("Enter the first number:")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter the second number:")
	fmt.Scanf("%d", &b)
	Add := a + b
	fmt.Println("The result is:", Add)
}
func Sub() {
	var a, b int
	fmt.Println("Enter the first number:")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter the second number:")
	fmt.Scanf("%d", &b)
	Sub := a - b
	fmt.Println("The result is:", Sub)
}
func Multi() {
	var a, b int
	fmt.Println("Enter the first number:")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter the second number:")
	fmt.Scanf("%d", &b)
	Multi := a * b
	fmt.Println("The result is:", Multi)
}
func Div() {
	var a, b int
	fmt.Println("Enter the first number:")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter the second number:")
	fmt.Scanf("%d", &b)
	Div := a / b
	fmt.Println("The result is:", Add)
}

func main() {
	menu()

}
