package repl

import (
	"HarshilBhatia/language/lexer"
	"HarshilBhatia/language/token"
	"bufio"
	"fmt"
)

const PROMPT = ">>"

func start(in io.reader, out io.writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
