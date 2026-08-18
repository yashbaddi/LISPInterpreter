package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yashbaddi/golisp/internal/evaluate"
	"github.com/yashbaddi/golisp/internal/lexer"
	"github.com/yashbaddi/golisp/internal/parser"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(out, err)
	}

	env := evaluate.NewGlobalEnv()

	for {
		fmt.Fprint(out, "lisp > ")

		if !scanner.Scan() {
			return
		}
		input := scanner.Text()

		lex := lexer.NewLexer(input)
		p := parser.NewParser(lex)

		parsed, err := p.Parse()
		if err != nil {
			fmt.Fprintln(out, err)
			continue
		}

		result, err := evaluate.Eval(parsed, env)
		if err != nil {
			fmt.Fprintln(out, err)
			continue
		}

		fmt.Fprintf(out, "%v\n", result)
	}
}
