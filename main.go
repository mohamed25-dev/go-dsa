package main

import (
	"fmt"

	"github.com/mohamed25-dev/go-dsa/stack"
)

func main() {
	res := stack.IsValid("(){{}}")
	fmt.Println("res: ", res)
}
