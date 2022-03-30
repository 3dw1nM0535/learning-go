package main

import "fmt"

const helloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloPrefix + name
}
