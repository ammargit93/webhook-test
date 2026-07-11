package main

import (
	"flag"
	"fmt"
	"time"
)

func greeting() string {
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		return "Good Morning"
	case hour < 18:
		return "Good Afternoon"
	default:
		return "Good Evening"
	}
}

func main() {
	name := flag.String("name", "World", "Name to greet")
	flag.Parse()

	fmt.Printf("%s, %s!\n", greeting(), *name)
}
