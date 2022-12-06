package main

import (
	"fmt"
	"os"
)

func main() {
	var domain string
	if len(os.Args) > 1 {
		domain = os.Args[len(os.Args)-1]
	}
	fmt.Println(domain)
}
