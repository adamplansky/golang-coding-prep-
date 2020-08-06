package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go count("hello", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("thing %d", i) //sending is blocking operation
		time.Sleep(time.Millisecond * 500)
	}
	close(c) //always close channel from sending site
}
