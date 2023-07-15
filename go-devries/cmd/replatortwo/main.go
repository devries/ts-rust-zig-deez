package main

import (
	"fmt"
	"os"

	"monkey/repl"
)

func main() {
	fmt.Println("Don't look now, there's a monkey on your back!")
	repl.Start(os.Stdin, os.Stdout)
}
