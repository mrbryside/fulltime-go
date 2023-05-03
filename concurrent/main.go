package main

import (
	"fmt"
	"time"
)

func main() {
	resultch := make(chan string) // -> unuffered channel

	go func() {
		result := <-resultch
		fmt.Println(result)
	}()

	resultch <- "foo"
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("results %d", n)
}
