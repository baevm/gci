package repl

import (
	"bufio"
	"fmt"
	"gci/lexer"
	"gci/token"
	"io"
)

const prompt = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, prompt)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tkn := l.NextToken(); tkn.Type != token.EOF; tkn = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tkn)
		}
	}
}
