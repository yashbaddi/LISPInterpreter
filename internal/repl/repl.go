package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yashbaddi/golisp/internal/evaluate"
	"github.com/yashbaddi/golisp/internal/lexer"
	"github.com/yashbaddi/golisp/internal/parser"
)

const (
	colorReset  = "\033[0m"
	colorPrompt = "\033[1;33m" // Bold Amber / Yellow
	colorOutput = "\033[36m"   // Cyan
	colorError  = "\033[1;31m" // Bold Red
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(out, "%s%v%s\n", colorError, err, colorReset)
	}

	env := evaluate.NewGlobalEnv()

	for {
		fmt.Fprintf(out, "%slisp > %s", colorPrompt, colorReset)

		if !scanner.Scan() {
			return
		}
		input := scanner.Text()

		lex := lexer.NewLexer(input)
		p := parser.NewParser(lex)

		parsed, err := p.Parse()
		if err != nil {
			fmt.Fprintf(out, "%s%v%s\n", colorError, err, colorReset)
			continue
		}

		result, err := evaluate.Eval(parsed, env)
		if err != nil {
			fmt.Fprintf(out, "%s%v%s\n", colorError, err, colorReset)
			continue
		}

		fmt.Fprintf(out, "%s%v%s\n", colorOutput, result, colorReset)
	}
}
