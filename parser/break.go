package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseBreak() *ast.Break {
	stmt := &ast.Break{Token: p.currToken}
	for p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
