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