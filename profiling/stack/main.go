package main

func main() {
	example(make([]string, 2, 3), "hello", 10)
}

func example(strings []string, s string, i int) {
	panic("want stack trace")
}

