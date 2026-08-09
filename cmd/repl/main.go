package main

import (
	"os"

	"github.com/yashbaddi/golisp/internal/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
