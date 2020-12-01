package visitor

import (
	"fmt"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/token"
	"io"
	"sort"
	"strconv"
	"strings"
)

type meta struct {
	singleNode bool
}

type Dump struct {
	writer io.Writer
	indent int
	depth  int
	stack  []meta
}

func NewDump(writer io.Writer) *Dump {
	return &Dump{writer: writer}
}

func (v *Dump) print(str string) {
	_, err := io.WriteString(v.writer, str)
	if err != nil {
		panic(err)
	}
}

func (v *Dump) printIndent(indentDepth int) {
	if indentDepth < 0 {
		indentDepth = 0
	}

	v.print(strings.Repeat("\t", indentDepth))
}

func (v *Dump) printIndentIfNotSingle(indentDepth int) {
	if indentDepth < 0 {
		indentDepth = 0
	}

	if !v.stack[v.depth-1].singleNode {
		v.print(strings.Repeat("\t", indentDepth))
	}
}

func (v *Dump) Enter(key string, singleNode bool) {
	if len(v.stack) < v.depth+1 {
		v.stack = append(v.stack, meta{})
	}

	v.stack[v.depth].singleNode = singleNode

	v.printIndent(v.indent)
	v.print(key)
	v.print(": ")

	if !singleNode {
		v.print("[]ast.Vertex{\n")
		v.indent++
	}
}

func (v *Dump) Leave(_ string, singleNode bool) {
	if !singleNode {
		v.indent--
		v.printIndent(v.indent)
		v.print("},\n")
	}
}

func (v *Dump) EnterNode(n ast.Vertex) bool {
	v.indent++
	v.depth++

	if len(v.stack) < v.depth {
		v.stack = append(v.stack, meta{})
	}

	n.Accept(v)

	return true
}

func (v *Dump) LeaveNode(_ ast.Vertex) {
	v.indent--
	v.depth--
	v.printIndent(v.indent)
	v.print("}")
	if v.depth != 0 {
		v.print(",")
	}
	v.print("\n")
}

func (v *Dump) printNode(n *ast.Node) {
	if n.Position == nil && n.Tokens == nil {
		return
	}

	v.printIndent(v.indent)
	v.print("Node: ast.Node{\n")

	if n.Tokens != nil {
		v.printIndent(v.indent + 1)
		v.print("Tokens: token.Collection{\n")

		keys := make([]int, 0, len(n.Tokens))
		for k := range n.Tokens {
			keys = append(keys, int(k))
		}
		sort.Ints(keys)

		for _, k := range keys {
			key := token.Position(k)

			v.printIndent(v.indent + 2)
			v.print("token." + key.String() + ": []*token.Token{\n")

			for _, tkn := range n.Tokens[key] {
				v.printIndent(v.indent + 3)
				v.print("{\n")

				v.printIndent(v.indent + 4)
				v.print("ID:    token." + tkn.ID.String() + ",\n")

				v.printIndent(v.indent + 4)
				v.print("Value: []byte(" + strconv.Quote(string(tkn.Value)) + "),\n")

				v.printIndent(v.indent + 3)
				v.print("},\n")
			}

			v.printIndent(v.indent + 2)
			v.print("},\n")
		}

		v.printIndent(v.indent + 1)
		v.print("},\n")
	}

	if n.Position != nil {
		v.printIndent(v.indent + 1)
		v.print("Position: &position.Position{\n")

		v.printIndent(v.indent + 2)
		v.print("StartLine: " + strconv.Itoa(n.Position.StartLine) + ",\n")

		v.printIndent(v.indent + 2)
		v.print("EndLine:   " + strconv.Itoa(n.Position.EndLine) + ",\n")

		v.printIndent(v.indent + 2)
		v.print("StartPos:  " + strconv.Itoa(n.Position.StartPos) + ",\n")

		v.printIndent(v.indent + 2)
		v.print("EndPos:    " + strconv.Itoa(n.Position.EndPos) + ",\n")

		v.printIndent(v.indent + 1)
		v.print("},\n")
	}

	v.printIndent(v.indent)
	v.print("},\n")
}

func (v *Dump) printToken(key string, t *token.Token) {
	if t == nil {
		return
	}

	v.printIndent(v.indent)
	v.print(key)
	v.print(": &token.Token{\n")

	v.printIndent(v.indent + 1)
	v.print("ID:      token." + t.ID.String() + ",\n")

	v.printIndent(v.indent + 1)
	v.print("Value:   []byte(" + strconv.Quote(string(t.Value)) + "),\n")

	v.printIndent(v.indent + 1)
	v.print("Skipped: []byte(" + strconv.Quote(string(t.Skipped)) + "),\n")

	v.printIndent(v.indent)
	v.print("},\n")
}

func (v *Dump) Root(n *ast.Root) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Root{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) Nullable(n *ast.Nullable) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Nullable{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) Reference(n *ast.Reference) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Reference{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) Variadic(n *ast.Variadic) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Variadic{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) Parameter(n *ast.Parameter) {
	v.printIndent(v.indent - 1)
	v.print("&ast.Parameter{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) Identifier(n *ast.Identifier) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Identifier{\n")
	v.printNode(n.GetNode())

	v.printIndent(v.indent)
	v.print(fmt.Sprintf("Value: []byte(%q),\n", n.Value))
}

func (v *Dump) ArgumentList(n *ast.ArgumentList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ArgumentList{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) Argument(n *ast.Argument) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Argument{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtBreak(n *ast.StmtBreak) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtBreak{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtCase(n *ast.StmtCase) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtCase{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtCatch(n *ast.StmtCatch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtCatch{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtClass(n *ast.StmtClass) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClass{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtClassConstList(n *ast.StmtClassConstList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClassConstList{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtClassExtends(n *ast.StmtClassExtends) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClassExtends{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtClassImplements(n *ast.StmtClassImplements) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClassImplements{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtClassMethod(n *ast.StmtClassMethod) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClassMethod{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtConstList(n *ast.StmtConstList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtConstList{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtConstant(n *ast.StmtConstant) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtConstant{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtContinue(n *ast.StmtContinue) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtContinue{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtDeclare(n *ast.StmtDeclare) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtDeclare{\n")
	v.printNode(n.GetNode())

	if n.Alt {
		v.printIndent(v.indent)
		v.print("Alt: true,\n")
	}
}

func (v *Dump) StmtDefault(n *ast.StmtDefault) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtDefault{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtDo(n *ast.StmtDo) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtDo{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtEcho(n *ast.StmtEcho) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtEcho{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtElse(n *ast.StmtElse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtElse{\n")
	v.printNode(n.GetNode())

	if n.Alt {
		v.printIndent(v.indent)
		v.print("Alt: true,\n")
	}
}

func (v *Dump) StmtElseIf(n *ast.StmtElseIf) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtElseIf{\n")
	v.printNode(n.GetNode())

	if n.Alt {
		v.printIndent(v.indent)
		v.print("Alt: true,\n")
	}
}

func (v *Dump) StmtExpression(n *ast.StmtExpression) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtExpression{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtFinally(n *ast.StmtFinally) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtFinally{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtFor(n *ast.StmtFor) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtFor{\n")
	v.printNode(n.GetNode())

	if n.Alt {
		v.printIndent(v.indent)
		v.print("Alt: true,\n")
	}
}

func (v *Dump) StmtForeach(n *ast.StmtForeach) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtForeach{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtFunction(n *ast.StmtFunction) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtFunction{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtGlobal(n *ast.StmtGlobal) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtGlobal{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtGoto(n *ast.StmtGoto) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtGoto{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtHaltCompiler(n *ast.StmtHaltCompiler) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtHaltCompiler{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtIf(n *ast.StmtIf) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtIf{\n")
	v.printNode(n.GetNode())

	if n.Alt {
		v.printIndent(v.indent)
		v.print("Alt: true,\n")
	}
}

func (v *Dump) StmtInlineHtml(n *ast.StmtInlineHtml) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtInlineHtml{\n")
	v.printNode(n.GetNode())

	v.printIndent(v.indent)
	v.print(fmt.Sprintf("Value: []byte(%q),\n", n.Value))
}

func (v *Dump) StmtInterface(n *ast.StmtInterface) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtInterface{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtInterfaceExtends(n *ast.StmtInterfaceExtends) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtInterfaceExtends{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtLabel(n *ast.StmtLabel) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtLabel{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtNamespace(n *ast.StmtNamespace) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtNamespace{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtNop(n *ast.StmtNop) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtNop{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtProperty(n *ast.StmtProperty) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtProperty{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtPropertyList(n *ast.StmtPropertyList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtPropertyList{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtReturn(n *ast.StmtReturn) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtReturn{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtStatic(n *ast.StmtStatic) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtStatic{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtStaticVar(n *ast.StmtStaticVar) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtStaticVar{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtStmtList(n *ast.StmtStmtList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtStmtList{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtSwitch(n *ast.StmtSwitch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtSwitch{\n")
	v.printNode(n.GetNode())

	if n.Alt {
		v.printIndent(v.indent)
		v.print("Alt: true,\n")
	}
}

func (v *Dump) StmtThrow(n *ast.StmtThrow) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtThrow{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtTrait(n *ast.StmtTrait) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTrait{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtTraitAdaptationList(n *ast.StmtTraitAdaptationList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitAdaptationList{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtTraitMethodRef(n *ast.StmtTraitMethodRef) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitMethodRef{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtTraitUse(n *ast.StmtTraitUse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitUse{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtTraitUseAlias(n *ast.StmtTraitUseAlias) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitUseAlias{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtTraitUsePrecedence(n *ast.StmtTraitUsePrecedence) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitUsePrecedence{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtTry(n *ast.StmtTry) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTry{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtUnset(n *ast.StmtUnset) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtUnset{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) StmtUse(n *ast.StmtUse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtUse{\n")
	v.printNode(n.GetNode())
	v.printToken("UseTkn", n.UseTkn)
	v.printToken("SemiColonTkn", n.SemiColonTkn)

}

func (v *Dump) StmtGroupUse(n *ast.StmtGroupUse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtGroupUse{\n")
	v.printNode(n.GetNode())
	v.printToken("UseTkn", n.UseTkn)
	v.printToken("LeadingNsSeparatorTkn", n.LeadingNsSeparatorTkn)
	v.printToken("NsSeparatorTkn", n.NsSeparatorTkn)
	v.printToken("OpenCurlyBracketTkn", n.OpenCurlyBracketTkn)
	v.printToken("CloseCurlyBracketTkn", n.CloseCurlyBracketTkn)
	v.printToken("SemiColonTkn", n.SemiColonTkn)
}

func (v *Dump) StmtUseDeclaration(n *ast.StmtUseDeclaration) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtUseDeclaration{\n")
	v.printNode(n.GetNode())
	v.printToken("NsSeparatorTkn", n.NsSeparatorTkn)
	v.printToken("AsTkn", n.AsTkn)
	v.printToken("CommaTkn", n.CommaTkn)
}

func (v *Dump) StmtWhile(n *ast.StmtWhile) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtWhile{\n")
	v.printNode(n.GetNode())

	if n.Alt {
		v.printIndent(v.indent)
		v.print("Alt: true,\n")
	}
}

func (v *Dump) ExprArray(n *ast.ExprArray) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprArray{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprArrayDimFetch(n *ast.ExprArrayDimFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprArrayDimFetch{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprArrayItem(n *ast.ExprArrayItem) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprArrayItem{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprArrowFunction(n *ast.ExprArrowFunction) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprArrowFunction{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBitwiseNot(n *ast.ExprBitwiseNot) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBitwiseNot{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBooleanNot(n *ast.ExprBooleanNot) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBooleanNot{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprClassConstFetch{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprClone(n *ast.ExprClone) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprClone{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprClosure(n *ast.ExprClosure) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprClosure{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprClosureUse(n *ast.ExprClosureUse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprClosureUse{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprConstFetch(n *ast.ExprConstFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprConstFetch{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprEmpty(n *ast.ExprEmpty) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprEmpty{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprErrorSuppress(n *ast.ExprErrorSuppress) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprErrorSuppress{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprEval(n *ast.ExprEval) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprEval{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprExit(n *ast.ExprExit) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprExit{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprFunctionCall(n *ast.ExprFunctionCall) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprFunctionCall{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprInclude(n *ast.ExprInclude) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprInclude{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprIncludeOnce(n *ast.ExprIncludeOnce) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprIncludeOnce{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprInstanceOf(n *ast.ExprInstanceOf) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprInstanceOf{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprIsset(n *ast.ExprIsset) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprIsset{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprList(n *ast.ExprList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprList{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprMethodCall(n *ast.ExprMethodCall) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprMethodCall{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprNew(n *ast.ExprNew) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprNew{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprPostDec(n *ast.ExprPostDec) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPostDec{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprPostInc(n *ast.ExprPostInc) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPostInc{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprPreDec(n *ast.ExprPreDec) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPreDec{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprPreInc(n *ast.ExprPreInc) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPreInc{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprPrint(n *ast.ExprPrint) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPrint{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprPropertyFetch(n *ast.ExprPropertyFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPropertyFetch{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprReference(n *ast.ExprReference) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprReference{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprRequire(n *ast.ExprRequire) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprRequire{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprRequireOnce(n *ast.ExprRequireOnce) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprRequireOnce{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprShellExec(n *ast.ExprShellExec) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprShellExec{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprStaticCall(n *ast.ExprStaticCall) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprStaticCall{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprStaticPropertyFetch{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprTernary(n *ast.ExprTernary) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprTernary{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprUnaryMinus(n *ast.ExprUnaryMinus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprUnaryMinus{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprUnaryPlus(n *ast.ExprUnaryPlus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprUnaryPlus{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprVariable(n *ast.ExprVariable) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprVariable{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprYield(n *ast.ExprYield) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprYield{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprYieldFrom(n *ast.ExprYieldFrom) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprYieldFrom{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssign(n *ast.ExprAssign) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssign{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignReference(n *ast.ExprAssignReference) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignReference{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignBitwiseAnd(n *ast.ExprAssignBitwiseAnd) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignBitwiseAnd{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignBitwiseOr(n *ast.ExprAssignBitwiseOr) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignBitwiseOr{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignBitwiseXor(n *ast.ExprAssignBitwiseXor) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignBitwiseXor{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignCoalesce(n *ast.ExprAssignCoalesce) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignCoalesce{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignConcat(n *ast.ExprAssignConcat) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignConcat{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignDiv(n *ast.ExprAssignDiv) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignDiv{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignMinus(n *ast.ExprAssignMinus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignMinus{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignMod(n *ast.ExprAssignMod) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignMod{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignMul(n *ast.ExprAssignMul) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignMul{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignPlus(n *ast.ExprAssignPlus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignPlus{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignPow(n *ast.ExprAssignPow) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignPow{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignShiftLeft(n *ast.ExprAssignShiftLeft) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignShiftLeft{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprAssignShiftRight(n *ast.ExprAssignShiftRight) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignShiftRight{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryBitwiseAnd(n *ast.ExprBinaryBitwiseAnd) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBitwiseAnd{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryBitwiseOr(n *ast.ExprBinaryBitwiseOr) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBitwiseOr{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryBitwiseXor(n *ast.ExprBinaryBitwiseXor) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBitwiseXor{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryBooleanAnd(n *ast.ExprBinaryBooleanAnd) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBooleanAnd{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryBooleanOr(n *ast.ExprBinaryBooleanOr) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBooleanOr{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryCoalesce(n *ast.ExprBinaryCoalesce) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryCoalesce{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryConcat(n *ast.ExprBinaryConcat) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryConcat{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryDiv(n *ast.ExprBinaryDiv) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryDiv{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryEqual(n *ast.ExprBinaryEqual) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryEqual{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryGreater(n *ast.ExprBinaryGreater) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryGreater{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryGreaterOrEqual(n *ast.ExprBinaryGreaterOrEqual) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryGreaterOrEqual{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryIdentical(n *ast.ExprBinaryIdentical) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryIdentical{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryLogicalAnd(n *ast.ExprBinaryLogicalAnd) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryLogicalAnd{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryLogicalOr(n *ast.ExprBinaryLogicalOr) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryLogicalOr{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryLogicalXor(n *ast.ExprBinaryLogicalXor) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryLogicalXor{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryMinus(n *ast.ExprBinaryMinus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryMinus{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryMod(n *ast.ExprBinaryMod) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryMod{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryMul(n *ast.ExprBinaryMul) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryMul{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryNotEqual(n *ast.ExprBinaryNotEqual) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryNotEqual{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryNotIdentical(n *ast.ExprBinaryNotIdentical) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryNotIdentical{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryPlus(n *ast.ExprBinaryPlus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryPlus{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryPow(n *ast.ExprBinaryPow) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryPow{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryShiftLeft(n *ast.ExprBinaryShiftLeft) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryShiftLeft{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinaryShiftRight(n *ast.ExprBinaryShiftRight) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryShiftRight{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinarySmaller(n *ast.ExprBinarySmaller) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinarySmaller{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinarySmallerOrEqual(n *ast.ExprBinarySmallerOrEqual) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinarySmallerOrEqual{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprBinarySpaceship(n *ast.ExprBinarySpaceship) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinarySpaceship{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprCastArray(n *ast.ExprCastArray) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastArray{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprCastBool(n *ast.ExprCastBool) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastBool{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprCastDouble(n *ast.ExprCastDouble) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastDouble{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprCastInt(n *ast.ExprCastInt) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastInt{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprCastObject(n *ast.ExprCastObject) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastObject{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprCastString(n *ast.ExprCastString) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastString{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ExprCastUnset(n *ast.ExprCastUnset) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastUnset{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ScalarDnumber(n *ast.ScalarDnumber) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarDnumber{\n")
	v.printNode(n.GetNode())

	v.printIndent(v.indent)
	v.print(fmt.Sprintf("Value: []byte(%q),\n", n.Value))
}

func (v *Dump) ScalarEncapsed(n *ast.ScalarEncapsed) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarEncapsed{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ScalarEncapsedStringPart(n *ast.ScalarEncapsedStringPart) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarEncapsedStringPart{\n")
	v.printNode(n.GetNode())

	v.printIndent(v.indent)
	v.print(fmt.Sprintf("Value: []byte(%q),\n", n.Value))
}

func (v *Dump) ScalarHeredoc(n *ast.ScalarHeredoc) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarHeredoc{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ScalarLnumber(n *ast.ScalarLnumber) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarLnumber{\n")
	v.printNode(n.GetNode())

	v.printIndent(v.indent)
	v.print(fmt.Sprintf("Value: []byte(%q),\n", n.Value))
}

func (v *Dump) ScalarMagicConstant(n *ast.ScalarMagicConstant) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarMagicConstant{\n")
	v.printNode(n.GetNode())

	v.printIndent(v.indent)
	v.print(fmt.Sprintf("Value: []byte(%q),\n", n.Value))
}

func (v *Dump) ScalarString(n *ast.ScalarString) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarString{\n")
	v.printNode(n.GetNode())

	v.printIndent(v.indent)
	v.print(fmt.Sprintf("Value: []byte(%q),\n", n.Value))
}

func (v *Dump) NameName(n *ast.NameName) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.NameName{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) NameFullyQualified(n *ast.NameFullyQualified) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.NameFullyQualified{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) NameRelative(n *ast.NameRelative) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.NameRelative{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) NameNamePart(n *ast.NameNamePart) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.NameNamePart{\n")
	v.printNode(n.GetNode())

	v.printIndent(v.indent)
	v.print(fmt.Sprintf("Value: []byte(%q),\n", n.Value))
}

func (v *Dump) ParserBrackets(n *ast.ParserBrackets) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ParserBrackets{\n")
	v.printNode(n.GetNode())
}

func (v *Dump) ParserSeparatedList(n *ast.ParserSeparatedList) {
	// do nothing
}
