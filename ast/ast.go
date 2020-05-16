package ast

type Node interface {
	TokenListeral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// root Node of every statement
type Program struct {
	Statements []Statement
}

func (p *Program) TokenListeral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenListeral()
	} else {
		return ""
	}
}
