package main

import (
	"flag"
	"fmt"
	"time"
)

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	yellow = "\033[33m"
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

func isWeekend() bool {
	day := time.Now().Weekday()
	return day == time.Saturday || day == time.Sunday
}

func buildMessage(name string) string {
	msg := fmt.Sprintf("%s, %s!", greeting(), name)

	if isWeekend() {
		msg += " 🎉 Enjoy your weekend!"
	}

	return msg
}

func printMessage(msg string, color bool) {
	if color {
		fmt.Println(green + msg + reset)
		return
	}
	fmt.Println(msg)
}

func main() {
	name := flag.String("name", "World", "Name to greet")
	color := flag.Bool("color", true, "Enable colored output")
	flag.Parse()

	printMessage(buildMessage(*name), *color)
}
