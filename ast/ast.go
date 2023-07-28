package ast

import (
	"bytes"

	"example/sawan/goInterpreter/token"
)

// Every interface must implement this node so that they implement the token literal and have a string
type Node interface {
	TokenLiteral() string
	String() string
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
	// It is an array of Statement interface, that implements the node and statementNode.
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

// Loops through the created buffer and stores the statements and returns said buffer
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// Currently does nothing specific.
func (i *Identifier) expressionNode() {}

// Currently does nothing specific. Returns the literal value inside the token
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Basic Return Type to hold return values
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// Currently does nothing specific.
func (rs *ReturnStatement) statementNode() {}

// Currently does nothing specific. Returns the literal value inside the token
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Basic Expression type to hold expression values
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// Currently does nothing specific.
func (es *ExpressionStatement) statementNode() {}

// Currently does nothing specific. Returns the literal value inside the token
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// just returns the value of the identifier
func (i *Identifier) String() string { return i.Value }

// Basic Integer Type to hold integer values
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type PrefixExpression struct {
  Token token.Token
  Operator string
  Right Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string { 
  var out bytes.Buffer

  out.WriteString("(")
  out.WriteString(pe.Operator)
  out.WriteString(pe.Right.String())
  out.WriteString(")")

  return out.String()
}

