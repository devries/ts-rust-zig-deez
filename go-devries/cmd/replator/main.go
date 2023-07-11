package main

import (
	"fmt"
	"os"

	"monkey"
)

func main() {
	fmt.Println("Don't look now, there's a monkey on your back!")
	monkey.RunRepl(os.Stdin, os.Stdout)
}
