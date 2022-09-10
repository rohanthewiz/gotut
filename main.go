package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	name := "John"
	fmt.Println("name is:", name)

	aNum := 1
	fmt.Println(aNum)

	var anotherNum int = 2
	println(anotherNum)

	nameN := "Noah"
	println(nameN)

	var name2 string
	name2 = "christian"
	println(name2)

	// Take some input and hash it
	var input string
	print("Please enter a word to hash ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("My error:", err)
		return
	}

	// Hash the input
	sum := sha256.Sum256([]byte(input + "sgyuff$$$!!!~)"))
	fmt.Printf("%x\n", sum)
	// fmt.Printf("Hash: %x\n", sum)

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
}
