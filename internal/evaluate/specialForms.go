package evaluate

import (
	"fmt"

	"github.com/yashbaddi/golisp/internal/node"
)

func isSpecialForm(sym node.Symbol) bool {
	switch sym {
	case "define", "if":
		return true
	default:
		return false
	}
}

func evalSpecialForm(sym node.Symbol, args node.List, env *Env) (any, error) {
	switch sym {
	case "define":
		return evalDefine(args, env)
	case "if":
		return evalIf(args, env)
	default:
		return nil, fmt.Errorf("unknown special form: %s", sym)
	}
}

func evalDefine(args node.List, env *Env) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("define requires exactly 2 arguments, got %d", len(args))
	}

	sym, ok := args[0].(node.Symbol)
	if !ok {
		return nil, fmt.Errorf("define target must be a symbol, got %T", args[0])
	}

	val, err := Eval(args[1], env)
	if err != nil {
		return nil, err
	}

	env.Set(sym, val)
	return val, nil
}

func isTruthy(val any) bool {
	if val == nil {
		return false
	}
	if b, ok := val.(bool); ok {
		return b
	}
	return true
}

func evalIf(args node.List, env *Env) (any, error) {
	if len(args) < 2 || len(args) > 3 {
		return nil, fmt.Errorf("if requires 2 or 3 arguments, got %d", len(args))
	}

	testVal, err := Eval(args[0], env)
	if err != nil {
		return nil, err
	}

	if isTruthy(testVal) {
		return Eval(args[1], env)
	}

	if len(args) == 3 {
		return Eval(args[2], env)
	}

	return nil, nil
}
