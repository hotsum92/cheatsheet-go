package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("arg %d: %s\n", i, arg)
	}
}
