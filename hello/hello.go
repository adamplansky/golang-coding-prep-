package hello

import "fmt"

func hello(text string) string {
	if text == "" {
		return "hello"
	}

	return fmt.Sprintf("hello %s", text)
}
