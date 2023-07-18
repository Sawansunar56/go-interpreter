package ast

import (
	"example/sawan/goInterpreter/token"
)

// Every interface must implement this node so that they implement the token literal and have a string
type Node interface {
	TokenLiteral() string
}

// Statements produce no value while Expressions produce value
// Expressions produce value but Statements don't
// e.g - For monkey Language, let is a statement, and inside it variable name is an Identifier and the value is the expression
type Statement interface {
	Node
	statementNode()
}

// Expressions produce value but Statements don't
// e.g - For monkey Language, let is a statement, and inside it variable name is an Identifier and the value is the expression
type Expression interface {
	Node
	expressionNode()
}

// Program node is going to be the root node of every AST our parser produces
// Every valid Monkey program is a series of statments. These statements are
// contained in the Program.Statements, which is a slice of AST nodes that implement
// the Satement interface.
type Program struct {
	Statements []Statement
}

// Returns the first statement from the program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
