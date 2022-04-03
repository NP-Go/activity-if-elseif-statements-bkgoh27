package main

import "fmt"

func main() {
	//Insert your code here
	//Hint if a:= ?? ; condition {  }
	a := 17

	if a%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if a >= 10 {
		fmt.Println("It has more than 1 digit!")
	} else {
		fmt.Println("It has only 1 digit!")
	}
}
