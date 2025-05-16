package main

import (
	"fmt"
)

func GenDisplaceFn( a, v, s float64) func(float64) float64 {
	fn := func (t float64) float64  {
		dis := (0.5 * a * t * t) + (v * t) + s
		return dis
	}
	return fn
}

func main() {
	var a, v, s, t float64
	fmt.Printf("Enter acceleration: ")
	fmt.Scanf("%f\n", &a)
	fmt.Printf("Enter initial velocity: ")
	fmt.Scanf("%f\n", &v)
	fmt.Printf("Enter initial displacement: ")
	fmt.Scanf("%f\n", &s)
	fmt.Printf("Enter time: ")
	fmt.Scanf("%f\n", &t)
	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(t))
}