package main

import "fmt"

func main() {

	for i := 1; i <= 6; i++ {

		if i == 6 {
			for j := 1; j <= i; j++ {
				fmt.Printf("@\t")
			}
			fmt.Println()
		} else if i%2 == 0 {
			for j := 1; j <= i; j++ {
				fmt.Printf("#\t")
			}
			fmt.Println()
		} else {
			for j := 1; j <= i; j++ {
				fmt.Printf("*\t")
			}
			fmt.Println()
		}
	}
}
