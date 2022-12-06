package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	for _, v := range strings.Split(msg, " ") {
		str := ""
		for ii, vv := range strings.Split(v,"") {
			str += strings.Repeat(vv,(ii+1))
		}
		print(str)
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}