package ast

import "miguenz-lang/token"

type Node interface {
	TokenLiteral() string
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
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Identifier implements Expression to simplify the reuse of variable name bind as an Expression after
type Identifier struct {
	Token token.Token
	Value string
}

// LetStatement Name holds the Name of the binding also the Expression that produces value
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}
