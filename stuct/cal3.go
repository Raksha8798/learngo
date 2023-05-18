package main

import "fmt"

type calculator struct {
	number1, number2 int
}
type calculation interface {
	add() int
	sub() int
	mul() int
	div() int
}

func (cc calculator) Add() int {
	return cc.number1 + cc.number2
}
func (cc calculator) Sub() int {
	return cc.number1 - cc.number2
}
func (cc calculator) Multi() int {
	return cc.number1 * cc.number2
}
func (cc calculator) Div() int {
	return cc.number1 / cc.number2
}
func main() {

	c := calculator{12, 10}
	fmt.Println(c.Add())
	fmt.Println(c.Sub())
	fmt.Println(c.Multi())
	fmt.Println(c.Div())

}
