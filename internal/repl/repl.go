package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yashbaddi/golisp/internal/lexer"
	"github.com/yashbaddi/golisp/internal/parser"
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
		p := parser.NewParser(lex)

		ast, err := p.Parse()
		if err != nil {
			fmt.Fprintln(out, err)
			break
		}

		fmt.Fprintf(out, "%q\n", ast)

	}
}
