package printer

import (
	"bytes"
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
	w         io.Writer
	s         printerState
	bufStart  string
	lastWrite []byte
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

func (p *Printer) write(b []byte) {
	p.lastWrite = b
	p.w.Write(b)
}

func (p *Printer) Print(n ast.Vertex) {
	_, isRoot := n.(*ast.Root)
	_, isInlineHtml := n.(*ast.StmtInlineHtml)
	if p.s == HtmlState && !isInlineHtml && !isRoot {
		if n.GetNode().Tokens.IsEmpty() {
			p.bufStart = "<?php "
		}
		p.SetState(PhpState)
	}

	p.printNode(n)
}

func (p *Printer) joinPrint(glue string, nn []ast.Vertex) {
	for k, n := range nn {
		if k > 0 {
			p.write([]byte(glue))
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

func (p *Printer) printSeparatedList(nodeList []ast.Vertex, separatorList []*token.Token, def string) {
	var separators []*token.Token

	if cap(separatorList) >= len(nodeList) {
		separators = separatorList[:len(nodeList)]
	} else {
		separators = make([]*token.Token, len(nodeList))
		copy(separators, separatorList)
	}

	for k, n := range nodeList {
		p.Print(n)
		if k < len(nodeList)-1 {
			p.printToken(separators[k], def)
		} else {
			p.printToken(separators[k], "")
		}
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
		p.write([]byte(def))
		return
	}

	for _, m := range n.GetNode().Tokens[pos] {
		p.write(m.Value)
	}
}

func (p *Printer) printToken(t *token.Token, def string) {
	if t != nil {
		p.write(t.Skipped)
		p.write(t.Value)
		p.bufStart = ""
		return
	}

	if def != "" {
		p.write([]byte(p.bufStart))
		p.bufStart = ""

		p.write([]byte(def))
		return
	}
}

func (p *Printer) printFreeFloating(n ast.Vertex, pos token.Position) {
	if n == nil {
		return
	}

	for _, m := range n.GetNode().Tokens[pos] {
		p.write(m.Value)
	}
}

func (p *Printer) printNode(n ast.Vertex) {
	switch n := n.(type) {

	// node

	case *ast.Root:
		p.printNodeRoot(n)
	case *ast.Identifier:
		p.printNodeIdentifier(n)
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

	p.write(nn.Value)

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
		p.write([]byte("="))
		p.Print(nn.DefaultValue)
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printNodeNullable(n ast.Vertex) {
	nn := n.(*ast.Nullable)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("?"))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printNodeArgument(n ast.Vertex) {
	nn := n.(*ast.Argument)
	p.printFreeFloating(nn, token.Start)

	if nn.AmpersandTkn != nil {
		p.write([]byte("&"))
	}
	p.printFreeFloating(nn, token.Ampersand)

	if nn.VariadicTkn != nil {
		p.write([]byte("..."))
	}

	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

// name

func (p *Printer) printNameNamePart(n *ast.NameNamePart) {
	p.printToken(n.StringTkn, string(n.Value))
}

func (p *Printer) printNameName(n *ast.NameName) {
	p.printFreeFloating(n, token.Start)

	p.joinPrintRefactored("\\", n.Parts)
}

func (p *Printer) printNameFullyQualified(n *ast.NameFullyQualified) {
	p.printFreeFloating(n, token.Start)
	p.printToken(n.NsSeparatorTkn, "\\")

	p.joinPrintRefactored("\\", n.Parts)
}

func (p *Printer) printNameRelative(n *ast.NameRelative) {
	p.printFreeFloating(n, token.Start)
	p.printToken(n.NsTkn, "namespace")
	p.printToken(n.NsSeparatorTkn, "\\")

	p.joinPrintRefactored("\\", n.Parts)
}

// scalar

func (p *Printer) printScalarLNumber(n ast.Vertex) {
	nn := n.(*ast.ScalarLnumber)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	p.write(nn.Value)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarDNumber(n ast.Vertex) {
	nn := n.(*ast.ScalarDnumber)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	p.write(nn.Value)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarString(n ast.Vertex) {
	nn := n.(*ast.ScalarString)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	p.write(nn.Value)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarEncapsedStringPart(n ast.Vertex) {
	nn := n.(*ast.ScalarEncapsedStringPart)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	p.write(nn.Value)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarEncapsed(n ast.Vertex) {
	nn := n.(*ast.ScalarEncapsed)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	p.write([]byte("\""))
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
	p.write([]byte("\""))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarHeredoc(n ast.Vertex) {
	nn := n.(*ast.ScalarHeredoc)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	p.write(nn.OpenHeredocTkn.Value)

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

	p.write([]byte(strings.Trim(string(nn.OpenHeredocTkn.Value), "<\"'\n")))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printScalarMagicConstant(n ast.Vertex) {
	nn := n.(*ast.ScalarMagicConstant)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	p.write(nn.Value)

	p.printFreeFloating(nn, token.End)
}

// Assign

func (p *Printer) printAssign(n ast.Vertex) {
	nn := n.(*ast.ExprAssign)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignReference(n ast.Vertex) {
	nn := n.(*ast.ExprAssignReference)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("="))
	p.printFreeFloating(nn, token.Equal)
	p.write([]byte("&"))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignBitwiseAnd(n ast.Vertex) {
	nn := n.(*ast.ExprAssignBitwiseAnd)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("&"))
	p.write([]byte("="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignBitwiseOr(n ast.Vertex) {
	nn := n.(*ast.ExprAssignBitwiseOr)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("|="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignBitwiseXor(n ast.Vertex) {
	nn := n.(*ast.ExprAssignBitwiseXor)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("^="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignCoalesce(n ast.Vertex) {
	nn := n.(*ast.ExprAssignCoalesce)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("??="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignConcat(n ast.Vertex) {
	nn := n.(*ast.ExprAssignConcat)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte(".="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignDiv(n ast.Vertex) {
	nn := n.(*ast.ExprAssignDiv)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("/="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignMinus(n ast.Vertex) {
	nn := n.(*ast.ExprAssignMinus)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("-="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignMod(n ast.Vertex) {
	nn := n.(*ast.ExprAssignMod)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("%="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignMul(n ast.Vertex) {
	nn := n.(*ast.ExprAssignMul)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("*="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignPlus(n ast.Vertex) {
	nn := n.(*ast.ExprAssignPlus)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("+="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignPow(n ast.Vertex) {
	nn := n.(*ast.ExprAssignPow)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("**="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignShiftLeft(n ast.Vertex) {
	nn := n.(*ast.ExprAssignShiftLeft)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("<<="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printAssignShiftRight(n ast.Vertex) {
	nn := n.(*ast.ExprAssignShiftRight)
	p.printFreeFloating(nn, token.Start)
	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte(">>="))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

// binary

func (p *Printer) printBinaryBitwiseAnd(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBitwiseAnd)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("&"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryBitwiseOr(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBitwiseOr)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("|"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryBitwiseXor(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBitwiseXor)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("^"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryBooleanAnd(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBooleanAnd)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("&&"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryBooleanOr(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryBooleanOr)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("||"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryCoalesce(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryCoalesce)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("??"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryConcat(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryConcat)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("."))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryDiv(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryDiv)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("/"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryEqual)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("=="))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryGreaterOrEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryGreaterOrEqual)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte(">="))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryGreater(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryGreater)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte(">"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryIdentical(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryIdentical)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("==="))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryLogicalAnd(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryLogicalAnd)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.write([]byte("and"))
	if nn.Right.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
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
		p.write([]byte(" "))
	}
	p.write([]byte("or"))
	if nn.Right.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
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
		p.write([]byte(" "))
	}
	p.write([]byte("xor"))
	if nn.Right.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryMinus(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryMinus)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("-"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryMod(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryMod)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("%"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryMul(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryMul)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("*"))
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
		p.write([]byte("!="))
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryNotIdentical(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryNotIdentical)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("!=="))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryPlus(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryPlus)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("+"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryPow(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryPow)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("**"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryShiftLeft(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryShiftLeft)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("<<"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinaryShiftRight(n ast.Vertex) {
	nn := n.(*ast.ExprBinaryShiftRight)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte(">>"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinarySmallerOrEqual(n ast.Vertex) {
	nn := n.(*ast.ExprBinarySmallerOrEqual)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("<="))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinarySmaller(n ast.Vertex) {
	nn := n.(*ast.ExprBinarySmaller)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("<"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBinarySpaceship(n ast.Vertex) {
	nn := n.(*ast.ExprBinarySpaceship)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, token.Expr)
	p.write([]byte("<=>"))
	p.Print(nn.Right)

	p.printFreeFloating(nn, token.End)
}

// cast

func (p *Printer) printArray(n ast.Vertex) {
	nn := n.(*ast.ExprCastArray)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("(array)"))
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printBool(n ast.Vertex) {
	nn := n.(*ast.ExprCastBool)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("(boolean)"))
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printDouble(n ast.Vertex) {
	nn := n.(*ast.ExprCastDouble)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("(float)"))
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printInt(n ast.Vertex) {
	nn := n.(*ast.ExprCastInt)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("(integer)"))
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printObject(n ast.Vertex) {
	nn := n.(*ast.ExprCastObject)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("(object)"))
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printString(n ast.Vertex) {
	nn := n.(*ast.ExprCastString)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("(string)"))
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printUnset(n ast.Vertex) {
	nn := n.(*ast.ExprCastUnset)
	p.printFreeFloating(nn, token.Start)

	p.printFreeFloating(nn, token.Cast)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("(unset)"))
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
		p.write([]byte("["))
	}
	p.Print(nn.Dim)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("]"))
	}
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprArrayDimFetchWithoutLeadingDollar(n ast.Vertex) {
	nn := n.(*ast.ExprArrayDimFetch)
	p.printFreeFloating(nn, token.Start)
	p.printExprVariableWithoutLeadingDollar(nn.Var)
	p.printFreeFloating(nn, token.Var)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("["))
	}
	p.Print(nn.Dim)
	p.printFreeFloating(nn, token.Expr)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte("]"))
	}
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprArrayItem(n ast.Vertex) {
	nn := n.(*ast.ExprArrayItem)
	p.printFreeFloating(nn, token.Start)

	if nn.EllipsisTkn != nil {
		p.write([]byte("..."))
	}

	if nn.Key != nil {
		p.Print(nn.Key)
		p.printFreeFloating(nn, token.Expr)
		p.write([]byte("=>"))
	}

	p.Print(nn.Val)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprArray(n ast.Vertex) {
	nn := n.(*ast.ExprArray)
	p.printFreeFloating(nn, token.Start)
	p.write([]byte("array"))
	p.printFreeFloating(nn, token.Array)
	p.write([]byte("("))
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, token.ArrayPairList)
	p.write([]byte(")"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprArrowFunction(n ast.Vertex) {
	nn := n.(*ast.ExprArrowFunction)
	p.printFreeFloating(nn, token.Start)

	if nn.StaticTkn != nil {
		p.write([]byte("static"))
	}
	p.printFreeFloating(nn, token.Static)
	if nn.StaticTkn != nil && n.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}

	p.write([]byte("fn"))
	p.printFreeFloating(nn, token.Function)

	if nn.AmpersandTkn != nil {
		p.write([]byte("&"))
	}
	p.printFreeFloating(nn, token.Ampersand)

	p.write([]byte("("))
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, token.ParameterList)
	p.write([]byte(")"))
	p.printFreeFloating(nn, token.Params)

	if nn.ReturnType != nil {
		p.bufStart = ":"
		p.Print(nn.ReturnType)
	}
	p.printFreeFloating(nn, token.ReturnType)

	p.write([]byte("=>"))

	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprBitwiseNot(n ast.Vertex) {
	nn := n.(*ast.ExprBitwiseNot)
	p.printFreeFloating(nn, token.Start)
	p.write([]byte("~"))
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprBooleanNot(n ast.Vertex) {
	nn := n.(*ast.ExprBooleanNot)
	p.printFreeFloating(nn, token.Start)
	p.write([]byte("!"))
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprClassConstFetch(n ast.Vertex) {
	nn := n.(*ast.ExprClassConstFetch)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Class)
	p.printFreeFloating(nn, token.Name)
	p.write([]byte("::"))
	p.Print(nn.ConstantName)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprClone(n ast.Vertex) {
	nn := n.(*ast.ExprClone)
	p.printFreeFloating(nn, token.Start)
	p.write([]byte("clone"))
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprClosureUse(n ast.Vertex) {
	nn := n.(*ast.ExprClosureUse)
	p.printFreeFloating(nn, token.Start)
	p.write([]byte("use"))
	p.printFreeFloating(nn, token.Use)
	p.write([]byte("("))
	p.joinPrint(",", nn.Uses)
	p.printFreeFloating(nn, token.LexicalVarList)
	p.write([]byte(")"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprClosure(n ast.Vertex) {
	nn := n.(*ast.ExprClosure)
	p.printFreeFloating(nn, token.Start)

	if nn.StaticTkn != nil {
		p.write([]byte("static"))
	}
	p.printFreeFloating(nn, token.Static)
	if nn.StaticTkn != nil && n.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}

	p.write([]byte("function"))
	p.printFreeFloating(nn, token.Function)

	if nn.AmpersandTkn != nil {
		p.write([]byte("&"))
	}
	p.printFreeFloating(nn, token.Ampersand)

	p.write([]byte("("))
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, token.ParameterList)
	p.write([]byte(")"))
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

	p.write([]byte("{"))
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	p.write([]byte("}"))

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
	p.write([]byte("empty"))

	if _, ok := nn.Expr.(*ast.ParserBrackets); !ok {
		p.write([]byte("("))
	}

	p.Print(nn.Expr)

	if _, ok := nn.Expr.(*ast.ParserBrackets); !ok {
		p.write([]byte(")"))
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprErrorSuppress(n ast.Vertex) {
	nn := n.(*ast.ExprErrorSuppress)
	p.printFreeFloating(nn, token.Start)
	p.write([]byte("@"))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprEval(n ast.Vertex) {
	nn := n.(*ast.ExprEval)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("eval"))

	if _, ok := nn.Expr.(*ast.ParserBrackets); !ok {
		p.write([]byte("("))
	}

	p.Print(nn.Expr)

	if _, ok := nn.Expr.(*ast.ParserBrackets); !ok {
		p.write([]byte(")"))
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprExit(n ast.Vertex) {
	nn := n.(*ast.ExprExit)
	p.printFreeFloating(nn, token.Start)

	if nn.DieTkn != nil {
		p.write(nn.DieTkn.Value)
	} else {
		p.write([]byte("exit"))
	}

	if nn.Expr != nil && nn.Expr.GetNode().Tokens.IsEmpty() && nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprFunctionCall(n ast.Vertex) {
	nn := n.(*ast.ExprFunctionCall)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Function)

	p.printToken(nn.OpenParenthesisTkn, "(")
	p.joinPrint(",", nn.Arguments)
	p.printToken(nn.CloseParenthesisTkn, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprInclude(n ast.Vertex) {
	nn := n.(*ast.ExprInclude)
	p.printFreeFloating(nn, token.Start)
	p.write([]byte("include"))
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprIncludeOnce(n ast.Vertex) {
	nn := n.(*ast.ExprIncludeOnce)
	p.printFreeFloating(nn, token.Start)
	p.write([]byte("include_once"))
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
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
		p.write([]byte(" "))
	}

	p.write([]byte("instanceof"))

	p.bufStart = " "
	p.Print(nn.Class)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprIsset(n ast.Vertex) {
	nn := n.(*ast.ExprIsset)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("isset"))
	p.printFreeFloating(nn, token.Isset)
	p.write([]byte("("))
	p.joinPrint(",", nn.Vars)
	p.printFreeFloating(nn, token.VarList)
	p.write([]byte(")"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprList(n ast.Vertex) {
	nn := n.(*ast.ExprList)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("list"))
	p.printFreeFloating(nn, token.List)
	p.write([]byte("("))
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, token.ArrayPairList)
	p.write([]byte(")"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprMethodCall(n ast.Vertex) {
	nn := n.(*ast.ExprMethodCall)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("->"))
	p.Print(nn.Method)

	p.printToken(nn.OpenParenthesisTkn, "(")
	p.joinPrint(",", nn.Arguments)
	p.printToken(nn.CloseParenthesisTkn, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprNew(n ast.Vertex) {
	nn := n.(*ast.ExprNew)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("new"))
	p.bufStart = " "
	p.Print(nn.Class)

	if nn.Arguments != nil {
		p.printToken(nn.OpenParenthesisTkn, "(")
		p.joinPrint(",", nn.Arguments)
		p.printToken(nn.CloseParenthesisTkn, ")")
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPostDec(n ast.Vertex) {
	nn := n.(*ast.ExprPostDec)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("--"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPostInc(n ast.Vertex) {
	nn := n.(*ast.ExprPostInc)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("++"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPreDec(n ast.Vertex) {
	nn := n.(*ast.ExprPreDec)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("--"))
	p.Print(nn.Var)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPreInc(n ast.Vertex) {
	nn := n.(*ast.ExprPreInc)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("++"))
	p.Print(nn.Var)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPrint(n ast.Vertex) {
	nn := n.(*ast.ExprPrint)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("print"))
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprPropertyFetch(n ast.Vertex) {
	nn := n.(*ast.ExprPropertyFetch)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)
	p.printFreeFloating(nn, token.Var)
	p.write([]byte("->"))
	p.Print(nn.Property)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprReference(n ast.Vertex) {
	nn := n.(*ast.ExprReference)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("&"))
	p.Print(nn.Var)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprRequire(n ast.Vertex) {
	nn := n.(*ast.ExprRequire)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("require"))
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprRequireOnce(n ast.Vertex) {
	nn := n.(*ast.ExprRequireOnce)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("require_once"))
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprShellExec(n ast.Vertex) {
	nn := n.(*ast.ExprShellExec)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("`"))
	p.joinPrint("", nn.Parts)
	p.write([]byte("`"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprStaticCall(n ast.Vertex) {
	nn := n.(*ast.ExprStaticCall)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Class)
	p.printFreeFloating(nn, token.Name)
	p.write([]byte("::"))
	p.Print(nn.Call)

	p.printToken(nn.OpenParenthesisTkn, "(")
	p.joinPrint(",", nn.Arguments)
	p.printToken(nn.CloseParenthesisTkn, ")")

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprStaticPropertyFetch(n ast.Vertex) {
	nn := n.(*ast.ExprStaticPropertyFetch)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Class)
	p.printFreeFloating(nn, token.Name)
	p.write([]byte("::"))
	p.Print(nn.Property)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprTernary(n ast.Vertex) {
	nn := n.(*ast.ExprTernary)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Condition)
	p.printFreeFloating(nn, token.Cond)
	p.write([]byte("?"))

	if nn.IfTrue != nil {
		p.Print(nn.IfTrue)
	}
	p.printFreeFloating(nn, token.True)

	p.write([]byte(":"))
	p.Print(nn.IfFalse)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprUnaryMinus(n ast.Vertex) {
	nn := n.(*ast.ExprUnaryMinus)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("-"))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprUnaryPlus(n ast.Vertex) {
	nn := n.(*ast.ExprUnaryPlus)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("+"))
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprVariable(n ast.Vertex) {
	nn := n.(*ast.ExprVariable)
	p.printFreeFloatingOrDefault(nn, token.Start, p.bufStart)
	p.bufStart = ""

	if _, ok := nn.VarName.(*ast.Identifier); !ok {
		p.write([]byte("$"))
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

	p.write([]byte("yield from"))
	if nn.Expr.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printExprYield(n ast.Vertex) {
	nn := n.(*ast.ExprYield)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("yield"))

	if nn.Key != nil {
		if nn.Key.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
		p.Print(nn.Key)
		p.printFreeFloating(nn, token.Expr)
		p.write([]byte("=>"))
	} else {
		if nn.Value.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
	}

	p.Print(nn.Value)

	p.printFreeFloating(nn, token.End)
}

// smtm

func (p *Printer) printStmtBreak(n *ast.StmtBreak) {
	p.printToken(n.BreakTkn, "break")

	if n.Expr != nil {
		p.bufStart = " "
	}

	p.Print(n.Expr)
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtCase(n *ast.StmtCase) {
	p.printToken(n.CaseTkn, "case")
	p.bufStart = " "
	p.Print(n.Cond)
	p.printToken(n.CaseSeparatorTkn, ":")
	p.printNodes(n.Stmts)
}

func (p *Printer) printStmtCatch(n *ast.StmtCatch) {
	p.printToken(n.CatchTkn, "catch")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.printSeparatedList(n.Types, n.SeparatorTkns, "|")
	p.Print(n.Var)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.OpenCurlyBracketTkn, "{")
	p.printNodes(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, "}")
}

func (p *Printer) printStmtClassMethod(n ast.Vertex) {
	nn := n.(*ast.StmtClassMethod)
	p.printFreeFloating(nn, token.Start)

	if nn.Modifiers != nil {
		for k, m := range nn.Modifiers {
			if k > 0 && m.GetNode().Tokens.IsEmpty() {
				p.write([]byte(" "))
			}
			p.Print(m)
		}

		if nn.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
	}
	p.printFreeFloating(nn, token.ModifierList)
	p.write([]byte("function"))
	p.printFreeFloating(nn, token.Function)

	if nn.AmpersandTkn != nil {
		if nn.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
		p.write([]byte("&"))
		p.printFreeFloating(nn, token.Ampersand)
	} else {
		if nn.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
	}

	p.Print(nn.MethodName)
	p.printFreeFloating(nn, token.Name)
	p.write([]byte("("))
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, token.ParameterList)
	p.write([]byte(")"))
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
				p.write([]byte(" "))
			}
			p.Print(m)
		}

		if nn.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
	}
	p.printFreeFloating(nn, token.ModifierList)
	p.write([]byte("class"))

	if nn.ClassName != nil {
		p.bufStart = " "
		p.Print(nn.ClassName)
	}

	if nn.Arguments != nil {
		p.printToken(nn.OpenParenthesisTkn, "(")
		p.joinPrint(",", nn.Arguments)
		p.printToken(nn.CloseParenthesisTkn, ")")
	}

	if nn.Extends != nil {
		p.printFreeFloating(nn.Extends, token.Start)
		if nn.Extends.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
		p.write([]byte("extends"))
		p.bufStart = " "
		p.Print(nn.Extends.(*ast.StmtClassExtends).ClassName)
	}

	if nn.Implements != nil {
		p.printFreeFloating(nn.Implements, token.Start)
		if nn.Implements.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
		p.write([]byte("implements"))
		p.bufStart = " "
		p.joinPrintRefactored(",", nn.Implements.(*ast.StmtClassImplements).InterfaceNames)

	}

	p.printFreeFloating(nn, token.Name)
	p.write([]byte("{"))
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	p.write([]byte("}"))

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
	p.printSeparatedList(n.Consts, n.SeparatorTkns, ",")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtConstant(n *ast.StmtConstant) {
	p.Print(n.Name)
	p.printToken(n.EqualTkn, "=")
	p.Print(n.Expr)
}

func (p *Printer) printStmtContinue(n *ast.StmtContinue) {
	p.printToken(n.ContinueTkn, "continue")

	if n.Expr != nil {
		p.bufStart = " "
	}

	p.Print(n.Expr)
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtDeclare(n *ast.StmtDeclare) {
	if n.Alt {
		p.printStmtAltDeclare(n)
		return
	}
	p.printToken(n.DeclareTkn, "declare")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.printSeparatedList(n.Consts, n.SeparatorTkns, ",")
	p.printToken(n.CloseParenthesisTkn, ")")
	p.Print(n.Stmt)
}

func (p *Printer) printStmtAltDeclare(n *ast.StmtDeclare) {
	p.printToken(n.DeclareTkn, "declare")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.printSeparatedList(n.Consts, n.SeparatorTkns, ",")
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.ColonTkn, ":")

	if stmtList, ok := n.Stmt.(*ast.StmtStmtList); ok {
		p.printNodes(stmtList.Stmts)
	} else {
		p.Print(n.Stmt)
	}

	p.printToken(n.EndDeclareTkn, "enddeclare")
	p.printToken(n.SemiColonTkn, ";")
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

func (p *Printer) printStmtEcho(n *ast.StmtEcho) {
	p.printToken(n.EchoTkn, "echo")
	p.bufStart = " "
	p.printSeparatedList(n.Exprs, n.SeparatorTkns, ",")
	p.printToken(n.SemiColonTkn, ";")
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
		p.write([]byte(";"))
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtFinally(n *ast.StmtFinally) {
	p.printToken(n.FinallyTkn, "finally")
	p.printToken(n.OpenCurlyBracketTkn, "{")
	p.printNodes(n.Stmts)
	p.printToken(n.CloseCurlyBracketTkn, "}")
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
		p.Print(n.Stmt)
	}

	p.printToken(n.EndForTkn, "endfor")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtForeach(n *ast.StmtForeach) {
	if n.Alt {
		p.printStmtAltForeach(n)
		return
	}

	p.printToken(n.ForeachTkn, "foreach")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Expr)
	p.bufStart = " "
	p.printToken(n.AsTkn, "as")
	p.bufStart = " "
	if n.Key != nil {
		p.Print(n.Key)
		p.printToken(n.DoubleArrowTkn, "=>")
	}
	p.Print(n.Var)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.Print(n.Stmt)
}

func (p *Printer) printStmtAltForeach(n *ast.StmtForeach) {
	p.printToken(n.ForeachTkn, "foreach")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.Print(n.Expr)
	p.bufStart = " "
	p.printToken(n.AsTkn, "as")
	p.bufStart = " "
	if n.Key != nil {
		p.Print(n.Key)
		p.printToken(n.DoubleArrowTkn, "=>")
	}
	p.Print(n.Var)
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.ColonTkn, ":")

	if stmtList, ok := n.Stmt.(*ast.StmtStmtList); ok {
		p.printNodes(stmtList.Stmts)
	} else {
		p.Print(n.Stmt)
	}

	p.printToken(n.EndForeachTkn, "endforeach")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtFunction(n ast.Vertex) {
	nn := n.(*ast.StmtFunction)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("function"))
	p.printFreeFloating(nn, token.Function)

	if nn.AmpersandTkn != nil {
		if nn.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
		p.write([]byte("&"))
	} else {
		if nn.FunctionName.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
	}

	p.Print(nn.FunctionName)
	p.printFreeFloating(nn, token.Name)

	p.write([]byte("("))
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, token.ParamList)
	p.write([]byte(")"))
	p.printFreeFloating(nn, token.Params)

	if nn.ReturnType != nil {
		p.bufStart = ":"
		p.Print(nn.ReturnType)
	}
	p.printFreeFloating(nn, token.ReturnType)

	p.write([]byte("{"))
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	p.write([]byte("}"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtGlobal(n *ast.StmtGlobal) {
	p.printToken(n.GlobalTkn, "global")
	p.bufStart = " "
	p.printSeparatedList(n.Vars, n.SeparatorTkns, ",")
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtGoto(n *ast.StmtGoto) {
	p.printToken(n.GotoTkn, "goto")
	p.bufStart = " "
	p.Print(n.Label)
	p.printToken(n.SemiColonTkn, ";")
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

func (p *Printer) printStmtInlineHTML(n *ast.StmtInlineHtml) {
	if p.s == PhpState && !bytes.Contains(p.lastWrite, []byte("?>")) {
		p.write([]byte("?>"))
	}
	p.SetState(HtmlState)

	p.printToken(n.InlineHtmlTkn, string(n.Value))
}

func (p *Printer) printStmtInterface(n ast.Vertex) {
	nn := n.(*ast.StmtInterface)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("interface"))

	if nn.InterfaceName.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}

	p.Print(nn.InterfaceName)

	if nn.Extends != nil {
		p.printFreeFloating(nn.Extends, token.Start)
		if nn.Extends.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
		p.write([]byte("extends"))
		p.bufStart = " "
		p.joinPrintRefactored(",", nn.Extends.(*ast.StmtInterfaceExtends).InterfaceNames)
	}

	p.printFreeFloating(nn, token.Name)
	p.write([]byte("{"))
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	p.write([]byte("}"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtLabel(n *ast.StmtLabel) {
	p.Print(n.LabelName)
	p.printToken(n.ColonTkn, ":")
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

func (p *Printer) printStmtNop(n *ast.StmtNop) {
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtPropertyList(n ast.Vertex) {
	nn := n.(*ast.StmtPropertyList)
	p.printFreeFloating(nn, token.Start)

	for k, m := range nn.Modifiers {
		if k > 0 && m.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
		p.Print(m)
	}

	if nn.Type != nil {
		p.bufStart = " "
		p.Print(nn.Type)
	}

	if nn.Properties[0].GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}

	p.joinPrint(",", nn.Properties)
	p.printFreeFloating(n, token.PropertyList)

	p.printFreeFloating(n, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		p.write([]byte(";"))
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtProperty(n ast.Vertex) {
	nn := n.(*ast.StmtProperty)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Var)

	if nn.Expr != nil {
		p.printFreeFloating(nn, token.Var)
		p.write([]byte("="))
		p.Print(nn.Expr)
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtReturn(n *ast.StmtReturn) {
	p.printToken(n.ReturnTkn, "return")

	if n.Expr != nil {
		p.bufStart = " "
	}
	p.Print(n.Expr)

	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtStaticVar(n *ast.StmtStaticVar) {
	p.Print(n.Var)

	if n.Expr != nil {
		p.printToken(n.EqualTkn, "=")
		p.Print(n.Expr)
	}
}

func (p *Printer) printStmtStatic(n *ast.StmtStatic) {
	p.printToken(n.StaticTkn, "static")
	p.bufStart = " "
	p.printSeparatedList(n.Vars, n.SeparatorTkns, ",")
	p.printToken(n.SemiColonTkn, ";")
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

func (p *Printer) printStmtThrow(n *ast.StmtThrow) {
	p.printToken(n.ThrowTkn, "throw")
	p.bufStart = " "
	p.Print(n.Expr)
	p.printToken(n.SemiColonTkn, ";")
}

func (p *Printer) printStmtTraitAdaptationList(n ast.Vertex) {
	nn := n.(*ast.StmtTraitAdaptationList)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("{"))
	p.printNodes(nn.Adaptations)
	p.printFreeFloating(nn, token.AdaptationList)
	p.write([]byte("}"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTraitMethodRef(n ast.Vertex) {
	nn := n.(*ast.StmtTraitMethodRef)
	p.printFreeFloating(nn, token.Start)

	if nn.Trait != nil {
		p.Print(nn.Trait)
		p.printFreeFloating(nn, token.Name)
		p.write([]byte("::"))
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
		p.write([]byte(" "))
	}
	p.write([]byte("as"))

	if nn.Modifier != nil {
		if nn.Modifier.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
		p.Print(nn.Modifier)
	}

	if nn.Alias != nil {
		if nn.Alias.GetNode().Tokens.IsEmpty() {
			p.write([]byte(" "))
		}
		p.Print(nn.Alias)
	}
	p.printFreeFloating(nn, token.Alias)

	p.printFreeFloating(nn, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		p.write([]byte(";"))
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTraitUsePrecedence(n ast.Vertex) {
	nn := n.(*ast.StmtTraitUsePrecedence)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Ref)
	p.printFreeFloating(nn, token.Ref)
	if nn.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}

	p.write([]byte("insteadof"))
	p.bufStart = " "
	p.joinPrint(",", nn.Insteadof)
	p.printFreeFloating(nn, token.NameList)

	p.printFreeFloating(nn, token.SemiColon)
	if n.GetNode().Tokens.IsEmpty() {
		p.write([]byte(";"))
	}

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTraitUse(n ast.Vertex) {
	nn := n.(*ast.StmtTraitUse)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("use"))
	p.bufStart = " "
	p.joinPrintRefactored(",", nn.Traits)

	p.Print(nn.Adaptations)

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTrait(n ast.Vertex) {
	nn := n.(*ast.StmtTrait)
	p.printFreeFloating(nn, token.Start)

	p.write([]byte("trait"))
	if nn.TraitName.GetNode().Tokens.IsEmpty() {
		p.write([]byte(" "))
	}
	p.Print(nn.TraitName)

	p.printFreeFloating(nn, token.Name)
	p.write([]byte("{"))
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, token.Stmts)
	p.write([]byte("}"))

	p.printFreeFloating(nn, token.End)
}

func (p *Printer) printStmtTry(n *ast.StmtTry) {
	p.printToken(n.TryTkn, "try")
	p.printToken(n.OpenCurlyBracket, "{")
	p.printNodes(n.Stmts)
	p.printToken(n.CloseCurlyBracket, "}")

	if n.Catches != nil {
		p.printNodes(n.Catches)
	}

	if n.Finally != nil {
		p.Print(n.Finally)
	}
}

func (p *Printer) printStmtUnset(n *ast.StmtUnset) {
	p.printToken(n.UnsetTkn, "unset")
	p.printToken(n.OpenParenthesisTkn, "(")
	p.printSeparatedList(n.Vars, n.SeparatorTkns, ",")
	p.printToken(n.CloseParenthesisTkn, ")")
	p.printToken(n.SemiColonTkn, ";")
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
		return
	}

	p.bufStart = " "
	p.printToken(n.AsTkn, "as")

	p.bufStart = " "
	p.Print(n.Alias)
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

func (p *Printer) printParserBrackets(n ast.Vertex) {
	nn := n.(*ast.ParserBrackets)
	p.printFreeFloating(nn, token.Start)

	p.Print(nn.Child)

	p.printFreeFloating(nn, token.End)
}
