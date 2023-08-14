package ast

import (
	"bytes"
	"strings"

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
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

/*
*Condition - Expression, that holds the condition

*Consequence - if condition block

*Alternative - else condtion block
 */
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
  var out bytes.Buffer
  args := []string{}

  for _, a := range ce.Arguments {
    args = append(args, a.String())
  }
  out.WriteString(ce.Function.String())
  out.WriteString("(")
  out.WriteString(strings.Join(args, ", "))
  out.WriteString(")")

  return out.String()
}
