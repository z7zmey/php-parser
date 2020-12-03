package printer

import (
	"io"
	"strings"

	"github.com/z7zmey/php-parser/pkg/ast"
)

type PrettyPrinter struct {
	w           io.Writer
	indentStr   string
	indentDepth int
}

// NewPrettyPrinter -  Constructor for PrettyPrinter
func NewPrettyPrinter(w io.Writer, indentStr string) *PrettyPrinter {
	return &PrettyPrinter{
		w:           w,
		indentStr:   indentStr,
		indentDepth: 0,
	}
}

func (p *PrettyPrinter) Print(n ast.Vertex) {
	p.printNode(n)
}

func (p *PrettyPrinter) joinPrint(glue string, nn []ast.Vertex) {
	for k, n := range nn {
		if k > 0 {
			io.WriteString(p.w, glue)
		}

		p.Print(n)
	}
}

func (p *PrettyPrinter) printNodes(nn []ast.Vertex) {
	p.indentDepth++
	l := len(nn) - 1
	for k, n := range nn {
		p.printIndent()
		p.Print(n)
		if k < l {
			io.WriteString(p.w, "\n")
		}
	}
	p.indentDepth--
}

func (p *PrettyPrinter) printIndent() {
	for i := 0; i < p.indentDepth; i++ {
		io.WriteString(p.w, p.indentStr)
	}
}

func (p *PrettyPrinter) printNode(n ast.Vertex) {
	switch n.(type) {

	// node

	case *ast.Root:
		p.printNodeRoot(n)
	case *ast.Identifier:
		p.printNodeIdentifier(n)
	case *ast.Reference:
		p.printNodeReference(n)
	case *ast.Variadic:
		p.printNodeVariadic(n)
	case *ast.Parameter:
		p.printNodeParameter(n)
	case *ast.Nullable:
		p.printNodeNullable(n)
	case *ast.Argument:
		p.printNodeArgument(n)

		// name

	case *ast.NameNamePart:
		p.printNameNamePart(n)
	case *ast.NameName:
		p.printNameName(n)
	case *ast.NameFullyQualified:
		p.printNameFullyQualified(n)
	case *ast.NameRelative:
		p.printNameRelative(n)

		// scalar

	case *ast.ScalarLnumber:
		p.printScalarLNumber(n)
	case *ast.ScalarDnumber:
		p.printScalarDNumber(n)
	case *ast.ScalarString:
		p.printScalarString(n)
	case *ast.ScalarEncapsedStringPart:
		p.printScalarEncapsedStringPart(n)
	case *ast.ScalarEncapsed:
		p.printScalarEncapsed(n)
	case *ast.ScalarHeredoc:
		p.printScalarHeredoc(n)
	case *ast.ScalarMagicConstant:
		p.printScalarMagicConstant(n)

		// assign

	case *ast.ExprAssign:
		p.printAssignAssign(n)
	case *ast.ExprAssignReference:
		p.printAssignReference(n)
	case *ast.ExprAssignBitwiseAnd:
		p.printAssignBitwiseAnd(n)
	case *ast.ExprAssignBitwiseOr:
		p.printAssignBitwiseOr(n)
	case *ast.ExprAssignBitwiseXor:
		p.printAssignBitwiseXor(n)
	case *ast.ExprAssignConcat:
		p.printAssignConcat(n)
	case *ast.ExprAssignDiv:
		p.printAssignDiv(n)
	case *ast.ExprAssignMinus:
		p.printAssignMinus(n)
	case *ast.ExprAssignMod:
		p.printAssignMod(n)
	case *ast.ExprAssignMul:
		p.printAssignMul(n)
	case *ast.ExprAssignPlus:
		p.printAssignPlus(n)
	case *ast.ExprAssignPow:
		p.printAssignPow(n)
	case *ast.ExprAssignShiftLeft:
		p.printAssignShiftLeft(n)
	case *ast.ExprAssignShiftRight:
		p.printAssignShiftRight(n)

		// binary

	case *ast.ExprBinaryBitwiseAnd:
		p.printBinaryBitwiseAnd(n)
	case *ast.ExprBinaryBitwiseOr:
		p.printBinaryBitwiseOr(n)
	case *ast.ExprBinaryBitwiseXor:
		p.printBinaryBitwiseXor(n)
	case *ast.ExprBinaryBooleanAnd:
		p.printBinaryBooleanAnd(n)
	case *ast.ExprBinaryBooleanOr:
		p.printBinaryBooleanOr(n)
	case *ast.ExprBinaryCoalesce:
		p.printBinaryCoalesce(n)
	case *ast.ExprBinaryConcat:
		p.printBinaryConcat(n)
	case *ast.ExprBinaryDiv:
		p.printBinaryDiv(n)
	case *ast.ExprBinaryEqual:
		p.printBinaryEqual(n)
	case *ast.ExprBinaryGreaterOrEqual:
		p.printBinaryGreaterOrEqual(n)
	case *ast.ExprBinaryGreater:
		p.printBinaryGreater(n)
	case *ast.ExprBinaryIdentical:
		p.printBinaryIdentical(n)
	case *ast.ExprBinaryLogicalAnd:
		p.printBinaryLogicalAnd(n)
	case *ast.ExprBinaryLogicalOr:
		p.printBinaryLogicalOr(n)
	case *ast.ExprBinaryLogicalXor:
		p.printBinaryLogicalXor(n)
	case *ast.ExprBinaryMinus:
		p.printBinaryMinus(n)
	case *ast.ExprBinaryMod:
		p.printBinaryMod(n)
	case *ast.ExprBinaryMul:
		p.printBinaryMul(n)
	case *ast.ExprBinaryNotEqual:
		p.printBinaryNotEqual(n)
	case *ast.ExprBinaryNotIdentical:
		p.printBinaryNotIdentical(n)
	case *ast.ExprBinaryPlus:
		p.printBinaryPlus(n)
	case *ast.ExprBinaryPow:
		p.printBinaryPow(n)
	case *ast.ExprBinaryShiftLeft:
		p.printBinaryShiftLeft(n)
	case *ast.ExprBinaryShiftRight:
		p.printBinaryShiftRight(n)
	case *ast.ExprBinarySmallerOrEqual:
		p.printBinarySmallerOrEqual(n)
	case *ast.ExprBinarySmaller:
		p.printBinarySmaller(n)
	case *ast.ExprBinarySpaceship:
		p.printBinarySpaceship(n)

		// cast

	case *ast.ExprCastArray:
		p.printArray(n)
	case *ast.ExprCastBool:
		p.printBool(n)
	case *ast.ExprCastDouble:
		p.printDouble(n)
	case *ast.ExprCastInt:
		p.printInt(n)
	case *ast.ExprCastObject:
		p.printObject(n)
	case *ast.ExprCastString:
		p.printString(n)
	case *ast.ExprCastUnset:
		p.printUnset(n)

		// expr

	case *ast.ExprArrayDimFetch:
		p.printExprArrayDimFetch(n)
	case *ast.ExprArrayItem:
		p.printExprArrayItem(n)
	case *ast.ExprArray:
		p.printExprArray(n)
	case *ast.ExprBitwiseNot:
		p.printExprBitwiseNot(n)
	case *ast.ExprBooleanNot:
		p.printExprBooleanNot(n)
	case *ast.ExprClassConstFetch:
		p.printExprClassConstFetch(n)
	case *ast.ExprClone:
		p.printExprClone(n)
	case *ast.ExprClosureUse:
		p.printExprClosureUse(n)
	case *ast.ExprClosure:
		p.printExprClosure(n)
	case *ast.ExprConstFetch:
		p.printExprConstFetch(n)
	case *ast.ExprEmpty:
		p.printExprEmpty(n)
	case *ast.ExprErrorSuppress:
		p.printExprErrorSuppress(n)
	case *ast.ExprEval:
		p.printExprEval(n)
	case *ast.ExprExit:
		p.printExprExit(n)
	case *ast.ExprFunctionCall:
		p.printExprFunctionCall(n)
	case *ast.ExprInclude:
		p.printExprInclude(n)
	case *ast.ExprIncludeOnce:
		p.printExprIncludeOnce(n)
	case *ast.ExprInstanceOf:
		p.printExprInstanceOf(n)
	case *ast.ExprIsset:
		p.printExprIsset(n)
	case *ast.ExprList:
		p.printExprList(n)
	case *ast.ExprMethodCall:
		p.printExprMethodCall(n)
	case *ast.ExprNew:
		p.printExprNew(n)
	case *ast.ExprPostDec:
		p.printExprPostDec(n)
	case *ast.ExprPostInc:
		p.printExprPostInc(n)
	case *ast.ExprPreDec:
		p.printExprPreDec(n)
	case *ast.ExprPreInc:
		p.printExprPreInc(n)
	case *ast.ExprPrint:
		p.printExprPrint(n)
	case *ast.ExprPropertyFetch:
		p.printExprPropertyFetch(n)
	case *ast.ExprReference:
		p.printExprReference(n)
	case *ast.ExprRequire:
		p.printExprRequire(n)
	case *ast.ExprRequireOnce:
		p.printExprRequireOnce(n)
	case *ast.ExprShellExec:
		p.printExprShellExec(n)
	case *ast.ExprStaticCall:
		p.printExprStaticCall(n)
	case *ast.ExprStaticPropertyFetch:
		p.printExprStaticPropertyFetch(n)
	case *ast.ExprTernary:
		p.printExprTernary(n)
	case *ast.ExprUnaryMinus:
		p.printExprUnaryMinus(n)
	case *ast.ExprUnaryPlus:
		p.printExprUnaryPlus(n)
	case *ast.ExprVariable:
		p.printExprVariable(n)
	case *ast.ExprYieldFrom:
		p.printExprYieldFrom(n)
	case *ast.ExprYield:
		p.printExprYield(n)

		// stmt

	case *ast.StmtBreak:
		p.printStmtBreak(n)
	case *ast.StmtCase:
		p.printStmtCase(n)
	case *ast.StmtCatch:
		p.printStmtCatch(n)
	case *ast.StmtClassMethod:
		p.printStmtClassMethod(n)
	case *ast.StmtClass:
		p.printStmtClass(n)
	case *ast.StmtClassConstList:
		p.printStmtClassConstList(n)
	case *ast.StmtConstant:
		p.printStmtConstant(n)
	case *ast.StmtContinue:
		p.printStmtContinue(n)
	case *ast.StmtDeclare:
		p.printStmtDeclare(n)
	case *ast.StmtDefault:
		p.printStmtDefault(n)
	case *ast.StmtDo:
		p.printStmtDo(n)
	case *ast.StmtEcho:
		p.printStmtEcho(n)
	case *ast.StmtElseIf:
		p.printStmtElseif(n)
	case *ast.StmtElse:
		p.printStmtElse(n)
	case *ast.StmtExpression:
		p.printStmtExpression(n)
	case *ast.StmtFinally:
		p.printStmtFinally(n)
	case *ast.StmtFor:
		p.printStmtFor(n)
	case *ast.StmtForeach:
		p.printStmtForeach(n)
	case *ast.StmtFunction:
		p.printStmtFunction(n)
	case *ast.StmtGlobal:
		p.printStmtGlobal(n)
	case *ast.StmtGoto:
		p.printStmtGoto(n)
	case *ast.StmtHaltCompiler:
		p.printStmtHaltCompiler(n)
	case *ast.StmtIf:
		p.printStmtIf(n)
	case *ast.StmtInlineHtml:
		p.printStmtInlineHTML(n)
	case *ast.StmtInterface:
		p.printStmtInterface(n)
	case *ast.StmtLabel:
		p.printStmtLabel(n)
	case *ast.StmtNamespace:
		p.printStmtNamespace(n)
	case *ast.StmtNop:
		p.printStmtNop(n)
	case *ast.StmtPropertyList:
		p.printStmtPropertyList(n)
	case *ast.StmtProperty:
		p.printStmtProperty(n)
	case *ast.StmtReturn:
		p.printStmtReturn(n)
	case *ast.StmtStaticVar:
		p.printStmtStaticVar(n)
	case *ast.StmtStatic:
		p.printStmtStatic(n)
	case *ast.StmtStmtList:
		p.printStmtStmtList(n)
	case *ast.StmtSwitch:
		p.printStmtSwitch(n)
	case *ast.StmtThrow:
		p.printStmtThrow(n)
	case *ast.StmtTraitMethodRef:
		p.printStmtTraitMethodRef(n)
	case *ast.StmtTraitUseAlias:
		p.printStmtTraitUseAlias(n)
	case *ast.StmtTraitUsePrecedence:
		p.printStmtTraitUsePrecedence(n)
	case *ast.StmtTraitUse:
		p.printStmtTraitUse(n)
	case *ast.StmtTrait:
		p.printStmtTrait(n)
	case *ast.StmtTry:
		p.printStmtTry(n)
	case *ast.StmtUnset:
		p.printStmtUnset(n)
	case *ast.StmtUse:
		p.printStmtUse(n)
	case *ast.StmtGroupUse:
		p.printStmtGroupUse(n)
	case *ast.StmtUseDeclaration:
		p.printStmtUseDeclaration(n)
	case *ast.StmtWhile:
		p.printStmtWhile(n)
	}
}

// node

func (p *PrettyPrinter) printNodeRoot(n ast.Vertex) {
	var stmts []ast.Vertex
	v := n.(*ast.Root)

	if len(v.Stmts) > 0 {
		firstStmt := v.Stmts[0]
		stmts = v.Stmts[1:]

		switch fs := firstStmt.(type) {
		case *ast.StmtInlineHtml:
			io.WriteString(p.w, string(fs.Value))
			io.WriteString(p.w, "<?php\n")
		default:
			io.WriteString(p.w, "<?php\n")
			p.printIndent()
			p.Print(fs)
			io.WriteString(p.w, "\n")
		}
	}
	p.indentDepth--
	p.printNodes(stmts)
	io.WriteString(p.w, "\n")
}

func (p *PrettyPrinter) printNodeIdentifier(n ast.Vertex) {
	v := string(n.(*ast.Identifier).Value)
	io.WriteString(p.w, v)
}

func (p *PrettyPrinter) printNodeReference(n ast.Vertex) {
	nn := n.(*ast.Reference)

	io.WriteString(p.w, "&")
	p.Print(nn.Var)
}

func (p *PrettyPrinter) printNodeVariadic(n ast.Vertex) {
	nn := n.(*ast.Variadic)

	io.WriteString(p.w, "...")
	p.Print(nn.Var)
}

func (p *PrettyPrinter) printNodeParameter(n ast.Vertex) {
	nn := n.(*ast.Parameter)

	if nn.Type != nil {
		p.Print(nn.Type)
		io.WriteString(p.w, " ")
	}

	p.Print(nn.Var)

	if nn.DefaultValue != nil {
		io.WriteString(p.w, " = ")
		p.Print(nn.DefaultValue)
	}
}

func (p *PrettyPrinter) printNodeNullable(n ast.Vertex) {
	nn := n.(*ast.Nullable)

	io.WriteString(p.w, "?")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printNodeArgument(n ast.Vertex) {
	nn := n.(*ast.Argument)

	if nn.AmpersandTkn != nil {
		io.WriteString(p.w, "&")
	}

	if nn.VariadicTkn != nil {
		io.WriteString(p.w, "...")
	}

	p.Print(nn.Expr)
}

// name

func (p *PrettyPrinter) printNameNamePart(n ast.Vertex) {
	v := string(n.(*ast.NameNamePart).Value)
	io.WriteString(p.w, v)
}

func (p *PrettyPrinter) printNameName(n ast.Vertex) {
	nn := n.(*ast.NameName)

	for k, part := range nn.Parts {
		if k > 0 {
			io.WriteString(p.w, "\\")
		}

		p.Print(part)
	}
}

func (p *PrettyPrinter) printNameFullyQualified(n ast.Vertex) {
	nn := n.(*ast.NameFullyQualified)

	for _, part := range nn.Parts {
		io.WriteString(p.w, "\\")
		p.Print(part)
	}
}

func (p *PrettyPrinter) printNameRelative(n ast.Vertex) {
	nn := n.(*ast.NameRelative)

	io.WriteString(p.w, "namespace")
	for _, part := range nn.Parts {
		io.WriteString(p.w, "\\")
		p.Print(part)
	}
}

// scalar

func (p *PrettyPrinter) printScalarLNumber(n ast.Vertex) {
	v := string(n.(*ast.ScalarLnumber).Value)
	io.WriteString(p.w, v)
}

func (p *PrettyPrinter) printScalarDNumber(n ast.Vertex) {
	v := string(n.(*ast.ScalarDnumber).Value)
	io.WriteString(p.w, v)
}

func (p *PrettyPrinter) printScalarString(n ast.Vertex) {
	v := string(n.(*ast.ScalarString).Value)

	io.WriteString(p.w, v)
}

func (p *PrettyPrinter) printScalarEncapsedStringPart(n ast.Vertex) {
	v := string(n.(*ast.ScalarEncapsedStringPart).Value)
	io.WriteString(p.w, v)
}

func (p *PrettyPrinter) printScalarEncapsed(n ast.Vertex) {
	nn := n.(*ast.ScalarEncapsed)
	io.WriteString(p.w, "\"")

	for _, part := range nn.Parts {
		switch part.(type) {
		case *ast.ScalarEncapsedStringPart:
			p.Print(part)
		default:
			io.WriteString(p.w, "{")
			p.Print(part)
			io.WriteString(p.w, "}")
		}
	}

	io.WriteString(p.w, "\"")
}

func (p *PrettyPrinter) printScalarHeredoc(n ast.Vertex) {
	nn := n.(*ast.ScalarHeredoc)

	io.WriteString(p.w, string(nn.OpenHeredocTkn.Value))

	for _, part := range nn.Parts {
		switch part.(type) {
		case *ast.ScalarEncapsedStringPart:
			p.Print(part)
		default:
			io.WriteString(p.w, "{")
			p.Print(part)
			io.WriteString(p.w, "}")
		}
	}

	io.WriteString(p.w, strings.Trim(string(nn.OpenHeredocTkn.Value), "<\"'\n"))
}

func (p *PrettyPrinter) printScalarMagicConstant(n ast.Vertex) {
	v := string(n.(*ast.ScalarMagicConstant).Value)
	io.WriteString(p.w, v)
}

// Assign

func (p *PrettyPrinter) printAssignAssign(n ast.Vertex) {
	nn := n.(*ast.ExprAssign)
	p.Print(nn.Var)
	io.WriteString(p.w, " = ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignReference(n ast.Vertex) {
	nn := n.(*ast.ExprAssignReference)
	p.Print(nn.Var)
	io.WriteString(p.w, " =& ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignBitwiseAnd(n ast.Vertex) {
	nn := n.(*ast.ExprAssignBitwiseAnd)
	p.Print(nn.Var)
	io.WriteString(p.w, " &= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignBitwiseOr(n ast.Vertex) {
	nn := n.(*ast.ExprAssignBitwiseOr)
	p.Print(nn.Var)
	io.WriteString(p.w, " |= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignBitwiseXor(n ast.Vertex) {
	nn := n.(*ast.ExprAssignBitwiseXor)
	p.Print(nn.Var)
	io.WriteString(p.w, " ^= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignConcat(n ast.Vertex) {
	nn := n.(*ast.ExprAssignConcat)
	p.Print(nn.Var)
	io.WriteString(p.w, " .= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignDiv(n ast.Vertex) {
	nn := n.(*ast.ExprAssignDiv)
	p.Print(nn.Var)
	io.WriteString(p.w, " /= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignMinus(n ast.Vertex) {
	nn := n.(*ast.ExprAssignMinus)
	p.Print(nn.Var)
	io.WriteString(p.w, " -= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignMod(n ast.Vertex) {
	nn := n.(*ast.ExprAssignMod)
	p.Print(nn.Var)
	io.WriteString(p.w, " %= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignMul(n ast.Vertex) {
	nn := n.(*ast.ExprAssignMul)
	p.Print(nn.Var)
	io.WriteString(p.w, " *= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignPlus(n ast.Vertex) {
	nn := n.(*ast.ExprAssignPlus)
	p.Print(nn.Var)
	io.WriteString(p.w, " += ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignPow(n ast.Vertex) {
	nn := n.(*ast.ExprAssignPow)
	p.Print(nn.Var)
	io.WriteString(p.w, " **= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignShiftLeft(n ast.Vertex) {
	nn := n.(*ast.ExprAssignShiftLeft)
	p.Print(nn.Var)
	io.WriteString(p.w, " <<= ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printAssignShiftRight(n ast.Vertex) {
	nn := n.(*ast.ExprAssignShiftRight)
	p.Print(nn.Var)
	io.WriteString(p.w, " >>= ")
	p.Print(nn.Expr)
}

// binary

func (p *PrettyPrinter) printBinaryBitwiseAnd(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBitwiseAnd)

	p.Print(nn.Left)
	io.WriteString(p.w, " & ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryBitwiseOr(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBitwiseOr)

	p.Print(nn.Left)
	io.WriteString(p.w, " | ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryBitwiseXor(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBitwiseXor)

	p.Print(nn.Left)
	io.WriteString(p.w, " ^ ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryBooleanAnd(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBooleanAnd)

	p.Print(nn.Left)
	io.WriteString(p.w, " && ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryBooleanOr(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBooleanOr)

	p.Print(nn.Left)
	io.WriteString(p.w, " || ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryCoalesce(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryCoalesce)

	p.Print(nn.Left)
	io.WriteString(p.w, " ?? ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryConcat(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryConcat)

	p.Print(nn.Left)
	io.WriteString(p.w, " . ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryDiv(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryDiv)

	p.Print(nn.Left)
	io.WriteString(p.w, " / ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryEqual)

	p.Print(nn.Left)
	io.WriteString(p.w, " == ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryGreaterOrEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryGreaterOrEqual)

	p.Print(nn.Left)
	io.WriteString(p.w, " >= ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryGreater(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryGreater)

	p.Print(nn.Left)
	io.WriteString(p.w, " > ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryIdentical(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryIdentical)

	p.Print(nn.Left)
	io.WriteString(p.w, " === ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryLogicalAnd(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryLogicalAnd)

	p.Print(nn.Left)
	io.WriteString(p.w, " and ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryLogicalOr(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryLogicalOr)

	p.Print(nn.Left)
	io.WriteString(p.w, " or ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryLogicalXor(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryLogicalXor)

	p.Print(nn.Left)
	io.WriteString(p.w, " xor ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryMinus(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryMinus)

	p.Print(nn.Left)
	io.WriteString(p.w, " - ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryMod(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryMod)

	p.Print(nn.Left)
	io.WriteString(p.w, " % ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryMul(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryMul)

	p.Print(nn.Left)
	io.WriteString(p.w, " * ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryNotEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryNotEqual)

	p.Print(nn.Left)
	io.WriteString(p.w, " != ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryNotIdentical(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryNotIdentical)

	p.Print(nn.Left)
	io.WriteString(p.w, " !== ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryPlus(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryPlus)

	p.Print(nn.Left)
	io.WriteString(p.w, " + ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryPow(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryPow)

	p.Print(nn.Left)
	io.WriteString(p.w, " ** ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryShiftLeft(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryShiftLeft)

	p.Print(nn.Left)
	io.WriteString(p.w, " << ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinaryShiftRight(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryShiftRight)

	p.Print(nn.Left)
	io.WriteString(p.w, " >> ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinarySmallerOrEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinarySmallerOrEqual)

	p.Print(nn.Left)
	io.WriteString(p.w, " <= ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinarySmaller(n ast.Vertex) {
	nn := n.(*ast.ExprBinarySmaller)

	p.Print(nn.Left)
	io.WriteString(p.w, " < ")
	p.Print(nn.Right)
}

func (p *PrettyPrinter) printBinarySpaceship(n ast.Vertex) {
	nn := n.(*ast.ExprBinarySpaceship)

	p.Print(nn.Left)
	io.WriteString(p.w, " <=> ")
	p.Print(nn.Right)
}

// cast

func (p *PrettyPrinter) printArray(n ast.Vertex) {
	nn := n.(*ast.ExprCastArray)

	io.WriteString(p.w, "(array)")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printBool(n ast.Vertex) {
	nn := n.(*ast.ExprCastBool)

	io.WriteString(p.w, "(bool)")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printDouble(n ast.Vertex) {
	nn := n.(*ast.ExprCastDouble)

	io.WriteString(p.w, "(float)")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printInt(n ast.Vertex) {
	nn := n.(*ast.ExprCastInt)

	io.WriteString(p.w, "(int)")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printObject(n ast.Vertex) {
	nn := n.(*ast.ExprCastObject)

	io.WriteString(p.w, "(object)")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printString(n ast.Vertex) {
	nn := n.(*ast.ExprCastString)

	io.WriteString(p.w, "(string)")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printUnset(n ast.Vertex) {
	nn := n.(*ast.ExprCastUnset)

	io.WriteString(p.w, "(unset)")
	p.Print(nn.Expr)
}

// expr

func (p *PrettyPrinter) printExprArrayDimFetch(n ast.Vertex) {
	nn := n.(*ast.ExprArrayDimFetch)
	p.Print(nn.Var)
	io.WriteString(p.w, "[")
	p.Print(nn.Dim)
	io.WriteString(p.w, "]")
}

func (p *PrettyPrinter) printExprArrayItem(n ast.Vertex) {
	nn := n.(*ast.ExprArrayItem)

	if nn.Key != nil {
		p.Print(nn.Key)
		io.WriteString(p.w, " => ")
	}

	p.Print(nn.Val)
}

func (p *PrettyPrinter) printExprArray(n ast.Vertex) {
	nn := n.(*ast.ExprArray)

	io.WriteString(p.w, "array(")
	p.joinPrint(", ", nn.Items)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprBitwiseNot(n ast.Vertex) {
	nn := n.(*ast.ExprBitwiseNot)
	io.WriteString(p.w, "~")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprBooleanNot(n ast.Vertex) {
	nn := n.(*ast.ExprBooleanNot)
	io.WriteString(p.w, "!")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprClassConstFetch(n ast.Vertex) {
	nn := n.(*ast.ExprClassConstFetch)

	p.Print(nn.Class)
	io.WriteString(p.w, "::")
	io.WriteString(p.w, string(nn.ConstantName.(*ast.Identifier).Value))
}

func (p *PrettyPrinter) printExprClone(n ast.Vertex) {
	nn := n.(*ast.ExprClone)

	io.WriteString(p.w, "clone ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprClosureUse(n ast.Vertex) {
	nn := n.(*ast.ExprClosureUse)

	io.WriteString(p.w, "use (")
	p.joinPrint(", ", nn.Uses)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprClosure(n ast.Vertex) {
	nn := n.(*ast.ExprClosure)

	if nn.StaticTkn != nil {
		io.WriteString(p.w, "static ")
	}

	io.WriteString(p.w, "function ")

	if nn.AmpersandTkn != nil {
		io.WriteString(p.w, "&")
	}

	io.WriteString(p.w, "(")
	p.joinPrint(", ", nn.Params)
	io.WriteString(p.w, ")")

	if nn.ClosureUse != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.ClosureUse)
	}

	if nn.ReturnType != nil {
		io.WriteString(p.w, ": ")
		p.Print(nn.ReturnType)
	}

	io.WriteString(p.w, " {\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printExprConstFetch(n ast.Vertex) {
	nn := n.(*ast.ExprConstFetch)

	p.Print(nn.Const)
}

func (p *PrettyPrinter) printExprEmpty(n ast.Vertex) {
	nn := n.(*ast.ExprEmpty)

	io.WriteString(p.w, "empty(")
	p.Print(nn.Expr)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprErrorSuppress(n ast.Vertex) {
	nn := n.(*ast.ExprErrorSuppress)

	io.WriteString(p.w, "@")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprEval(n ast.Vertex) {
	nn := n.(*ast.ExprEval)

	io.WriteString(p.w, "eval(")
	p.Print(nn.Expr)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprExit(n ast.Vertex) {
	nn := n.(*ast.ExprExit)

	io.WriteString(p.w, "exit(")
	p.Print(nn.Expr)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprFunctionCall(n ast.Vertex) {
	nn := n.(*ast.ExprFunctionCall)

	p.Print(nn.Function)
	io.WriteString(p.w, "(")
	p.joinPrint(", ", nn.Arguments)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprInclude(n ast.Vertex) {
	nn := n.(*ast.ExprInclude)

	io.WriteString(p.w, "include ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprIncludeOnce(n ast.Vertex) {
	nn := n.(*ast.ExprIncludeOnce)

	io.WriteString(p.w, "include_once ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprInstanceOf(n ast.Vertex) {
	nn := n.(*ast.ExprInstanceOf)

	p.Print(nn.Expr)
	io.WriteString(p.w, " instanceof ")
	p.Print(nn.Class)
}

func (p *PrettyPrinter) printExprIsset(n ast.Vertex) {
	nn := n.(*ast.ExprIsset)

	io.WriteString(p.w, "isset(")
	p.joinPrint(", ", nn.Vars)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprList(n ast.Vertex) {
	nn := n.(*ast.ExprList)

	io.WriteString(p.w, "list(")
	p.joinPrint(", ", nn.Items)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprMethodCall(n ast.Vertex) {
	nn := n.(*ast.ExprMethodCall)

	p.Print(nn.Var)
	io.WriteString(p.w, "->")
	p.Print(nn.Method)
	io.WriteString(p.w, "(")
	p.joinPrint(", ", nn.Arguments)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprNew(n ast.Vertex) {
	nn := n.(*ast.ExprNew)

	io.WriteString(p.w, "new ")
	p.Print(nn.Class)

	if nn.Arguments != nil {
		io.WriteString(p.w, "(")
		p.joinPrint(", ", nn.Arguments)
		io.WriteString(p.w, ")")
	}
}

func (p *PrettyPrinter) printExprPostDec(n ast.Vertex) {
	nn := n.(*ast.ExprPostDec)

	p.Print(nn.Var)
	io.WriteString(p.w, "--")
}

func (p *PrettyPrinter) printExprPostInc(n ast.Vertex) {
	nn := n.(*ast.ExprPostInc)

	p.Print(nn.Var)
	io.WriteString(p.w, "++")
}

func (p *PrettyPrinter) printExprPreDec(n ast.Vertex) {
	nn := n.(*ast.ExprPreDec)

	io.WriteString(p.w, "--")
	p.Print(nn.Var)
}

func (p *PrettyPrinter) printExprPreInc(n ast.Vertex) {
	nn := n.(*ast.ExprPreInc)

	io.WriteString(p.w, "++")
	p.Print(nn.Var)
}

func (p *PrettyPrinter) printExprPrint(n ast.Vertex) {
	nn := n.(*ast.ExprPrint)

	io.WriteString(p.w, "print(")
	p.Print(nn.Expr)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprPropertyFetch(n ast.Vertex) {
	nn := n.(*ast.ExprPropertyFetch)

	p.Print(nn.Var)
	io.WriteString(p.w, "->")
	p.Print(nn.Property)
}

func (p *PrettyPrinter) printExprReference(n ast.Vertex) {
	nn := n.(*ast.ExprReference)

	io.WriteString(p.w, "&")
	p.Print(nn.Var)
}

func (p *PrettyPrinter) printExprRequire(n ast.Vertex) {
	nn := n.(*ast.ExprRequire)

	io.WriteString(p.w, "require ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprRequireOnce(n ast.Vertex) {
	nn := n.(*ast.ExprRequireOnce)

	io.WriteString(p.w, "require_once ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprShellExec(n ast.Vertex) {
	nn := n.(*ast.ExprShellExec)

	io.WriteString(p.w, "`")
	for _, part := range nn.Parts {
		switch part.(type) {
		case *ast.ScalarEncapsedStringPart:
			p.Print(part)
		default:
			io.WriteString(p.w, "{")
			p.Print(part)
			io.WriteString(p.w, "}")
		}
	}
	io.WriteString(p.w, "`")
}

func (p *PrettyPrinter) printExprStaticCall(n ast.Vertex) {
	nn := n.(*ast.ExprStaticCall)

	p.Print(nn.Class)
	io.WriteString(p.w, "::")
	p.Print(nn.Call)
	io.WriteString(p.w, "(")
	p.joinPrint(", ", nn.Arguments)
	io.WriteString(p.w, ")")
}

func (p *PrettyPrinter) printExprStaticPropertyFetch(n ast.Vertex) {
	nn := n.(*ast.ExprStaticPropertyFetch)

	p.Print(nn.Class)
	io.WriteString(p.w, "::")
	p.Print(nn.Property)
}

func (p *PrettyPrinter) printExprTernary(n ast.Vertex) {
	nn := n.(*ast.ExprTernary)

	p.Print(nn.Condition)
	io.WriteString(p.w, " ?")

	if nn.IfTrue != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.IfTrue)
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, ": ")
	p.Print(nn.IfFalse)
}

func (p *PrettyPrinter) printExprUnaryMinus(n ast.Vertex) {
	nn := n.(*ast.ExprUnaryMinus)

	io.WriteString(p.w, "-")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprUnaryPlus(n ast.Vertex) {
	nn := n.(*ast.ExprUnaryPlus)

	io.WriteString(p.w, "+")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprVariable(n ast.Vertex) {
	nn := n.(*ast.ExprVariable)
	io.WriteString(p.w, "$")
	p.Print(nn.VarName)
}

func (p *PrettyPrinter) printExprYieldFrom(n ast.Vertex) {
	nn := n.(*ast.ExprYieldFrom)

	io.WriteString(p.w, "yield from ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printExprYield(n ast.Vertex) {
	nn := n.(*ast.ExprYield)

	io.WriteString(p.w, "yield ")

	if nn.Key != nil {
		p.Print(nn.Key)
		io.WriteString(p.w, " => ")
	}

	p.Print(nn.Value)
}

// smtm

func (p *PrettyPrinter) printStmtAltElseIf(n ast.Vertex) {
	nn := n.(*ast.StmtElseIf)

	io.WriteString(p.w, "elseif (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ") :")

	if s := nn.Stmt.(*ast.StmtStmtList).Stmts; len(s) > 0 {
		io.WriteString(p.w, "\n")
		p.printNodes(s)
	}
}

func (p *PrettyPrinter) printStmtAltElse(n ast.Vertex) {
	nn := n.(*ast.StmtElse)

	io.WriteString(p.w, "else :")

	if s := nn.Stmt.(*ast.StmtStmtList).Stmts; len(s) > 0 {
		io.WriteString(p.w, "\n")
		p.printNodes(s)
	}
}

func (p *PrettyPrinter) printStmtAltIf(n ast.Vertex) {
	nn := n.(*ast.StmtIf)

	io.WriteString(p.w, "if (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ") :\n")

	s := nn.Stmt.(*ast.StmtStmtList)
	p.printNodes(s.Stmts)

	for _, elseif := range nn.ElseIf {
		io.WriteString(p.w, "\n")
		p.printIndent()
		p.Print(elseif)
	}

	if nn.Else != nil {
		io.WriteString(p.w, "\n")
		p.printIndent()
		p.Print(nn.Else)
	}

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "endif;")
}

func (p *PrettyPrinter) printStmtBreak(n ast.Vertex) {
	nn := n.(*ast.StmtBreak)

	io.WriteString(p.w, "break")
	if nn.Expr != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.Expr)
	}

	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtCase(n ast.Vertex) {
	nn := n.(*ast.StmtCase)

	io.WriteString(p.w, "case ")
	p.Print(nn.Cond)
	io.WriteString(p.w, ":")

	if len(nn.Stmts) > 0 {
		io.WriteString(p.w, "\n")
		p.printNodes(nn.Stmts)
	}
}

func (p *PrettyPrinter) printStmtCatch(n ast.Vertex) {
	nn := n.(*ast.StmtCatch)

	io.WriteString(p.w, "catch (")
	p.joinPrint(" | ", nn.Types)
	io.WriteString(p.w, " ")
	p.Print(nn.Var)
	io.WriteString(p.w, ") {\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printStmtClassMethod(n ast.Vertex) {
	nn := n.(*ast.StmtClassMethod)

	if nn.Modifiers != nil {
		p.joinPrint(" ", nn.Modifiers)
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "function ")

	if nn.AmpersandTkn != nil {
		io.WriteString(p.w, "&")
	}

	p.Print(nn.MethodName)
	io.WriteString(p.w, "(")
	p.joinPrint(", ", nn.Params)
	io.WriteString(p.w, ")")

	if nn.ReturnType != nil {
		io.WriteString(p.w, ": ")
		p.Print(nn.ReturnType)
	}

	switch s := nn.Stmt.(type) {
	case *ast.StmtStmtList:
		io.WriteString(p.w, "\n")
		p.printIndent()
		io.WriteString(p.w, "{\n")
		p.printNodes(s.Stmts)
		io.WriteString(p.w, "\n")
		p.printIndent()
		io.WriteString(p.w, "}")
	default:
		p.Print(s)
	}
}

func (p *PrettyPrinter) printStmtClass(n ast.Vertex) {
	nn := n.(*ast.StmtClass)

	if nn.Modifiers != nil {
		p.joinPrint(" ", nn.Modifiers)
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "class")

	if nn.ClassName != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.ClassName)
	}

	if nn.ArgumentList != nil {
		io.WriteString(p.w, "(")
		p.joinPrint(", ", nn.ArgumentList.(*ast.ArgumentList).Arguments)
		io.WriteString(p.w, ")")
	}

	if nn.Extends != nil {
		io.WriteString(p.w, " extends ")
		p.Print(nn.Extends.(*ast.StmtClassExtends).ClassName)
	}

	if nn.Implements != nil {
		io.WriteString(p.w, " implements ")
		p.joinPrint(", ", nn.Implements.(*ast.StmtClassImplements).InterfaceNames)
	}

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "{\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printStmtClassConstList(n ast.Vertex) {
	nn := n.(*ast.StmtClassConstList)

	if nn.Modifiers != nil {
		p.joinPrint(" ", nn.Modifiers)
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "const ")

	p.joinPrint(", ", nn.Consts)

	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtConstant(n ast.Vertex) {
	nn := n.(*ast.StmtConstant)

	p.Print(nn.Name)
	io.WriteString(p.w, " = ")
	p.Print(nn.Expr)
}

func (p *PrettyPrinter) printStmtContinue(n ast.Vertex) {
	nn := n.(*ast.StmtContinue)

	io.WriteString(p.w, "continue")
	if nn.Expr != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.Expr)
	}

	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtDeclare(n ast.Vertex) {
	nn := n.(*ast.StmtDeclare)

	io.WriteString(p.w, "declare(")
	p.joinPrint(", ", nn.Consts)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *ast.StmtNop:
		p.Print(s)
		break
	case *ast.StmtStmtList:
		io.WriteString(p.w, " ")
		p.Print(s)
	default:
		io.WriteString(p.w, "\n")
		p.indentDepth++
		p.printIndent()
		p.Print(s)
		p.indentDepth--
	}
}

func (p *PrettyPrinter) printStmtDefault(n ast.Vertex) {
	nn := n.(*ast.StmtDefault)
	io.WriteString(p.w, "default:")

	if len(nn.Stmts) > 0 {
		io.WriteString(p.w, "\n")
		p.printNodes(nn.Stmts)
	}
}

func (p *PrettyPrinter) printStmtDo(n ast.Vertex) {
	nn := n.(*ast.StmtDo)
	io.WriteString(p.w, "do")

	switch s := nn.Stmt.(type) {
	case *ast.StmtStmtList:
		io.WriteString(p.w, " ")
		p.Print(s)
		io.WriteString(p.w, " ")
	default:
		io.WriteString(p.w, "\n")
		p.indentDepth++
		p.printIndent()
		p.Print(s)
		p.indentDepth--
		io.WriteString(p.w, "\n")
		p.printIndent()
	}

	io.WriteString(p.w, "while (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ");")
}

func (p *PrettyPrinter) printStmtEcho(n ast.Vertex) {
	nn := n.(*ast.StmtEcho)
	io.WriteString(p.w, "echo ")
	p.joinPrint(", ", nn.Exprs)
	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtElseif(n ast.Vertex) {
	nn := n.(*ast.StmtElseIf)

	if nn.Alt {
		p.printStmtAltElseIf(nn)
		return
	}

	io.WriteString(p.w, "elseif (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *ast.StmtNop:
		p.Print(s)
		break
	case *ast.StmtStmtList:
		io.WriteString(p.w, " ")
		p.Print(s)
	default:
		io.WriteString(p.w, "\n")
		p.indentDepth++
		p.printIndent()
		p.Print(s)
		p.indentDepth--
	}
}

func (p *PrettyPrinter) printStmtElse(n ast.Vertex) {
	nn := n.(*ast.StmtElse)

	if nn.Alt {
		p.printStmtAltElse(nn)
		return
	}

	io.WriteString(p.w, "else")

	switch s := nn.Stmt.(type) {
	case *ast.StmtNop:
		p.Print(s)
		break
	case *ast.StmtStmtList:
		io.WriteString(p.w, " ")
		p.Print(s)
	default:
		io.WriteString(p.w, "\n")
		p.indentDepth++
		p.printIndent()
		p.Print(s)
		p.indentDepth--
	}
}

func (p *PrettyPrinter) printStmtExpression(n ast.Vertex) {
	nn := n.(*ast.StmtExpression)

	p.Print(nn.Expr)

	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtFinally(n ast.Vertex) {
	nn := n.(*ast.StmtFinally)

	io.WriteString(p.w, "finally {\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printStmtFor(n ast.Vertex) {
	nn := n.(*ast.StmtFor)

	if nn.Alt {
		p.printStmtAltFor(nn)
		return
	}

	io.WriteString(p.w, "for (")
	p.joinPrint(", ", nn.Init)
	io.WriteString(p.w, "; ")
	p.joinPrint(", ", nn.Cond)
	io.WriteString(p.w, "; ")
	p.joinPrint(", ", nn.Loop)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *ast.StmtNop:
		p.Print(s)
		break
	case *ast.StmtStmtList:
		io.WriteString(p.w, " ")
		p.Print(s)
	default:
		io.WriteString(p.w, "\n")
		p.indentDepth++
		p.printIndent()
		p.Print(s)
		p.indentDepth--
	}
}

func (p *PrettyPrinter) printStmtAltFor(n ast.Vertex) {
	nn := n.(*ast.StmtFor)

	io.WriteString(p.w, "for (")
	p.joinPrint(", ", nn.Init)
	io.WriteString(p.w, "; ")
	p.joinPrint(", ", nn.Cond)
	io.WriteString(p.w, "; ")
	p.joinPrint(", ", nn.Loop)
	io.WriteString(p.w, ") :\n")

	s := nn.Stmt.(*ast.StmtStmtList)
	p.printNodes(s.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()

	io.WriteString(p.w, "endfor;")
}

func (p *PrettyPrinter) printStmtForeach(n ast.Vertex) {
	nn := n.(*ast.StmtForeach)

	if nn.Alt {
		p.printStmtAltForeach(n)
		return
	}

	io.WriteString(p.w, "foreach (")
	p.Print(nn.Expr)
	io.WriteString(p.w, " as ")

	if nn.Key != nil {
		p.Print(nn.Key)
		io.WriteString(p.w, " => ")
	}

	p.Print(nn.Var)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *ast.StmtNop:
		p.Print(s)
		break
	case *ast.StmtStmtList:
		io.WriteString(p.w, " ")
		p.Print(s)
	default:
		io.WriteString(p.w, "\n")
		p.indentDepth++
		p.printIndent()
		p.Print(s)
		p.indentDepth--
	}
}

func (p *PrettyPrinter) printStmtAltForeach(n ast.Vertex) {
	nn := n.(*ast.StmtForeach)

	io.WriteString(p.w, "foreach (")
	p.Print(nn.Expr)
	io.WriteString(p.w, " as ")

	if nn.Key != nil {
		p.Print(nn.Key)
		io.WriteString(p.w, " => ")
	}

	p.Print(nn.Var)

	io.WriteString(p.w, ") :\n")

	s := nn.Stmt.(*ast.StmtStmtList)
	p.printNodes(s.Stmts)

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "endforeach;")
}

func (p *PrettyPrinter) printStmtFunction(n ast.Vertex) {
	nn := n.(*ast.StmtFunction)

	io.WriteString(p.w, "function ")

	if nn.AmpersandTkn != nil {
		io.WriteString(p.w, "&")
	}

	p.Print(nn.FunctionName)

	io.WriteString(p.w, "(")
	p.joinPrint(", ", nn.Params)
	io.WriteString(p.w, ")")

	if nn.ReturnType != nil {
		io.WriteString(p.w, ": ")
		p.Print(nn.ReturnType)
	}

	io.WriteString(p.w, " {\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printStmtGlobal(n ast.Vertex) {
	nn := n.(*ast.StmtGlobal)

	io.WriteString(p.w, "global ")
	p.joinPrint(", ", nn.Vars)
	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtGoto(n ast.Vertex) {
	nn := n.(*ast.StmtGoto)

	io.WriteString(p.w, "goto ")
	p.Print(nn.Label)
	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtHaltCompiler(n ast.Vertex) {
	io.WriteString(p.w, "__halt_compiler();")
}

func (p *PrettyPrinter) printStmtIf(n ast.Vertex) {
	nn := n.(*ast.StmtIf)

	if nn.Alt {
		p.printStmtAltIf(nn)
		return
	}

	io.WriteString(p.w, "if (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *ast.StmtNop:
		p.Print(s)
		break
	case *ast.StmtStmtList:
		io.WriteString(p.w, " ")
		p.Print(s)
	default:
		io.WriteString(p.w, "\n")
		p.indentDepth++
		p.printIndent()
		p.Print(s)
		p.indentDepth--
	}

	if nn.ElseIf != nil {
		io.WriteString(p.w, "\n")
		p.indentDepth--
		p.printNodes(nn.ElseIf)
		p.indentDepth++
	}

	if nn.Else != nil {
		io.WriteString(p.w, "\n")
		p.printIndent()
		p.Print(nn.Else)
	}
}

func (p *PrettyPrinter) printStmtInlineHTML(n ast.Vertex) {
	nn := n.(*ast.StmtInlineHtml)

	io.WriteString(p.w, "?>")
	io.WriteString(p.w, string(nn.Value))
	io.WriteString(p.w, "<?php")
}

func (p *PrettyPrinter) printStmtInterface(n ast.Vertex) {
	nn := n.(*ast.StmtInterface)

	io.WriteString(p.w, "interface")

	if nn.InterfaceName != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.InterfaceName)
	}

	if nn.Extends != nil {
		io.WriteString(p.w, " extends ")
		p.joinPrint(", ", nn.Extends.(*ast.StmtInterfaceExtends).InterfaceNames)
	}

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "{\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printStmtLabel(n ast.Vertex) {
	nn := n.(*ast.StmtLabel)

	p.Print(nn.LabelName)
	io.WriteString(p.w, ":")
}

func (p *PrettyPrinter) printStmtNamespace(n ast.Vertex) {
	nn := n.(*ast.StmtNamespace)

	io.WriteString(p.w, "namespace")

	if nn.Name != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.Name)
	}

	if nn.Stmts != nil {
		io.WriteString(p.w, " {\n")
		p.printNodes(nn.Stmts)
		io.WriteString(p.w, "\n")
		p.printIndent()
		io.WriteString(p.w, "}")
	} else {
		io.WriteString(p.w, ";")
	}
}

func (p *PrettyPrinter) printStmtNop(n ast.Vertex) {
	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtPropertyList(n ast.Vertex) {
	nn := n.(*ast.StmtPropertyList)

	p.joinPrint(" ", nn.Modifiers)
	io.WriteString(p.w, " ")
	p.joinPrint(", ", nn.Properties)
	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtProperty(n ast.Vertex) {
	nn := n.(*ast.StmtProperty)

	p.Print(nn.Var)

	if nn.Expr != nil {
		io.WriteString(p.w, " = ")
		p.Print(nn.Expr)
	}
}

func (p *PrettyPrinter) printStmtReturn(n ast.Vertex) {
	nn := n.(*ast.StmtReturn)

	io.WriteString(p.w, "return ")
	p.Print(nn.Expr)
	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtStaticVar(n ast.Vertex) {
	nn := n.(*ast.StmtStaticVar)
	p.Print(nn.Var)

	if nn.Expr != nil {
		io.WriteString(p.w, " = ")
		p.Print(nn.Expr)
	}
}

func (p *PrettyPrinter) printStmtStatic(n ast.Vertex) {
	nn := n.(*ast.StmtStatic)

	io.WriteString(p.w, "static ")
	p.joinPrint(", ", nn.Vars)
	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtStmtList(n ast.Vertex) {
	nn := n.(*ast.StmtStmtList)

	io.WriteString(p.w, "{\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printStmtSwitch(n ast.Vertex) {
	nn := n.(*ast.StmtSwitch)

	if nn.Alt {
		p.printStmtAltSwitch(n)
		return
	}

	io.WriteString(p.w, "switch (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ")")

	io.WriteString(p.w, " {\n")
	p.printNodes(nn.CaseList)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printStmtAltSwitch(n ast.Vertex) {
	nn := n.(*ast.StmtSwitch)

	io.WriteString(p.w, "switch (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ") :\n")

	s := nn.CaseList
	p.printNodes(s)

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "endswitch;")
}

func (p *PrettyPrinter) printStmtThrow(n ast.Vertex) {
	nn := n.(*ast.StmtThrow)

	io.WriteString(p.w, "throw ")
	p.Print(nn.Expr)
	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtTraitMethodRef(n ast.Vertex) {
	nn := n.(*ast.StmtTraitMethodRef)

	p.Print(nn.Trait)
	io.WriteString(p.w, "::")
	p.Print(nn.Method)
}

func (p *PrettyPrinter) printStmtTraitUseAlias(n ast.Vertex) {
	nn := n.(*ast.StmtTraitUseAlias)

	p.Print(nn.Ref)
	io.WriteString(p.w, " as")

	if nn.Modifier != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.Modifier)
	}

	if nn.Alias != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.Alias)
	}

	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtTraitUsePrecedence(n ast.Vertex) {
	nn := n.(*ast.StmtTraitUsePrecedence)

	p.Print(nn.Ref)
	io.WriteString(p.w, " insteadof ")
	p.joinPrint(", ", nn.Insteadof)

	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtTraitUse(n ast.Vertex) {
	nn := n.(*ast.StmtTraitUse)

	io.WriteString(p.w, "use ")
	p.joinPrint(", ", nn.Traits)

	if adaptationList, ok := nn.Adaptations.(*ast.StmtTraitAdaptationList); ok {
		adaptations := adaptationList.Adaptations
		io.WriteString(p.w, " {\n")
		p.printNodes(adaptations)
		io.WriteString(p.w, "\n")
		p.printIndent()
		io.WriteString(p.w, "}")
	} else {
		io.WriteString(p.w, ";")
	}
}

func (p *PrettyPrinter) printStmtTrait(n ast.Vertex) {
	nn := n.(*ast.StmtTrait)

	io.WriteString(p.w, "trait ")
	p.Print(nn.TraitName)

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "{\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printStmtTry(n ast.Vertex) {
	nn := n.(*ast.StmtTry)

	io.WriteString(p.w, "try {\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")

	if nn.Catches != nil {
		io.WriteString(p.w, "\n")
		p.indentDepth--
		p.printNodes(nn.Catches)
		p.indentDepth++
	}

	if nn.Finally != nil {
		io.WriteString(p.w, "\n")
		p.printIndent()
		p.Print(nn.Finally)
	}
}

func (p *PrettyPrinter) printStmtUnset(n ast.Vertex) {
	nn := n.(*ast.StmtUnset)

	io.WriteString(p.w, "unset(")
	p.joinPrint(", ", nn.Vars)
	io.WriteString(p.w, ");")
}

func (p *PrettyPrinter) printStmtUse(n ast.Vertex) {
	nn := n.(*ast.StmtUse)

	io.WriteString(p.w, "use ")

	if nn.Type != nil {
		p.Print(nn.Type)
		io.WriteString(p.w, " ")
	}

	p.joinPrint(", ", nn.UseDeclarations)

	io.WriteString(p.w, ";")
}

func (p *PrettyPrinter) printStmtGroupUse(n ast.Vertex) {
	nn := n.(*ast.StmtGroupUse)

	io.WriteString(p.w, "use ")

	if nn.Type != nil {
		p.Print(nn.Type)
		io.WriteString(p.w, " ")
	}

	p.Print(nn.Prefix)

	io.WriteString(p.w, "\\{")
	p.joinPrint(", ", nn.UseDeclarations)
	io.WriteString(p.w, "}")
}

func (p *PrettyPrinter) printStmtUseDeclaration(n ast.Vertex) {
	nn := n.(*ast.StmtUseDeclaration)

	if nn.Type != nil {
		p.Print(nn.Type)
		io.WriteString(p.w, " ")
	}

	p.Print(nn.Use)

	if nn.Alias != nil {
		io.WriteString(p.w, " as ")
		p.Print(nn.Alias)
	}
}

func (p *PrettyPrinter) printStmtWhile(n ast.Vertex) {
	nn := n.(*ast.StmtWhile)

	if nn.Alt {
		p.printStmtAltWhile(nn)
		return
	}

	io.WriteString(p.w, "while (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *ast.StmtNop:
		p.Print(s)
		break
	case *ast.StmtStmtList:
		io.WriteString(p.w, " ")
		p.Print(s)
	default:
		io.WriteString(p.w, "\n")
		p.indentDepth++
		p.printIndent()
		p.Print(s)
		p.indentDepth--
	}
}

func (p *PrettyPrinter) printStmtAltWhile(n ast.Vertex) {
	nn := n.(*ast.StmtWhile)

	io.WriteString(p.w, "while (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ") :\n")

	s := nn.Stmt.(*ast.StmtStmtList)
	p.printNodes(s.Stmts)

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "endwhile;")
}
