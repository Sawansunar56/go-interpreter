package parser

import (
	"example/sawan/goInterpreter/ast"
	"example/sawan/goInterpreter/lexer"
	"example/sawan/goInterpreter/token"
	"fmt"
)

/*
L - is the lexer in which next token is called repeatedly
curtoken - holds the postion of the current current token
peekToken - holds the positon of the next token
*/
type Parser struct {
	l *lexer.Lexer

  errors []string

	curToken  token.Token
	peekToken token.Token
}

// creates a new parser with default values and returns it
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
  return p.errors
}


// helper function to move forward in a parser
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Makes the root program and then adds every statement into it.
//
// Iterates over the tokens until it meets and EOF token and goes through the token
// the nextToken Method
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
  case token.RETURN:
    return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// checks whether the current value is the provided token
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
    p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
  msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
  p.errors = append(p.errors, msg)
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
  stmt := &ast.ReturnStatement{Token: p.curToken}

  p.nextToken()

  for !p.curTokenIs(token.SEMICOLON) {
    p.nextToken()
  }

  return stmt
}
