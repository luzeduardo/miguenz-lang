package parser

import (
	"miguenz-lang/ast"
	"miguenz-lang/lexer"
	"miguenz-lang/token"
)

// The main idea here is to parse each token generated by the Lexer
// with a context free grammar (CFG) as input and parse it using a recursive descent parser
// This simple parser example uses top down operator precedence also called Prat Parser
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	return program
}