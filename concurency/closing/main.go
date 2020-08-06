package main

import "fmt"

func main() {
	c := make(chan int)
	close(c)
	v, ok := <-c
	fmt.Println(v, ok) //0, false
	//receiver always returns two values
	//0 as it is the zero value of int
	//false because "no more data", "returned value is not valid"
}
