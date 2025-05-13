package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main(){
	sli := make([]int, 0, 3)
	for {
		var input string
		fmt.Printf("Enter Number to be inserted:")
		fmt.Scanf("%s", &input)
		if input == "X"{
			break
		}
		num, _ := strconv.Atoi(input)
		sli = append(sli, num)
		slices.Sort(sli)
		fmt.Println("slice:", sli)
	}
}