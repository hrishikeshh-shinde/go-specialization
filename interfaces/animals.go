/* Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

Animal Food eaten Locomtion method Spoken sound
cow		grass		walk			moo
bird	worms		fly 			peep
snake	mice		slither			hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.
Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.
Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.
Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
*/ 

package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {}
type Bird struct {}
type Snake struct {}

func (a *Cow) Eat() {fmt.Println("grass")}
func (a *Cow) Move() {fmt.Println("walk")}
func (a *Cow) Speak() {fmt.Println("moo")}

func (a *Bird) Eat() {fmt.Println("worms")}
func (a *Bird) Move() {fmt.Println("fly")}
func (a *Bird) Speak() {fmt.Println("peep")}

func (a *Snake) Eat() {fmt.Println("mice")}
func (a *Snake) Move() {fmt.Println("slither")}
func (a *Snake) Speak() {fmt.Println("hsss")}


func main() {
	animals := map[string]Animal {}
	for {
		fmt.Printf(">")
		var w1,w2,w3 string
		fmt.Scanf("%s %s %s", &w1, &w2, &w3)
		switch w1{
		case "newanimal":
			switch w3{
			case "cow":
				animals[w2] = &Cow{}
			case "bird":
				animals[w2] = &Bird{}
			case "snake":
				animals[w2] = &Snake{}
			}
			fmt.Println("Created it!")
		case "query":
			_, ok := animals[w2]
			if !ok {
				fmt.Println("Animal not found")
				break
			}
			switch w3{
			case "eat":
				animals[w2].Eat()
			case "move":
				animals[w2].Move()
			case "speak":
				animals[w2].Speak()
			default:
				fmt.Println("Unknown query")
			}
		}

	}
}