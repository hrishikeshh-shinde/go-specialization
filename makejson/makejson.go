// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
	"strings"
)

func main() {
	fmt.Printf("Enter Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Enter Address: ")
	addr, _ := reader.ReadString('\n')
	addr = strings.TrimSpace(addr)
	m := make(map[string]string)
	m["name"] = name
	m["address"] = addr

	mp, _ := json.Marshal(m)
	fmt.Println(string(mp))
}