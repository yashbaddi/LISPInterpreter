package ast

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

func (NumberLiteral) String() string {
	return ""
}
func (Identifier) String() string {
	return ""

}
func (List) String() string {
	return ""

}
