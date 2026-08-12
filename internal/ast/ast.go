package ast

import (
	"strconv"
	"strings"
)

type Node interface {
	String() string
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Expression Expression
}

func (p Program) String() string {
	if p.Expression != nil {
		return p.Expression.String()
	}
	return ""
}

type List struct {
	Elements []Expression
}

type NumberLiteral struct {
	Value int
}

type Identifier struct {
	Value string
}

func (NumberLiteral) expressionNode() {}
func (Identifier) expressionNode()    {}
func (List) expressionNode()          {}

func (n NumberLiteral) String() string {
	return strconv.Itoa(n.Value)
}

func (i Identifier) String() string {
	return i.Value
}

func (l List) String() string {
	var out strings.Builder
	out.WriteString("(")

	elems := make([]string, len(l.Elements))
	for idx, el := range l.Elements {
		if el != nil {
			elems[idx] = el.String()
		}
	}
	out.WriteString(strings.Join(elems, " "))

	out.WriteString(")")
	return out.String()
}
