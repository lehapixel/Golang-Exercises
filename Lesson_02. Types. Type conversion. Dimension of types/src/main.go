/* Lesson_02. Types. Type conversion. Dimension of types. */

package main

import (
	"fmt"
	"unsafe" // Variable size in bytes
)

func main() {
	/*
		// string (Initial value = "")
		//var hello string
		var hello string = "hello"
		fmt.Println(hello)
		//hello = "hello"
		fmt.Printf("Type: %T Value: %v\n", hello, hello)
	*/

	/*
		// bool (Initial value = false)
		//var ourBool bool
		ourBool := true
		fmt.Println(ourBool)
		ourBool = true
		fmt.Printf("Type: %T Value: %v\n", ourBool, ourBool)
	*/

	/*
		// fmt
		name := "World"
		hello := fmt.Sprintf("hello %s", name)
		//_ = hello // Blank identifier!!!
		fmt.Println(hello)
		fmt.Printf("Type: %T Value: %v\n", name, name)
	*/

	/*
		// Type conversion
		var num1 int64 = 15
		var num2 int = 15
		fmt.Println(num1 + int64(num2))
	*/

	// Dimension of types
	var num1 uint8 = 1
	var num2 uint64 = 1
	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
}
