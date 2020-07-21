package main

import (
	"fmt"
	"showmax/cgo/rand"
)

func main() {
	rand.Seed(100)
	fmt.Println(rand.Random())
}
