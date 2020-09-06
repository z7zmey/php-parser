package printer

import (
	"io"
	"strings"

	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/token"
)

type printerState int

const (
	PhpState printerState = iota
	HtmlState
)

type Printer struct {
	w        io.Writer
	s        printerState
	bufStart string
}

// NewPrinter - Constructor for Printer
func NewPrinter(w io.Writer) *Printer {
	return &Printer{
		w: w,
	}
}

func (p *Printer) SetState(s printerState) {
	p.s = s
}

func (p *Printer) Print(n ast.Vertex) {
	_, isRoot := n.(*ast.Root)
	_, isInlineHtml := n.(*ast.StmtInlineHtml)
	if p.s == HtmlState && !isInlineHtml && !isRoot {
		if n.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, "<?php ")
		}
		p.SetState(PhpState)
	}

	p.printNode(n)
}

func (p *Printer) joinPrint(glue string, nn []ast.Vertex) {
	for k, n := range nn {
		if k > 0 {
			io.WriteString(p.w, glue)
		}

		p.Print(n)
	}
}

func (p *Printer) joinPrintRefactored(glue string, nn []ast.Vertex) {
	for k, n := range nn {
		if k > 0 {
			p.bufStart = glue
		}

		p.Print(n)
	}
}

func (p *Printer) printNodes(nn []ast.Vertex) {
	for _, n := range nn {
		p.Print(n)
	}
}

func (p *Printer) printFreeFloatingOrDefault(n ast.Vertex, pos token.Position, def string) {
	if n == nil {
		return
	}

	if len(n.GetNode().Tokens[pos]) == 0 {
		io.WriteString(p.w, def)
		return
	}

	for _, m := range n.GetNode().Tokens[pos] {
		io.WriteString(p.w, string(m.Value))
	}
}

func (p *Printer) printToken(t *token.Token, def string) {
	if t != nil {
		p.w.Write(t.Skipped)
		p.w.Write(t.Value)
		p.bufStart = ""
		return
	}

	if def != "" {
		p.w.Write([]byte(p.bufStart))
		p.bufStart = ""

		p.w.Write([]byte(def))
		return
	}
}

func (p *Printer) printFreeFloating(n ast.Vertex, pos token.Position) {
	if n == nil {
		return
	}

	for _, m := range n.GetNode().Tokens[pos] {
		io.WriteString(p.w, string(m.Value))
	}
}

func (p *Printer) printNode(n ast.Vertex) {
	switch n := n.(type) {

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
		p.printAssign(n)
	case *ast.ExprAssignReference:
		p.printAssignReference(n)
	case *ast.ExprAssignBitwiseAnd:
		p.printAssignBitwiseAnd(n)
	case *ast.ExprAssignBitwiseOr:
		p.printAssignBitwiseOr(n)
	case *ast.ExprAssignBitwiseXor:
		p.printAssignBitwiseXor(n)
	case *ast.ExprAssignCoalesce:
		p.printAssignCoalesce(n)
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
	case *ast.ExprArrowFunction:
		p.printExprArrowFunction(n)
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
	case *ast.ExprShortArray:
		p.printExprShortArray(n)
	case *ast.ExprShortList:
		p.printExprShortList(n)
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

	case *ast.StmtAltForeach:
		p.printStmtAltForeach(n)
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
	case *ast.StmtConstList:
		p.printStmtConstList(n)
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
	case *ast.StmtTraitAdaptationList:
		p.printStmtTraitAdaptationList(n)
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
	case *ast.ParserAs:
		p.printParserAs(n)
	case *ast.ParserNsSeparator:
		p.printParserNsSeparator(n)
	case *ast.ParserBrackets:
		p.printParserBrackets(n)
	}
}

// node

func (p *Printer) printNodeRoot(n ast.Vertex) {
	nn := n.(*ast.Root)
	p.SetState(HtmlState)
	p.printFreeFloating(nn, token.Start)
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printNodeIdentifier(n ast.Vertex) {
	nn := n.(*ast.Identifier)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	io.WriteString(p.w, string(nn.Value))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printNodeReference(n ast.Vertex) {
	nn := n.(*ast.Reference)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "&")
	p.Print(nn.Var)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printNodeVariadic(n ast.Vertex) {
	nn := n.(*ast.Variadic)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "...")
	p.Print(nn.Var)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printNodeParameter(n ast.Vertex) {
	nn := n.(*ast.Parameter)
	p.printFreeFloating(nn, token.Start)

	if nn.Type != nil {
		p.Print(nn.Type)
	}

	p.Print(nn.Var)

	if nn.DefaultValue != nil {
		io.WriteString(p.w, "=")
		p.Print(nn.DefaultValue)
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printNodeNullable(n ast.Vertex) {
	nn := n.(*ast.Nullable)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "?")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printNodeArgument(n ast.Vertex) {
	nn := n.(*ast.Argument)
	p.printFreeFloating(nn, token.Start)

	if nn.IsReference {
		io.WriteString(p.w, "&")
	}
	p.printFreeFloating(nn, token.Ampersand)

	if nn.Variadic {
		io.WriteString(p.w, "...")
	}

	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

// name

func (p *Printer) printNameNamePart(n *ast.NameNamePart) {
	p.printToken(n.NsSeparatorTkn, "")
	p.printToken(n.StringTkn, string(n.Value))
}

func (p *Printer) printNameName(n *ast.NameName) {
	p.printFreeFloating(n, token.Start)

	p.joinPrintRefactored("\\", n.Parts)

	p.printToken(n.ListSeparatorTkn, "")
}

func (p *Printer) printNameFullyQualified(n *ast.NameFullyQualified) {
	p.printFreeFloating(n, token.Start)
	p.printToken(n.NsSeparatorTkn, "\\")

	p.joinPrintRefactored("\\", n.Parts)

	p.printToken(n.ListSeparatorTkn, "")
}

func (p *Printer) printNameRelative(n *ast.NameRelative) {
	p.printFreeFloating(n, token.Start)
	p.printToken(n.NsTkn, "namespace")
	p.printToken(n.NsSeparatorTkn, "\\")

	p.joinPrintRefactored("\\", n.Parts)

	p.printToken(n.ListSeparatorTkn, "")
}

// scalar

func (p *Printer) printScalarLNumber(n ast.Vertex) {
	nn := n.(*ast.ScalarLnumber)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	io.WriteString(p.w, string(nn.Value))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarDNumber(n ast.Vertex) {
	nn := n.(*ast.ScalarDnumber)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	io.WriteString(p.w, string(nn.Value))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarString(n ast.Vertex) {
	nn := n.(*ast.ScalarString)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	io.WriteString(p.w, string(nn.Value))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarEncapsedStringPart(n ast.Vertex) {
	nn := n.(*ast.ScalarEncapsedStringPart)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	io.WriteString(p.w, string(nn.Value))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarEncapsed(n ast.Vertex) {
	nn := n.(*ast.ScalarEncapsed)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	io.WriteString(p.w, "\"")
	for _, part := range nn.Parts {
		switch part.(type) {
		case *ast.ExprArrayDimFetch:
			s := part.GetNode().Tokens[token.Start]
			if len(s) > 0 && string(s[0].Value) == "${" {
				p.printExprArrayDimFetchWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		case *ast.ExprVariable:
			s := part.GetNode().Tokens[token.Start]
			if len(s) > 0 && string(s[0].Value) == "${" {
				p.printExprVariableWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		default:
			p.Print(part)
		}
	}
	io.WriteString(p.w, "\"")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarHeredoc(n ast.Vertex) {
	nn := n.(*ast.ScalarHeredoc)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	io.WriteString(p.w, string(nn.Label))

	for _, part := range nn.Parts {
		switch part.(type) {
		case *ast.ExprArrayDimFetch:
			s := part.GetNode().Tokens[token.Start]
			if len(s) > 0 && string(s[0].Value) == "${" {
				p.printExprArrayDimFetchWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		case *ast.ExprVariable:
			s := part.GetNode().Tokens[token.Start]
			if len(s) > 0 && string(s[0].Value) == "${" {
				p.printExprVariableWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		default:
			p.Print(part)
		}
	}

	io.WriteString(p.w, strings.Trim(string(nn.Label), "<\"'\n"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarMagicConstant(n ast.Vertex) {
	nn := n.(*ast.ScalarMagicConstant)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	io.WriteString(p.w, string(nn.Value))

	p.printFreeFloating(nn, token.End)
}

// Assign

func (p *Printer) printAssign(n ast.Vertex) {
	nn := n.(*ast.ExprAssign)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignReference(n ast.Vertex) {
	nn := n.(*ast.ExprAssignReference)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "=")
	p.printFreeFloating(nn, token.Equal)
	io.WriteString(p.w, "&")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignBitwiseAnd(n ast.Vertex) {
	nn := n.(*ast.ExprAssignBitwiseAnd)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "&")
	io.WriteString(p.w, "=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignBitwiseOr(n ast.Vertex) {
	nn := n.(*ast.ExprAssignBitwiseOr)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "|=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignBitwiseXor(n ast.Vertex) {
	nn := n.(*ast.ExprAssignBitwiseXor)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "^=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignCoalesce(n ast.Vertex) {
	nn := n.(*ast.ExprAssignCoalesce)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "??=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignConcat(n ast.Vertex) {
	nn := n.(*ast.ExprAssignConcat)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, ".=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignDiv(n ast.Vertex) {
	nn := n.(*ast.ExprAssignDiv)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "/=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignMinus(n ast.Vertex) {
	nn := n.(*ast.ExprAssignMinus)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "-=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignMod(n ast.Vertex) {
	nn := n.(*ast.ExprAssignMod)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "%=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignMul(n ast.Vertex) {
	nn := n.(*ast.ExprAssignMul)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "*=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignPlus(n ast.Vertex) {
	nn := n.(*ast.ExprAssignPlus)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "+=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignPow(n ast.Vertex) {
	nn := n.(*ast.ExprAssignPow)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "**=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignShiftLeft(n ast.Vertex) {
	nn := n.(*ast.ExprAssignShiftLeft)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "<<=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignShiftRight(n ast.Vertex) {
	nn := n.(*ast.ExprAssignShiftRight)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, ">>=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

// binary

func (p *Printer) printBinaryBitwiseAnd(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBitwiseAnd)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "&")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryBitwiseOr(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBitwiseOr)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "|")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryBitwiseXor(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBitwiseXor)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "^")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryBooleanAnd(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBooleanAnd)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "&&")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryBooleanOr(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBooleanOr)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "||")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryCoalesce(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryCoalesce)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "??")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryConcat(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryConcat)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, ".")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryDiv(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryDiv)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "/")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryEqual)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "==")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryGreaterOrEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryGreaterOrEqual)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, ">=")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryGreater(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryGreater)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, ">")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryIdentical(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryIdentical)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "===")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryLogicalAnd(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryLogicalAnd)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "and")
	if nn.Right.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryLogicalOr(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryLogicalOr)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "or")
	if nn.Right.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryLogicalXor(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryLogicalXor)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "xor")
	if nn.Right.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryMinus(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryMinus)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "-")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryMod(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryMod)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "%")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryMul(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryMul)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "*")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryNotEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryNotEqual)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.printFreeFloating(nn, token.Equal)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "!=")
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryNotIdentical(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryNotIdentical)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "!==")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryPlus(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryPlus)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "+")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryPow(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryPow)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "**")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryShiftLeft(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryShiftLeft)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "<<")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryShiftRight(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryShiftRight)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, ">>")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinarySmallerOrEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinarySmallerOrEqual)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "<=")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinarySmaller(n ast.Vertex) {
	nn := n.(*ast.ExprBinarySmaller)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "<")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinarySpaceship(n ast.Vertex) {
	nn := n.(*ast.ExprBinarySpaceship)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	io.WriteString(p.w, "<=>")
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

// cast

func (p *Printer) printArray(n ast.Vertex) {
	nn := n.(*ast.ExprCastArray)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "(array)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBool(n ast.Vertex) {
	nn := n.(*ast.ExprCastBool)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "(boolean)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printDouble(n ast.Vertex) {
	nn := n.(*ast.ExprCastDouble)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "(float)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printInt(n ast.Vertex) {
	nn := n.(*ast.ExprCastInt)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "(integer)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printObject(n ast.Vertex) {
	nn := n.(*ast.ExprCastObject)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "(object)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printString(n ast.Vertex) {
	nn := n.(*ast.ExprCastString)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "(string)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printUnset(n ast.Vertex) {
	nn := n.(*ast.ExprCastUnset)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "(unset)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

// expr

func (p *Printer) printExprArrayDimFetch(n ast.Vertex) {
	nn := n.(*ast.ExprArrayDimFetch)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "[")
	}
	p.Print(nn.Dim)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "]")
	}
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprArrayDimFetchWithoutLeadingDollar(n ast.Vertex) {
	nn := n.(*ast.ExprArrayDimFetch)
	p.printFreeFloating(nn, token.Start)
	p.printExprVariableWithoutLeadingDollar(nn.Var)
	p.printFreeFloating(nn, token.Var)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "[")
	}
	p.Print(nn.Dim)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "]")
	}
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprArrayItem(n ast.Vertex) {
	nn := n.(*ast.ExprArrayItem)
	p.printFreeFloating(nn, token.Start)

	if nn.Unpack {
		io.WriteString(p.w, "...")
	}

	if nn.Key != nil {
		p.Print(nn.Key)
		p.printFreeFloating(nn, token.Expr)
		io.WriteString(p.w, "=>")
	}

	p.Print(nn.Val)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprArray(n ast.Vertex) {
	nn := n.(*ast.ExprArray)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "array")
	p.printFreeFloating(nn, token.Array)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, token.ArrayPairList)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprArrowFunction(n ast.Vertex) {
	nn := n.(*ast.ExprArrowFunction)
	p.printFreeFloating(nn, token.Start)

	if nn.Static {
		io.WriteString(p.w, "static")
	}
	p.printFreeFloating(nn, token.Static)
	if nn.Static && n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "fn")
	p.printFreeFloating(nn, token.Function)

	if nn.ReturnsRef {
		io.WriteString(p.w, "&")
	}
	p.printFreeFloating(nn, token.Ampersand)

	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, token.ParameterList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, token.Params)

	if nn.ReturnType != nil {
		p.bufStart = ":"
		p.Print(nn.ReturnType)
	}
	p.printFreeFloating(nn, token.ReturnType)

	io.WriteString(p.w, "=>")

	p.printNode(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprBitwiseNot(n ast.Vertex) {
	nn := n.(*ast.ExprBitwiseNot)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "~")
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprBooleanNot(n ast.Vertex) {
	nn := n.(*ast.ExprBooleanNot)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "!")
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprClassConstFetch(n ast.Vertex) {
	nn := n.(*ast.ExprClassConstFetch)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Class)
	p.printFreeFloating(nn, token.Name)
	io.WriteString(p.w, "::")
	p.Print(nn.ConstantName)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprClone(n ast.Vertex) {
	nn := n.(*ast.ExprClone)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "clone")
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprClosureUse(n ast.Vertex) {
	nn := n.(*ast.ExprClosureUse)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "use")
	p.printFreeFloating(nn, token.Use)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Uses)
	p.printFreeFloating(nn, token.LexicalVarList)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprClosure(n ast.Vertex) {
	nn := n.(*ast.ExprClosure)
	p.printFreeFloating(nn, token.Start)

	if nn.Static {
		io.WriteString(p.w, "static")
	}
	p.printFreeFloating(nn, token.Static)
	if nn.Static && n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "function")
	p.printFreeFloating(nn, token.Function)

	if nn.ReturnsRef {
		io.WriteString(p.w, "&")
	}
	p.printFreeFloating(nn, token.Ampersand)

	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, token.ParameterList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, token.Params)

	if nn.ClosureUse != nil {
		p.Print(nn.ClosureUse)
	}
	p.printFreeFloating(nn, token.LexicalVars)

	if nn.ReturnType != nil {
		p.bufStart = ":"
		p.Print(nn.ReturnType)
	}
	p.printFreeFloating(nn, token.ReturnType)

	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprConstFetch(n ast.Vertex) {
	nn := n.(*ast.ExprConstFetch)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Const)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprEmpty(n ast.Vertex) {
	nn := n.(*ast.ExprEmpty)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "empty")

	if _, ok := nn.Expr.(*ast.ParserBrackets); !ok {
		io.WriteString(p.w, "(")
	}

	p.Print(nn.Expr)

	if _, ok := nn.Expr.(*ast.ParserBrackets); !ok {
		io.WriteString(p.w, ")")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprErrorSuppress(n ast.Vertex) {
	nn := n.(*ast.ExprErrorSuppress)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "@")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprEval(n ast.Vertex) {
	nn := n.(*ast.ExprEval)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "eval")

	if _, ok := nn.Expr.(*ast.ParserBrackets); !ok {
		io.WriteString(p.w, "(")
	}

	p.Print(nn.Expr)

	if _, ok := nn.Expr.(*ast.ParserBrackets); !ok {
		io.WriteString(p.w, ")")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprExit(n ast.Vertex) {
	nn := n.(*ast.ExprExit)
	p.printFreeFloating(nn, token.Start)

	if nn.Die {
		io.WriteString(p.w, "die")
	} else {
		io.WriteString(p.w, "exit")
	}

	if nn.Expr != nil && nn.Expr.GetNode().Tokens.IsEmpty() && nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprFunctionCall(n ast.Vertex) {
	nn := n.(*ast.ExprFunctionCall)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Function)

	p.printFreeFloatingOrDefault(nn.ArgumentList, token.Start, "(")
	p.joinPrint(",", nn.ArgumentList.Arguments)
	p.printFreeFloatingOrDefault(nn.ArgumentList, token.End, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprInclude(n ast.Vertex) {
	nn := n.(*ast.ExprInclude)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "include")
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprIncludeOnce(n ast.Vertex) {
	nn := n.(*ast.ExprIncludeOnce)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "include_once")
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprInstanceOf(n ast.Vertex) {
	nn := n.(*ast.ExprInstanceOf)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "instanceof")

	p.bufStart = " "
	p.Print(nn.Class)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprIsset(n ast.Vertex) {
	nn := n.(*ast.ExprIsset)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "isset")
	p.printFreeFloating(nn, token.Isset)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Vars)
	p.printFreeFloating(nn, token.VarList)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprList(n ast.Vertex) {
	nn := n.(*ast.ExprList)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "list")
	p.printFreeFloating(nn, token.List)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, token.ArrayPairList)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprMethodCall(n ast.Vertex) {
	nn := n.(*ast.ExprMethodCall)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "->")
	p.Print(nn.Method)

	p.printFreeFloatingOrDefault(nn.ArgumentList, token.Start, "(")
	p.joinPrint(",", nn.ArgumentList.Arguments)
	p.printFreeFloatingOrDefault(nn.ArgumentList, token.End, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprNew(n ast.Vertex) {
	nn := n.(*ast.ExprNew)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "new")
	p.bufStart = " "
	p.Print(nn.Class)

	if nn.ArgumentList != nil {
		p.printFreeFloatingOrDefault(nn.ArgumentList, token.Start, "(")
		p.joinPrint(",", nn.ArgumentList.Arguments)
		p.printFreeFloatingOrDefault(nn.ArgumentList, token.End, ")")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPostDec(n ast.Vertex) {
	nn := n.(*ast.ExprPostDec)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "--")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPostInc(n ast.Vertex) {
	nn := n.(*ast.ExprPostInc)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "++")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPreDec(n ast.Vertex) {
	nn := n.(*ast.ExprPreDec)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "--")
	p.Print(nn.Var)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPreInc(n ast.Vertex) {
	nn := n.(*ast.ExprPreInc)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "++")
	p.Print(nn.Var)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPrint(n ast.Vertex) {
	nn := n.(*ast.ExprPrint)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "print")
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPropertyFetch(n ast.Vertex) {
	nn := n.(*ast.ExprPropertyFetch)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, "->")
	p.Print(nn.Property)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprReference(n ast.Vertex) {
	nn := n.(*ast.ExprReference)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "&")
	p.Print(nn.Var)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprRequire(n ast.Vertex) {
	nn := n.(*ast.ExprRequire)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "require")
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprRequireOnce(n ast.Vertex) {
	nn := n.(*ast.ExprRequireOnce)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "require_once")
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprShellExec(n ast.Vertex) {
	nn := n.(*ast.ExprShellExec)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "`")
	p.joinPrint("", nn.Parts)
	io.WriteString(p.w, "`")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprShortArray(n ast.Vertex) {
	nn := n.(*ast.ExprShortArray)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "[")
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, token.ArrayPairList)
	io.WriteString(p.w, "]")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprShortList(n ast.Vertex) {
	nn := n.(*ast.ExprShortList)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "[")
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, token.ArrayPairList)
	io.WriteString(p.w, "]")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprStaticCall(n ast.Vertex) {
	nn := n.(*ast.ExprStaticCall)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Class)
	p.printFreeFloating(nn, token.Name)
	io.WriteString(p.w, "::")
	p.Print(nn.Call)

	p.printFreeFloatingOrDefault(nn.ArgumentList, token.Start, "(")
	p.joinPrint(",", nn.ArgumentList.Arguments)
	p.printFreeFloatingOrDefault(nn.ArgumentList, token.End, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprStaticPropertyFetch(n ast.Vertex) {
	nn := n.(*ast.ExprStaticPropertyFetch)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Class)
	p.printFreeFloating(nn, token.Name)
	io.WriteString(p.w, "::")
	p.Print(nn.Property)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprTernary(n ast.Vertex) {
	nn := n.(*ast.ExprTernary)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Condition)
	p.printFreeFloating(nn, token.Cond)
	io.WriteString(p.w, "?")

	if nn.IfTrue != nil {
		p.Print(nn.IfTrue)
	}
	p.printFreeFloating(nn, token.True)

	io.WriteString(p.w, ":")
	p.Print(nn.IfFalse)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprUnaryMinus(n ast.Vertex) {
	nn := n.(*ast.ExprUnaryMinus)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "-")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprUnaryPlus(n ast.Vertex) {
	nn := n.(*ast.ExprUnaryPlus)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "+")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprVariable(n ast.Vertex) {
	nn := n.(*ast.ExprVariable)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	if _, ok := nn.VarName.(*ast.Identifier); !ok {
		io.WriteString(p.w, "$")
	}

	p.Print(nn.VarName)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprVariableWithoutLeadingDollar(n ast.Vertex) {
	nn := n.(*ast.ExprVariable)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.VarName)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprYieldFrom(n ast.Vertex) {
	nn := n.(*ast.ExprYieldFrom)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "yield from")
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprYield(n ast.Vertex) {
	nn := n.(*ast.ExprYield)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "yield")

	if nn.Key != nil {
		if nn.Key.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Key)
		p.printFreeFloating(nn, token.Expr)
		io.WriteString(p.w, "=>")
	} else {
		if nn.Value.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Value)

	p.printFreeFloating(nn, token.End)
}

// smtm

func (p *Printer) printStmtAltForeach(n ast.Vertex) {
	nn := n.(*ast.StmtAltForeach)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "foreach")
	p.printFreeFloating(nn, token.Foreach)
	io.WriteString(p.w, "(")
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "as")

	if nn.Key != nil {
		if nn.Key.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Key)
		p.printFreeFloating(nn, token.Key)
		io.WriteString(p.w, "=>")
	} else {
		if nn.Var.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)

	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, token.Cond)

	io.WriteString(p.w, ":")
	s := nn.Stmt.(*ast.StmtStmtList)
	p.printNodes(s.Stmts)
	p.printFreeFloating(nn, token.Stmts)

	io.WriteString(p.w, "endforeach")
	p.printFreeFloating(nn, token.AltEnd)
	p.printFreeFloating(nn, token.SemiColon)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtBreak(n ast.Vertex) {
	nn := n.(*ast.StmtBreak)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "break")
	if nn.Expr != nil {
		if nn.Expr.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Expr)
	}
	p.printFreeFloating(nn, token.Expr)

	p.printFreeFloating(nn, token.SemiColon)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtCase(n *ast.StmtCase) {
	p.printToken(n.CaseTkn, "case")
	p.bufStart = " "
	p.Print(n.Cond)
	p.printToken(n.CaseSeparatorTkn, ":")
	p.printNodes(n.Stmts)
}

func (p *Printer) printStmtCatch(n ast.Vertex) {
	nn := n.(*ast.StmtCatch)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "catch")
	p.printFreeFloating(nn, token.Catch)
	io.WriteString(p.w, "(")

	p.joinPrintRefactored("|", nn.Types)

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, token.Cond)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtClassMethod(n ast.Vertex) {
	nn := n.(*ast.StmtClassMethod)
	p.printFreeFloating(nn, token.Start)

	if nn.Modifiers != nil {
		for k, m := range nn.Modifiers {
			if k > 0 && m.GetNode().Tokens.IsEmpty() {
				io.WriteString(p.w, " ")
			}
			p.Print(m)
		}

		if nn.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}
	p.printFreeFloating(nn, token.ModifierList)
	io.WriteString(p.w, "function")
	p.printFreeFloating(nn, token.Function)

	if nn.ReturnsRef {
		if nn.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "&")
		p.printFreeFloating(nn, token.Ampersand)
	} else {
		if nn.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.MethodName)
	p.printFreeFloating(nn, token.Name)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, token.ParameterList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, token.Params)

	if nn.ReturnType != nil {
		p.bufStart = ":"
		p.Print(nn.ReturnType)
	}

	p.Print(nn.Stmt)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtClass(n ast.Vertex) {
	nn := n.(*ast.StmtClass)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	if nn.Modifiers != nil {
		for k, m := range nn.Modifiers {
			if k > 0 && m.GetNode().Tokens.IsEmpty() {
				io.WriteString(p.w, " ")
			}
			p.Print(m)
		}

		if nn.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}
	p.printFreeFloating(nn, token.ModifierList)
	io.WriteString(p.w, "class")

	if nn.ClassName != nil {
		p.bufStart = " "
		p.Print(nn.ClassName)
	}

	if nn.ArgumentList != nil {
		p.printFreeFloatingOrDefault(nn.ArgumentList, token.Start, "(")
		p.joinPrint(",", nn.ArgumentList.Arguments)
		p.printFreeFloatingOrDefault(nn.ArgumentList, token.End, ")")
	}

	if nn.Extends != nil {
		p.printFreeFloating(nn.Extends, token.Start)
		if nn.Extends.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "extends")
		p.bufStart = " "
		p.Print(nn.Extends.ClassName)
	}

	if nn.Implements != nil {
		p.printFreeFloating(nn.Implements, token.Start)
		if nn.Implements.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "implements")
		p.bufStart = " "
		p.joinPrintRefactored(",", nn.Implements.InterfaceNames)

	}

	p.printFreeFloating(nn, token.Name)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtClassConstList(n *ast.StmtClassConstList) {
	p.joinPrintRefactored(" ", n.Modifiers)
	p.bufStart = " "
	p.printToken(n.ConstTkn, "const")
	p.bufStart = " "
	p.joinPrintRefactored(",", n.Consts)
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtConstList(n *ast.StmtConstList) {
	p.printToken(n.ConstTkn, "const")
	p.bufStart = " "
	p.joinPrintRefactored(",", n.Consts)
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtConstant(n *ast.StmtConstant) {
	p.Print(n.Name)
	p.printToken(n.EqualTkn, "=")
	p.Print(n.Expr)
	p.printToken(n.CommaTkn, "")
}

func (p *Printer) printStmtContinue(n ast.Vertex) {
	nn := n.(*ast.StmtContinue)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "continue")

	if nn.Expr != nil {
		if nn.Expr.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Expr)
	}
	p.printFreeFloating(nn, token.Expr)

	p.printFreeFloating(nn, token.SemiColon)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtDeclare(n ast.Vertex) {
	nn := n.(*ast.StmtDeclare)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "declare")
	p.printFreeFloating(nn, token.Declare)
	io.WriteString(p.w, "(")
	p.joinPrintRefactored(",", nn.Consts)
	p.printFreeFloating(nn, token.ConstList)
	io.WriteString(p.w, ")")

	if nn.Alt {
		p.printFreeFloating(nn, token.Cond)
		io.WriteString(p.w, ":")

		s := nn.Stmt.(*ast.StmtStmtList)
		p.printNodes(s.Stmts)
		p.printFreeFloating(nn, token.Stmts)

		io.WriteString(p.w, "enddeclare")
		p.printFreeFloating(nn, token.AltEnd)

		p.printFreeFloating(nn, token.SemiColon)
		if nn.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, ";")
		}
	} else {
		p.Print(nn.Stmt)
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtDefault(n *ast.StmtDefault) {
	p.printToken(n.DefaultTkn, "default")
	p.printToken(n.CaseSeparatorTkn, ":")
	p.printNodes(n.Stmts)
}

func (p *Printer) printStmtDo(n *ast.StmtDo) {
	p.printToken(n.DoTkn, "do")
	p.bufStart = " "

	p.Print(n.Stmt)

	p.printToken(n.WhileTkn, "while")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Cond)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtEcho(n ast.Vertex) {
	nn := n.(*ast.StmtEcho)

	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "echo")
	}
	if nn.Exprs[0].GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}

	p.printFreeFloating(nn, token.Start)
	p.printFreeFloating(nn, token.Echo)

	p.joinPrint(",", nn.Exprs)
	p.printFreeFloating(nn, token.Expr)

	p.printFreeFloating(nn, token.SemiColon)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtElseif(n *ast.StmtElseIf) {
	if n.Alt {
		p.printStmtAltElseIf(n)
		return
	}

	p.printToken(n.ElseIfTkn, "elseif")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Cond)
	p.printToken(n.CloseParenthesisTkn, ")")

	p.Print(n.Stmt)
}

func (p *Printer) printStmtAltElseIf(n *ast.StmtElseIf) {
	p.printToken(n.ElseIfTkn, "elseif")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Cond)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.ColonTkn, ":")

	if stmtList, ok := n.Stmt.(*ast.StmtStmtList); ok {
		p.printNodes(stmtList.Stmts)
	} else {
		p.Print(n.Stmt)
	}
}

func (p *Printer) printStmtElse(n *ast.StmtElse) {
	if n.Alt {
		p.printStmtAltElse(n)
		return
	}

	p.printToken(n.ElseTkn, "else")
	p.bufStart = " "
	p.Print(n.Stmt)
}

func (p *Printer) printStmtAltElse(n *ast.StmtElse) {
	p.printToken(n.ElseTkn, "else")
	p.printToken(n.ColonTkn, ":")

	if stmtList, ok := n.Stmt.(*ast.StmtStmtList); ok {
		p.printNodes(stmtList.Stmts)
	} else {
		p.Print(n.Stmt)
	}
}

func (p *Printer) printStmtExpression(n ast.Vertex) {
	nn := n.(*ast.StmtExpression)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.Expr)

	p.printFreeFloating(nn, token.SemiColon)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtFinally(n ast.Vertex) {
	nn := n.(*ast.StmtFinally)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "finally")
	p.printFreeFloating(nn, token.Finally)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtFor(n *ast.StmtFor) {
	if n.Alt {
		p.printStmtAltFor(n)
		return
	}

	p.printToken(n.ForTkn, "for")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.joinPrint(",", n.Init)
	p.printToken(n.InitSemiColonTkn, ";")
	p.joinPrint(",", n.Cond)
	p.printToken(n.CondSemiColonTkn, ";")
	p.joinPrint(",", n.Loop)
	p.printToken(n.CloseParenthesisTkn, ")")

	p.Print(n.Stmt)
}

func (p *Printer) printStmtAltFor(n *ast.StmtFor) {
	p.printToken(n.ForTkn, "for")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.joinPrint(",", n.Init)
	p.printToken(n.InitSemiColonTkn, ";")
	p.joinPrint(",", n.Cond)
	p.printToken(n.CondSemiColonTkn, ";")
	p.joinPrint(",", n.Loop)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.ColonTkn, ":")

	if stmtList, ok := n.Stmt.(*ast.StmtStmtList); ok {
		p.printNodes(stmtList.Stmts)
	} else {
		p.printNode(n.Stmt)
	}

	p.printToken(n.EndForTkn, "endfor")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtForeach(n ast.Vertex) {
	nn := n.(*ast.StmtForeach)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "foreach")
	p.printFreeFloating(nn, token.Foreach)
	io.WriteString(p.w, "(")

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "as")

	if nn.Key != nil {
		if nn.Key.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Key)
		p.printFreeFloating(nn, token.Key)
		io.WriteString(p.w, "=>")
	} else {
		if nn.Var.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)

	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtFunction(n ast.Vertex) {
	nn := n.(*ast.StmtFunction)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "function")
	p.printFreeFloating(nn, token.Function)

	if nn.ReturnsRef {
		if nn.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "&")
	} else {
		if nn.FunctionName.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.FunctionName)
	p.printFreeFloating(nn, token.Name)

	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, token.ParamList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, token.Params)

	if nn.ReturnType != nil {
		p.bufStart = ":"
		p.Print(nn.ReturnType)
	}
	p.printFreeFloating(nn, token.ReturnType)

	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtGlobal(n ast.Vertex) {
	nn := n.(*ast.StmtGlobal)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "global")
	p.joinPrint(",", nn.Vars)
	p.printFreeFloating(nn, token.VarList)

	p.printFreeFloating(nn, token.SemiColon)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtGoto(n ast.Vertex) {
	nn := n.(*ast.StmtGoto)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "goto")
	if nn.Label.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Label)
	p.printFreeFloating(nn, token.Label)

	p.printFreeFloating(nn, token.SemiColon)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtHaltCompiler(n *ast.StmtHaltCompiler) {
	p.printToken(n.HaltCompilerTkn, "__halt_compiler")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtIf(n *ast.StmtIf) {
	if n.Alt {
		p.printStmtAltIf(n)
		return
	}

	p.printToken(n.IfTkn, "if")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Cond)
	p.printToken(n.CloseParenthesisTkn, ")")

	p.Print(n.Stmt)
	p.printNodes(n.ElseIf)
	p.Print(n.Else)
}

func (p *Printer) printStmtAltIf(n *ast.StmtIf) {
	p.printToken(n.IfTkn, "if")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Cond)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.ColonTkn, ":")

	if stmtList, ok := n.Stmt.(*ast.StmtStmtList); ok {
		p.printNodes(stmtList.Stmts)
	} else {
		p.Print(n.Stmt)
	}

	p.printNodes(n.ElseIf)
	p.Print(n.Else)

	p.printToken(n.EndIfTkn, "endif")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtInlineHTML(n ast.Vertex) {
	nn := n.(*ast.StmtInlineHtml)
	p.printFreeFloating(nn, token.Start)

	if p.s == PhpState && nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, "?>")
	}
	p.SetState(HtmlState)

	io.WriteString(p.w, string(nn.Value))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtInterface(n ast.Vertex) {
	nn := n.(*ast.StmtInterface)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "interface")

	if nn.InterfaceName.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}

	p.Print(nn.InterfaceName)

	if nn.Extends != nil {
		p.printFreeFloating(nn.Extends, token.Start)
		if nn.Extends.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "extends")
		p.bufStart = " "
		p.joinPrintRefactored(",", nn.Extends.InterfaceNames)
	}

	p.printFreeFloating(nn, token.Name)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtLabel(n ast.Vertex) {
	nn := n.(*ast.StmtLabel)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.LabelName)
	p.printFreeFloating(nn, token.Label)

	io.WriteString(p.w, ":")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtNamespace(n *ast.StmtNamespace) {
	p.printToken(n.NsTkn, "namespace")

	if n.Name != nil {
		p.bufStart = " "
		p.Print(n.Name)
	}

	if n.Stmts != nil {
		p.printToken(n.OpenCurlyBracket, "{")
		p.printNodes(n.Stmts)
		p.printToken(n.CloseCurlyBracket, "}")
		return
	}

	if n.OpenCurlyBracket != nil {
		p.printToken(n.OpenCurlyBracket, "{")
		p.printToken(n.CloseCurlyBracket, "}")
		return
	}

	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtNop(n ast.Vertex) {
	p.printFreeFloatingOrDefault(n, token.Start, p.bufStart)
	p.printFreeFloating(n, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}
	p.printFreeFloating(n, token.End)
}

func (p *Printer) printStmtPropertyList(n ast.Vertex) {
	nn := n.(*ast.StmtPropertyList)
	p.printFreeFloating(nn, token.Start)

	for k, m := range nn.Modifiers {
		if k > 0 && m.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(m)
	}

	if nn.Type != nil {
		p.bufStart = " "
		p.Print(nn.Type)
	}

	if nn.Properties[0].GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}

	p.joinPrint(",", nn.Properties)
	p.printFreeFloating(n, token.PropertyList)

	p.printFreeFloating(n, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtProperty(n ast.Vertex) {
	nn := n.(*ast.StmtProperty)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)

	if nn.Expr != nil {
		p.printFreeFloating(nn, token.Var)
		io.WriteString(p.w, "=")
		p.Print(nn.Expr)
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtReturn(n ast.Vertex) {
	nn := n.(*ast.StmtReturn)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "return")
	p.bufStart = " "
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.Expr)

	p.printFreeFloating(nn, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtStaticVar(n ast.Vertex) {
	nn := n.(*ast.StmtStaticVar)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)

	if nn.Expr != nil {
		p.printFreeFloating(nn, token.Var)
		io.WriteString(p.w, "=")
		p.Print(nn.Expr)
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtStatic(n ast.Vertex) {
	nn := n.(*ast.StmtStatic)
	p.printFreeFloating(nn, token.Start)
	io.WriteString(p.w, "static")

	p.joinPrint(",", nn.Vars)
	p.printFreeFloating(nn, token.VarList)

	p.printFreeFloating(nn, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtStmtList(n *ast.StmtStmtList) {
	p.printToken(n.OpenCurlyBracket, "{")
	p.printNodes(n.Stmts)
	p.printToken(n.CloseCurlyBracket, "}")
}

func (p *Printer) printStmtSwitch(n *ast.StmtSwitch) {
	if n.Alt {
		p.printStmtAltSwitch(n)
		return
	}

	p.printToken(n.SwitchTkn, "switch")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Cond)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.OpenCurlyBracketTkn, "{")
	p.printToken(n.CaseSeparatorTkn, "")
	p.printNodes(n.CaseList)
	p.printToken(n.CloseCurlyBracketTkn, "}")
}

func (p *Printer) printStmtAltSwitch(n *ast.StmtSwitch) {
	p.printToken(n.SwitchTkn, "switch")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Cond)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.ColonTkn, ":")
	p.printToken(n.CaseSeparatorTkn, "")
	p.printNodes(n.CaseList)
	p.printToken(n.EndSwitchTkn, "endswitch")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtThrow(n ast.Vertex) {
	nn := n.(*ast.StmtThrow)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "throw")
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.Expr)

	p.printFreeFloating(nn, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTraitAdaptationList(n ast.Vertex) {
	nn := n.(*ast.StmtTraitAdaptationList)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "{")
	p.printNodes(nn.Adaptations)
	p.printFreeFloating(nn, token.AdaptationList)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTraitMethodRef(n ast.Vertex) {
	nn := n.(*ast.StmtTraitMethodRef)
	p.printFreeFloating(nn, token.Start)

	if nn.Trait != nil {
		p.Print(nn.Trait)
		p.printFreeFloating(nn, token.Name)
		io.WriteString(p.w, "::")
	}

	p.Print(nn.Method)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTraitUseAlias(n ast.Vertex) {
	nn := n.(*ast.StmtTraitUseAlias)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Ref)
	p.printFreeFloating(nn, token.Ref)

	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "as")

	if nn.Modifier != nil {
		if nn.Modifier.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Modifier)
	}

	if nn.Alias != nil {
		if nn.Alias.GetNode().Tokens.IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Alias)
	}
	p.printFreeFloating(nn, token.Alias)

	p.printFreeFloating(nn, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTraitUsePrecedence(n ast.Vertex) {
	nn := n.(*ast.StmtTraitUsePrecedence)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Ref)
	p.printFreeFloating(nn, token.Ref)
	if nn.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "insteadof")
	p.bufStart = " "
	p.joinPrint(",", nn.Insteadof)
	p.printFreeFloating(nn, token.NameList)

	p.printFreeFloating(nn, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTraitUse(n ast.Vertex) {
	nn := n.(*ast.StmtTraitUse)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "use")
	p.bufStart = " "
	p.joinPrintRefactored(",", nn.Traits)

	p.Print(nn.TraitAdaptationList)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTrait(n ast.Vertex) {
	nn := n.(*ast.StmtTrait)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "trait")
	if nn.TraitName.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.TraitName)

	p.printFreeFloating(nn, token.Name)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTry(n ast.Vertex) {
	nn := n.(*ast.StmtTry)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "try")
	p.printFreeFloating(nn, token.Try)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	io.WriteString(p.w, "}")

	if nn.Catches != nil {
		p.printNodes(nn.Catches)
	}

	if nn.Finally != nil {
		p.Print(nn.Finally)
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtUnset(n ast.Vertex) {
	nn := n.(*ast.StmtUnset)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "unset")
	p.printFreeFloating(nn, token.Unset)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Vars)
	p.printFreeFloating(nn, token.VarList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, token.CloseParenthesisToken)

	p.printFreeFloating(nn, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtUse(n *ast.StmtUse) {
	p.printToken(n.UseTkn, "use")

	if n.Type != nil {
		p.bufStart = " "
		p.Print(n.Type)
	}

	p.bufStart = " "
	p.joinPrintRefactored(",", n.UseDeclarations)

	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtGroupUse(n *ast.StmtGroupUse) {
	p.printToken(n.UseTkn, "use")

	p.bufStart = " "
	p.Print(n.Type)

	p.bufStart = " "
	p.printToken(n.LeadingNsSeparatorTkn, "")

	p.Print(n.Prefix)
	p.printToken(n.NsSeparatorTkn, "\\")
	p.printToken(n.OpenCurlyBracketTkn, "{")

	p.joinPrintRefactored(",", n.UseDeclarations)

	p.printToken(n.CloseCurlyBracketTkn, "}")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtUseDeclaration(n *ast.StmtUseDeclaration) {
	p.Print(n.Type)

	if n.Type != nil {
		p.bufStart = " "
	}

	p.printToken(n.NsSeparatorTkn, "")

	p.Print(n.Use)

	if n.Alias == nil {
		p.printToken(n.CommaTkn, "")
		return
	}

	p.bufStart = " "
	p.printToken(n.AsTkn, "as")

	p.bufStart = " "
	p.Print(n.Alias)

	p.printToken(n.CommaTkn, "")
}

func (p *Printer) printStmtWhile(n *ast.StmtWhile) {
	if n.Alt {
		p.printStmtAltWhile(n)
		return
	}

	p.printToken(n.WhileTkn, "while")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Cond)
	p.printToken(n.CloseParenthesisTkn, ")")

	p.Print(n.Stmt)
}

func (p *Printer) printStmtAltWhile(n *ast.StmtWhile) {
	p.printToken(n.WhileTkn, "while")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Cond)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.ColonTkn, ":")

	if stmtList, ok := n.Stmt.(*ast.StmtStmtList); ok {
		p.printNodes(stmtList.Stmts)
	} else {
		p.Print(n.Stmt)
	}

	p.printToken(n.EndWhileTkn, "endwhile")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printParserAs(n ast.Vertex) {
	nn := n.(*ast.ParserAs)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "as")
	p.Print(nn.Child)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printParserNsSeparator(n ast.Vertex) {
	nn := n.(*ast.ParserNsSeparator)
	p.printFreeFloating(nn, token.Start)

	io.WriteString(p.w, "\\")
	p.Print(nn.Child)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printParserBrackets(n ast.Vertex) {
	nn := n.(*ast.ParserBrackets)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Child)

	p.printFreeFloating(nn, token.End)
}
