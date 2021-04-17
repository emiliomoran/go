/*
Go routine is a separately running process or thread. It's some
function that is executing concurrently with the rest of a Go program

A channel is Go's fundamental communication mechanism. Go has the
ability to pass messages between Go routines

chan string, means that in the channel I go to pass a string

c <- word, means that the word is pass to the channel
*/

package main

import "fmt"

func emit(c chan string) {
	words := []string{"The", "quick", "brown", "fox"}

	for _, word := range words {
		c <- word
	}

	close(c)
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)

	for word := range wordChannel {
		fmt.Printf("Word: %s\n", word)
	}
}
