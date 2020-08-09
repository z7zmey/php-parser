package visitor

import (
	"github.com/z7zmey/php-parser/pkg/ast"
)

type FilterParserNodes struct {
	Null
}

func (v *FilterParserNodes) EnterNode(n ast.Vertex) bool {
	n.Accept(v)
	return true
}

func (v *FilterParserNodes) StmtGroupUseList(n *ast.StmtGroupUseList) {
	if nn, ok := n.Prefix.(*ast.ParserNsSeparator); ok {
		n.Prefix = nn.Child
	}

	if nn, ok := n.UseList.(*ast.ParserNsSeparator); ok {
		n.UseList = nn.Child
	}

	if nn, ok := n.UseList.(*ast.ParserBrackets); ok {
		n.UseList = nn.Child
	}
}

func (v *FilterParserNodes) StmtUseList(n *ast.StmtUseList) {
	for k, v := range n.UseDeclarations {
		if nn, ok := v.(*ast.ParserNsSeparator); ok {
			n.UseDeclarations[k] = nn.Child
		}
	}
}

func (v *FilterParserNodes) StmtUseDeclaration(n *ast.StmtUseDeclaration) {
	if nn, ok := n.Alias.(*ast.ParserAs); ok {
		n.Alias = nn.Child
	}
}

func (v *FilterParserNodes) StmtAltIf(n *ast.StmtAltIf) {
	if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
		n.Cond = nn.Child
	}

	if nn, ok := n.Stmt.(*ast.ParserBrackets); ok {
		n.Stmt = nn.Child
	}
}

func (v *FilterParserNodes) StmtAltElseIf(n *ast.StmtAltElseIf) {
	if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
		n.Cond = nn.Child
	}

	if nn, ok := n.Stmt.(*ast.ParserBrackets); ok {
		n.Stmt = nn.Child
	}
}

func (v *FilterParserNodes) StmtAltElse(n *ast.StmtAltElse) {
	if nn, ok := n.Stmt.(*ast.ParserBrackets); ok {
		n.Stmt = nn.Child
	}
}
