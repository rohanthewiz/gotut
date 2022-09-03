package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	/*
		var a int
		var b int
		var c int
		var f int

		a = 5
		b = 7
		c = 99
		f = 21
	*/
	// fmt.Println("The sum is:", a+b+c+f)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("My error:", err)
		return
	}

	// Hash the input
	sum := sha256.Sum256([]byte(input + "sgyuff$$$!!!~)"))
	fmt.Printf("%x\n", sum)
	// fmt.Printf("Hash: %x\n", sum)
}
