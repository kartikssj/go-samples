package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
)

func roll() (int, float32) {
	return (rand.Intn(5) + 1), 1.0/6.0
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " <number-of-dice>")
	} else {
		t,e := strconv.Atoi(os.Args[1])
		if e != nil {
			fmt.Println("Error: ", e)
		} else {
			var prob float32 = 1.0
			for i:=0; i<t; i++ {
				i,p := roll()
				prob *= p
				fmt.Printf("%d ",i)
			}
			fmt.Printf("%f",prob)
		}
	}
}
