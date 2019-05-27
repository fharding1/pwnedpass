package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fharding1/pwnedpass"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <password>\n", os.Args[0])
		os.Exit(1)
	}

	count, err := pwnedpass.Count(context.Background(), os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(count)
}
