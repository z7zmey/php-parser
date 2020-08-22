package visitor

import (
	"github.com/z7zmey/php-parser/pkg/ast"
)

type FilterTokens struct {
	Null
}

func (v *FilterTokens) EnterNode(n ast.Vertex) bool {
	n.GetNode().Tokens = nil
	n.Accept(v)
	return true
}

func (v *FilterTokens) StmtUse(n *ast.StmtUse) {
	n.UseTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtGroupUse(n *ast.StmtGroupUse) {
	n.UseTkn = nil
	n.LeadingNsSeparatorTkn = nil
	n.NsSeparatorTkn = nil
	n.OpenCurlyBracketTkn = nil
	n.CloseCurlyBracketTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtUseDeclaration(n *ast.StmtUseDeclaration) {
	n.NsSeparatorTkn = nil
	n.AsTkn = nil
	n.CommaTkn = nil
}
