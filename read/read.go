// Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line. 
// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).
// Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
)

type name struct {
	fname string
	lname string
}

func truncate(s string, max int) string {
	s = strings.TrimSpace(s)
	if len(s) >  max {
		return s[:max]
	}
	return s
}

func main() {
	var Names []name
	fmt.Printf("Enter File Name: ")
	reader := bufio.NewReader(os.Stdin)
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader = bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		words := strings.Fields(line)
		if len(words) < 2 {
			continue
		}
		firstname := truncate(words[0], 20)
		lastname := truncate(words[1], 20)
		Names = append(Names, name{fname: firstname, lname: lastname})
	}

	for _, person := range Names{
		fmt.Printf("%s %s \n", person.fname, person.lname)
	}

}