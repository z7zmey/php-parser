package visitor

import (
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
	"io"
	"strconv"
	"strings"

	"github.com/z7zmey/php-parser/pkg/ast"
)

type Dumper struct {
	writer        io.Writer
	indent        int
	withTokens    bool
	withPositions bool
}

func NewDumper(writer io.Writer) *Dumper {
	return &Dumper{writer: writer}
}

func (v *Dumper) WithTokens() *Dumper {
	v.withTokens = true
	return v
}

func (v *Dumper) WithPositions() *Dumper {
	v.withPositions = true
	return v
}

func (v *Dumper) Dump(n ast.Vertex) {
	n.Accept(v)
}

func (v *Dumper) print(indent int, str string) {
	_, err := io.WriteString(v.writer, strings.Repeat("\t", indent))
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(v.writer, str)
	if err != nil {
		panic(err)
	}
}

func (v *Dumper) dumpVertex(key string, node ast.Vertex) {
	if node == nil {
		return
	}

	v.print(v.indent, key+": ")
	node.Accept(v)
}

func (v *Dumper) dumpVertexList(key string, list []ast.Vertex) {
	if list == nil {
		return
	}

	if len(list) == 0 {
		v.print(v.indent, key+": []ast.Vertex{},\n")
		return
	}

	v.print(v.indent, key+": []ast.Vertex{\n")
	v.indent++

	for _, nn := range list {
		v.print(v.indent, "")
		nn.Accept(v)
	}

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) dumpToken(key string, tok *token.Token) {
	if !v.withTokens {
		return
	}

	if tok == nil {
		return
	}

	if key == "" {
		v.print(v.indent, "{\n")
	} else {
		v.print(v.indent, key+": &token.Token{\n")
	}

	v.indent++

	if tok.ID > 0 {
		v.print(v.indent, "ID: token."+tok.ID.String()+",\n")
	}
	if tok.Value != nil {
		v.print(v.indent, "Value: []byte("+strconv.Quote(string(tok.Value))+"),\n")
	}
	v.dumpPosition(tok.Position)
	v.dumpTokenList("FreeFloating", tok.FreeFloating)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) dumpTokenList(key string, list []*token.Token) {
	if !v.withTokens {
		return
	}

	if list == nil {
		return
	}

	if len(list) == 0 {
		v.print(v.indent, key+": []*token.Token{},\n")
		return
	}

	v.print(v.indent, key+": []*token.Token{\n")
	v.indent++

	for _, tok := range list {
		v.dumpToken("", tok)
	}

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) dumpPosition(pos *position.Position) {
	if !v.withPositions {
		return
	}

	if pos == nil {
		return
	}

	v.print(v.indent, "Position: &position.Position{\n")
	v.indent++

	v.print(v.indent, "StartLine: "+strconv.Itoa(pos.StartLine)+",\n")
	v.print(v.indent, "EndLine:   "+strconv.Itoa(pos.EndLine)+",\n")
	v.print(v.indent, "StartPos:  "+strconv.Itoa(pos.StartPos)+",\n")
	v.print(v.indent, "EndPos:    "+strconv.Itoa(pos.EndPos)+",\n")

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) dumpValue(key string, val []byte) {
	if val == nil {
		return
	}

	v.print(v.indent, key+": []byte("+strconv.Quote(string(val))+"),\n")

}

func (v *Dumper) Root(n *ast.Root) {
	v.print(0, "&ast.Root{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("EndTkn", n.EndTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) Nullable(n *ast.Nullable) {
	v.print(0, "&ast.Nullable{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("QuestionTkn", n.QuestionTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) Parameter(n *ast.Parameter) {
	v.print(0, "&ast.Parameter{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Type", n.Type)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpToken("VariadicTkn", n.VariadicTkn)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("DefaultValue", n.DefaultValue)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) Identifier(n *ast.Identifier) {
	v.print(0, "&ast.Identifier{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("IdentifierTkn", n.IdentifierTkn)
	v.dumpValue("Value", n.Value)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) Argument(n *ast.Argument) {
	v.print(0, "&ast.Argument{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpToken("VariadicTkn", n.VariadicTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtBreak(n *ast.StmtBreak) {
	v.print(0, "&ast.StmtBreak{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("BreakTkn", n.BreakTkn)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtCase(n *ast.StmtCase) {
	v.print(0, "&ast.StmtCase{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CaseTkn", n.CaseTkn)
	v.dumpVertex("Cond", n.Cond)
	v.dumpToken("CaseSeparatorTkn", n.CaseSeparatorTkn)
	v.dumpVertexList("Stmts", n.Stmts)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtCatch(n *ast.StmtCatch) {
	v.print(0, "&ast.StmtCatch{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CatchTkn", n.CatchTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Types", n.Types)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtClass(n *ast.StmtClass) {
	v.print(0, "&ast.StmtClass{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertexList("Modifiers", n.Modifiers)
	v.dumpToken("ClassTkn", n.ClassTkn)
	v.dumpVertex("ClassName", n.ClassName)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Arguments", n.Arguments)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpVertex("Extends", n.Extends)
	v.dumpVertex("Implements", n.Implements)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtClassConstList(n *ast.StmtClassConstList) {
	v.print(0, "&ast.StmtClassConstList{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertexList("Modifiers", n.Modifiers)
	v.dumpToken("ConstTkn", n.ConstTkn)
	v.dumpVertexList("Consts", n.Consts)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtClassExtends(n *ast.StmtClassExtends) {
	v.print(0, "&ast.StmtClassExtends{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ExtendTkn", n.ExtendTkn)
	v.dumpVertex("ClassName", n.ClassName)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtClassImplements(n *ast.StmtClassImplements) {
	v.print(0, "&ast.StmtClassImplements{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ImplementsTkn", n.ImplementsTkn)
	v.dumpVertexList("InterfaceNames", n.InterfaceNames)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtClassMethod(n *ast.StmtClassMethod) {
	v.print(0, "&ast.StmtClassMethod{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertexList("Modifiers", n.Modifiers)
	v.dumpToken("FunctionTkn", n.FunctionTkn)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpVertex("MethodName", n.MethodName)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Params", n.Params)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("ReturnType", n.ReturnType)
	v.dumpVertex("Stmt", n.Stmt)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtConstList(n *ast.StmtConstList) {
	v.print(0, "&ast.StmtConstList{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ConstTkn", n.ConstTkn)
	v.dumpVertexList("Consts", n.Consts)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtConstant(n *ast.StmtConstant) {
	v.print(0, "&ast.StmtConstant{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Name", n.Name)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtContinue(n *ast.StmtContinue) {
	v.print(0, "&ast.StmtContinue{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ContinueTkn", n.ContinueTkn)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtDeclare(n *ast.StmtDeclare) {
	v.print(0, "&ast.StmtDeclare{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("DeclareTkn", n.DeclareTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Consts", n.Consts)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("Stmt", n.Stmt)
	v.dumpToken("EndDeclareTkn", n.EndDeclareTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtDefault(n *ast.StmtDefault) {
	v.print(0, "&ast.StmtDefault{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("DefaultTkn", n.DefaultTkn)
	v.dumpToken("CaseSeparatorTkn", n.CaseSeparatorTkn)
	v.dumpVertexList("Stmts", n.Stmts)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtDo(n *ast.StmtDo) {
	v.print(0, "&ast.StmtDo{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("DoTkn", n.DoTkn)
	v.dumpVertex("Stmt", n.Stmt)
	v.dumpToken("WhileTkn", n.WhileTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertex("Cond", n.Cond)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")

}

func (v *Dumper) StmtEcho(n *ast.StmtEcho) {
	v.print(0, "&ast.StmtEcho{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("EchoTkn", n.EchoTkn)
	v.dumpVertexList("Exprs", n.Exprs)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtElse(n *ast.StmtElse) {
	v.print(0, "&ast.StmtElse{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ElseTkn", n.ElseTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("Stmt", n.Stmt)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtElseIf(n *ast.StmtElseIf) {
	v.print(0, "&ast.StmtElseIf{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ElseIfTkn", n.ElseIfTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertex("Cond", n.Cond)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("Stmt", n.Stmt)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtExpression(n *ast.StmtExpression) {
	v.print(0, "&ast.StmtExpression{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtFinally(n *ast.StmtFinally) {
	v.print(0, "&ast.StmtFinally{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("FinallyTkn", n.FinallyTkn)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtFor(n *ast.StmtFor) {
	v.print(0, "&ast.StmtFor{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ForTkn", n.ForTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Init", n.Init)
	v.dumpTokenList("InitSeparatorTkns", n.InitSeparatorTkns)
	v.dumpToken("InitSemiColonTkn", n.InitSemiColonTkn)
	v.dumpVertexList("Cond", n.Cond)
	v.dumpTokenList("CondSeparatorTkns", n.CondSeparatorTkns)
	v.dumpToken("CondSemiColonTkn", n.CondSemiColonTkn)
	v.dumpVertexList("Loop", n.Loop)
	v.dumpTokenList("LoopSeparatorTkns", n.LoopSeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("Stmt", n.Stmt)
	v.dumpToken("EndForTkn", n.EndForTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtForeach(n *ast.StmtForeach) {
	v.print(0, "&ast.StmtForeach{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ForeachTkn", n.ForeachTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("AsTkn", n.AsTkn)
	v.dumpVertex("Key", n.Key)
	v.dumpToken("DoubleArrowTkn", n.DoubleArrowTkn)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("Stmt", n.Stmt)
	v.dumpToken("EndForeachTkn", n.EndForeachTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtFunction(n *ast.StmtFunction) {
	v.print(0, "&ast.StmtFunction{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("FunctionTkn", n.FunctionTkn)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpVertex("FunctionName", n.FunctionName)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Params", n.Params)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("ReturnType", n.ReturnType)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtGlobal(n *ast.StmtGlobal) {
	v.print(0, "&ast.StmtGlobal{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("GlobalTkn", n.GlobalTkn)
	v.dumpVertexList("Vars", n.Vars)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtGoto(n *ast.StmtGoto) {
	v.print(0, "&ast.StmtGoto{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("GotoTkn", n.GotoTkn)
	v.dumpVertex("Label", n.Label)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtHaltCompiler(n *ast.StmtHaltCompiler) {
	v.print(0, "&ast.StmtHaltCompiler{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("HaltCompilerTkn", n.HaltCompilerTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtIf(n *ast.StmtIf) {
	v.print(0, "&ast.StmtIf{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("IfTkn", n.IfTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertex("Cond", n.Cond)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("Stmt", n.Stmt)
	v.dumpVertexList("ElseIf", n.ElseIf)
	v.dumpVertex("Else", n.Else)
	v.dumpToken("EndIfTkn", n.EndIfTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtInlineHtml(n *ast.StmtInlineHtml) {
	v.print(0, "&ast.StmtInlineHtml{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("InlineHtmlTkn", n.InlineHtmlTkn)
	v.dumpValue("Value", n.Value)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtInterface(n *ast.StmtInterface) {
	v.print(0, "&ast.StmtInterface{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("InterfaceTkn", n.InterfaceTkn)
	v.dumpVertex("InterfaceName", n.InterfaceName)
	v.dumpVertex("Extends", n.Extends)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtInterfaceExtends(n *ast.StmtInterfaceExtends) {
	v.print(0, "&ast.StmtInterfaceExtends{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ExtendsTkn", n.ExtendsTkn)
	v.dumpVertexList("InterfaceNames", n.InterfaceNames)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtLabel(n *ast.StmtLabel) {
	v.print(0, "&ast.StmtLabel{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("LabelName", n.LabelName)
	v.dumpToken("ColonTkn", n.ColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtNamespace(n *ast.StmtNamespace) {
	v.print(0, "&ast.StmtNamespace{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("NsTkn", n.NsTkn)
	v.dumpVertex("Name", n.Name)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtNop(n *ast.StmtNop) {
	v.print(0, "&ast.StmtNop{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtProperty(n *ast.StmtProperty) {
	v.print(0, "&ast.StmtProperty{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtPropertyList(n *ast.StmtPropertyList) {
	v.print(0, "&ast.StmtPropertyList{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertexList("Modifiers", n.Modifiers)
	v.dumpVertex("Type", n.Type)
	v.dumpVertexList("Properties", n.Properties)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtReturn(n *ast.StmtReturn) {
	v.print(0, "&ast.StmtReturn{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ReturnTkn", n.ReturnTkn)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtStatic(n *ast.StmtStatic) {
	v.print(0, "&ast.StmtStatic{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("StaticTkn", n.StaticTkn)
	v.dumpVertexList("Vars", n.Vars)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtStaticVar(n *ast.StmtStaticVar) {
	v.print(0, "&ast.StmtStaticVar{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtStmtList(n *ast.StmtStmtList) {
	v.print(0, "&ast.StmtStmtList{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtSwitch(n *ast.StmtSwitch) {
	v.print(0, "&ast.StmtSwitch{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("SwitchTkn", n.SwitchTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertex("Cond", n.Cond)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpToken("CaseSeparatorTkn", n.CaseSeparatorTkn)
	v.dumpVertexList("CaseList", n.CaseList)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)
	v.dumpToken("EndSwitchTkn", n.EndSwitchTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtThrow(n *ast.StmtThrow) {
	v.print(0, "&ast.StmtThrow{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ThrowTkn", n.ThrowTkn)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtTrait(n *ast.StmtTrait) {
	v.print(0, "&ast.StmtTrait{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("TraitTkn", n.TraitTkn)
	v.dumpVertex("TraitName", n.TraitName)
	v.dumpVertex("Extends", n.Extends)
	v.dumpVertex("Implements", n.Implements)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtTraitMethodRef(n *ast.StmtTraitMethodRef) {
	v.print(0, "&ast.StmtTraitMethodRef{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Trait", n.Trait)
	v.dumpToken("DoubleColonTkn", n.DoubleColonTkn)
	v.dumpVertex("Method", n.Method)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtTraitUse(n *ast.StmtTraitUse) {
	v.print(0, "&ast.StmtTraitUse{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("UseTkn", n.UseTkn)
	v.dumpVertexList("Traits", n.Traits)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Adaptations", n.Adaptations)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtTraitUseAlias(n *ast.StmtTraitUseAlias) {
	v.print(0, "&ast.StmtTraitUseAlias{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Ref", n.Ref)
	v.dumpToken("AsTkn", n.AsTkn)
	v.dumpVertex("Modifier", n.Modifier)
	v.dumpVertex("Alias", n.Alias)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtTraitUsePrecedence(n *ast.StmtTraitUsePrecedence) {
	v.print(0, "&ast.StmtTraitUsePrecedence{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Ref", n.Ref)
	v.dumpToken("InsteadofTkn", n.InsteadofTkn)
	v.dumpVertexList("Insteadof", n.Insteadof)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtTry(n *ast.StmtTry) {
	v.print(0, "&ast.StmtTry{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("TryTkn", n.TryTkn)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)
	v.dumpVertexList("Catches", n.Catches)
	v.dumpVertex("Finally", n.Finally)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtUnset(n *ast.StmtUnset) {
	v.print(0, "&ast.StmtUnset{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("UnsetTkn", n.UnsetTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Vars", n.Vars)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtUse(n *ast.StmtUse) {
	v.print(0, "&ast.StmtUse{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("UseTkn", n.UseTkn)
	v.dumpVertex("Type", n.Type)
	v.dumpVertexList("UseDeclarations", n.UseDeclarations)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtGroupUse(n *ast.StmtGroupUse) {
	v.print(0, "&ast.StmtGroupUse{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("UseTkn", n.UseTkn)
	v.dumpVertex("Type", n.Type)
	v.dumpToken("LeadingNsSeparatorTkn", n.LeadingNsSeparatorTkn)
	v.dumpVertex("Prefix", n.Prefix)
	v.dumpToken("NsSeparatorTkn", n.NsSeparatorTkn)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("UseDeclarations", n.UseDeclarations)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtUseDeclaration(n *ast.StmtUseDeclaration) {
	v.print(0, "&ast.StmtUseDeclaration{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Type", n.Type)
	v.dumpToken("NsSeparatorTkn", n.NsSeparatorTkn)
	v.dumpVertex("Use", n.Use)
	v.dumpToken("AsTkn", n.AsTkn)
	v.dumpVertex("Alias", n.Alias)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) StmtWhile(n *ast.StmtWhile) {
	v.print(0, "&ast.StmtWhile{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("WhileTkn", n.WhileTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertex("Cond", n.Cond)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("Stmt", n.Stmt)
	v.dumpToken("EndWhileTkn", n.EndWhileTkn)
	v.dumpToken("SemiColonTkn", n.SemiColonTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprArray(n *ast.ExprArray) {
	v.print(0, "&ast.ExprArray{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ArrayTkn", n.ArrayTkn)
	v.dumpToken("OpenBracketTkn", n.OpenBracketTkn)
	v.dumpVertexList("Items", n.Items)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseBracketTkn", n.CloseBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprArrayDimFetch(n *ast.ExprArrayDimFetch) {
	v.print(0, "&ast.ExprArrayDimFetch{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("OpenBracketTkn", n.OpenBracketTkn)
	v.dumpVertex("Dim", n.Dim)
	v.dumpToken("CloseBracketTkn", n.CloseBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprArrayItem(n *ast.ExprArrayItem) {
	v.print(0, "&ast.ExprArrayItem{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("EllipsisTkn", n.EllipsisTkn)
	v.dumpVertex("Key", n.Key)
	v.dumpToken("DoubleArrowTkn", n.DoubleArrowTkn)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpVertex("Val", n.Val)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprArrowFunction(n *ast.ExprArrowFunction) {
	v.print(0, "&ast.ExprArrowFunction{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("StaticTkn", n.StaticTkn)
	v.dumpToken("FnTkn", n.FnTkn)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Params", n.Params)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("ReturnType", n.ReturnType)
	v.dumpToken("DoubleArrowTkn", n.DoubleArrowTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBitwiseNot(n *ast.ExprBitwiseNot) {
	v.print(0, "&ast.ExprBitwiseNot{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("TildaTkn", n.TildaTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBooleanNot(n *ast.ExprBooleanNot) {
	v.print(0, "&ast.ExprBooleanNot{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ExclamationTkn", n.ExclamationTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	v.print(0, "&ast.ExprClassConstFetch{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Class", n.Class)
	v.dumpToken("DoubleColonTkn", n.DoubleColonTkn)
	v.dumpVertex("ConstantName", n.ConstantName)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprClone(n *ast.ExprClone) {
	v.print(0, "&ast.ExprClone{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CloneTkn", n.CloneTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprClosure(n *ast.ExprClosure) {
	v.print(0, "&ast.ExprClosure{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("StaticTkn", n.StaticTkn)
	v.dumpToken("FunctionTkn", n.FunctionTkn)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Params", n.Params)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)
	v.dumpToken("UseTkn", n.UseTkn)
	v.dumpToken("UseOpenParenthesisTkn", n.UseOpenParenthesisTkn)
	v.dumpVertexList("Use", n.Use)
	v.dumpTokenList("UseSeparatorTkns", n.UseSeparatorTkns)
	v.dumpToken("UseCloseParenthesisTkn", n.UseCloseParenthesisTkn)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("ReturnType", n.ReturnType)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertexList("Stmts", n.Stmts)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprClosureUse(n *ast.ExprClosureUse) {
	v.print(0, "&ast.ExprClosureUse{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpVertex("Var", n.Var)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprConstFetch(n *ast.ExprConstFetch) {
	v.print(0, "&ast.ExprConstFetch{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Const", n.Const)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprEmpty(n *ast.ExprEmpty) {
	v.print(0, "&ast.ExprEmpty{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("EmptyTkn", n.EmptyTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprErrorSuppress(n *ast.ExprErrorSuppress) {
	v.print(0, "&ast.ExprErrorSuppress{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("AtTkn", n.AtTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprEval(n *ast.ExprEval) {
	v.print(0, "&ast.ExprEval{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("EvalTkn", n.EvalTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprExit(n *ast.ExprExit) {
	v.print(0, "&ast.ExprExit{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ExitTkn", n.ExitTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprFunctionCall(n *ast.ExprFunctionCall) {
	v.print(0, "&ast.ExprFunctionCall{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Function", n.Function)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Arguments", n.Arguments)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprInclude(n *ast.ExprInclude) {
	v.print(0, "&ast.ExprInclude{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("IncludeOnceTkn", n.IncludeTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprIncludeOnce(n *ast.ExprIncludeOnce) {
	v.print(0, "&ast.ExprIncludeOnce{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("IncludeOnceTkn", n.IncludeOnceTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprInstanceOf(n *ast.ExprInstanceOf) {
	v.print(0, "&ast.ExprInstanceOf{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Expr", n.Expr)
	v.dumpToken("InstanceOfTkn", n.InstanceOfTkn)
	v.dumpVertex("Class", n.Class)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprIsset(n *ast.ExprIsset) {
	v.print(0, "&ast.ExprIsset{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("IssetTkn", n.IssetTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Vars", n.Vars)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprList(n *ast.ExprList) {
	v.print(0, "&ast.ExprList{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("ListTkn", n.ListTkn)
	v.dumpToken("OpenBracketTkn", n.OpenBracketTkn)
	v.dumpVertexList("Items", n.Items)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseBracketTkn", n.CloseBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprMethodCall(n *ast.ExprMethodCall) {
	v.print(0, "&ast.ExprMethodCall{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("ObjectOperatorTkn", n.ObjectOperatorTkn)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertex("Method", n.Method)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Arguments", n.Arguments)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprNew(n *ast.ExprNew) {
	v.print(0, "&ast.ExprNew{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("NewTkn", n.NewTkn)
	v.dumpVertex("Class", n.Class)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Arguments", n.Arguments)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprPostDec(n *ast.ExprPostDec) {
	v.print(0, "&ast.ExprPostDec{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("DecTkn", n.DecTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprPostInc(n *ast.ExprPostInc) {
	v.print(0, "&ast.ExprPostInc{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("IncTkn", n.IncTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprPreDec(n *ast.ExprPreDec) {
	v.print(0, "&ast.ExprPreDec{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("DecTkn", n.DecTkn)
	v.dumpVertex("Var", n.Var)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprPreInc(n *ast.ExprPreInc) {
	v.print(0, "&ast.ExprPreInc{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("IncTkn", n.IncTkn)
	v.dumpVertex("Var", n.Var)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprPrint(n *ast.ExprPrint) {
	v.print(0, "&ast.ExprPrint{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("PrintTkn", n.PrintTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprPropertyFetch(n *ast.ExprPropertyFetch) {
	v.print(0, "&ast.ExprPropertyFetch{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("ObjectOperatorTkn", n.ObjectOperatorTkn)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertex("Property", n.Property)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprRequire(n *ast.ExprRequire) {
	v.print(0, "&ast.ExprRequire{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("RequireTkn", n.RequireTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprRequireOnce(n *ast.ExprRequireOnce) {
	v.print(0, "&ast.ExprRequireOnce{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("RequireOnceTkn", n.RequireOnceTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprShellExec(n *ast.ExprShellExec) {
	v.print(0, "&ast.ExprShellExec{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("OpenBacktickTkn", n.OpenBacktickTkn)
	v.dumpVertexList("Parts", n.Parts)
	v.dumpToken("CloseBacktickTkn", n.CloseBacktickTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprStaticCall(n *ast.ExprStaticCall) {
	v.print(0, "&ast.ExprStaticCall{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Class", n.Class)
	v.dumpToken("DoubleColonTkn", n.DoubleColonTkn)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertex("Call", n.Call)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)
	v.dumpToken("OpenParenthesisTkn", n.OpenParenthesisTkn)
	v.dumpVertexList("Arguments", n.Arguments)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)
	v.dumpToken("CloseParenthesisTkn", n.CloseParenthesisTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch) {
	v.print(0, "&ast.ExprStaticPropertyFetch{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Class", n.Class)
	v.dumpToken("DoubleColonTkn", n.DoubleColonTkn)
	v.dumpVertex("Property", n.Property)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprTernary(n *ast.ExprTernary) {
	v.print(0, "&ast.ExprTernary{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Condition", n.Condition)
	v.dumpToken("QuestionTkn", n.QuestionTkn)
	v.dumpVertex("IfTrue", n.IfTrue)
	v.dumpToken("ColonTkn", n.ColonTkn)
	v.dumpVertex("IfFalse", n.IfFalse)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprUnaryMinus(n *ast.ExprUnaryMinus) {
	v.print(0, "&ast.ExprUnaryMinus{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("MinusTkn", n.MinusTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprUnaryPlus(n *ast.ExprUnaryPlus) {
	v.print(0, "&ast.ExprUnaryPlus{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("PlusTkn", n.PlusTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprVariable(n *ast.ExprVariable) {
	v.print(0, "&ast.ExprVariable{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("DollarTkn", n.DollarTkn)
	v.dumpToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.dumpVertex("VarName", n.VarName)
	v.dumpToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprYield(n *ast.ExprYield) {
	v.print(0, "&ast.ExprYield{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("YieldTkn", n.YieldTkn)
	v.dumpVertex("Key", n.Key)
	v.dumpToken("DoubleArrowTkn", n.DoubleArrowTkn)
	v.dumpVertex("Value", n.Value)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprYieldFrom(n *ast.ExprYieldFrom) {
	v.print(0, "&ast.ExprYieldFrom{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("YieldFromTkn", n.YieldFromTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssign(n *ast.ExprAssign) {
	v.print(0, "&ast.ExprAssign{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignReference(n *ast.ExprAssignReference) {
	v.print(0, "&ast.ExprAssignReference{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpToken("AmpersandTkn", n.AmpersandTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignBitwiseAnd(n *ast.ExprAssignBitwiseAnd) {
	v.print(0, "&ast.ExprAssignBitwiseAnd{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignBitwiseOr(n *ast.ExprAssignBitwiseOr) {
	v.print(0, "&ast.ExprAssignBitwiseOr{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignBitwiseXor(n *ast.ExprAssignBitwiseXor) {
	v.print(0, "&ast.ExprAssignBitwiseXor{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignCoalesce(n *ast.ExprAssignCoalesce) {
	v.print(0, "&ast.ExprAssignCoalesce{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignConcat(n *ast.ExprAssignConcat) {
	v.print(0, "&ast.ExprAssignConcat{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignDiv(n *ast.ExprAssignDiv) {
	v.print(0, "&ast.ExprAssignDiv{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignMinus(n *ast.ExprAssignMinus) {
	v.print(0, "&ast.ExprAssignMinus{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignMod(n *ast.ExprAssignMod) {
	v.print(0, "&ast.ExprAssignMod{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignMul(n *ast.ExprAssignMul) {
	v.print(0, "&ast.ExprAssignMul{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignPlus(n *ast.ExprAssignPlus) {
	v.print(0, "&ast.ExprAssignPlus{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignPow(n *ast.ExprAssignPow) {
	v.print(0, "&ast.ExprAssignPow{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignShiftLeft(n *ast.ExprAssignShiftLeft) {
	v.print(0, "&ast.ExprAssignShiftLeft{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprAssignShiftRight(n *ast.ExprAssignShiftRight) {
	v.print(0, "&ast.ExprAssignShiftRight{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Var", n.Var)
	v.dumpToken("EqualTkn", n.EqualTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryBitwiseAnd(n *ast.ExprBinaryBitwiseAnd) {
	v.print(0, "&ast.ExprBinaryBitwiseAnd{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryBitwiseOr(n *ast.ExprBinaryBitwiseOr) {
	v.print(0, "&ast.ExprBinaryBitwiseOr{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryBitwiseXor(n *ast.ExprBinaryBitwiseXor) {
	v.print(0, "&ast.ExprBinaryBitwiseXor{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryBooleanAnd(n *ast.ExprBinaryBooleanAnd) {
	v.print(0, "&ast.ExprBinaryBooleanAnd{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryBooleanOr(n *ast.ExprBinaryBooleanOr) {
	v.print(0, "&ast.ExprBinaryBooleanOr{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryCoalesce(n *ast.ExprBinaryCoalesce) {
	v.print(0, "&ast.ExprBinaryCoalesce{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryConcat(n *ast.ExprBinaryConcat) {
	v.print(0, "&ast.ExprBinaryConcat{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryDiv(n *ast.ExprBinaryDiv) {
	v.print(0, "&ast.ExprBinaryDiv{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryEqual(n *ast.ExprBinaryEqual) {
	v.print(0, "&ast.ExprBinaryEqual{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryGreater(n *ast.ExprBinaryGreater) {
	v.print(0, "&ast.ExprBinaryGreater{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryGreaterOrEqual(n *ast.ExprBinaryGreaterOrEqual) {
	v.print(0, "&ast.ExprBinaryGreaterOrEqual{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryIdentical(n *ast.ExprBinaryIdentical) {
	v.print(0, "&ast.ExprBinaryIdentical{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryLogicalAnd(n *ast.ExprBinaryLogicalAnd) {
	v.print(0, "&ast.ExprBinaryLogicalAnd{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryLogicalOr(n *ast.ExprBinaryLogicalOr) {
	v.print(0, "&ast.ExprBinaryLogicalOr{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryLogicalXor(n *ast.ExprBinaryLogicalXor) {
	v.print(0, "&ast.ExprBinaryLogicalXor{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryMinus(n *ast.ExprBinaryMinus) {
	v.print(0, "&ast.ExprBinaryMinus{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryMod(n *ast.ExprBinaryMod) {
	v.print(0, "&ast.ExprBinaryMod{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryMul(n *ast.ExprBinaryMul) {
	v.print(0, "&ast.ExprBinaryMul{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryNotEqual(n *ast.ExprBinaryNotEqual) {
	v.print(0, "&ast.ExprBinaryNotEqual{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryNotIdentical(n *ast.ExprBinaryNotIdentical) {
	v.print(0, "&ast.ExprBinaryNotIdentical{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryPlus(n *ast.ExprBinaryPlus) {
	v.print(0, "&ast.ExprBinaryPlus{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryPow(n *ast.ExprBinaryPow) {
	v.print(0, "&ast.ExprBinaryPow{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryShiftLeft(n *ast.ExprBinaryShiftLeft) {
	v.print(0, "&ast.ExprBinaryShiftLeft{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinaryShiftRight(n *ast.ExprBinaryShiftRight) {
	v.print(0, "&ast.ExprBinaryShiftRight{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinarySmaller(n *ast.ExprBinarySmaller) {
	v.print(0, "&ast.ExprBinarySmaller{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinarySmallerOrEqual(n *ast.ExprBinarySmallerOrEqual) {
	v.print(0, "&ast.ExprBinarySmallerOrEqual{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprBinarySpaceship(n *ast.ExprBinarySpaceship) {
	v.print(0, "&ast.ExprBinarySpaceship{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertex("Left", n.Left)
	v.dumpToken("OpTkn", n.OpTkn)
	v.dumpVertex("Right", n.Right)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprCastArray(n *ast.ExprCastArray) {
	v.print(0, "&ast.ExprCastArray{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CastTkn", n.CastTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprCastBool(n *ast.ExprCastBool) {
	v.print(0, "&ast.ExprCastBool{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CastTkn", n.CastTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprCastDouble(n *ast.ExprCastDouble) {
	v.print(0, "&ast.ExprCastDouble{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CastTkn", n.CastTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprCastInt(n *ast.ExprCastInt) {
	v.print(0, "&ast.ExprCastInt{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CastTkn", n.CastTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprCastObject(n *ast.ExprCastObject) {
	v.print(0, "&ast.ExprCastObject{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CastTkn", n.CastTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprCastString(n *ast.ExprCastString) {
	v.print(0, "&ast.ExprCastString{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CastTkn", n.CastTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ExprCastUnset(n *ast.ExprCastUnset) {
	v.print(0, "&ast.ExprCastUnset{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("CastTkn", n.CastTkn)
	v.dumpVertex("Expr", n.Expr)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ScalarDnumber(n *ast.ScalarDnumber) {
	v.print(0, "&ast.ScalarDnumber{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("NumberTkn", n.NumberTkn)
	v.dumpValue("Value", n.Value)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ScalarEncapsed(n *ast.ScalarEncapsed) {
	v.print(0, "&ast.ScalarEncapsed{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("OpenQuoteTkn", n.OpenQuoteTkn)
	v.dumpVertexList("Parts", n.Parts)
	v.dumpToken("CloseQuoteTkn", n.CloseQuoteTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ScalarEncapsedStringPart(n *ast.ScalarEncapsedStringPart) {
	v.print(0, "&ast.ScalarEncapsedStringPart{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("EncapsedStrTkn", n.EncapsedStrTkn)
	v.dumpValue("Value", n.Value)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ScalarHeredoc(n *ast.ScalarHeredoc) {
	v.print(0, "&ast.ScalarHeredoc{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("OpenHeredocTkn", n.OpenHeredocTkn)
	v.dumpVertexList("Parts", n.Parts)
	v.dumpToken("CloseHeredocTkn", n.CloseHeredocTkn)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ScalarLnumber(n *ast.ScalarLnumber) {
	v.print(0, "&ast.ScalarLnumber{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("NumberTkn", n.NumberTkn)
	v.dumpValue("Value", n.Value)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ScalarMagicConstant(n *ast.ScalarMagicConstant) {
	v.print(0, "&ast.ScalarMagicConstant{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("MagicConstTkn", n.MagicConstTkn)
	v.dumpValue("Value", n.Value)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ScalarString(n *ast.ScalarString) {
	v.print(0, "&ast.ScalarString{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("MinusTkn", n.MinusTkn)
	v.dumpToken("StringTkn", n.StringTkn)
	v.dumpValue("Value", n.Value)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) NameName(n *ast.NameName) {
	v.print(0, "&ast.NameName{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpVertexList("Parts", n.Parts)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) NameFullyQualified(n *ast.NameFullyQualified) {
	v.print(0, "&ast.NameFullyQualified{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("NsSeparatorTkn", n.NsSeparatorTkn)
	v.dumpVertexList("Parts", n.Parts)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) NameRelative(n *ast.NameRelative) {
	v.print(0, "&ast.NameRelative{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("NsTkn", n.NsTkn)
	v.dumpToken("NsSeparatorTkn", n.NsSeparatorTkn)
	v.dumpVertexList("Parts", n.Parts)
	v.dumpTokenList("SeparatorTkns", n.SeparatorTkns)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) NameNamePart(n *ast.NameNamePart) {
	v.print(0, "&ast.NameNamePart{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("StringTkn", n.StringTkn)
	v.dumpValue("Value", n.Value)

	v.indent--
	v.print(v.indent, "},\n")
}

func (v *Dumper) ParserBrackets(n *ast.ParserBrackets) {
	v.print(0, "&ast.ParserBrackets{\n")
	v.indent++

	v.dumpPosition(n.Position)
	v.dumpToken("OpenBracketTkn", n.OpenBracketTkn)
	v.dumpVertex("Child", n.Child)
	v.dumpToken("CloseBracketTkn", n.CloseBracketTkn)

	v.indent--
	v.print(v.indent, "},\n")
}
