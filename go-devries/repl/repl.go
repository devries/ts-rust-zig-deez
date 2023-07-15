package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/parser"
)

var Prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	for {
		fmt.Fprintf(writer, Prompt)
		writer.Flush()

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors) != 0 {
			printParseErrors(writer, p.Errors)
			continue
		}

		io.WriteString(writer, program.String())
		io.WriteString(writer, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
