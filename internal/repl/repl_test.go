package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestREPL(t *testing.T) {
	t.Run("should evaluate expression and output result with colors", func(t *testing.T) {
		input := "(+ 1 2)\n"
		var out bytes.Buffer

		Start(strings.NewReader(input), &out)

		expectedPrompt := "\033[1;33mlisp > \033[0m"
		expectedOutput := "\033[36m3\033[0m\n"

		res := out.String()
		if !strings.Contains(res, expectedPrompt) {
			t.Errorf("expected prompt %q in output, got %q", expectedPrompt, res)
		}
		if !strings.Contains(res, expectedOutput) {
			t.Errorf("expected output %q in output, got %q", expectedOutput, res)
		}
	})

	t.Run("should print parse and evaluation errors with error colors", func(t *testing.T) {
		input := "(+\nunknown-var\n"
		var out bytes.Buffer

		Start(strings.NewReader(input), &out)

		res := out.String()
		if !strings.Contains(res, colorError) {
			t.Errorf("expected error color %q in output, got %q", colorError, res)
		}
	})
}
