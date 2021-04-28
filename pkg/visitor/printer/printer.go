package printer

import (
	"bytes"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/token"
	"io"
)

type printerState int

const (
	PrinterStateHTML printerState = iota
	PrinterStatePHP
)

type printer struct {
	output io.Writer
	state  printerState
	last   []byte
	// it will does not print the FreeFloating tokens when it's true
	withoutFreeFloating bool
}

func NewPrinter(output io.Writer) *printer {
	return &printer{
		output: output,
	}
}

func (p *printer) WithState(state printerState) *printer {
	p.state = state
	return p
}

func (p *printer) WithoutFreeFloating() *printer {
	p.withoutFreeFloating = true
	return p
}

func isValidVarName(r byte) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' || r >= 0x80
}

func (p *printer) write(b []byte) {
	if len(b) == 0 {
		return
	}

	if p.state == PrinterStateHTML {
		if !bytes.HasPrefix(b, []byte("<?")) {
			p.output.Write([]byte("<?php "))
		}
		p.state = PrinterStatePHP
	}

	if p.last != nil && isValidVarName(p.last[len(p.last)-1]) && isValidVarName(b[0]) {
		p.output.Write([]byte(" "))
	}

	p.last = b
	p.output.Write(b)
}

func (p *printer) printNode(n ast.Vertex) {
	if n != nil {
		n.Accept(p)
	}
}

func (p *printer) printList(list []ast.Vertex) {
	for _, nn := range list {
		p.printNode(nn)
	}
}

func (p *printer) printSeparatedList(list []ast.Vertex, separators []*token.Token, defaultSeparator []byte) {
	for k, nn := range list {
		p.printNode(nn)
		if k < len(separators) {
			p.printToken(separators[k], defaultSeparator)
		} else if k < len(list)-1 {
			p.write(defaultSeparator)
		}
	}
}

func (p *printer) printToken(t *token.Token, def []byte) {
	if t == nil && def == nil {
		return
	}

	if t == nil {
		p.write(def)
		return
	}

	if !p.withoutFreeFloating {
		for _, ff := range t.FreeFloating {
			p.write(ff.Value)
		}
	}
	p.write(t.Value)
}

func (p *printer) ifNode(n ast.Vertex, val []byte) []byte {
	if n == nil {
		return nil
	}

	return val
}

func (p *printer) ifNodeList(n []ast.Vertex, val []byte) []byte {
	if n == nil {
		return nil
	}

	return val
}

func (p *printer) ifNotNodeList(n []ast.Vertex, val []byte) []byte {
	if n != nil {
		return nil
	}

	return val
}

func (p *printer) ifToken(t *token.Token, true []byte, false []byte) []byte {
	if t == nil {
		return false
	}

	return true
}

func (p *printer) ifNotToken(t *token.Token, val []byte) []byte {
	if t != nil {
		return nil
	}

	return val
}

func (p *printer) Root(n *ast.Root) {
	p.printList(n.Stmts)
	p.printToken(n.EndTkn, nil)
}

func (p *printer) Nullable(n *ast.Nullable) {
	p.printToken(n.QuestionTkn, []byte("?"))
	p.printNode(n.Expr)
}

func (p *printer) Parameter(n *ast.Parameter) {
	p.printNode(n.Type)
	p.printToken(n.AmpersandTkn, nil)
	p.printToken(n.VariadicTkn, nil)
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, p.ifNode(n.DefaultValue, []byte("=")))
	p.printNode(n.DefaultValue)
}

func (p *printer) Identifier(n *ast.Identifier) {
	p.printToken(n.IdentifierTkn, n.Value)
}

func (p *printer) Argument(n *ast.Argument) {
	p.printToken(n.VariadicTkn, nil)
	p.printToken(n.AmpersandTkn, nil)
	p.printNode(n.Expr)
}

func (p *printer) StmtBreak(n *ast.StmtBreak) {
	p.printToken(n.BreakTkn, []byte("break"))
	p.printNode(n.Expr)
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtCase(n *ast.StmtCase) {
	p.printToken(n.CaseTkn, []byte("case"))
	p.printNode(n.Cond)
	p.printToken(n.CaseSeparatorTkn, []byte(":"))
	p.printList(n.Stmts)
}

func (p *printer) StmtCatch(n *ast.StmtCatch) {
	p.printToken(n.CatchTkn, []byte("catch"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Types, n.SeparatorTkns, []byte("|"))
	p.printNode(n.Var)
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) StmtClass(n *ast.StmtClass) {
	p.printList(n.Modifiers)
	p.printToken(n.ClassTkn, []byte("class"))
	p.printNode(n.Name)
	p.printToken(n.OpenParenthesisTkn, p.ifNodeList(n.Args, []byte("(")))
	p.printSeparatedList(n.Args, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, p.ifNodeList(n.Args, []byte(")")))
	p.printToken(n.ExtendsTkn, p.ifNode(n.Extends, []byte("extends")))
	p.printNode(n.Extends)
	p.printToken(n.ImplementsTkn, p.ifNodeList(n.Implements, []byte("implements")))
	p.printSeparatedList(n.Implements, n.ImplementsSeparatorTkns, []byte(","))
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) StmtClassConstList(n *ast.StmtClassConstList) {
	p.printList(n.Modifiers)
	p.printToken(n.ConstTkn, []byte("const"))
	p.printSeparatedList(n.Consts, n.SeparatorTkns, []byte(","))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtClassMethod(n *ast.StmtClassMethod) {
	p.printList(n.Modifiers)
	p.printToken(n.FunctionTkn, []byte("function"))
	p.printToken(n.AmpersandTkn, nil)
	p.printNode(n.Name)
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Params, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, p.ifNode(n.ReturnType, []byte(":")))
	p.printNode(n.ReturnType)
	p.printNode(n.Stmt)
}

func (p *printer) StmtConstList(n *ast.StmtConstList) {
	p.printToken(n.ConstTkn, []byte("const"))
	p.printSeparatedList(n.Consts, n.SeparatorTkns, []byte(","))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtConstant(n *ast.StmtConstant) {
	p.printNode(n.Name)
	p.printToken(n.EqualTkn, []byte("="))
	p.printNode(n.Expr)
}

func (p *printer) StmtContinue(n *ast.StmtContinue) {
	p.printToken(n.ContinueTkn, []byte("continue"))
	p.printNode(n.Expr)
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtDeclare(n *ast.StmtDeclare) {
	p.printToken(n.DeclareTkn, []byte("declare"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Consts, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, nil)
	if stmt, ok := n.Stmt.(*ast.StmtStmtList); ok && n.ColonTkn != nil {
		p.printToken(stmt.OpenCurlyBracketTkn, nil)
		p.printList(stmt.Stmts)
		p.printToken(stmt.CloseCurlyBracketTkn, nil)
	} else {
		p.printNode(n.Stmt)
	}
	p.printToken(n.EndDeclareTkn, p.ifToken(n.ColonTkn, []byte("enddeclare"), nil))
	p.printToken(n.SemiColonTkn, p.ifToken(n.ColonTkn, []byte(";"), nil))
}

func (p *printer) StmtDefault(n *ast.StmtDefault) {
	p.printToken(n.DefaultTkn, []byte("default"))
	p.printToken(n.CaseSeparatorTkn, []byte(":"))
	p.printList(n.Stmts)
}

func (p *printer) StmtDo(n *ast.StmtDo) {
	p.printToken(n.DoTkn, []byte("do"))
	p.printNode(n.Stmt)
	p.printToken(n.WhileTkn, []byte("while"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printNode(n.Cond)
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.SemiColonTkn, []byte(";"))

}

func (p *printer) StmtEcho(n *ast.StmtEcho) {
	p.printToken(n.EchoTkn, []byte("echo"))
	p.printSeparatedList(n.Exprs, n.SeparatorTkns, []byte(","))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtElse(n *ast.StmtElse) {
	p.printToken(n.ElseTkn, []byte("else"))
	p.printToken(n.ColonTkn, nil)
	if stmt, ok := n.Stmt.(*ast.StmtStmtList); ok && n.ColonTkn != nil {
		p.printToken(stmt.OpenCurlyBracketTkn, nil)
		p.printList(stmt.Stmts)
		p.printToken(stmt.CloseCurlyBracketTkn, nil)
	} else {
		p.printNode(n.Stmt)
	}
}

func (p *printer) StmtElseIf(n *ast.StmtElseIf) {
	p.printToken(n.ElseIfTkn, []byte("elseif"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printNode(n.Cond)
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, nil)
	if stmt, ok := n.Stmt.(*ast.StmtStmtList); ok && n.ColonTkn != nil {
		p.printToken(stmt.OpenCurlyBracketTkn, nil)
		p.printList(stmt.Stmts)
		p.printToken(stmt.CloseCurlyBracketTkn, nil)
	} else {
		p.printNode(n.Stmt)
	}
}

func (p *printer) StmtExpression(n *ast.StmtExpression) {
	p.printNode(n.Expr)
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtFinally(n *ast.StmtFinally) {
	p.printToken(n.FinallyTkn, []byte("finally"))
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) StmtFor(n *ast.StmtFor) {
	p.printToken(n.ForTkn, []byte("for"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Init, n.InitSeparatorTkns, []byte(","))
	p.printToken(n.InitSemiColonTkn, []byte(";"))
	p.printSeparatedList(n.Cond, n.CondSeparatorTkns, []byte(","))
	p.printToken(n.CondSemiColonTkn, []byte(";"))
	p.printSeparatedList(n.Loop, n.LoopSeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, nil)
	if stmt, ok := n.Stmt.(*ast.StmtStmtList); ok && n.ColonTkn != nil {
		p.printToken(stmt.OpenCurlyBracketTkn, nil)
		p.printList(stmt.Stmts)
		p.printToken(stmt.CloseCurlyBracketTkn, nil)
	} else {
		p.printNode(n.Stmt)
	}
	p.printToken(n.EndForTkn, p.ifToken(n.ColonTkn, []byte("endfor"), nil))
	p.printToken(n.SemiColonTkn, p.ifToken(n.ColonTkn, []byte(";"), nil))
}

func (p *printer) StmtForeach(n *ast.StmtForeach) {
	p.printToken(n.ForeachTkn, []byte("foreach"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printNode(n.Expr)
	p.printToken(n.AsTkn, []byte("as"))
	p.printNode(n.Key)
	p.printToken(n.DoubleArrowTkn, p.ifNode(n.Key, []byte("=>")))
	p.printToken(n.AmpersandTkn, nil)
	p.printNode(n.Var)
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, nil)
	if stmt, ok := n.Stmt.(*ast.StmtStmtList); ok && n.ColonTkn != nil {
		p.printToken(stmt.OpenCurlyBracketTkn, nil)
		p.printList(stmt.Stmts)
		p.printToken(stmt.CloseCurlyBracketTkn, nil)
	} else {
		p.printNode(n.Stmt)
	}
	p.printToken(n.EndForeachTkn, p.ifToken(n.ColonTkn, []byte("endforeach"), nil))
	p.printToken(n.SemiColonTkn, p.ifToken(n.ColonTkn, []byte(";"), nil))
}

func (p *printer) StmtFunction(n *ast.StmtFunction) {
	p.printToken(n.FunctionTkn, []byte("function"))
	p.printToken(n.AmpersandTkn, nil)
	p.printNode(n.Name)
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Params, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, p.ifNode(n.ReturnType, []byte(":")))
	p.printNode(n.ReturnType)
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) StmtGlobal(n *ast.StmtGlobal) {
	p.printToken(n.GlobalTkn, []byte("global"))
	p.printSeparatedList(n.Vars, n.SeparatorTkns, []byte(","))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtGoto(n *ast.StmtGoto) {
	p.printToken(n.GotoTkn, []byte("goto"))
	p.printNode(n.Label)
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtHaltCompiler(n *ast.StmtHaltCompiler) {
	p.printToken(n.HaltCompilerTkn, []byte("__halt_compiler"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtIf(n *ast.StmtIf) {
	p.printToken(n.IfTkn, []byte("if"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printNode(n.Cond)
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, nil)
	if stmt, ok := n.Stmt.(*ast.StmtStmtList); ok && n.ColonTkn != nil {
		p.printToken(stmt.OpenCurlyBracketTkn, nil)
		p.printList(stmt.Stmts)
		p.printToken(stmt.CloseCurlyBracketTkn, nil)
	} else {
		p.printNode(n.Stmt)
	}
	p.printList(n.ElseIf)
	p.printNode(n.Else)
	p.printToken(n.EndIfTkn, p.ifToken(n.ColonTkn, []byte("endif"), nil))
	p.printToken(n.SemiColonTkn, p.ifToken(n.ColonTkn, []byte(";"), nil))
}

func (p *printer) StmtInlineHtml(n *ast.StmtInlineHtml) {
	p.state = PrinterStatePHP
	if p.last != nil && !bytes.HasSuffix(p.last, []byte("?>")) && !bytes.HasSuffix(p.last, []byte("?>\n")) {
		p.write([]byte("?>"))
	}

	p.printToken(n.InlineHtmlTkn, n.Value)
	p.state = PrinterStateHTML
}

func (p *printer) StmtInterface(n *ast.StmtInterface) {
	p.printToken(n.InterfaceTkn, []byte("interface"))
	p.printNode(n.Name)
	p.printToken(n.ExtendsTkn, p.ifNodeList(n.Extends, []byte("extends")))
	p.printSeparatedList(n.Extends, n.ExtendsSeparatorTkns, []byte(","))
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) StmtLabel(n *ast.StmtLabel) {
	p.printNode(n.Name)
	p.printToken(n.ColonTkn, []byte(":"))
}

func (p *printer) StmtNamespace(n *ast.StmtNamespace) {
	p.printToken(n.NsTkn, []byte("namespace"))
	p.printNode(n.Name)
	p.printToken(n.OpenCurlyBracketTkn, p.ifNodeList(n.Stmts, []byte("{")))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, p.ifNodeList(n.Stmts, []byte("}")))
	p.printToken(n.SemiColonTkn, p.ifNotNodeList(n.Stmts, []byte(";")))
}

func (p *printer) StmtNop(n *ast.StmtNop) {
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtProperty(n *ast.StmtProperty) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, p.ifNode(n.Expr, []byte("=")))
	p.printNode(n.Expr)
}

func (p *printer) StmtPropertyList(n *ast.StmtPropertyList) {
	p.printList(n.Modifiers)
	p.printNode(n.Type)
	p.printSeparatedList(n.Props, n.SeparatorTkns, []byte(","))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtReturn(n *ast.StmtReturn) {
	p.printToken(n.ReturnTkn, []byte("return"))
	p.printNode(n.Expr)
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtStatic(n *ast.StmtStatic) {
	p.printToken(n.StaticTkn, []byte("static"))
	p.printSeparatedList(n.Vars, n.SeparatorTkns, []byte(","))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtStaticVar(n *ast.StmtStaticVar) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, p.ifNode(n.Expr, []byte("=")))
	p.printNode(n.Expr)
}

func (p *printer) StmtStmtList(n *ast.StmtStmtList) {
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) StmtSwitch(n *ast.StmtSwitch) {
	p.printToken(n.SwitchTkn, []byte("switch"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printNode(n.Cond)
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, nil)
	p.printToken(n.OpenCurlyBracketTkn, p.ifNotToken(n.ColonTkn, []byte("{")))
	p.printToken(n.CaseSeparatorTkn, nil)
	p.printList(n.Cases)
	p.printToken(n.CloseCurlyBracketTkn, p.ifNotToken(n.ColonTkn, []byte("}")))
	p.printToken(n.EndSwitchTkn, p.ifToken(n.ColonTkn, []byte("endswitch"), nil))
	p.printToken(n.SemiColonTkn, p.ifToken(n.ColonTkn, []byte(";"), nil))
}

func (p *printer) StmtThrow(n *ast.StmtThrow) {
	p.printToken(n.ThrowTkn, []byte("throw"))
	p.printNode(n.Expr)
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtTrait(n *ast.StmtTrait) {
	p.printToken(n.TraitTkn, []byte("trait"))
	p.printNode(n.Name)
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) StmtTraitUse(n *ast.StmtTraitUse) {
	p.printToken(n.UseTkn, []byte("use"))
	p.printSeparatedList(n.Traits, n.SeparatorTkns, []byte(","))
	p.printToken(n.OpenCurlyBracketTkn, p.ifNodeList(n.Adaptations, []byte("{")))
	p.printList(n.Adaptations)
	p.printToken(n.CloseCurlyBracketTkn, p.ifNodeList(n.Adaptations, []byte("}")))
	p.printToken(n.SemiColonTkn, p.ifNotToken(n.OpenCurlyBracketTkn, p.ifNotNodeList(n.Adaptations, []byte(";"))))
}

func (p *printer) StmtTraitUseAlias(n *ast.StmtTraitUseAlias) {
	p.printNode(n.Trait)
	p.printToken(n.DoubleColonTkn, p.ifNode(n.Trait, []byte("::")))
	p.printNode(n.Method)
	p.printToken(n.AsTkn, []byte("as"))
	p.printNode(n.Modifier)
	p.printNode(n.Alias)
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtTraitUsePrecedence(n *ast.StmtTraitUsePrecedence) {
	p.printNode(n.Trait)
	p.printToken(n.DoubleColonTkn, p.ifNode(n.Trait, []byte("::")))
	p.printNode(n.Method)
	p.printToken(n.InsteadofTkn, []byte("insteadof"))
	p.printSeparatedList(n.Insteadof, n.SeparatorTkns, []byte(","))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtTry(n *ast.StmtTry) {
	p.printToken(n.TryTkn, []byte("try"))
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
	p.printList(n.Catches)
	p.printNode(n.Finally)
}

func (p *printer) StmtUnset(n *ast.StmtUnset) {
	p.printToken(n.UnsetTkn, []byte("unset"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Vars, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtUse(n *ast.StmtUseList) {
	p.printToken(n.UseTkn, []byte("use"))
	p.printNode(n.Type)
	p.printSeparatedList(n.Uses, n.SeparatorTkns, []byte(","))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtGroupUse(n *ast.StmtGroupUseList) {
	p.printToken(n.UseTkn, []byte("use"))
	p.printNode(n.Type)
	p.printToken(n.LeadingNsSeparatorTkn, nil)
	p.printNode(n.Prefix)
	p.printToken(n.NsSeparatorTkn, []byte("\\"))
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printSeparatedList(n.Uses, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
	p.printToken(n.SemiColonTkn, []byte(";"))
}

func (p *printer) StmtUseDeclaration(n *ast.StmtUse) {
	p.printNode(n.Type)
	p.printToken(n.NsSeparatorTkn, nil)
	p.printNode(n.Use)
	p.printToken(n.AsTkn, p.ifNode(n.Alias, []byte("as")))
	p.printNode(n.Alias)
}

func (p *printer) StmtWhile(n *ast.StmtWhile) {
	p.printToken(n.WhileTkn, []byte("while"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printNode(n.Cond)
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, nil)
	if stmt, ok := n.Stmt.(*ast.StmtStmtList); ok && n.ColonTkn != nil {
		p.printToken(stmt.OpenCurlyBracketTkn, nil)
		p.printList(stmt.Stmts)
		p.printToken(stmt.CloseCurlyBracketTkn, nil)
	} else {
		p.printNode(n.Stmt)
	}
	p.printToken(n.EndWhileTkn, p.ifToken(n.ColonTkn, []byte("endwhile"), nil))
	p.printToken(n.SemiColonTkn, p.ifToken(n.ColonTkn, []byte(";"), nil))
}

func (p *printer) ExprArray(n *ast.ExprArray) {
	p.printToken(n.ArrayTkn, nil)
	p.printToken(n.OpenBracketTkn, p.ifToken(n.ArrayTkn, []byte("("), []byte("[")))
	p.printSeparatedList(n.Items, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseBracketTkn, p.ifToken(n.ArrayTkn, []byte(")"), []byte("]")))
}

func (p *printer) ExprArrayDimFetch(n *ast.ExprArrayDimFetch) {
	p.printNode(n.Var)
	p.printToken(n.OpenBracketTkn, []byte("["))
	p.printNode(n.Dim)
	p.printToken(n.CloseBracketTkn, []byte("]"))
}

func (p *printer) ExprArrayItem(n *ast.ExprArrayItem) {
	p.printToken(n.EllipsisTkn, nil)
	p.printNode(n.Key)
	p.printToken(n.DoubleArrowTkn, p.ifNode(n.Key, []byte("=>")))
	p.printToken(n.AmpersandTkn, nil)
	p.printNode(n.Val)
}

func (p *printer) ExprArrowFunction(n *ast.ExprArrowFunction) {
	p.printToken(n.StaticTkn, nil)
	p.printToken(n.FnTkn, []byte("fn"))
	p.printToken(n.AmpersandTkn, nil)
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Params, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.ColonTkn, p.ifNode(n.ReturnType, []byte(":")))
	p.printNode(n.ReturnType)
	p.printToken(n.DoubleArrowTkn, []byte("=>"))
	p.printNode(n.Expr)
}

func (p *printer) ExprBitwiseNot(n *ast.ExprBitwiseNot) {
	p.printToken(n.TildaTkn, []byte("~"))
	p.printNode(n.Expr)
}

func (p *printer) ExprBooleanNot(n *ast.ExprBooleanNot) {
	p.printToken(n.ExclamationTkn, []byte("!"))
	p.printNode(n.Expr)
}

func (p *printer) ExprBrackets(n *ast.ExprBrackets) {
	p.printToken(n.OpenParenthesisTkn, nil)
	p.printNode(n.Expr)
	p.printToken(n.CloseParenthesisTkn, nil)
}

func (p *printer) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	p.printNode(n.Class)
	p.printToken(n.DoubleColonTkn, []byte("::"))
	p.printNode(n.Const)
}

func (p *printer) ExprClone(n *ast.ExprClone) {
	p.printToken(n.CloneTkn, []byte("clone"))
	p.printNode(n.Expr)
}

func (p *printer) ExprClosure(n *ast.ExprClosure) {
	p.printToken(n.StaticTkn, nil)
	p.printToken(n.FunctionTkn, []byte("function"))
	p.printToken(n.AmpersandTkn, nil)
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Params, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
	p.printToken(n.UseTkn, p.ifNodeList(n.Uses, []byte("use")))
	p.printToken(n.UseOpenParenthesisTkn, p.ifNodeList(n.Uses, []byte("(")))
	p.printSeparatedList(n.Uses, n.UseSeparatorTkns, []byte(","))
	p.printToken(n.UseCloseParenthesisTkn, p.ifNodeList(n.Uses, []byte(")")))
	p.printToken(n.ColonTkn, p.ifNode(n.ReturnType, []byte(":")))
	p.printNode(n.ReturnType)
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printList(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) ExprClosureUse(n *ast.ExprClosureUse) {
	p.printToken(n.AmpersandTkn, nil)
	p.printNode(n.Var)
}

func (p *printer) ExprConstFetch(n *ast.ExprConstFetch) {
	p.printNode(n.Const)
}

func (p *printer) ExprEmpty(n *ast.ExprEmpty) {
	p.printToken(n.EmptyTkn, []byte("empty"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printNode(n.Expr)
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
}

func (p *printer) ExprErrorSuppress(n *ast.ExprErrorSuppress) {
	p.printToken(n.AtTkn, []byte("@"))
	p.printNode(n.Expr)
}

func (p *printer) ExprEval(n *ast.ExprEval) {
	p.printToken(n.EvalTkn, []byte("eval"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printNode(n.Expr)
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
}

func (p *printer) ExprExit(n *ast.ExprExit) {
	p.printToken(n.ExitTkn, []byte("exit"))
	p.printToken(n.OpenParenthesisTkn, nil)
	p.printNode(n.Expr)
	p.printToken(n.CloseParenthesisTkn, p.ifToken(n.OpenParenthesisTkn, []byte(")"), nil))
}

func (p *printer) ExprFunctionCall(n *ast.ExprFunctionCall) {
	p.printNode(n.Function)
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Args, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
}

func (p *printer) ExprInclude(n *ast.ExprInclude) {
	p.printToken(n.IncludeTkn, []byte("include"))
	p.printNode(n.Expr)
}

func (p *printer) ExprIncludeOnce(n *ast.ExprIncludeOnce) {
	p.printToken(n.IncludeOnceTkn, []byte("include_once"))
	p.printNode(n.Expr)
}

func (p *printer) ExprInstanceOf(n *ast.ExprInstanceOf) {
	p.printNode(n.Expr)
	p.printToken(n.InstanceOfTkn, []byte("instanceof"))
	p.printNode(n.Class)
}

func (p *printer) ExprIsset(n *ast.ExprIsset) {
	p.printToken(n.IssetTkn, []byte("isset"))
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Vars, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
}

func (p *printer) ExprList(n *ast.ExprList) {
	p.printToken(n.ListTkn, p.ifToken(n.OpenBracketTkn, nil, []byte("list")))
	p.printToken(n.OpenBracketTkn, []byte("("))
	p.printSeparatedList(n.Items, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseBracketTkn, []byte(")"))
}

func (p *printer) ExprMethodCall(n *ast.ExprMethodCall) {
	p.printNode(n.Var)
	p.printToken(n.ObjectOperatorTkn, []byte("->"))
	p.printToken(n.OpenCurlyBracketTkn, nil)
	p.printNode(n.Method)
	p.printToken(n.CloseCurlyBracketTkn, nil)
	p.printToken(n.OpenParenthesisTkn, []byte("("))
	p.printSeparatedList(n.Args, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, []byte(")"))
}

func (p *printer) ExprNew(n *ast.ExprNew) {
	p.printToken(n.NewTkn, []byte("new"))
	p.printNode(n.Class)
	p.printToken(n.OpenParenthesisTkn, p.ifNodeList(n.Args, []byte("(")))
	p.printSeparatedList(n.Args, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, p.ifNodeList(n.Args, []byte(")")))
}

func (p *printer) ExprPostDec(n *ast.ExprPostDec) {
	p.printNode(n.Var)
	p.printToken(n.DecTkn, []byte("--"))
}

func (p *printer) ExprPostInc(n *ast.ExprPostInc) {
	p.printNode(n.Var)
	p.printToken(n.IncTkn, []byte("++"))
}

func (p *printer) ExprPreDec(n *ast.ExprPreDec) {
	p.printToken(n.DecTkn, []byte("--"))
	p.printNode(n.Var)
}

func (p *printer) ExprPreInc(n *ast.ExprPreInc) {
	p.printToken(n.IncTkn, []byte("++"))
	p.printNode(n.Var)
}

func (p *printer) ExprPrint(n *ast.ExprPrint) {
	p.printToken(n.PrintTkn, []byte("print"))
	p.printNode(n.Expr)
}

func (p *printer) ExprPropertyFetch(n *ast.ExprPropertyFetch) {
	p.printNode(n.Var)
	p.printToken(n.ObjectOperatorTkn, []byte("->"))
	p.printToken(n.OpenCurlyBracketTkn, nil)
	p.printNode(n.Prop)
	p.printToken(n.CloseCurlyBracketTkn, nil)
}

func (p *printer) ExprRequire(n *ast.ExprRequire) {
	p.printToken(n.RequireTkn, []byte("require"))
	p.printNode(n.Expr)
}

func (p *printer) ExprRequireOnce(n *ast.ExprRequireOnce) {
	p.printToken(n.RequireOnceTkn, []byte("require_once"))
	p.printNode(n.Expr)
}

func (p *printer) ExprShellExec(n *ast.ExprShellExec) {
	p.printToken(n.OpenBacktickTkn, []byte("`"))
	p.printList(n.Parts)
	p.printToken(n.CloseBacktickTkn, []byte("`"))
}

func (p *printer) ExprStaticCall(n *ast.ExprStaticCall) {
	p.printNode(n.Class)
	p.printToken(n.DoubleColonTkn, []byte("::"))
	p.printToken(n.OpenCurlyBracketTkn, nil)
	p.printNode(n.Call)
	p.printToken(n.CloseCurlyBracketTkn, nil)
	p.printToken(n.OpenParenthesisTkn, p.ifNodeList(n.Args, []byte("(")))
	p.printSeparatedList(n.Args, n.SeparatorTkns, []byte(","))
	p.printToken(n.CloseParenthesisTkn, p.ifNodeList(n.Args, []byte(")")))
}

func (p *printer) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch) {
	p.printNode(n.Class)
	p.printToken(n.DoubleColonTkn, []byte("::"))
	p.printNode(n.Prop)
}

func (p *printer) ExprTernary(n *ast.ExprTernary) {
	p.printNode(n.Cond)
	p.printToken(n.QuestionTkn, []byte("?"))
	p.printNode(n.IfTrue)
	p.printToken(n.ColonTkn, []byte(":"))
	p.printNode(n.IfFalse)
}

func (p *printer) ExprUnaryMinus(n *ast.ExprUnaryMinus) {
	p.printToken(n.MinusTkn, []byte("-"))
	p.printNode(n.Expr)
}

func (p *printer) ExprUnaryPlus(n *ast.ExprUnaryPlus) {
	p.printToken(n.PlusTkn, []byte("+"))
	p.printNode(n.Expr)
}

func (p *printer) ExprVariable(n *ast.ExprVariable) {
	p.printToken(n.DollarTkn, nil)
	p.printToken(n.OpenCurlyBracketTkn, nil)
	p.printNode(n.Name)
	p.printToken(n.CloseCurlyBracketTkn, nil)
}

func (p *printer) ExprYield(n *ast.ExprYield) {
	p.printToken(n.YieldTkn, []byte("yield"))
	p.printNode(n.Key)
	p.printToken(n.DoubleArrowTkn, p.ifNode(n.Key, []byte("=>")))
	p.printNode(n.Val)
}

func (p *printer) ExprYieldFrom(n *ast.ExprYieldFrom) {
	p.printToken(n.YieldFromTkn, []byte("yield from"))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssign(n *ast.ExprAssign) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignReference(n *ast.ExprAssignReference) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("="))
	p.printToken(n.AmpersandTkn, []byte("&"))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignBitwiseAnd(n *ast.ExprAssignBitwiseAnd) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("&="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignBitwiseOr(n *ast.ExprAssignBitwiseOr) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("|="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignBitwiseXor(n *ast.ExprAssignBitwiseXor) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("^="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignCoalesce(n *ast.ExprAssignCoalesce) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("??="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignConcat(n *ast.ExprAssignConcat) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte(".="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignDiv(n *ast.ExprAssignDiv) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("/="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignMinus(n *ast.ExprAssignMinus) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("-="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignMod(n *ast.ExprAssignMod) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("%="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignMul(n *ast.ExprAssignMul) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("*="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignPlus(n *ast.ExprAssignPlus) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("+="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignPow(n *ast.ExprAssignPow) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("**="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignShiftLeft(n *ast.ExprAssignShiftLeft) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte("<<="))
	p.printNode(n.Expr)
}

func (p *printer) ExprAssignShiftRight(n *ast.ExprAssignShiftRight) {
	p.printNode(n.Var)
	p.printToken(n.EqualTkn, []byte(">>="))
	p.printNode(n.Expr)
}

func (p *printer) ExprBinaryBitwiseAnd(n *ast.ExprBinaryBitwiseAnd) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("&"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryBitwiseOr(n *ast.ExprBinaryBitwiseOr) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("|"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryBitwiseXor(n *ast.ExprBinaryBitwiseXor) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("^"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryBooleanAnd(n *ast.ExprBinaryBooleanAnd) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("&&"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryBooleanOr(n *ast.ExprBinaryBooleanOr) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("||"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryCoalesce(n *ast.ExprBinaryCoalesce) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("??"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryConcat(n *ast.ExprBinaryConcat) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("."))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryDiv(n *ast.ExprBinaryDiv) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("/"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryEqual(n *ast.ExprBinaryEqual) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("=="))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryGreater(n *ast.ExprBinaryGreater) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte(">"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryGreaterOrEqual(n *ast.ExprBinaryGreaterOrEqual) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte(">="))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryIdentical(n *ast.ExprBinaryIdentical) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("==="))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryLogicalAnd(n *ast.ExprBinaryLogicalAnd) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("and"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryLogicalOr(n *ast.ExprBinaryLogicalOr) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("or"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryLogicalXor(n *ast.ExprBinaryLogicalXor) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("xor"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryMinus(n *ast.ExprBinaryMinus) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("-"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryMod(n *ast.ExprBinaryMod) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("%"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryMul(n *ast.ExprBinaryMul) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("*"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryNotEqual(n *ast.ExprBinaryNotEqual) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("!="))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryNotIdentical(n *ast.ExprBinaryNotIdentical) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("!=="))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryPlus(n *ast.ExprBinaryPlus) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("+"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryPow(n *ast.ExprBinaryPow) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("**"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryShiftLeft(n *ast.ExprBinaryShiftLeft) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("<<"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinaryShiftRight(n *ast.ExprBinaryShiftRight) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte(">>"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinarySmaller(n *ast.ExprBinarySmaller) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("<"))
	p.printNode(n.Right)
}

func (p *printer) ExprBinarySmallerOrEqual(n *ast.ExprBinarySmallerOrEqual) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("<="))
	p.printNode(n.Right)
}

func (p *printer) ExprBinarySpaceship(n *ast.ExprBinarySpaceship) {
	p.printNode(n.Left)
	p.printToken(n.OpTkn, []byte("<=>"))
	p.printNode(n.Right)
}

func (p *printer) ExprCastArray(n *ast.ExprCastArray) {
	p.printToken(n.CastTkn, []byte("(array)"))
	p.printNode(n.Expr)
}

func (p *printer) ExprCastBool(n *ast.ExprCastBool) {
	p.printToken(n.CastTkn, []byte("(bool)"))
	p.printNode(n.Expr)
}

func (p *printer) ExprCastDouble(n *ast.ExprCastDouble) {
	p.printToken(n.CastTkn, []byte("(float)"))
	p.printNode(n.Expr)
}

func (p *printer) ExprCastInt(n *ast.ExprCastInt) {
	p.printToken(n.CastTkn, []byte("(int)"))
	p.printNode(n.Expr)
}

func (p *printer) ExprCastObject(n *ast.ExprCastObject) {
	p.printToken(n.CastTkn, []byte("(object)"))
	p.printNode(n.Expr)
}

func (p *printer) ExprCastString(n *ast.ExprCastString) {
	p.printToken(n.CastTkn, []byte("(string)"))
	p.printNode(n.Expr)
}

func (p *printer) ExprCastUnset(n *ast.ExprCastUnset) {
	p.printToken(n.CastTkn, []byte("(unset)"))
	p.printNode(n.Expr)
}

func (p *printer) ScalarDnumber(n *ast.ScalarDnumber) {
	p.printToken(n.NumberTkn, n.Value)
}

func (p *printer) ScalarEncapsed(n *ast.ScalarEncapsed) {
	p.printToken(n.OpenQuoteTkn, []byte("\""))
	p.printList(n.Parts)
	p.printToken(n.CloseQuoteTkn, []byte("\""))
}

func (p *printer) ScalarEncapsedStringPart(n *ast.ScalarEncapsedStringPart) {
	p.printToken(n.EncapsedStrTkn, n.Value)
}

func (p *printer) ScalarEncapsedStringVar(n *ast.ScalarEncapsedStringVar) {
	p.printToken(n.DollarOpenCurlyBracketTkn, []byte("${"))
	p.printNode(n.Name)
	p.printToken(n.OpenSquareBracketTkn, p.ifNode(n.Dim, []byte("[")))
	p.printNode(n.Dim)
	p.printToken(n.CloseSquareBracketTkn, p.ifNode(n.Dim, []byte("]")))
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) ScalarEncapsedStringBrackets(n *ast.ScalarEncapsedStringBrackets) {
	p.printToken(n.OpenCurlyBracketTkn, []byte("{"))
	p.printNode(n.Var)
	p.printToken(n.CloseCurlyBracketTkn, []byte("}"))
}

func (p *printer) ScalarHeredoc(n *ast.ScalarHeredoc) {
	p.printToken(n.OpenHeredocTkn, []byte("<<<EOT\n"))
	p.printList(n.Parts)
	p.printToken(n.CloseHeredocTkn, []byte("EOT"))
}

func (p *printer) ScalarLnumber(n *ast.ScalarLnumber) {
	p.printToken(n.NumberTkn, n.Value)
}

func (p *printer) ScalarMagicConstant(n *ast.ScalarMagicConstant) {
	p.printToken(n.MagicConstTkn, n.Value)
}

func (p *printer) ScalarString(n *ast.ScalarString) {
	p.printToken(n.MinusTkn, nil)
	p.printToken(n.StringTkn, n.Value)
}

func (p *printer) NameName(n *ast.Name) {
	p.printSeparatedList(n.Parts, n.SeparatorTkns, []byte("\\"))
}

func (p *printer) NameFullyQualified(n *ast.NameFullyQualified) {
	p.printToken(n.NsSeparatorTkn, []byte("\\"))
	p.printSeparatedList(n.Parts, n.SeparatorTkns, []byte("\\"))
}

func (p *printer) NameRelative(n *ast.NameRelative) {
	p.printToken(n.NsTkn, []byte("namespace"))
	p.printToken(n.NsSeparatorTkn, []byte("\\"))
	p.printSeparatedList(n.Parts, n.SeparatorTkns, []byte("\\"))
}

func (p *printer) NameNamePart(n *ast.NamePart) {
	p.printToken(n.StringTkn, n.Value)
}
