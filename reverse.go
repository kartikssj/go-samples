package main

import (
	"fmt"
)

func main() {
	var list [3]string
	var i int = 0
	var line string = ""
	for {
		if i == len(list) { break }
		fmt.Scanf("%s",&line)
		if line == "END" {
			break
		} else {
			list[i] = line
			i++
		}
	}
	i--
	for ;i>=0;i-- {
		fmt.Println(list[i])
	}
}

