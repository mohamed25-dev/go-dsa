package main

import (
	"fmt"

	"github.com/mohamed25-dev/go-dsa/queue"
)

func main() {

	var q queue.Queue

	for i := 1; i <= 3; i++ {
		q.Enqueue(i)
	}

	for i := 1; i <= 5; i++ {
		front, ok := q.Peek()
		q.Dequeue()
		size := q.Size()
		fmt.Println(ok, front, size)
	}
}
