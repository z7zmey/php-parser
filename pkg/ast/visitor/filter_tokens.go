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

func (v *FilterTokens) NameNamePart(n *ast.NameNamePart) {
	n.NsSeparatorTkn = nil
	n.StringTkn = nil
}

func (v *FilterTokens) NameName(n *ast.NameName) {
	n.ListSeparatorTkn = nil
}

func (v *FilterTokens) NameFullyQualified(n *ast.NameFullyQualified) {
	n.NsSeparatorTkn = nil
	n.ListSeparatorTkn = nil
}

func (v *FilterTokens) NameRelative(n *ast.NameRelative) {
	n.NsTkn = nil
	n.NsSeparatorTkn = nil
	n.ListSeparatorTkn = nil
}

func (v *FilterTokens) StmtNamespace(n *ast.StmtNamespace) {
	n.NsTkn = nil
	n.OpenCurlyBracket = nil
	n.CloseCurlyBracket = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtHaltCompiler(n *ast.StmtHaltCompiler) {
	n.HaltCompilerTkn = nil
	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtConstList(n *ast.StmtConstList) {
	n.ConstTkn = nil
	n.SeparatorTkns = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtClassConstList(n *ast.StmtClassConstList) {
	n.ConstTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtConstant(n *ast.StmtConstant) {
	n.EqualTkn = nil
	n.CommaTkn = nil
}

func (v *FilterTokens) StmtStmtList(n *ast.StmtStmtList) {
	n.OpenCurlyBracket = nil
	n.CloseCurlyBracket = nil
}

func (v *FilterTokens) StmtIf(n *ast.StmtIf) {
	n.IfTkn = nil
	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	n.ColonTkn = nil
	n.EndIfTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtElseIf(n *ast.StmtElseIf) {
	n.ElseIfTkn = nil
	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	n.ColonTkn = nil
}

func (v *FilterTokens) StmtElse(n *ast.StmtElse) {
	n.ElseTkn = nil
	n.ColonTkn = nil
}

func (v *FilterTokens) ParserBrackets(n *ast.ParserBrackets) {
	n.OpenBracketTkn = nil
	n.CloseBracketTkn = nil
}

func (v *FilterTokens) StmtWhile(n *ast.StmtWhile) {
	n.WhileTkn = nil
	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	n.ColonTkn = nil
	n.EndWhileTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtDo(n *ast.StmtDo) {
	n.DoTkn = nil
	n.WhileTkn = nil
	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtFor(n *ast.StmtFor) {
	n.ForTkn = nil
	n.OpenParenthesisTkn = nil
	n.InitSemiColonTkn = nil
	n.CondSemiColonTkn = nil
	n.CloseParenthesisTkn = nil
	n.ColonTkn = nil
	n.EndForTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtSwitch(n *ast.StmtSwitch) {
	n.SwitchTkn = nil
	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	n.OpenCurlyBracketTkn = nil
	n.CaseSeparatorTkn = nil
	n.ColonTkn = nil
	n.CloseCurlyBracketTkn = nil
	n.EndSwitchTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtCase(n *ast.StmtCase) {
	n.CaseTkn = nil
	n.CaseSeparatorTkn = nil
}

func (v *FilterTokens) StmtDefault(n *ast.StmtDefault) {
	n.DefaultTkn = nil
	n.CaseSeparatorTkn = nil
}

func (v *FilterTokens) StmtBreak(n *ast.StmtBreak) {
	n.BreakTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtContinue(n *ast.StmtContinue) {
	n.ContinueTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtReturn(n *ast.StmtReturn) {
	n.ReturnTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtGlobal(n *ast.StmtGlobal) {
	n.GlobalTkn = nil
	n.SeparatorTkns = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtStatic(n *ast.StmtStatic) {
	n.StaticTkn = nil
	n.SeparatorTkns = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtStaticVar(n *ast.StmtStaticVar) {
	n.EqualTkn = nil
}

func (v *FilterTokens) StmtEcho(n *ast.StmtEcho) {
	n.EchoTkn = nil
	n.SeparatorTkns = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtInlineHtml(n *ast.StmtInlineHtml) {
	n.InlineHtmlTkn = nil
}

func (v *FilterTokens) StmtUnset(n *ast.StmtUnset) {
	n.UnsetTkn = nil
	n.OpenParenthesisTkn = nil
	n.SeparatorTkns = nil
	n.CloseParenthesisTkn = nil
	n.SemiColonTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtForeach(n *ast.StmtForeach) {
	n.ForeachTkn = nil
	n.OpenParenthesisTkn = nil
	n.AsTkn = nil
	n.DoubleArrowTkn = nil
	n.CloseParenthesisTkn = nil
	n.ColonTkn = nil
	n.EndForeachTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtDeclare(n *ast.StmtDeclare) {
	n.DeclareTkn = nil
	n.OpenParenthesisTkn = nil
	n.SeparatorTkns = nil
	n.CloseParenthesisTkn = nil
	n.ColonTkn = nil
	n.EndDeclareTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtNop(n *ast.StmtNop) {
	n.SemiColonTkn = nil
}
