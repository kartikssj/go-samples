package main

import "fmt"
import "os"
import "strconv"

func factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <number>", os.Args[0])
	} else {
		i,e := strconv.Atoi(os.Args[1]);
		if e != nil {
			fmt.Printf("Error: %s", e)
		} else {
			fmt.Printf("%d\n",factorial(i))
		}
	}
}
