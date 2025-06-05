// Copyright (c) 2025. guimochila.com. Continuous Learning.

package ast

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

type Program struct {
	Statemments []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statemments) > 0 {
		return p.Statemments[0].TokenLiteral()
	} else {
		return ""
	}
}
