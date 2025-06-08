package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func BubbleSort(sli []int, wg *sync.WaitGroup) {
	n := len(sli)
	var swapped bool

	for range n-1 {
		swapped = false
		for j := 0; j < n-1; j++ {
			if(sli[j] > sli[j+1]){
				Swap(sli, j)
				swapped = true
			}
		}
		
		if !swapped {
			break
		}
	}
	wg.Done()
}

func min(a, b, c, d int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	if d < m {
		m = d
	}
	return m
}

func main () {
	fmt.Printf("Enter Sequence of Integers:")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	numbers := []int{}
	for _, token := range strings.Split(userInput, " ") {
		number, _ := strconv.Atoi(token)
		numbers = append(numbers, number)
	}
	var wg sync.WaitGroup
	n := len(numbers)
	part1 := numbers[0:n/4]
	part2 := numbers[n/4:n/2]
	part3 := numbers[n/2:3*n/4]
	part4 := numbers[3*n/4:]
	wg.Add(1)
	go BubbleSort(part1, &wg)
	wg.Add(1)
	go BubbleSort(part2, &wg)
	wg.Add(1)
	go BubbleSort(part3, &wg)
	wg.Add(1)
	go BubbleSort(part4, &wg)
	wg.Wait()
	var res []int
	i, j, k, l := 0, 0, 0, 0
	n1, n2, n3, n4 := len(part1), len(part2), len(part3), len(part4)
	a, b, c, d := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	if i < n1 { a = part1[i] }
	if j < n2 { b = part2[j] }
	if k < n3 { c = part3[k] }
	if l < n4 { d = part4[l] }
	for i<n1 || j<n2 || k<n3 || l<n4 {
		tmp := min(a, b, c, d)
		if tmp == a {
			i += 1
			if i < n1{
				a = part1[i]
			} else {
				a = math.MaxInt
			}
			
		} else if tmp == b {
			j += 1
			if j < n2{
				b = part2[j]
			} else {
				b = math.MaxInt
			}
			
		} else if tmp == c {
			k += 1
			if k < n3{
				c = part3[k]
			} else {
				c = math.MaxInt
			}
			
		} else if tmp == d {
			l += 1
			if l < n4{
				d = part4[l]
			} else {
				d = math.MaxInt
			}
			
		}
		res = append(res, tmp)
	}
	fmt.Println(res)
}