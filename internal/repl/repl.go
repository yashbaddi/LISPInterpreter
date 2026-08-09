package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yashbaddi/golisp/internal/lexer"
	"github.com/yashbaddi/golisp/internal/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(out, err)
	}

	for {
		fmt.Fprint(out, "lisp > ")

		if !scanner.Scan() {
			return
		}
		input := scanner.Text()

		lex := lexer.NewLexer(input)

		for {
			tok, err := lex.NextToken()
			if err != nil {
				fmt.Println(out, err)
				break
			}
			if tok.Type == token.EOF {
				break
			}

			fmt.Fprintf(out, "%v %q\n", tok.Type, tok.Literal)
		}
	}
}
