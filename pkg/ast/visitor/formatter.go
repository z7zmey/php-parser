package visitor

import (
	"bytes"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/token"
)

type formatterState int

const (
	FormatterStateHTML formatterState = iota
	FormatterStatePHP
)

type formatter struct {
	state        formatterState
	indent       int
	freeFloating []*token.Token

	lastSemiColon *token.Token
}

func NewFormatter() *formatter {
	return &formatter{}
}

func (f *formatter) WithState(state formatterState) *formatter {
	f.state = state
	return f
}

func (f *formatter) WithIndent(indent int) *formatter {
	f.indent = indent
	return f
}

func (f *formatter) addFreeFloating(id token.ID, val []byte) {
	f.freeFloating = append(f.freeFloating, &token.Token{
		ID:    id,
		Value: val,
	})
}

func (f *formatter) addIndent() {
	if f.indent < 1 {
		return
	}

	f.freeFloating = append(f.freeFloating, &token.Token{
		ID:    token.T_WHITESPACE,
		Value: bytes.Repeat([]byte("    "), f.indent),
	})
}

func (f *formatter) resetFreeFloating() {
	f.freeFloating = nil
}

func (f *formatter) getFreeFloating() []*token.Token {
	defer f.resetFreeFloating()

	if f.state == FormatterStateHTML {
		t := &token.Token{
			ID:    token.T_OPEN_TAG,
			Value: []byte("<?php "),
		}
		f.freeFloating = append([]*token.Token{t}, f.freeFloating...)

		f.state = FormatterStatePHP
	}

	return f.freeFloating
}

func (f *formatter) newToken(id token.ID, val []byte) *token.Token {
	return &token.Token{
		ID:           id,
		Value:        val,
		FreeFloating: f.getFreeFloating(),
	}
}

func (f *formatter) formatList(nodes []ast.Vertex, separator byte) []*token.Token {
	separatorTkns := make([]*token.Token, len(nodes)-1)
	for i, v := range nodes {
		v.Accept(f)

		if i != len(nodes)-1 {
			separatorTkns[i] = f.newToken(token.ID(separator), []byte{separator})
			f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		}
	}

	return separatorTkns
}

func (f *formatter) formatStmts(list *[]ast.Vertex) {
	var insertCounter int

	for i, stmt := range *list {
		f.lastSemiColon = nil

		if _, ok := stmt.(*ast.StmtInlineHtml); ok {
			if f.lastSemiColon != nil {
				f.lastSemiColon.Value = append(f.lastSemiColon.Value, '?', '>')
			} else {
				*list = insert(*list, i+insertCounter, &ast.StmtNop{
					SemiColonTkn: &token.Token{
						Value: []byte("?>"),
					},
				})
				insertCounter++
			}
		} else {
			f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
			f.addIndent()
		}

		stmt.Accept(f)
	}
}

func (f *formatter) newSemicolonTkn() *token.Token {
	f.lastSemiColon = f.newToken(';', []byte(";"))
	return f.lastSemiColon
}

func insert(s []ast.Vertex, k int, vs ...ast.Vertex) []ast.Vertex {
	if n := len(s) + len(vs); n <= cap(s) {
		s2 := s[:n]
		copy(s2[k+len(vs):], s[k:])
		copy(s2[k:], vs)
		return s2
	}
	s2 := make([]ast.Vertex, len(s)+len(vs))
	copy(s2, s[:k])
	copy(s2[k:], vs)
	copy(s2[k+len(vs):], s[k:])
	return s2
}

func (f *formatter) Root(n *ast.Root) {
	f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
	f.addIndent()

	f.formatStmts(&n.Stmts)
}

func (f *formatter) Nullable(n *ast.Nullable) {
	n.QuestionTkn = f.newToken('?', []byte("?"))
	n.Expr.Accept(f)
}

func (f *formatter) Parameter(n *ast.Parameter) {
	if n.Type != nil {
		n.Type.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	if n.AmpersandTkn != nil {
		n.AmpersandTkn = f.newToken('&', []byte("&"))
	}

	if n.VariadicTkn != nil {
		n.VariadicTkn = f.newToken(token.T_ELLIPSIS, []byte("..."))
	}

	n.Var.Accept(f)

	if n.DefaultValue != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.EqualTkn = f.newToken('=', []byte("="))
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.DefaultValue.Accept(f)
	}
}

func (f *formatter) Identifier(n *ast.Identifier) {
	if n.IdentifierTkn == nil {
		n.IdentifierTkn = f.newToken(token.T_STRING, n.Value)
	} else {
		n.IdentifierTkn.FreeFloating = f.getFreeFloating()
	}
}

func (f *formatter) Argument(n *ast.Argument) {
	if n.VariadicTkn != nil {
		n.VariadicTkn = f.newToken(token.T_ELLIPSIS, []byte("..."))
	}

	if n.AmpersandTkn != nil {
		n.AmpersandTkn = f.newToken('&', []byte("&"))
	}

	n.Expr.Accept(f)
}

func (f *formatter) StmtBreak(n *ast.StmtBreak) {
	n.BreakTkn = f.newToken(token.T_BREAK, []byte("break"))

	if n.Expr != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Expr.Accept(f)
	}

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtCase(n *ast.StmtCase) {
	n.CaseTkn = f.newToken(token.T_CASE, []byte("case"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Cond.Accept(f)

	n.CaseSeparatorTkn = f.newToken(':', []byte(":"))

	f.indent++
	f.formatStmts(&n.Stmts)
	f.indent--
}

func (f *formatter) StmtCatch(n *ast.StmtCatch) {
	n.CatchTkn = f.newToken(token.T_CATCH, []byte("catch"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpenParenthesisTkn = f.newToken('(', []byte("("))

	n.SeparatorTkns = make([]*token.Token, len(n.Types)-1)
	for i, t := range n.Types {
		t.Accept(f)

		if i != len(n.Types)-1 {
			f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
			n.SeparatorTkns[i] = f.newToken('|', []byte("|"))
			f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		}
	}

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Var.Accept(f)

	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

	if len(n.Stmts) > 0 {
		f.indent++
		f.formatStmts(&n.Stmts)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}

	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
}

func (f *formatter) StmtClass(n *ast.StmtClass) {
	for _, m := range n.Modifiers {
		m.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.ClassTkn = f.newToken(token.T_CLASS, []byte("class"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.ClassName.Accept(f)

	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	if len(n.Arguments) > 0 {
		n.OpenParenthesisTkn = f.newToken('(', []byte("("))

		n.SeparatorTkns = f.formatList(n.Arguments, ',')

		n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
	}

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	if n.Extends != nil {
		n.Extends.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	if n.Implements != nil {
		n.Implements.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

	if len(n.Stmts) > 0 {
		f.indent++
		f.formatStmts(&n.Stmts)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}

	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
}

func (f *formatter) StmtClassConstList(n *ast.StmtClassConstList) {
	for _, m := range n.Modifiers {
		m.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.ConstTkn = f.newToken(token.T_CONST, []byte("const"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.SeparatorTkns = f.formatList(n.Consts, ',')

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtClassExtends(n *ast.StmtClassExtends) {
	n.ExtendTkn = f.newToken(token.T_EXTENDS, []byte("extends"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.ClassName.Accept(f)
}

func (f *formatter) StmtClassImplements(n *ast.StmtClassImplements) {
	n.ImplementsTkn = f.newToken(token.T_IMPLEMENTS, []byte("implements"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.SeparatorTkns = f.formatList(n.InterfaceNames, ',')
}

func (f *formatter) StmtClassMethod(n *ast.StmtClassMethod) {
	for _, m := range n.Modifiers {
		m.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.FunctionTkn = f.newToken(token.T_FUNCTION, []byte("function"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if n.AmpersandTkn != nil {
		n.AmpersandTkn = f.newToken('&', []byte("&"))
	}

	n.MethodName.Accept(f)

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))

	n.SeparatorTkns = nil
	if len(n.Params) > 0 {
		n.SeparatorTkns = f.formatList(n.Params, ',')
	}

	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	if n.ReturnType != nil {
		n.ColonTkn = f.newToken(':', []byte(":"))

		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.ReturnType.Accept(f)
	}

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Stmt.Accept(f)
}

func (f *formatter) StmtConstList(n *ast.StmtConstList) {
	n.ConstTkn = f.newToken(token.T_CONST, []byte("const"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.SeparatorTkns = f.formatList(n.Consts, ',')

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtConstant(n *ast.StmtConstant) {
	n.Name.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken('=', []byte("="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) StmtContinue(n *ast.StmtContinue) {
	n.ContinueTkn = f.newToken(token.T_CONTINUE, []byte("continue"))

	if n.Expr != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Expr.Accept(f)
	}

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtDeclare(n *ast.StmtDeclare) {
	n.ColonTkn = nil
	n.EndDeclareTkn = nil
	n.SemiColonTkn = nil

	n.DeclareTkn = f.newToken(token.T_DECLARE, []byte("declare"))
	n.OpenParenthesisTkn = f.newToken('(', []byte("("))

	n.SeparatorTkns = nil
	if len(n.Consts) > 0 {
		n.SeparatorTkns = f.formatList(n.Consts, ',')
	}

	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Stmt.Accept(f)
}

func (f *formatter) StmtDefault(n *ast.StmtDefault) {
	n.DefaultTkn = f.newToken(token.T_DEFAULT, []byte("default"))

	n.CaseSeparatorTkn = f.newToken(':', []byte(":"))

	f.indent++
	f.formatStmts(&n.Stmts)
	f.indent--
}

func (f *formatter) StmtDo(n *ast.StmtDo) {
	n.DoTkn = f.newToken(token.T_DO, []byte("do"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Stmt.Accept(f)
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.WhileTkn = f.newToken(token.T_WHILE, []byte("while"))

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.Cond.Accept(f)
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtEcho(n *ast.StmtEcho) {
	n.EchoTkn = f.newToken(token.T_ECHO, []byte("echo"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.SeparatorTkns = nil
	if len(n.Exprs) > 0 {
		n.SeparatorTkns = f.formatList(n.Exprs, ',')
	}

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtElse(n *ast.StmtElse) {
	n.ColonTkn = nil

	n.ElseTkn = f.newToken(token.T_ELSE, []byte("else"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Stmt.Accept(f)
}

func (f *formatter) StmtElseIf(n *ast.StmtElseIf) {
	n.ColonTkn = nil

	n.ElseIfTkn = f.newToken(token.T_ELSEIF, []byte("elseif"))

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.Cond.Accept(f)
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Stmt.Accept(f)
}

func (f *formatter) StmtExpression(n *ast.StmtExpression) {
	n.Expr.Accept(f)
	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtFinally(n *ast.StmtFinally) {
	n.FinallyTkn = f.newToken(token.T_FINALLY, []byte("finally"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

	if len(n.Stmts) > 0 {
		f.indent++
		f.formatStmts(&n.Stmts)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}

	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
}

func (f *formatter) StmtFor(n *ast.StmtFor) {
	n.ColonTkn = nil
	n.EndForTkn = nil
	n.SemiColonTkn = nil

	n.ForTkn = f.newToken(token.T_FOR, []byte("for"))
	n.OpenParenthesisTkn = f.newToken('(', []byte("("))

	n.InitSeparatorTkns = nil
	if len(n.Init) > 0 {
		n.InitSeparatorTkns = f.formatList(n.Init, ',')
	}

	n.InitSemiColonTkn = f.newSemicolonTkn()
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.CondSeparatorTkns = nil
	if len(n.Cond) > 0 {
		n.CondSeparatorTkns = f.formatList(n.Cond, ',')
	}

	n.CondSemiColonTkn = f.newSemicolonTkn()
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.LoopSeparatorTkns = nil
	if len(n.Loop) > 0 {
		n.LoopSeparatorTkns = f.formatList(n.Loop, ',')
	}

	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Stmt.Accept(f)
}

func (f *formatter) StmtForeach(n *ast.StmtForeach) {
	n.ColonTkn = nil
	n.EndForeachTkn = nil
	n.SemiColonTkn = nil

	n.ForeachTkn = f.newToken(token.T_FOREACH, []byte("foreach"))

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))

	n.Expr.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.AsTkn = f.newToken(token.T_AS, []byte("as"))

	if n.Key != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Key.Accept(f)

		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.DoubleArrowTkn = f.newToken(token.T_DOUBLE_ARROW, []byte("=>"))
	}

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	if n.AmpersandTkn != nil {
		n.AmpersandTkn = f.newToken('&', []byte("&"))
	}
	n.Var.Accept(f)

	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Stmt.Accept(f)
}

func (f *formatter) StmtFunction(n *ast.StmtFunction) {
	n.FunctionTkn = f.newToken(token.T_FUNCTION, []byte("function"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if n.AmpersandTkn != nil {
		n.AmpersandTkn = f.newToken('&', []byte("&"))
	}

	n.FunctionName.Accept(f)

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))

	n.SeparatorTkns = nil
	if len(n.Params) > 0 {
		n.SeparatorTkns = f.formatList(n.Params, ',')
	}

	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	if n.ReturnType != nil {
		n.ColonTkn = f.newToken(':', []byte(":"))

		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.ReturnType.Accept(f)
	}

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

	if len(n.Stmts) > 0 {
		f.indent++
		f.formatStmts(&n.Stmts)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}

	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
}

func (f *formatter) StmtGlobal(n *ast.StmtGlobal) {
	n.GlobalTkn = f.newToken(token.T_GLOBAL, []byte("global"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.SeparatorTkns = nil
	if len(n.Vars) > 0 {
		n.SeparatorTkns = f.formatList(n.Vars, ',')
	}

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtGoto(n *ast.StmtGoto) {
	n.GotoTkn = f.newToken(token.T_GOTO, []byte("goto"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Label.Accept(f)

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtHaltCompiler(n *ast.StmtHaltCompiler) {
	n.HaltCompilerTkn = f.newToken(token.T_HALT_COMPILER, []byte("__halt_compiler"))
	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtIf(n *ast.StmtIf) {
	n.ColonTkn = nil
	n.EndIfTkn = nil
	n.SemiColonTkn = nil

	n.IfTkn = f.newToken(token.T_IF, []byte("if"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.Cond.Accept(f)
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Stmt.Accept(f)

	if len(n.ElseIf) > 0 {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		f.formatList(n.ElseIf, ' ')
	}

	if n.Else != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Else.Accept(f)
	}
}

func (f *formatter) StmtInlineHtml(n *ast.StmtInlineHtml) {
	n.InlineHtmlTkn = f.newToken(token.T_STRING, n.Value)
	f.state = FormatterStateHTML
}

func (f *formatter) StmtInterface(n *ast.StmtInterface) {
	n.InterfaceTkn = f.newToken(token.T_INTERFACE, []byte("interface"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.InterfaceName.Accept(f)
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if n.Extends != nil {
		n.Extends.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

	if len(n.Stmts) > 0 {
		f.indent++
		f.formatStmts(&n.Stmts)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}

	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
}

func (f *formatter) StmtInterfaceExtends(n *ast.StmtInterfaceExtends) {
	n.ExtendsTkn = f.newToken(token.T_EXTENDS, []byte("extends"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.SeparatorTkns = nil
	if len(n.InterfaceNames) > 0 {
		n.SeparatorTkns = f.formatList(n.InterfaceNames, ',')
	}
}

func (f *formatter) StmtLabel(n *ast.StmtLabel) {
	n.LabelName.Accept(f)
	n.ColonTkn = f.newToken(':', []byte(":"))
}

func (f *formatter) StmtNamespace(n *ast.StmtNamespace) {
	n.OpenCurlyBracketTkn = nil
	n.CloseCurlyBracketTkn = nil
	n.SemiColonTkn = nil

	n.NsTkn = f.newToken(token.T_NAMESPACE, []byte("namespace"))

	if n.Name != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Name.Accept(f)
	}

	if len(n.Stmts) > 0 {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))
		if len(n.Stmts) > 0 {
			f.indent++
			f.formatStmts(&n.Stmts)
			f.indent--

			f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
			f.addIndent()
		}
		n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
	} else {
		n.SemiColonTkn = f.newSemicolonTkn()
	}

}

func (f *formatter) StmtNop(n *ast.StmtNop) {
	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtProperty(n *ast.StmtProperty) {
	n.Var.Accept(f)

	if n.Expr != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.EqualTkn = f.newToken('=', []byte("="))
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

		n.Expr.Accept(f)
	}
}

func (f *formatter) StmtPropertyList(n *ast.StmtPropertyList) {
	for _, m := range n.Modifiers {
		m.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	if n.Type != nil {
		n.Type.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.SeparatorTkns = f.formatList(n.Properties, ',')

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtReturn(n *ast.StmtReturn) {
	n.ReturnTkn = f.newToken(token.T_RETURN, []byte("return"))

	if n.Expr != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Expr.Accept(f)
	}

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtStatic(n *ast.StmtStatic) {
	n.StaticTkn = f.newToken(token.T_STATIC, []byte("static"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.SeparatorTkns = nil
	if len(n.Vars) > 0 {
		n.SeparatorTkns = f.formatList(n.Vars, ',')
	}

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtStaticVar(n *ast.StmtStaticVar) {
	n.Var.Accept(f)

	if n.Expr != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.EqualTkn = f.newToken('=', []byte("="))
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

		n.Expr.Accept(f)
	}
}

func (f *formatter) StmtStmtList(n *ast.StmtStmtList) {
	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

	if len(n.Stmts) > 0 {
		f.indent++
		f.formatStmts(&n.Stmts)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}

	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
}

func (f *formatter) StmtSwitch(n *ast.StmtSwitch) {
	n.CaseSeparatorTkn = nil
	n.ColonTkn = nil
	n.EndSwitchTkn = nil
	n.SemiColonTkn = nil

	n.SwitchTkn = f.newToken(token.T_SWITCH, []byte("switch"))

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.Cond.Accept(f)
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

	if len(n.CaseList) > 0 {
		f.indent++
		f.formatStmts(&n.CaseList)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}

	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
}

func (f *formatter) StmtThrow(n *ast.StmtThrow) {
	n.ThrowTkn = f.newToken(token.T_THROW, []byte("throw"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtTrait(n *ast.StmtTrait) {
	n.TraitTkn = f.newToken(token.T_TRAIT, []byte("trait"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.TraitName.Accept(f)
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if n.Extends != nil {
		n.Extends.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	if n.Implements != nil {
		n.Implements.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

	if len(n.Stmts) > 0 {
		f.indent++
		f.formatStmts(&n.Stmts)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}

	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
}

func (f *formatter) StmtTraitMethodRef(n *ast.StmtTraitMethodRef) {
	if n.Trait != nil {
		n.Trait.Accept(f)
		n.DoubleColonTkn = f.newToken(token.T_PAAMAYIM_NEKUDOTAYIM, []byte("::"))
	}

	n.Method.Accept(f)
}

func (f *formatter) StmtTraitUse(n *ast.StmtTraitUse) {
	n.UseTkn = f.newToken(token.T_USE, []byte("use"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.SeparatorTkns = f.formatList(n.Traits, ',')

	n.OpenCurlyBracketTkn = nil
	n.CloseCurlyBracketTkn = nil
	n.SemiColonTkn = nil

	if len(n.Adaptations) > 0 {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

		if len(n.Adaptations) > 0 {
			f.indent++
			f.formatStmts(&n.Adaptations)
			f.indent--

			f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
			f.addIndent()
		}

		n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
	} else {
		n.SemiColonTkn = f.newToken(';', []byte(";"))
	}
}

func (f *formatter) StmtTraitUseAlias(n *ast.StmtTraitUseAlias) {
	n.Ref.Accept(f)
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.AsTkn = f.newToken(token.T_AS, []byte("as"))

	if n.Modifier != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Modifier.Accept(f)
	}

	if n.Alias != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Alias.Accept(f)
	}

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtTraitUsePrecedence(n *ast.StmtTraitUsePrecedence) {
	n.Ref.Accept(f)
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.InsteadofTkn = f.newToken(token.T_INSTEADOF, []byte("insteadof"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.SeparatorTkns = f.formatList(n.Insteadof, ',')

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtTry(n *ast.StmtTry) {
	n.TryTkn = f.newToken(token.T_TRY, []byte("try"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))

	if len(n.Stmts) > 0 {
		f.indent++
		f.formatStmts(&n.Stmts)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}

	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))

	for _, catch := range n.Catches {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		catch.Accept(f)
	}

	if n.Finally != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Finally.Accept(f)
	}
}

func (f *formatter) StmtUnset(n *ast.StmtUnset) {
	n.UnsetTkn = f.newToken(token.T_UNSET, []byte("unset"))

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.SeparatorTkns = f.formatList(n.Vars, ',')
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtUse(n *ast.StmtUse) {
	n.UseTkn = f.newToken(token.T_USE, []byte("use"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if n.Type != nil {
		n.Type.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.SeparatorTkns = f.formatList(n.UseDeclarations, ',')

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtGroupUse(n *ast.StmtGroupUse) {
	n.UseTkn = f.newToken(token.T_USE, []byte("use"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if n.Type != nil {
		n.Type.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.LeadingNsSeparatorTkn = nil

	n.Prefix.Accept(f)
	n.NsSeparatorTkn = f.newToken(token.T_NS_SEPARATOR, []byte("\\"))

	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))
	n.SeparatorTkns = f.formatList(n.UseDeclarations, ',')
	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))

	n.SemiColonTkn = f.newSemicolonTkn()
}

func (f *formatter) StmtUseDeclaration(n *ast.StmtUseDeclaration) {
	if n.Type != nil {
		n.Type.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.NsSeparatorTkn = nil

	n.Use.Accept(f)

	if n.Alias != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.AsTkn = f.newToken(token.T_AS, []byte("as"))
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.Alias.Accept(f)
	}
}

func (f *formatter) StmtWhile(n *ast.StmtWhile) {
	n.ColonTkn = nil
	n.EndWhileTkn = nil
	n.SemiColonTkn = nil

	n.WhileTkn = f.newToken(token.T_WHILE, []byte("while"))
	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.Cond.Accept(f)
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Stmt.Accept(f)
}

func (f *formatter) ExprArray(n *ast.ExprArray) {
	n.ArrayTkn = f.newToken(token.T_ARRAY, []byte("array"))
	n.OpenBracketTkn = f.newToken('(', []byte("("))
	n.SeparatorTkns = f.formatList(n.Items, ',')
	n.CloseBracketTkn = f.newToken(')', []byte(")"))
}

func (f *formatter) ExprArrayDimFetch(n *ast.ExprArrayDimFetch) {
	n.Var.Accept(f)
	n.OpenBracketTkn = f.newToken('[', []byte("["))
	n.Dim.Accept(f)
	n.CloseBracketTkn = f.newToken(']', []byte("]"))
}

func (f *formatter) ExprArrayItem(n *ast.ExprArrayItem) {
	if n.EllipsisTkn != nil {
		n.EllipsisTkn = f.newToken(token.T_ELLIPSIS, []byte("..."))
	}

	if n.Key != nil {
		n.Key.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.DoubleArrowTkn = f.newToken(token.T_DOUBLE_ARROW, []byte("=>"))
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.Val.Accept(f)
}

func (f *formatter) ExprArrowFunction(n *ast.ExprArrowFunction) {
	if n.StaticTkn != nil {
		n.StaticTkn = f.newToken(token.T_STATIC, []byte("static"))
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.FnTkn = f.newToken(token.T_FN, []byte("fn"))

	if n.AmpersandTkn != nil {
		n.AmpersandTkn = f.newToken('&', []byte("&"))
	}

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.SeparatorTkns = nil
	if len(n.Params) > 0 {
		n.SeparatorTkns = f.formatList(n.Params, ',')
	}
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	n.ColonTkn = nil
	if n.ReturnType != nil {
		n.ColonTkn = f.newToken(':', []byte(":"))

		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.ReturnType.Accept(f)
	}

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.DoubleArrowTkn = f.newToken(token.T_DOUBLE_ARROW, []byte("=>"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprBitwiseNot(n *ast.ExprBitwiseNot) {
	n.TildaTkn = f.newToken('~', []byte("~"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprBooleanNot(n *ast.ExprBooleanNot) {
	n.ExclamationTkn = f.newToken('!', []byte("!"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	n.Class.Accept(f)
	n.DoubleColonTkn = f.newToken(token.T_PAAMAYIM_NEKUDOTAYIM, []byte("::"))
	n.ConstantName.Accept(f)
}

func (f *formatter) ExprClone(n *ast.ExprClone) {
	n.CloneTkn = f.newToken(token.T_CLONE, []byte("clone"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Expr.Accept(f)
}

func (f *formatter) ExprClosure(n *ast.ExprClosure) {
	if n.StaticTkn != nil {
		n.StaticTkn = f.newToken(token.T_STATIC, []byte("static"))
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.FunctionTkn = f.newToken(token.T_FN, []byte("function"))

	if n.AmpersandTkn != nil {
		n.AmpersandTkn = f.newToken('&', []byte("&"))
	}

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.SeparatorTkns = nil
	if len(n.Params) > 0 {
		n.SeparatorTkns = f.formatList(n.Params, ',')
	}
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))

	n.UseTkn = nil
	n.UseOpenParenthesisTkn = nil
	n.UseCloseParenthesisTkn = nil
	n.UseSeparatorTkns = nil
	if len(n.Use) > 0 {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.UseTkn = f.newToken(token.T_USE, []byte("use"))
		n.OpenParenthesisTkn = f.newToken('(', []byte("("))
		n.SeparatorTkns = f.formatList(n.Use, ',')
		n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
	}

	n.ColonTkn = nil
	if n.ReturnType != nil {
		n.ColonTkn = f.newToken(':', []byte(":"))

		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.ReturnType.Accept(f)
	}

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpenCurlyBracketTkn = f.newToken('{', []byte("{"))
	if len(n.Stmts) > 0 {
		f.indent++
		f.formatStmts(&n.Stmts)
		f.indent--

		f.addFreeFloating(token.T_WHITESPACE, []byte("\n"))
		f.addIndent()
	}
	n.CloseCurlyBracketTkn = f.newToken('}', []byte("}"))
}

func (f *formatter) ExprClosureUse(n *ast.ExprClosureUse) {
	if n.AmpersandTkn != nil {
		n.AmpersandTkn = f.newToken('&', []byte("&"))
	}

	n.Var.Accept(f)
}

func (f *formatter) ExprConstFetch(n *ast.ExprConstFetch) {
	n.Const.Accept(f)
}

func (f *formatter) ExprEmpty(n *ast.ExprEmpty) {
	n.EmptyTkn = f.newToken(token.T_EMPTY, []byte("empty"))
	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.Expr.Accept(f)
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
}

func (f *formatter) ExprErrorSuppress(n *ast.ExprErrorSuppress) {
	n.AtTkn = f.newToken('@', []byte("@"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprEval(n *ast.ExprEval) {
	n.EvalTkn = f.newToken(token.T_EVAL, []byte("eval"))
	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.Expr.Accept(f)
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
}

func (f *formatter) ExprExit(n *ast.ExprExit) {
	n.DieTkn = f.newToken(token.T_EVAL, []byte("exit"))

	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	if n.Expr != nil {
		n.OpenParenthesisTkn = f.newToken('(', []byte("("))
		n.Expr.Accept(f)
		n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
	}
}

func (f *formatter) ExprFunctionCall(n *ast.ExprFunctionCall) {
	n.Function.Accept(f)
	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.SeparatorTkns = nil
	if len(n.Arguments) > 0 {
		n.SeparatorTkns = f.formatList(n.Arguments, ',')
	}
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
}

func (f *formatter) ExprInclude(n *ast.ExprInclude) {
	n.IncludeTkn = f.newToken(token.T_INCLUDE, []byte("include"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Expr.Accept(f)
}

func (f *formatter) ExprIncludeOnce(n *ast.ExprIncludeOnce) {
	n.IncludeTkn = f.newToken(token.T_INCLUDE_ONCE, []byte("include_once"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Expr.Accept(f)
}

func (f *formatter) ExprInstanceOf(n *ast.ExprInstanceOf) {
	n.Expr.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.InstanceOfTkn = f.newToken(token.T_INSTANCEOF, []byte("instanceof"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Class.Accept(f)
}

func (f *formatter) ExprIsset(n *ast.ExprIsset) {
	n.IssetTkn = f.newToken(token.T_ISSET, []byte("isset"))
	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.SeparatorTkns = f.formatList(n.Vars, ',')
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
}

func (f *formatter) ExprList(n *ast.ExprList) {
	n.ListTkn = f.newToken(token.T_LIST, []byte("list"))
	n.OpenBracketTkn = f.newToken('(', []byte("("))
	n.SeparatorTkns = f.formatList(n.Items, ',')
	n.CloseBracketTkn = f.newToken(')', []byte(")"))
}

func (f *formatter) ExprMethodCall(n *ast.ExprMethodCall) {
	n.Var.Accept(f)
	n.ObjectOperatorTkn = f.newToken(token.T_OBJECT_OPERATOR, []byte("->"))
	n.Method.Accept(f)

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.SeparatorTkns = nil
	if len(n.Arguments) > 0 {
		n.SeparatorTkns = f.formatList(n.Arguments, ',')
	}
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
}

func (f *formatter) ExprNew(n *ast.ExprNew) {
	n.NewTkn = f.newToken(token.T_NEW, []byte("new"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Class.Accept(f)

	n.SeparatorTkns = nil
	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	if len(n.Arguments) > 0 {
		n.OpenParenthesisTkn = f.newToken('(', []byte("("))
		n.SeparatorTkns = f.formatList(n.Arguments, ',')
		n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
	}
}

func (f *formatter) ExprPostDec(n *ast.ExprPostDec) {
	n.Var.Accept(f)
	n.DecTkn = f.newToken(token.T_DEC, []byte("--"))
}

func (f *formatter) ExprPostInc(n *ast.ExprPostInc) {
	n.Var.Accept(f)
	n.IncTkn = f.newToken(token.T_INC, []byte("++"))
}

func (f *formatter) ExprPreDec(n *ast.ExprPreDec) {
	n.DecTkn = f.newToken(token.T_DEC, []byte("--"))
	n.Var.Accept(f)
}

func (f *formatter) ExprPreInc(n *ast.ExprPreInc) {
	n.IncTkn = f.newToken(token.T_INC, []byte("++"))
	n.Var.Accept(f)
}

func (f *formatter) ExprPrint(n *ast.ExprPrint) {
	n.PrintTkn = f.newToken(token.T_PRINT, []byte("print"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprPropertyFetch(n *ast.ExprPropertyFetch) {
	n.Var.Accept(f)
	n.ObjectOperatorTkn = f.newToken(token.T_OBJECT_OPERATOR, []byte("->"))
	n.Property.Accept(f)
}

func (f *formatter) ExprRequire(n *ast.ExprRequire) {
	n.RequireTkn = f.newToken(token.T_REQUIRE, []byte("require"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Expr.Accept(f)
}

func (f *formatter) ExprRequireOnce(n *ast.ExprRequireOnce) {
	n.RequireOnceTkn = f.newToken(token.T_REQUIRE_ONCE, []byte("require_once"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.Expr.Accept(f)
}

func (f *formatter) ExprShellExec(n *ast.ExprShellExec) {
	n.OpenBacktickTkn = f.newToken('`', []byte("`"))
	for _, p := range n.Parts {
		p.Accept(f)
	}
	n.CloseBacktickTkn = f.newToken('`', []byte("`"))
}

func (f *formatter) ExprStaticCall(n *ast.ExprStaticCall) {
	n.Class.Accept(f)
	n.DoubleColonTkn = f.newToken(token.T_PAAMAYIM_NEKUDOTAYIM, []byte("::"))
	n.Call.Accept(f)

	n.OpenParenthesisTkn = f.newToken('(', []byte("("))
	n.SeparatorTkns = nil
	if len(n.Arguments) > 0 {
		n.SeparatorTkns = f.formatList(n.Arguments, ',')
	}
	n.CloseParenthesisTkn = f.newToken(')', []byte(")"))
}

func (f *formatter) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch) {
	n.Class.Accept(f)
	n.DoubleColonTkn = f.newToken(token.T_PAAMAYIM_NEKUDOTAYIM, []byte("::"))
	n.Property.Accept(f)
}

func (f *formatter) ExprTernary(n *ast.ExprTernary) {
	n.Condition.Accept(f)
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.QuestionTkn = f.newToken('?', []byte("?"))
	if n.IfTrue != nil {
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.IfTrue.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}
	n.ColonTkn = f.newToken(':', []byte(":"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.IfFalse.Accept(f)
}

func (f *formatter) ExprUnaryMinus(n *ast.ExprUnaryMinus) {
	n.MinusTkn = f.newToken('-', []byte("-"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprUnaryPlus(n *ast.ExprUnaryPlus) {
	n.PlusTkn = f.newToken('+', []byte("+"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprVariable(n *ast.ExprVariable) {
	if _, ok := n.VarName.(*ast.Identifier); !ok {
		n.DollarTkn = f.newToken('$', []byte("$"))
	}

	n.VarName.Accept(f)
}

func (f *formatter) ExprYield(n *ast.ExprYield) {
	n.YieldTkn = f.newToken(token.T_YIELD, []byte("yield"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	if n.Key != nil {
		n.Key.Accept(f)
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
		n.DoubleArrowTkn = f.newToken(token.T_DOUBLE_ARROW, []byte("=>"))
		f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	}

	n.Value.Accept(f)
}

func (f *formatter) ExprYieldFrom(n *ast.ExprYieldFrom) {
	n.YieldFromTkn = f.newToken(token.T_YIELD_FROM, []byte("yield from"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssign(n *ast.ExprAssign) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken('=', []byte("="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignReference(n *ast.ExprAssignReference) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken('=', []byte("="))
	n.AmpersandTkn = f.newToken('&', []byte("&"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignBitwiseAnd(n *ast.ExprAssignBitwiseAnd) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_AND_EQUAL, []byte("&="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignBitwiseOr(n *ast.ExprAssignBitwiseOr) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_OR_EQUAL, []byte("|="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignBitwiseXor(n *ast.ExprAssignBitwiseXor) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_XOR_EQUAL, []byte("^="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignCoalesce(n *ast.ExprAssignCoalesce) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_COALESCE_EQUAL, []byte("??="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignConcat(n *ast.ExprAssignConcat) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_CONCAT_EQUAL, []byte(".="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignDiv(n *ast.ExprAssignDiv) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_DIV_EQUAL, []byte("/="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignMinus(n *ast.ExprAssignMinus) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_MINUS_EQUAL, []byte("-="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignMod(n *ast.ExprAssignMod) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_MOD_EQUAL, []byte("%="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignMul(n *ast.ExprAssignMul) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_MUL_EQUAL, []byte("*="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignPlus(n *ast.ExprAssignPlus) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_PLUS_EQUAL, []byte("+="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignPow(n *ast.ExprAssignPow) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_POW_EQUAL, []byte("**="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignShiftLeft(n *ast.ExprAssignShiftLeft) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_SL_EQUAL, []byte("<<="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprAssignShiftRight(n *ast.ExprAssignShiftRight) {
	n.Var.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.EqualTkn = f.newToken(token.T_SR_EQUAL, []byte(">>="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Expr.Accept(f)
}

func (f *formatter) ExprBinaryBitwiseAnd(n *ast.ExprBinaryBitwiseAnd) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('&', []byte("&"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryBitwiseOr(n *ast.ExprBinaryBitwiseOr) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('|', []byte("|"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryBitwiseXor(n *ast.ExprBinaryBitwiseXor) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('^', []byte("^"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryBooleanAnd(n *ast.ExprBinaryBooleanAnd) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_BOOLEAN_AND, []byte("&&"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryBooleanOr(n *ast.ExprBinaryBooleanOr) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_BOOLEAN_OR, []byte("||"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryCoalesce(n *ast.ExprBinaryCoalesce) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_COALESCE, []byte("??"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryConcat(n *ast.ExprBinaryConcat) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('.', []byte("."))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryDiv(n *ast.ExprBinaryDiv) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('/', []byte("/"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryEqual(n *ast.ExprBinaryEqual) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_IS_EQUAL, []byte("=="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryGreater(n *ast.ExprBinaryGreater) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('>', []byte(">"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryGreaterOrEqual(n *ast.ExprBinaryGreaterOrEqual) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_IS_GREATER_OR_EQUAL, []byte(">="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryIdentical(n *ast.ExprBinaryIdentical) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_IS_IDENTICAL, []byte("==="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryLogicalAnd(n *ast.ExprBinaryLogicalAnd) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_LOGICAL_AND, []byte("and"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryLogicalOr(n *ast.ExprBinaryLogicalOr) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_LOGICAL_OR, []byte("or"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryLogicalXor(n *ast.ExprBinaryLogicalXor) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_LOGICAL_XOR, []byte("xor"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryMinus(n *ast.ExprBinaryMinus) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('-', []byte("-"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryMod(n *ast.ExprBinaryMod) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('%', []byte("%"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryMul(n *ast.ExprBinaryMul) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('*', []byte("*"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryNotEqual(n *ast.ExprBinaryNotEqual) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_IS_NOT_EQUAL, []byte("!="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryNotIdentical(n *ast.ExprBinaryNotIdentical) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_IS_NOT_IDENTICAL, []byte("!=="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryPlus(n *ast.ExprBinaryPlus) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('+', []byte("+"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryPow(n *ast.ExprBinaryPow) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_POW, []byte("**"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryShiftLeft(n *ast.ExprBinaryShiftLeft) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_SL, []byte("<<"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinaryShiftRight(n *ast.ExprBinaryShiftRight) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_SR, []byte(">>"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinarySmaller(n *ast.ExprBinarySmaller) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken('<', []byte("<"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinarySmallerOrEqual(n *ast.ExprBinarySmallerOrEqual) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_IS_SMALLER_OR_EQUAL, []byte("<="))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprBinarySpaceship(n *ast.ExprBinarySpaceship) {
	n.Left.Accept(f)

	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))
	n.OpTkn = f.newToken(token.T_SPACESHIP, []byte("<=>"))
	f.addFreeFloating(token.T_WHITESPACE, []byte(" "))

	n.Right.Accept(f)
}

func (f *formatter) ExprCastArray(n *ast.ExprCastArray) {
	n.CastTkn = f.newToken(token.T_ARRAY_CAST, []byte("(array)"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprCastBool(n *ast.ExprCastBool) {
	n.CastTkn = f.newToken(token.T_BOOL_CAST, []byte("(bool)"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprCastDouble(n *ast.ExprCastDouble) {
	n.CastTkn = f.newToken(token.T_DOUBLE_CAST, []byte("(float)"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprCastInt(n *ast.ExprCastInt) {
	n.CastTkn = f.newToken(token.T_INT_CAST, []byte("(integer)"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprCastObject(n *ast.ExprCastObject) {
	n.CastTkn = f.newToken(token.T_OBJECT_CAST, []byte("(object)"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprCastString(n *ast.ExprCastString) {
	n.CastTkn = f.newToken(token.T_STRING_CAST, []byte("(string)"))
	n.Expr.Accept(f)
}

func (f *formatter) ExprCastUnset(n *ast.ExprCastUnset) {
	n.CastTkn = f.newToken(token.T_UNSET_CAST, []byte("(unset)"))
	n.Expr.Accept(f)
}

func (f *formatter) ScalarDnumber(n *ast.ScalarDnumber) {
	if n.NumberTkn == nil {
		n.NumberTkn = f.newToken(token.T_STRING, n.Value)
	} else {
		n.NumberTkn.FreeFloating = f.getFreeFloating()
	}
}

func (f *formatter) ScalarEncapsed(n *ast.ScalarEncapsed) {
	n.OpenQuoteTkn = f.newToken('"', []byte("\""))
	for _, p := range n.Parts {
		p.Accept(f)
	}
	n.CloseQuoteTkn = f.newToken('"', []byte("\""))
}

func (f *formatter) ScalarEncapsedStringPart(n *ast.ScalarEncapsedStringPart) {
	if n.EncapsedStrTkn == nil {
		n.EncapsedStrTkn = f.newToken(token.T_STRING, n.Value)
	} else {
		n.EncapsedStrTkn.FreeFloating = f.getFreeFloating()
	}
}

func (f *formatter) ScalarHeredoc(n *ast.ScalarHeredoc) {
	n.OpenHeredocTkn = f.newToken(token.T_START_HEREDOC, []byte("<<<EOT\n"))
	for _, p := range n.Parts {
		p.Accept(f)
	}
	n.CloseHeredocTkn = f.newToken(token.T_START_HEREDOC, []byte("EOT"))
}

func (f *formatter) ScalarLnumber(n *ast.ScalarLnumber) {
	if n.NumberTkn == nil {
		n.NumberTkn = f.newToken(token.T_STRING, n.Value)
	} else {
		n.NumberTkn.FreeFloating = f.getFreeFloating()
	}
}

func (f *formatter) ScalarMagicConstant(n *ast.ScalarMagicConstant) {
	if n.MagicConstTkn == nil {
		n.MagicConstTkn = f.newToken(token.T_STRING, n.Value)
	} else {
		n.MagicConstTkn.FreeFloating = f.getFreeFloating()
	}
}

func (f *formatter) ScalarString(n *ast.ScalarString) {
	if n.StringTkn == nil {
		n.StringTkn = f.newToken(token.T_STRING, n.Value)
	} else {
		n.StringTkn.FreeFloating = f.getFreeFloating()
	}
}

func (f *formatter) NameName(n *ast.NameName) {
	separatorTkns := make([]*token.Token, len(n.Parts)-1)
	for i, v := range n.Parts {
		v.Accept(f)

		if i != len(n.Parts)-1 {
			separatorTkns[i] = f.newToken(token.T_NS_SEPARATOR, []byte("\\"))
		}
	}
}

func (f *formatter) NameFullyQualified(n *ast.NameFullyQualified) {
	n.NsSeparatorTkn = f.newToken(token.T_NS_SEPARATOR, []byte("\\"))

	separatorTkns := make([]*token.Token, len(n.Parts)-1)
	for i, v := range n.Parts {
		v.Accept(f)

		if i != len(n.Parts)-1 {
			separatorTkns[i] = f.newToken(token.T_NS_SEPARATOR, []byte("\\"))
		}
	}
}

func (f *formatter) NameRelative(n *ast.NameRelative) {
	n.NsTkn = f.newToken(token.T_NAMESPACE, []byte("namespace"))
	n.NsSeparatorTkn = f.newToken(token.T_NS_SEPARATOR, []byte("\\"))

	separatorTkns := make([]*token.Token, len(n.Parts)-1)
	for i, v := range n.Parts {
		v.Accept(f)

		if i != len(n.Parts)-1 {
			separatorTkns[i] = f.newToken(token.T_NS_SEPARATOR, []byte("\\"))
		}
	}
}

func (f *formatter) NameNamePart(n *ast.NameNamePart) {
	if n.StringTkn == nil {
		n.StringTkn = f.newToken(token.T_STRING, n.Value)
	} else {
		n.StringTkn.FreeFloating = f.getFreeFloating()
	}
}

func (f *formatter) ParserBrackets(n *ast.ParserBrackets) {
	// TODO
}
