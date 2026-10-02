package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var previous int = 1
	var previousPrevious int = -2
	var number int
	
	return func() int {
		if previousPrevious == -2 {
			previousPrevious = -1
			return 0
		} else if previousPrevious == -1 {
			previousPrevious = 0
			return 1
		} else {
			number = previousPrevious + previous
			previousPrevious = previous
			previous = number
			return number
		
		}
		
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}