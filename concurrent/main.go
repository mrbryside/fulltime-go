package main

import (
	"fmt"
	"time"
)

func main() {
	resultch := make(chan string, 1) // -> unuffered channel
	resultch <- "foo"                // -> is now FULL -> IT WILL BLOCK -> BLOCK HERE

	// this code below will never execute!!!!
	result := <-resultch
	fmt.Println(result)
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("results %d", n)
}
