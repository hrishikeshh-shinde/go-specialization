package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter String:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	startsWithI := strings.HasPrefix(input, "i")
	endsWithN := strings.HasSuffix(input, "n")
	containsA := strings.Contains(input, "a")
	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}