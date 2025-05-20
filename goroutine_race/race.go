//Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

package main

import (
	"fmt"
	"time"
)

func addOne(x *int){
	for i:=0; i<5; i++ {
		fmt.Printf("x=%d\n", *x)
		time.Sleep(100 * time.Millisecond)
		*x += 1
	}
}

func addTwo(x *int){
	for i:=0; i<5; i++ {
		fmt.Printf("x=%d\n", *x)
		time.Sleep(100 * time.Millisecond)
		*x += 2
	}
}

func main(){
	x := 1
	go addOne(&x)
	go addTwo(&x)

	time.Sleep(1 * time.Second)
	fmt.Println("Main Done")
}

/* Explaination:
This program contains a race condition because two goroutines, addOne and addTwo, 
access and modify the same shared variable x concurrently without any synchronization.

Each goroutine:
-> Reads the value of x
->Waits for 100 milliseconds
->Then increments x (by 1 or 2)

Since these operations (read, modify, and write) are not atomic and happen at the same time in both goroutines, 
they can interfere with each other. For example:
-> One goroutine might read x while the other is modifying it
->Both may read the same value of x, increment it separately, and write back conflicting results
This causes the final value of x to be unpredictable and different every time you run the program. 
The race condition leads to incorrect updates and inconsistent output.

*/