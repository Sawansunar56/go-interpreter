package parser

import (
	"example/sawan/goInterpreter/ast"
	"example/sawan/goInterpreter/lexer"
	"example/sawan/goInterpreter/token"
)

/*
L - is the lexer in which next token is called repeatedly
curtoken - holds the postion of the current current token
peekToken - holds the positon of the next token
*/
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// creates a new parser with default values
func New(l *lexer.Lexer) *Parser {
  p := &Parser{l: l}

  p.nextToken()
  p.nextToken()
  
  return p
}

// helper function to move forward in a parser
func (p *Parser) nextToken() {
  p.curToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  return nil
}
