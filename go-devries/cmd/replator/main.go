package main

import (
	"fmt"
	"os"

	"monkey/lexer"
)

func main() {
	fmt.Println("Don't look now, there's a monkey on your back!")
	lexer.RunRepl(os.Stdin, os.Stdout)
}
