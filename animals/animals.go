package main

import (
	"fmt"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	animals := map[string]Animal {
		"cow" : cow,
		"bird" : bird,
		"snake" : snake,
	}

	for {
		fmt.Printf(">")
		var name, request string
		fmt.Scanf("%s %s", &name, &request)
		a, ok := animals[name]
		if !ok {
			fmt.Println("Unknown animal")
			continue
		}
		switch request{
		case "eat":
			fmt.Println(a.Eat())
		case "move":
			fmt.Println(a.Move())
		case "speak":
			fmt.Println(a.Speak())
		default:
			fmt.Println("Unknown request")
		}
	}
}