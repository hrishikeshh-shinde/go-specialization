// Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

package main

import "fmt"

func main(){
	var num1 float64
	var num2 int
	for {
		fmt.Printf("Enter number: ")
		fmt.Scanf("%g", &num1)
		num2 = int(num1)
		fmt.Printf("Truncated number:%d\n", num2)
	}
}