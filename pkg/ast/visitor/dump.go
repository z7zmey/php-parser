package visitor

import (
	"fmt"
	"github.com/z7zmey/php-parser/pkg/ast"
	"io"
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
	v.print(strings.Repeat("\t", indentDepth))
}

func (v *Dump) printIndentIfNotSingle(indentDepth int) {
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

func (v *Dump) Root(_ *ast.Root) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Root{\n")
}

func (v *Dump) Nullable(_ *ast.Nullable) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Nullable{\n")
}

func (v *Dump) Parameter(n *ast.Parameter) {
	v.printIndent(v.indent - 1)
	v.print("&ast.Parameter{\n")

	if n.ByRef {
		v.printIndent(v.indent)
		v.print("ByRef: true,\n")
	}

	if n.Variadic {
		v.printIndent(v.indent)
		v.print("Variadic: true,\n")
	}
}

func (v *Dump) Identifier(n *ast.Identifier) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Identifier{\n")

	v.printIndentIfNotSingle(v.indent)
	v.print(fmt.Sprintf("Value: %q,\n", n.Value))
}

func (v *Dump) ArgumentList(_ *ast.ArgumentList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ArgumentList{\n")
}

func (v *Dump) Argument(n *ast.Argument) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.Argument{\n")

	if n.Variadic {
		v.printIndent(v.indent)
		v.print("Variadic: true,\n")
	}

	if n.IsReference {
		v.printIndent(v.indent)
		v.print("IsReference: true,\n")
	}
}

func (v *Dump) StmtAltElse(_ *ast.StmtAltElse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtAltElse{\n")
}

func (v *Dump) StmtAltElseIf(_ *ast.StmtAltElseIf) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtAltElseIf{\n")
}

func (v *Dump) StmtAltFor(_ *ast.StmtAltFor) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtAltFor{\n")
}

func (v *Dump) StmtAltForeach(_ *ast.StmtAltForeach) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtAltForeach{\n")
}

func (v *Dump) StmtAltIf(_ *ast.StmtAltIf) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtAltIf{\n")
}

func (v *Dump) StmtAltSwitch(_ *ast.StmtAltSwitch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtAltSwitch{\n")
}

func (v *Dump) StmtAltWhile(_ *ast.StmtAltWhile) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtAltWhile{\n")
}

func (v *Dump) StmtBreak(_ *ast.StmtBreak) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtBreak{\n")
}

func (v *Dump) StmtCase(_ *ast.StmtCase) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtCase{\n")
}

func (v *Dump) StmtCaseList(_ *ast.StmtCaseList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtCaseList{\n")
}

func (v *Dump) StmtCatch(_ *ast.StmtCatch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtCatch{\n")
}

func (v *Dump) StmtClass(_ *ast.StmtClass) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClass{\n")
}

func (v *Dump) StmtClassConstList(_ *ast.StmtClassConstList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClassConstList{\n")
}

func (v *Dump) StmtClassExtends(_ *ast.StmtClassExtends) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClassExtends{\n")
}

func (v *Dump) StmtClassImplements(_ *ast.StmtClassImplements) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClassImplements{\n")
}

func (v *Dump) StmtClassMethod(n *ast.StmtClassMethod) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtClassMethod{\n")

	if n.ReturnsRef {
		v.printIndent(v.indent)
		v.print("ReturnsRef: true,\n")
	}
}

func (v *Dump) StmtConstList(_ *ast.StmtConstList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtConstList{\n")
}

func (v *Dump) StmtConstant(_ *ast.StmtConstant) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtConstant{\n")
}

func (v *Dump) StmtContinue(_ *ast.StmtContinue) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtContinue{\n")
}

func (v *Dump) StmtDeclare(n *ast.StmtDeclare) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtDeclare{\n")

	if n.Alt {
		v.printIndent(v.indent)
		v.print("Alt: true,\n")
	}
}

func (v *Dump) StmtDefault(_ *ast.StmtDefault) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtDefault{\n")
}

func (v *Dump) StmtDo(_ *ast.StmtDo) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtDo{\n")
}

func (v *Dump) StmtEcho(_ *ast.StmtEcho) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtEcho{\n")
}

func (v *Dump) StmtElse(_ *ast.StmtElse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtElse{\n")
}

func (v *Dump) StmtElseIf(_ *ast.StmtElseIf) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtElseIf{\n")
}

func (v *Dump) StmtExpression(_ *ast.StmtExpression) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtExpression{\n")
}

func (v *Dump) StmtFinally(_ *ast.StmtFinally) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtFinally{\n")
}

func (v *Dump) StmtFor(_ *ast.StmtFor) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtFor{\n")
}

func (v *Dump) StmtForeach(_ *ast.StmtForeach) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtForeach{\n")
}

func (v *Dump) StmtFunction(n *ast.StmtFunction) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtFunction{\n")

	if n.ReturnsRef {
		v.printIndent(v.indent)
		v.print("ReturnsRef: true,\n")
	}
}

func (v *Dump) StmtGlobal(_ *ast.StmtGlobal) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtGlobal{\n")
}

func (v *Dump) StmtGoto(_ *ast.StmtGoto) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtGoto{\n")
}

func (v *Dump) StmtGroupUse(_ *ast.StmtGroupUse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtGroupUse{\n")
}

func (v *Dump) StmtHaltCompiler(_ *ast.StmtHaltCompiler) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtHaltCompiler{\n")
}

func (v *Dump) StmtIf(_ *ast.StmtIf) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtIf{\n")
}

func (v *Dump) StmtInlineHtml(n *ast.StmtInlineHtml) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtInlineHtml{\n")

	v.printIndentIfNotSingle(v.indent)
	v.print(fmt.Sprintf("Value: %q,\n", n.Value))
}

func (v *Dump) StmtInterface(_ *ast.StmtInterface) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtInterface{\n")
}

func (v *Dump) StmtInterfaceExtends(_ *ast.StmtInterfaceExtends) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtInterfaceExtends{\n")
}

func (v *Dump) StmtLabel(_ *ast.StmtLabel) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtLabel{\n")
}

func (v *Dump) StmtNamespace(_ *ast.StmtNamespace) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtNamespace{\n")
}

func (v *Dump) StmtNop(_ *ast.StmtNop) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtNop{\n")
}

func (v *Dump) StmtProperty(_ *ast.StmtProperty) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtProperty{\n")
}

func (v *Dump) StmtPropertyList(_ *ast.StmtPropertyList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtPropertyList{\n")
}

func (v *Dump) StmtReturn(_ *ast.StmtReturn) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtReturn{\n")
}

func (v *Dump) StmtStatic(_ *ast.StmtStatic) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtStatic{\n")
}

func (v *Dump) StmtStaticVar(_ *ast.StmtStaticVar) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtStaticVar{\n")
}

func (v *Dump) StmtStmtList(_ *ast.StmtStmtList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtStmtList{\n")
}

func (v *Dump) StmtSwitch(_ *ast.StmtSwitch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtSwitch{\n")
}

func (v *Dump) StmtThrow(_ *ast.StmtThrow) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtThrow{\n")
}

func (v *Dump) StmtTrait(_ *ast.StmtTrait) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTrait{\n")
}

func (v *Dump) StmtTraitAdaptationList(_ *ast.StmtTraitAdaptationList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitAdaptationList{\n")
}

func (v *Dump) StmtTraitMethodRef(_ *ast.StmtTraitMethodRef) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitMethodRef{\n")
}

func (v *Dump) StmtTraitUse(_ *ast.StmtTraitUse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitUse{\n")
}

func (v *Dump) StmtTraitUseAlias(_ *ast.StmtTraitUseAlias) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitUseAlias{\n")
}

func (v *Dump) StmtTraitUsePrecedence(_ *ast.StmtTraitUsePrecedence) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTraitUsePrecedence{\n")
}

func (v *Dump) StmtTry(_ *ast.StmtTry) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtTry{\n")
}

func (v *Dump) StmtUnset(_ *ast.StmtUnset) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtUnset{\n")
}

func (v *Dump) StmtUse(_ *ast.StmtUse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtUse{\n")
}

func (v *Dump) StmtUseList(_ *ast.StmtUseList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtUseList{\n")
}

func (v *Dump) StmtWhile(_ *ast.StmtWhile) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.StmtWhile{\n")
}

func (v *Dump) ExprArray(_ *ast.ExprArray) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprArray{\n")
}

func (v *Dump) ExprArrayDimFetch(_ *ast.ExprArrayDimFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprArrayDimFetch{\n")
}

func (v *Dump) ExprArrayItem(n *ast.ExprArrayItem) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprArrayItem{\n")

	if n.Unpack {
		v.printIndent(v.indent)
		v.print("Unpack: true,\n")
	}
}

func (v *Dump) ExprArrowFunction(n *ast.ExprArrowFunction) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprArrowFunction{\n")

	if n.ReturnsRef {
		v.printIndent(v.indent)
		v.print("ReturnsRef: true,\n")
	}

	if n.Static {
		v.printIndent(v.indent)
		v.print("Static: true,\n")
	}
}

func (v *Dump) ExprBitwiseNot(_ *ast.ExprBitwiseNot) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBitwiseNot{\n")
}

func (v *Dump) ExprBooleanNot(_ *ast.ExprBooleanNot) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBooleanNot{\n")
}

func (v *Dump) ExprClassConstFetch(_ *ast.ExprClassConstFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprClassConstFetch{\n")
}

func (v *Dump) ExprClone(_ *ast.ExprClone) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprClone{\n")
}

func (v *Dump) ExprClosure(n *ast.ExprClosure) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprClosure{\n")

	if n.ReturnsRef {
		v.printIndent(v.indent)
		v.print("ReturnsRef: true,\n")
	}

	if n.Static {
		v.printIndent(v.indent)
		v.print("Static: true,\n")
	}
}

func (v *Dump) ExprClosureUse(_ *ast.ExprClosureUse) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprClosureUse{\n")
}

func (v *Dump) ExprConstFetch(_ *ast.ExprConstFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprConstFetch{\n")
}

func (v *Dump) ExprEmpty(_ *ast.ExprEmpty) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprEmpty{\n")
}

func (v *Dump) ExprErrorSuppress(_ *ast.ExprErrorSuppress) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprErrorSuppress{\n")
}

func (v *Dump) ExprEval(_ *ast.ExprEval) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprEval{\n")
}

func (v *Dump) ExprExit(n *ast.ExprExit) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprExit{\n")

	if n.Die {
		v.printIndent(v.indent)
		v.print("Die: true,\n")
	}
}

func (v *Dump) ExprFunctionCall(_ *ast.ExprFunctionCall) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprFunctionCall{\n")
}

func (v *Dump) ExprInclude(_ *ast.ExprInclude) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprInclude{\n")
}

func (v *Dump) ExprIncludeOnce(_ *ast.ExprIncludeOnce) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprIncludeOnce{\n")
}

func (v *Dump) ExprInstanceOf(_ *ast.ExprInstanceOf) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprInstanceOf{\n")
}

func (v *Dump) ExprIsset(_ *ast.ExprIsset) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprIsset{\n")
}

func (v *Dump) ExprList(_ *ast.ExprList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprList{\n")
}

func (v *Dump) ExprMethodCall(_ *ast.ExprMethodCall) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprMethodCall{\n")
}

func (v *Dump) ExprNew(_ *ast.ExprNew) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprNew{\n")
}

func (v *Dump) ExprPostDec(_ *ast.ExprPostDec) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPostDec{\n")
}

func (v *Dump) ExprPostInc(_ *ast.ExprPostInc) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPostInc{\n")
}

func (v *Dump) ExprPreDec(_ *ast.ExprPreDec) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPreDec{\n")
}

func (v *Dump) ExprPreInc(_ *ast.ExprPreInc) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPreInc{\n")
}

func (v *Dump) ExprPrint(_ *ast.ExprPrint) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPrint{\n")
}

func (v *Dump) ExprPropertyFetch(_ *ast.ExprPropertyFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprPropertyFetch{\n")
}

func (v *Dump) ExprReference(_ *ast.ExprReference) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprReference{\n")
}

func (v *Dump) ExprRequire(_ *ast.ExprRequire) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprRequire{\n")
}

func (v *Dump) ExprRequireOnce(_ *ast.ExprRequireOnce) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprRequireOnce{\n")
}

func (v *Dump) ExprShellExec(_ *ast.ExprShellExec) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprShellExec{\n")
}

func (v *Dump) ExprShortArray(_ *ast.ExprShortArray) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprShortArray{\n")
}

func (v *Dump) ExprShortList(_ *ast.ExprShortList) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprShortList{\n")
}

func (v *Dump) ExprStaticCall(_ *ast.ExprStaticCall) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprStaticCall{\n")
}

func (v *Dump) ExprStaticPropertyFetch(_ *ast.ExprStaticPropertyFetch) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprStaticPropertyFetch{\n")
}

func (v *Dump) ExprTernary(_ *ast.ExprTernary) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprTernary{\n")
}

func (v *Dump) ExprUnaryMinus(_ *ast.ExprUnaryMinus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprUnaryMinus{\n")
}

func (v *Dump) ExprUnaryPlus(_ *ast.ExprUnaryPlus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprUnaryPlus{\n")
}

func (v *Dump) ExprVariable(_ *ast.ExprVariable) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprVariable{\n")
}

func (v *Dump) ExprYield(_ *ast.ExprYield) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprYield{\n")
}

func (v *Dump) ExprYieldFrom(_ *ast.ExprYieldFrom) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprYieldFrom{\n")
}

func (v *Dump) ExprAssign(_ *ast.ExprAssign) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssign{\n")
}

func (v *Dump) ExprAssignReference(_ *ast.ExprAssignReference) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignReference{\n")
}

func (v *Dump) ExprAssignBitwiseAnd(_ *ast.ExprAssignBitwiseAnd) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignBitwiseAnd{\n")
}

func (v *Dump) ExprAssignBitwiseOr(_ *ast.ExprAssignBitwiseOr) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignBitwiseOr{\n")
}

func (v *Dump) ExprAssignBitwiseXor(_ *ast.ExprAssignBitwiseXor) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignBitwiseXor{\n")
}

func (v *Dump) ExprAssignCoalesce(_ *ast.ExprAssignCoalesce) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignCoalesce{\n")
}

func (v *Dump) ExprAssignConcat(_ *ast.ExprAssignConcat) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignConcat{\n")
}

func (v *Dump) ExprAssignDiv(_ *ast.ExprAssignDiv) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignDiv{\n")
}

func (v *Dump) ExprAssignMinus(_ *ast.ExprAssignMinus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignMinus{\n")
}

func (v *Dump) ExprAssignMod(_ *ast.ExprAssignMod) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignMod{\n")
}

func (v *Dump) ExprAssignMul(_ *ast.ExprAssignMul) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignMul{\n")
}

func (v *Dump) ExprAssignPlus(_ *ast.ExprAssignPlus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignPlus{\n")
}

func (v *Dump) ExprAssignPow(_ *ast.ExprAssignPow) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignPow{\n")
}

func (v *Dump) ExprAssignShiftLeft(_ *ast.ExprAssignShiftLeft) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignShiftLeft{\n")
}

func (v *Dump) ExprAssignShiftRight(_ *ast.ExprAssignShiftRight) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprAssignShiftRight{\n")
}

func (v *Dump) ExprBinaryBitwiseAnd(_ *ast.ExprBinaryBitwiseAnd) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBitwiseAnd{\n")
}

func (v *Dump) ExprBinaryBitwiseOr(_ *ast.ExprBinaryBitwiseOr) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBitwiseOr{\n")
}

func (v *Dump) ExprBinaryBitwiseXor(_ *ast.ExprBinaryBitwiseXor) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBitwiseXor{\n")
}

func (v *Dump) ExprBinaryBooleanAnd(_ *ast.ExprBinaryBooleanAnd) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBooleanAnd{\n")
}

func (v *Dump) ExprBinaryBooleanOr(_ *ast.ExprBinaryBooleanOr) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryBooleanOr{\n")
}

func (v *Dump) ExprBinaryCoalesce(_ *ast.ExprBinaryCoalesce) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryCoalesce{\n")
}

func (v *Dump) ExprBinaryConcat(_ *ast.ExprBinaryConcat) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryConcat{\n")
}

func (v *Dump) ExprBinaryDiv(_ *ast.ExprBinaryDiv) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryDiv{\n")
}

func (v *Dump) ExprBinaryEqual(_ *ast.ExprBinaryEqual) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryEqual{\n")
}

func (v *Dump) ExprBinaryGreater(_ *ast.ExprBinaryGreater) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryGreater{\n")
}

func (v *Dump) ExprBinaryGreaterOrEqual(_ *ast.ExprBinaryGreaterOrEqual) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryGreaterOrEqual{\n")
}

func (v *Dump) ExprBinaryIdentical(_ *ast.ExprBinaryIdentical) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryIdentical{\n")
}

func (v *Dump) ExprBinaryLogicalAnd(_ *ast.ExprBinaryLogicalAnd) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryLogicalAnd{\n")
}

func (v *Dump) ExprBinaryLogicalOr(_ *ast.ExprBinaryLogicalOr) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryLogicalOr{\n")
}

func (v *Dump) ExprBinaryLogicalXor(_ *ast.ExprBinaryLogicalXor) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryLogicalXor{\n")
}

func (v *Dump) ExprBinaryMinus(_ *ast.ExprBinaryMinus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryMinus{\n")
}

func (v *Dump) ExprBinaryMod(_ *ast.ExprBinaryMod) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryMod{\n")
}

func (v *Dump) ExprBinaryMul(_ *ast.ExprBinaryMul) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryMul{\n")
}

func (v *Dump) ExprBinaryNotEqual(_ *ast.ExprBinaryNotEqual) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryNotEqual{\n")
}

func (v *Dump) ExprBinaryNotIdentical(_ *ast.ExprBinaryNotIdentical) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryNotIdentical{\n")
}

func (v *Dump) ExprBinaryPlus(_ *ast.ExprBinaryPlus) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryPlus{\n")
}

func (v *Dump) ExprBinaryPow(_ *ast.ExprBinaryPow) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryPow{\n")
}

func (v *Dump) ExprBinaryShiftLeft(_ *ast.ExprBinaryShiftLeft) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryShiftLeft{\n")
}

func (v *Dump) ExprBinaryShiftRight(_ *ast.ExprBinaryShiftRight) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinaryShiftRight{\n")
}

func (v *Dump) ExprBinarySmaller(_ *ast.ExprBinarySmaller) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinarySmaller{\n")
}

func (v *Dump) ExprBinarySmallerOrEqual(_ *ast.ExprBinarySmallerOrEqual) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinarySmallerOrEqual{\n")
}

func (v *Dump) ExprBinarySpaceship(_ *ast.ExprBinarySpaceship) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprBinarySpaceship{\n")
}

func (v *Dump) ExprCastArray(_ *ast.ExprCastArray) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastArray{\n")
}

func (v *Dump) ExprCastBool(_ *ast.ExprCastBool) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastBool{\n")
}

func (v *Dump) ExprCastDouble(_ *ast.ExprCastDouble) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastDouble{\n")
}

func (v *Dump) ExprCastInt(_ *ast.ExprCastInt) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastInt{\n")
}

func (v *Dump) ExprCastObject(_ *ast.ExprCastObject) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastObject{\n")
}

func (v *Dump) ExprCastString(_ *ast.ExprCastString) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastString{\n")
}

func (v *Dump) ExprCastUnset(_ *ast.ExprCastUnset) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ExprCastUnset{\n")
}

func (v *Dump) ScalarDnumber(n *ast.ScalarDnumber) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarDnumber{\n")

	v.printIndentIfNotSingle(v.indent)
	v.print(fmt.Sprintf("Value: %q,\n", n.Value))
}

func (v *Dump) ScalarEncapsed(_ *ast.ScalarEncapsed) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarEncapsed{\n")
}

func (v *Dump) ScalarEncapsedStringPart(n *ast.ScalarEncapsedStringPart) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarEncapsedStringPart{\n")

	v.printIndentIfNotSingle(v.indent)
	v.print(fmt.Sprintf("Value: %q,\n", n.Value))
}

func (v *Dump) ScalarHeredoc(n *ast.ScalarHeredoc) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarHeredoc{\n")

	v.printIndentIfNotSingle(v.indent)
	v.print(fmt.Sprintf("Label: %q,\n", n.Label))
}

func (v *Dump) ScalarLnumber(n *ast.ScalarLnumber) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarLnumber{\n")

	v.printIndentIfNotSingle(v.indent)
	v.print(fmt.Sprintf("Value: %q,\n", n.Value))
}

func (v *Dump) ScalarMagicConstant(n *ast.ScalarMagicConstant) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarMagicConstant{\n")

	v.printIndentIfNotSingle(v.indent)
	v.print(fmt.Sprintf("Value: %q,\n", n.Value))
}

func (v *Dump) ScalarString(n *ast.ScalarString) {
	v.printIndentIfNotSingle(v.indent - 1)
	v.print("&ast.ScalarString{\n")

	v.printIndentIfNotSingle(v.indent)
	v.print(fmt.Sprintf("Value: %q,\n", n.Value))
}
