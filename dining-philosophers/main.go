// Implement the dining philosopher’s problem with the following constraints/modifications.

// There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// The host allows no more than 2 philosophers to eat concurrently.
// Each philosopher is numbered, 1 through 5.
// When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
// When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup, num int, host chan struct{}) {
	for i:=0; i<3; i++{
		host <- struct{}{}
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ",num)
		fmt.Println("finishing eating ",num)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		<- host
	}
	wg.Done()
}

func main(){
	fmt.Println("Dining Philosophers Problem")
	CSticks := make([]*ChopS, 5)
	for i:=0; i<5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i:=0; i<5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}
	host := make(chan struct{}, 2)
	var wg sync.WaitGroup
	for i:=0; i<5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg, i+1, host)
	}
	wg.Wait()
}