package lexer

import (
	"bufio"
	"fmt"
	"io"
)

// PROMPT is the REPL prompt
const PROMPT = ">> "

// RunRepl runs a REPL for the given input and output.
func RunRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	fmt.Fprintf(writer, PROMPT)
	writer.Flush()
	for scanner.Scan() {
		line := scanner.Text()
		l := NewLexer(line)

		for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
			fmt.Fprintf(writer, "%+v\n", tok)
		}
		fmt.Fprintf(writer, PROMPT)
		writer.Flush()
	}
}
