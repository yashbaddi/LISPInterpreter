package parser

import (
	"reflect"
	"testing"

	"github.com/yashbaddi/golisp/internal/lexer"
	"github.com/yashbaddi/golisp/internal/node"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{"123", 123},
		{"3.14", 3.14},
		{"-42.5", -42.5},
		{"foo", node.Symbol("foo")},
		{"( + 1 2)", node.List{node.Symbol("+"), 1, 2}},
		{"( * 2.5 4)", node.List{node.Symbol("*"), 2.5, 4}},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		res, err := p.Parse()
		if err != nil {
			t.Fatalf("Parse(%q) returned error: %v", tt.input, err)
		}
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("Parse(%q) = %v; want %v", tt.input, res, tt.expected)
		}
	}
}
