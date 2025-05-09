package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Muhammad-Sabir/monkeyplus-go/lexer"
	"github.com/Muhammad-Sabir/monkeyplus-go/token"
)

const PROMPT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Fprint(output, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.NewLexer(line)

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Fprintf(output, "%+v\n", tok)
		}
	}
}
