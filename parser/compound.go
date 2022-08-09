package parser

import "pilang/ast"

func (p *Parser) parseCompound(left ast.Expression) ast.Expression {
	expression := &ast.Compound{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
		Left:     left,
	}
	precedence := p.currPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}
