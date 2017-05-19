package main

import (
	"fmt"
	"./boring"
)


func main() {
	c := make(chan string)
	go boring.Boring10("boring!", c)
	for i := 0; ; i++ {
		fmt.Printf("You say: %q\n", <-c) 
	}
	fmt.Println("You're boring; I'm leaving.")
}
