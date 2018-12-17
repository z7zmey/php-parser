package printer

import (
	"io"
	"strings"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/node/stmt"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/expr/cast"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
)

type printerState int

const (
	PhpState printerState = iota
	HtmlState
)

type Printer struct {
	w io.Writer
	s printerState
}

// NewPrinter -  Constructor for Printer
func NewPrinter(w io.Writer) *Printer {
	return &Printer{
		w: w,
	}
}

func (p *Printer) SetState(s printerState) {
	p.s = s
}

func (p *Printer) Print(n node.Node) {
	_, isRoot := n.(*node.Root)
	_, isInlineHtml := n.(*stmt.InlineHtml)
	_, isEcho := n.(*stmt.Echo)
	if p.s == HtmlState && !isInlineHtml && !isRoot && !isEcho {
		if len((*n.GetMeta())) == 0 {
			io.WriteString(p.w, "<?php")
		}

		if len((*n.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.SetState(PhpState)
	}

	p.printNode(n)
}

func (p *Printer) joinPrint(glue string, nn []node.Node) {
	for k, n := range nn {
		if k > 0 {
			io.WriteString(p.w, glue)
		}

		p.Print(n)
	}
}

func (p *Printer) printNodes(nn []node.Node) {
	for _, n := range nn {
		p.Print(n)
	}
}

func (p *Printer) printMeta(n node.Node, tn meta.TokenName) bool {
	if n == nil {
		return false
	}

	r := false

	for _, m := range *n.GetMeta() {
		if m.TokenName == tn {
			io.WriteString(p.w, m.String())
			r = true
		}
	}

	return r
}

func (p *Printer) printNode(n node.Node) {
	switch n.(type) {

	// node

	case *node.Root:
		p.printNodeRoot(n)
	case *node.Identifier:
		p.printNodeIdentifier(n)
	case *node.Parameter:
		p.printNodeParameter(n)
	case *node.Nullable:
		p.printNodeNullable(n)
	case *node.Argument:
		p.printNodeArgument(n)

	// name

	case *name.NamePart:
		p.printNameNamePart(n)
	case *name.Name:
		p.printNameName(n)
	case *name.FullyQualified:
		p.printNameFullyQualified(n)
	case *name.Relative:
		p.printNameRelative(n)

	// scalar

	case *scalar.Lnumber:
		p.printScalarLNumber(n)
	case *scalar.Dnumber:
		p.printScalarDNumber(n)
	case *scalar.String:
		p.printScalarString(n)
	case *scalar.EncapsedStringPart:
		p.printScalarEncapsedStringPart(n)
	case *scalar.Encapsed:
		p.printScalarEncapsed(n)
	case *scalar.Heredoc:
		p.printScalarHeredoc(n)
	case *scalar.MagicConstant:
		p.printScalarMagicConstant(n)

	// assign

	case *assign.Assign:
		p.printAssign(n)
	case *assign.Reference:
		p.printAssignReference(n)
	case *assign.BitwiseAnd:
		p.printAssignBitwiseAnd(n)
	case *assign.BitwiseOr:
		p.printAssignBitwiseOr(n)
	case *assign.BitwiseXor:
		p.printAssignBitwiseXor(n)
	case *assign.Concat:
		p.printAssignConcat(n)
	case *assign.Div:
		p.printAssignDiv(n)
	case *assign.Minus:
		p.printAssignMinus(n)
	case *assign.Mod:
		p.printAssignMod(n)
	case *assign.Mul:
		p.printAssignMul(n)
	case *assign.Plus:
		p.printAssignPlus(n)
	case *assign.Pow:
		p.printAssignPow(n)
	case *assign.ShiftLeft:
		p.printAssignShiftLeft(n)
	case *assign.ShiftRight:
		p.printAssignShiftRight(n)

	// binary

	case *binary.BitwiseAnd:
		p.printBinaryBitwiseAnd(n)
	case *binary.BitwiseOr:
		p.printBinaryBitwiseOr(n)
	case *binary.BitwiseXor:
		p.printBinaryBitwiseXor(n)
	case *binary.BooleanAnd:
		p.printBinaryBooleanAnd(n)
	case *binary.BooleanOr:
		p.printBinaryBooleanOr(n)
	case *binary.Coalesce:
		p.printBinaryCoalesce(n)
	case *binary.Concat:
		p.printBinaryConcat(n)
	case *binary.Div:
		p.printBinaryDiv(n)
	case *binary.Equal:
		p.printBinaryEqual(n)
	case *binary.GreaterOrEqual:
		p.printBinaryGreaterOrEqual(n)
	case *binary.Greater:
		p.printBinaryGreater(n)
	case *binary.Identical:
		p.printBinaryIdentical(n)
	case *binary.LogicalAnd:
		p.printBinaryLogicalAnd(n)
	case *binary.LogicalOr:
		p.printBinaryLogicalOr(n)
	case *binary.LogicalXor:
		p.printBinaryLogicalXor(n)
	case *binary.Minus:
		p.printBinaryMinus(n)
	case *binary.Mod:
		p.printBinaryMod(n)
	case *binary.Mul:
		p.printBinaryMul(n)
	case *binary.NotEqual:
		p.printBinaryNotEqual(n)
	case *binary.NotIdentical:
		p.printBinaryNotIdentical(n)
	case *binary.Plus:
		p.printBinaryPlus(n)
	case *binary.Pow:
		p.printBinaryPow(n)
	case *binary.ShiftLeft:
		p.printBinaryShiftLeft(n)
	case *binary.ShiftRight:
		p.printBinaryShiftRight(n)
	case *binary.SmallerOrEqual:
		p.printBinarySmallerOrEqual(n)
	case *binary.Smaller:
		p.printBinarySmaller(n)
	case *binary.Spaceship:
		p.printBinarySpaceship(n)

	// cast

	case *cast.Array:
		p.printArray(n)
	case *cast.Bool:
		p.printBool(n)
	case *cast.Double:
		p.printDouble(n)
	case *cast.Int:
		p.printInt(n)
	case *cast.Object:
		p.printObject(n)
	case *cast.String:
		p.printString(n)
	case *cast.Unset:
		p.printUnset(n)

	// expr

	case *expr.ArrayDimFetch:
		p.printExprArrayDimFetch(n)
	case *expr.ArrayItem:
		p.printExprArrayItem(n)
	case *expr.Array:
		p.printExprArray(n)
	case *expr.BitwiseNot:
		p.printExprBitwiseNot(n)
	case *expr.BooleanNot:
		p.printExprBooleanNot(n)
	case *expr.ClassConstFetch:
		p.printExprClassConstFetch(n)
	case *expr.Clone:
		p.printExprClone(n)
	case *expr.ClosureUse:
		p.printExprClosureUse(n)
	case *expr.Closure:
		p.printExprClosure(n)
	case *expr.ConstFetch:
		p.printExprConstFetch(n)
	case *expr.Empty:
		p.printExprEmpty(n)
	case *expr.ErrorSuppress:
		p.printExprErrorSuppress(n)
	case *expr.Eval:
		p.printExprEval(n)
	case *expr.Exit:
		p.printExprExit(n)
	case *expr.FunctionCall:
		p.printExprFunctionCall(n)
	case *expr.Include:
		p.printExprInclude(n)
	case *expr.IncludeOnce:
		p.printExprIncludeOnce(n)
	case *expr.InstanceOf:
		p.printExprInstanceOf(n)
	case *expr.Isset:
		p.printExprIsset(n)
	case *expr.List:
		p.printExprList(n)
	case *expr.MethodCall:
		p.printExprMethodCall(n)
	case *expr.New:
		p.printExprNew(n)
	case *expr.PostDec:
		p.printExprPostDec(n)
	case *expr.PostInc:
		p.printExprPostInc(n)
	case *expr.PreDec:
		p.printExprPreDec(n)
	case *expr.PreInc:
		p.printExprPreInc(n)
	case *expr.Print:
		p.printExprPrint(n)
	case *expr.PropertyFetch:
		p.printExprPropertyFetch(n)
	case *expr.Reference:
		p.printExprReference(n)
	case *expr.Require:
		p.printExprRequire(n)
	case *expr.RequireOnce:
		p.printExprRequireOnce(n)
	case *expr.ShellExec:
		p.printExprShellExec(n)
	case *expr.ShortArray:
		p.printExprShortArray(n)
	case *expr.ShortList:
		p.printExprShortList(n)
	case *expr.StaticCall:
		p.printExprStaticCall(n)
	case *expr.StaticPropertyFetch:
		p.printExprStaticPropertyFetch(n)
	case *expr.Ternary:
		p.printExprTernary(n)
	case *expr.UnaryMinus:
		p.printExprUnaryMinus(n)
	case *expr.UnaryPlus:
		p.printExprUnaryPlus(n)
	case *expr.Variable:
		p.printExprVariable(n)
	case *expr.YieldFrom:
		p.printExprYieldFrom(n)
	case *expr.Yield:
		p.printExprYield(n)

	// stmt

	case *stmt.AltElseIf:
		p.printStmtAltElseIf(n)
	case *stmt.AltElse:
		p.printStmtAltElse(n)
	case *stmt.AltFor:
		p.printStmtAltFor(n)
	case *stmt.AltForeach:
		p.printStmtAltForeach(n)
	case *stmt.AltIf:
		p.printStmtAltIf(n)
	case *stmt.AltSwitch:
		p.printStmtAltSwitch(n)
	case *stmt.AltWhile:
		p.printStmtAltWhile(n)
	case *stmt.Break:
		p.printStmtBreak(n)
	case *stmt.Case:
		p.printStmtCase(n)
	case *stmt.Catch:
		p.printStmtCatch(n)
	case *stmt.ClassMethod:
		p.printStmtClassMethod(n)
	case *stmt.Class:
		p.printStmtClass(n)
	case *stmt.ClassConstList:
		p.printStmtClassConstList(n)
	case *stmt.ConstList:
		p.printStmtConstList(n)
	case *stmt.Constant:
		p.printStmtConstant(n)
	case *stmt.Continue:
		p.printStmtContinue(n)
	case *stmt.Declare:
		p.printStmtDeclare(n)
	case *stmt.Default:
		p.printStmtDefault(n)
	case *stmt.Do:
		p.printStmtDo(n)
	case *stmt.Echo:
		p.printStmtEcho(n)
	case *stmt.ElseIf:
		p.printStmtElseif(n)
	case *stmt.Else:
		p.printStmtElse(n)
	case *stmt.Expression:
		p.printStmtExpression(n)
	case *stmt.Finally:
		p.printStmtFinally(n)
	case *stmt.For:
		p.printStmtFor(n)
	case *stmt.Foreach:
		p.printStmtForeach(n)
	case *stmt.Function:
		p.printStmtFunction(n)
	case *stmt.Global:
		p.printStmtGlobal(n)
	case *stmt.Goto:
		p.printStmtGoto(n)
	case *stmt.GroupUse:
		p.printStmtGroupUse(n)
	case *stmt.HaltCompiler:
		p.printStmtHaltCompiler(n)
	case *stmt.If:
		p.printStmtIf(n)
	case *stmt.InlineHtml:
		p.printStmtInlineHTML(n)
	case *stmt.Interface:
		p.printStmtInterface(n)
	case *stmt.Label:
		p.printStmtLabel(n)
	case *stmt.Namespace:
		p.printStmtNamespace(n)
	case *stmt.Nop:
		p.printStmtNop(n)
	case *stmt.PropertyList:
		p.printStmtPropertyList(n)
	case *stmt.Property:
		p.printStmtProperty(n)
	case *stmt.Return:
		p.printStmtReturn(n)
	case *stmt.StaticVar:
		p.printStmtStaticVar(n)
	case *stmt.Static:
		p.printStmtStatic(n)
	case *stmt.StmtList:
		p.printStmtStmtList(n)
	case *stmt.Switch:
		p.printStmtSwitch(n)
	case *stmt.Throw:
		p.printStmtThrow(n)
	case *stmt.TraitAdaptationList:
		p.printStmtTraitAdaptationList(n)
	case *stmt.TraitMethodRef:
		p.printStmtTraitMethodRef(n)
	case *stmt.TraitUseAlias:
		p.printStmtTraitUseAlias(n)
	case *stmt.TraitUsePrecedence:
		p.printStmtTraitUsePrecedence(n)
	case *stmt.TraitUse:
		p.printStmtTraitUse(n)
	case *stmt.Trait:
		p.printStmtTrait(n)
	case *stmt.Try:
		p.printStmtTry(n)
	case *stmt.Unset:
		p.printStmtUnset(n)
	case *stmt.UseList:
		p.printStmtUseList(n)
	case *stmt.Use:
		p.printStmtUse(n)
	case *stmt.While:
		p.printStmtWhile(n)
	}
}

// node

func (p *Printer) printNodeRoot(n node.Node) {
	nn := n.(*node.Root)
	p.SetState(HtmlState)
	p.printMeta(nn, meta.NodeStart)
	p.printNodes(nn.Stmts)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printNodeIdentifier(n node.Node) {
	nn := n.(*node.Identifier)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, nn.Value)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printNodeParameter(n node.Node) {
	nn := n.(*node.Parameter)
	p.printMeta(nn, meta.NodeStart)

	if nn.VariableType != nil {
		p.Print(nn.VariableType)
	}

	if nn.ByRef {
		p.printMeta(nn, meta.AmpersandToken)
		io.WriteString(p.w, "&")
	}

	if nn.Variadic {
		p.printMeta(nn, meta.EllipsisToken)
		io.WriteString(p.w, "...")
	}

	p.Print(nn.Variable)

	if nn.DefaultValue != nil {
		p.printMeta(nn, meta.EqualToken)
		io.WriteString(p.w, "=")
		p.Print(nn.DefaultValue)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printNodeNullable(n node.Node) {
	nn := n.(*node.Nullable)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "?")
	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printNodeArgument(n node.Node) {
	nn := n.(*node.Argument)
	p.printMeta(nn, meta.NodeStart)

	if nn.IsReference {
		p.printMeta(nn, meta.AmpersandToken)
		io.WriteString(p.w, "&")
	}

	if nn.Variadic {
		p.printMeta(nn, meta.EllipsisToken)
		io.WriteString(p.w, "...")
	}

	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

// name

func (p *Printer) printNameNamePart(n node.Node) {
	nn := n.(*name.NamePart)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, nn.Value)

	p.printMeta(nn, meta.NsSeparatorToken)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printNameName(n node.Node) {
	nn := n.(*name.Name)
	p.printMeta(nn, meta.NodeStart)

	for k, part := range nn.Parts {
		if k > 0 {
			io.WriteString(p.w, "\\")
		}

		p.Print(part)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printNameFullyQualified(n node.Node) {
	nn := n.(*name.FullyQualified)
	p.printMeta(nn, meta.NodeStart)

	for _, part := range nn.Parts {
		io.WriteString(p.w, "\\")
		p.Print(part)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printNameRelative(n node.Node) {
	nn := n.(*name.Relative)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "namespace")
	p.printMeta(nn, meta.NsSeparatorToken)

	for _, part := range nn.Parts {
		io.WriteString(p.w, "\\")
		p.Print(part)
	}

	p.printMeta(nn, meta.NodeEnd)
}

// scalar

func (p *Printer) printScalarLNumber(n node.Node) {
	nn := n.(*scalar.Lnumber)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, nn.Value)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printScalarDNumber(n node.Node) {
	nn := n.(*scalar.Dnumber)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, nn.Value)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printScalarString(n node.Node) {
	nn := n.(*scalar.String)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, nn.Value)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printScalarEncapsedStringPart(n node.Node) {
	nn := n.(*scalar.EncapsedStringPart)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, nn.Value)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printScalarEncapsed(n node.Node) {
	nn := n.(*scalar.Encapsed)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "\"")
	for _, part := range nn.Parts {
		switch part.(type) {
		case *expr.ArrayDimFetch:
			if len(part.GetMeta().FindBy(meta.ValueFilter("${"))) == 1 {
				p.printExprArrayDimFetchWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		case *expr.Variable:
			if len(part.GetMeta().FindBy(meta.ValueFilter("${"))) == 1 {
				p.printExprVariableWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		default:
			p.Print(part)
		}
	}
	io.WriteString(p.w, "\"")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printScalarHeredoc(n node.Node) {
	nn := n.(*scalar.Heredoc)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "<<<")
	io.WriteString(p.w, nn.Label)
	io.WriteString(p.w, "\n")

	for _, part := range nn.Parts {
		switch part.(type) {
		case *expr.ArrayDimFetch:
			if len(part.GetMeta().FindBy(meta.ValueFilter("${"))) == 1 {
				p.printExprArrayDimFetchWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		case *expr.Variable:
			if len(part.GetMeta().FindBy(meta.ValueFilter("${"))) == 1 {
				p.printExprVariableWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		default:
			p.Print(part)
		}
	}

	io.WriteString(p.w, "\n")
	io.WriteString(p.w, strings.Trim(nn.Label, "\"'"))

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printScalarMagicConstant(n node.Node) {
	nn := n.(*scalar.MagicConstant)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, nn.Value)
	p.printMeta(nn, meta.NodeEnd)
}

// Assign

func (p *Printer) printAssign(n node.Node) {
	nn := n.(*assign.Assign)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.EqualToken)
	io.WriteString(p.w, "=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignReference(n node.Node) {
	nn := n.(*assign.Reference)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.EqualToken)
	io.WriteString(p.w, "=")
	p.printMeta(nn, meta.AmpersandToken)
	io.WriteString(p.w, "&")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignBitwiseAnd(n node.Node) {
	nn := n.(*assign.BitwiseAnd)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.AndEqualToken)
	io.WriteString(p.w, "&")
	io.WriteString(p.w, "=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignBitwiseOr(n node.Node) {
	nn := n.(*assign.BitwiseOr)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.OrEqualToken)
	io.WriteString(p.w, "|=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignBitwiseXor(n node.Node) {
	nn := n.(*assign.BitwiseXor)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.XorEqualToken)
	io.WriteString(p.w, "^=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignConcat(n node.Node) {
	nn := n.(*assign.Concat)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.ConcatEqualToken)
	io.WriteString(p.w, ".=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignDiv(n node.Node) {
	nn := n.(*assign.Div)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.DivEqualToken)
	io.WriteString(p.w, "/=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignMinus(n node.Node) {
	nn := n.(*assign.Minus)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.MinusEqualToken)
	io.WriteString(p.w, "-=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignMod(n node.Node) {
	nn := n.(*assign.Mod)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.ModEqualToken)
	io.WriteString(p.w, "%=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignMul(n node.Node) {
	nn := n.(*assign.Mul)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.MulEqualToken)
	io.WriteString(p.w, "*=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignPlus(n node.Node) {
	nn := n.(*assign.Plus)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.PlusEqualToken)
	io.WriteString(p.w, "+=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignPow(n node.Node) {
	nn := n.(*assign.Pow)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.PowEqualToken)
	io.WriteString(p.w, "**=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignShiftLeft(n node.Node) {
	nn := n.(*assign.ShiftLeft)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.SlEqualToken)
	io.WriteString(p.w, "<<=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printAssignShiftRight(n node.Node) {
	nn := n.(*assign.ShiftRight)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.SrEqualToken)
	io.WriteString(p.w, ">>=")
	p.Print(nn.Expression)

	p.printMeta(nn, meta.NodeEnd)
}

// binary

func (p *Printer) printBinaryBitwiseAnd(n node.Node) {
	nn := n.(*binary.BitwiseAnd)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.AmpersandToken)
	io.WriteString(p.w, "&")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryBitwiseOr(n node.Node) {
	nn := n.(*binary.BitwiseOr)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.VerticalBarToken)
	io.WriteString(p.w, "|")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryBitwiseXor(n node.Node) {
	nn := n.(*binary.BitwiseXor)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.CaretToken)
	io.WriteString(p.w, "^")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryBooleanAnd(n node.Node) {
	nn := n.(*binary.BooleanAnd)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.BooleanAndToken)
	io.WriteString(p.w, "&&")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryBooleanOr(n node.Node) {
	nn := n.(*binary.BooleanOr)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.BooleanOrToken)
	io.WriteString(p.w, "||")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryCoalesce(n node.Node) {
	nn := n.(*binary.Coalesce)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.CoalesceToken)
	io.WriteString(p.w, "??")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryConcat(n node.Node) {
	nn := n.(*binary.Concat)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.DotToken)
	io.WriteString(p.w, ".")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryDiv(n node.Node) {
	nn := n.(*binary.Div)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.SlashToken)
	io.WriteString(p.w, "/")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryEqual(n node.Node) {
	nn := n.(*binary.Equal)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.IsEqualToken)
	io.WriteString(p.w, "==")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryGreaterOrEqual(n node.Node) {
	nn := n.(*binary.GreaterOrEqual)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.IsGreaterOrEqualToken)
	io.WriteString(p.w, ">=")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryGreater(n node.Node) {
	nn := n.(*binary.Greater)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.GreaterToken)
	io.WriteString(p.w, ">")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryIdentical(n node.Node) {
	nn := n.(*binary.Identical)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.IsIdenticalToken)
	io.WriteString(p.w, "===")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryLogicalAnd(n node.Node) {
	nn := n.(*binary.LogicalAnd)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.printMeta(nn, meta.LogicalAndToken)
	io.WriteString(p.w, "and")
	if len((*nn.Right.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryLogicalOr(n node.Node) {
	nn := n.(*binary.LogicalOr)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.printMeta(nn, meta.LogicalOrToken)
	io.WriteString(p.w, "or")
	if len((*nn.Right.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryLogicalXor(n node.Node) {
	nn := n.(*binary.LogicalXor)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.printMeta(nn, meta.LogicalXorToken)
	io.WriteString(p.w, "xor")
	if len((*nn.Right.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryMinus(n node.Node) {
	nn := n.(*binary.Minus)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.MinusToken)
	io.WriteString(p.w, "-")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryMod(n node.Node) {
	nn := n.(*binary.Mod)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.PercentToken)
	io.WriteString(p.w, "%")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryMul(n node.Node) {
	nn := n.(*binary.Mul)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.AsteriskToken)
	io.WriteString(p.w, "*")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryNotEqual(n node.Node) {
	nn := n.(*binary.NotEqual)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "!=")
	}
	p.printMeta(nn, meta.IsNotEqualToken)
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryNotIdentical(n node.Node) {
	nn := n.(*binary.NotIdentical)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.IsNotIdenticalToken)
	io.WriteString(p.w, "!==")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryPlus(n node.Node) {
	nn := n.(*binary.Plus)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.PlusToken)
	io.WriteString(p.w, "+")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryPow(n node.Node) {
	nn := n.(*binary.Pow)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.PowToken)
	io.WriteString(p.w, "**")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryShiftLeft(n node.Node) {
	nn := n.(*binary.ShiftLeft)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.SlToken)
	io.WriteString(p.w, "<<")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinaryShiftRight(n node.Node) {
	nn := n.(*binary.ShiftRight)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.SrToken)
	io.WriteString(p.w, ">>")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinarySmallerOrEqual(n node.Node) {
	nn := n.(*binary.SmallerOrEqual)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.IsSmallerOrEqualToken)
	io.WriteString(p.w, "<=")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinarySmaller(n node.Node) {
	nn := n.(*binary.Smaller)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.LessToken)
	io.WriteString(p.w, "<")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBinarySpaceship(n node.Node) {
	nn := n.(*binary.Spaceship)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Left)
	p.printMeta(nn, meta.SpaceshipToken)
	io.WriteString(p.w, "<=>")
	p.Print(nn.Right)

	p.printMeta(nn, meta.NodeEnd)
}

// cast

func (p *Printer) printArray(n node.Node) {
	nn := n.(*cast.Array)
	p.printMeta(nn, meta.NodeStart)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "(array)")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printBool(n node.Node) {
	nn := n.(*cast.Bool)
	p.printMeta(nn, meta.NodeStart)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "(boolean)")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printDouble(n node.Node) {
	nn := n.(*cast.Double)
	p.printMeta(nn, meta.NodeStart)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "(float)")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printInt(n node.Node) {
	nn := n.(*cast.Int)
	p.printMeta(nn, meta.NodeStart)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "(integer)")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printObject(n node.Node) {
	nn := n.(*cast.Object)
	p.printMeta(nn, meta.NodeStart)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "(object)")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printString(n node.Node) {
	nn := n.(*cast.String)
	p.printMeta(nn, meta.NodeStart)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "(string)")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printUnset(n node.Node) {
	nn := n.(*cast.Unset)
	p.printMeta(nn, meta.NodeStart)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "(unset)")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

// expr

func (p *Printer) printExprArrayDimFetch(n node.Node) {
	nn := n.(*expr.ArrayDimFetch)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.OpenSquareBracket)
	p.printMeta(nn, meta.OpenCurlyBracesToken)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "[")
	}
	p.Print(nn.Dim)
	p.printMeta(nn, meta.CloseSquareBracket)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "]")
	}
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprArrayDimFetchWithoutLeadingDollar(n node.Node) {
	nn := n.(*expr.ArrayDimFetch)
	p.printMeta(nn, meta.NodeStart)
	p.printExprVariableWithoutLeadingDollar(nn.Variable)
	p.printMeta(nn, meta.OpenSquareBracket)
	p.printMeta(nn, meta.OpenCurlyBracesToken)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "[")
	}
	p.Print(nn.Dim)
	p.printMeta(nn, meta.CloseSquareBracket)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "]")
	}
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprArrayItem(n node.Node) {
	nn := n.(*expr.ArrayItem)
	p.printMeta(nn, meta.NodeStart)

	if nn.Key != nil {
		p.Print(nn.Key)
		p.printMeta(nn, meta.DoubleArrowToken)
		io.WriteString(p.w, "=>")
	}

	p.Print(nn.Val)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprArray(n node.Node) {
	nn := n.(*expr.Array)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "array")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Items)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprBitwiseNot(n node.Node) {
	nn := n.(*expr.BitwiseNot)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "~")
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprBooleanNot(n node.Node) {
	nn := n.(*expr.BooleanNot)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "!")
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprClassConstFetch(n node.Node) {
	nn := n.(*expr.ClassConstFetch)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Class)
	p.printMeta(nn, meta.PaamayimNekudotayimToken)
	io.WriteString(p.w, "::")
	p.Print(nn.ConstantName)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprClone(n node.Node) {
	nn := n.(*expr.Clone)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "clone")
	if len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprClosureUse(n node.Node) {
	nn := n.(*expr.ClosureUse)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "use")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Uses)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprClosure(n node.Node) {
	nn := n.(*expr.Closure)
	p.printMeta(nn, meta.NodeStart)

	if nn.Static {
		io.WriteString(p.w, "static")
	}

	if nn.Static && len(n.GetMeta().FindBy(meta.TypeFilter(meta.WhiteSpaceType))) == 0 {
		io.WriteString(p.w, " ")
	}

	p.printMeta(nn, meta.FunctionToken)
	io.WriteString(p.w, "function")

	if nn.ReturnsRef {
		p.printMeta(nn, meta.AmpersandToken)
		io.WriteString(p.w, "&")
	}

	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	if nn.ClosureUse != nil {
		p.Print(nn.ClosureUse)
	}

	if nn.ReturnType != nil {
		p.printMeta(nn.ReturnType, meta.ColonToken)
		io.WriteString(p.w, ":")
		p.Print(nn.ReturnType)
	}

	p.printMeta(nn, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprConstFetch(n node.Node) {
	nn := n.(*expr.ConstFetch)
	p.printMeta(nn, meta.NodeStart)
	p.Print(nn.Constant)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprEmpty(n node.Node) {
	nn := n.(*expr.Empty)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "empty")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Expr)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprErrorSuppress(n node.Node) {
	nn := n.(*expr.ErrorSuppress)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "@")
	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprEval(n node.Node) {
	nn := n.(*expr.Eval)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "eval")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Expr)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprExit(n node.Node) {
	nn := n.(*expr.Exit)
	p.printMeta(nn, meta.NodeStart)

	if nn.Die {
		io.WriteString(p.w, "die")
	} else {
		io.WriteString(p.w, "exit")
	}

	p.printMeta(nn, meta.OpenParenthesisToken)
	if len((*nn.GetMeta())) == 0 && nn.Expr != nil && len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.CloseParenthesisToken)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprFunctionCall(n node.Node) {
	nn := n.(*expr.FunctionCall)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Function)
	p.printMeta(nn.ArgumentList, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.ArgumentList.Arguments)
	p.printMeta(nn.ArgumentList, meta.CommaToken)
	p.printMeta(nn.ArgumentList, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprInclude(n node.Node) {
	nn := n.(*expr.Include)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "include")
	if len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprIncludeOnce(n node.Node) {
	nn := n.(*expr.IncludeOnce)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "include_once")
	if len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprInstanceOf(n node.Node) {
	nn := n.(*expr.InstanceOf)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Expr)
	p.printMeta(nn, meta.InstanceofToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "instanceof")
	if len((*nn.Class.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Class)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprIsset(n node.Node) {
	nn := n.(*expr.Isset)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "isset")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Variables)
	p.printMeta(nn, meta.CommaToken)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprList(n node.Node) {
	nn := n.(*expr.List)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "list")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Items)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprMethodCall(n node.Node) {
	nn := n.(*expr.MethodCall)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Variable)
	p.printMeta(nn, meta.ObjectOperatorToken)
	io.WriteString(p.w, "->")
	p.Print(nn.Method)
	p.printMeta(nn.ArgumentList, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.ArgumentList.Arguments)
	p.printMeta(nn.ArgumentList, meta.CommaToken)
	p.printMeta(nn.ArgumentList, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprNew(n node.Node) {
	nn := n.(*expr.New)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "new")
	if len((*nn.Class.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Class)

	if nn.ArgumentList != nil {
		p.printMeta(nn.ArgumentList, meta.OpenParenthesisToken)
		io.WriteString(p.w, "(")
		p.joinPrint(",", nn.ArgumentList.Arguments)
		p.printMeta(nn.ArgumentList, meta.CommaToken)
		p.printMeta(nn.ArgumentList, meta.CloseParenthesisToken)
		io.WriteString(p.w, ")")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprPostDec(n node.Node) {
	nn := n.(*expr.PostDec)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Variable)
	p.printMeta(nn, meta.DecToken)
	io.WriteString(p.w, "--")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprPostInc(n node.Node) {
	nn := n.(*expr.PostInc)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Variable)
	p.printMeta(nn, meta.IncToken)
	io.WriteString(p.w, "++")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprPreDec(n node.Node) {
	nn := n.(*expr.PreDec)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "--")
	p.Print(nn.Variable)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprPreInc(n node.Node) {
	nn := n.(*expr.PreInc)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "++")
	p.Print(nn.Variable)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprPrint(n node.Node) {
	nn := n.(*expr.Print)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "print")
	if len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprPropertyFetch(n node.Node) {
	nn := n.(*expr.PropertyFetch)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Variable)
	p.printMeta(nn, meta.ObjectOperatorToken)
	io.WriteString(p.w, "->")
	p.Print(nn.Property)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprReference(n node.Node) {
	nn := n.(*expr.Reference)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "&")
	p.Print(nn.Variable)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprRequire(n node.Node) {
	nn := n.(*expr.Require)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "require")
	if len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprRequireOnce(n node.Node) {
	nn := n.(*expr.RequireOnce)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "require_once")
	if len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprShellExec(n node.Node) {
	nn := n.(*expr.ShellExec)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "`")
	for _, part := range nn.Parts {
		p.Print(part)
	}
	io.WriteString(p.w, "`")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprShortArray(n node.Node) {
	nn := n.(*expr.ShortArray)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "[")
	p.joinPrint(",", nn.Items)
	p.printMeta(nn, meta.CloseSquareBracket)
	io.WriteString(p.w, "]")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprShortList(n node.Node) {
	nn := n.(*expr.ShortList)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "[")
	p.joinPrint(",", nn.Items)
	p.printMeta(nn, meta.CloseSquareBracket)
	io.WriteString(p.w, "]")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprStaticCall(n node.Node) {
	nn := n.(*expr.StaticCall)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Class)
	p.printMeta(nn, meta.PaamayimNekudotayimToken)
	io.WriteString(p.w, "::")
	p.Print(nn.Call)
	p.printMeta(nn.ArgumentList, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.ArgumentList.Arguments)
	p.printMeta(nn.ArgumentList, meta.CommaToken)
	p.printMeta(nn.ArgumentList, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprStaticPropertyFetch(n node.Node) {
	nn := n.(*expr.StaticPropertyFetch)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Class)
	p.printMeta(nn, meta.PaamayimNekudotayimToken)
	io.WriteString(p.w, "::")
	p.Print(nn.Property)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprTernary(n node.Node) {
	nn := n.(*expr.Ternary)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Condition)
	p.printMeta(nn, meta.QuestionMarkToken)
	io.WriteString(p.w, "?")

	if nn.IfTrue != nil {
		p.Print(nn.IfTrue)
	}

	p.printMeta(nn, meta.ColonToken)
	io.WriteString(p.w, ":")
	p.Print(nn.IfFalse)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprUnaryMinus(n node.Node) {
	nn := n.(*expr.UnaryMinus)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "-")
	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprUnaryPlus(n node.Node) {
	nn := n.(*expr.UnaryPlus)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "+")
	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprVariable(n node.Node) {
	nn := n.(*expr.Variable)
	p.printMeta(nn, meta.NodeStart)

	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, "$")
	}

	p.Print(nn.VarName)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprVariableWithoutLeadingDollar(n node.Node) {
	nn := n.(*expr.Variable)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.VarName)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprYieldFrom(n node.Node) {
	nn := n.(*expr.YieldFrom)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "yield from")
	if len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printExprYield(n node.Node) {
	nn := n.(*expr.Yield)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "yield")

	if nn.Key != nil {
		if len((*nn.Key.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Key)
		p.printMeta(nn, meta.DoubleArrowToken)
		io.WriteString(p.w, "=>")
	} else {
		if len((*nn.Value.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Value)

	p.printMeta(nn, meta.NodeEnd)
}

// smtm

func (p *Printer) printStmtAltElseIf(n node.Node) {
	nn := n.(*stmt.AltElseIf)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "elseif")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.ColonToken)
	io.WriteString(p.w, ":")

	if s := nn.Stmt.(*stmt.StmtList).Stmts; len(s) > 0 {
		p.printNodes(s)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtAltElse(n node.Node) {
	nn := n.(*stmt.AltElse)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "else")
	p.printMeta(nn, meta.ColonToken)
	io.WriteString(p.w, ":")

	if s := nn.Stmt.(*stmt.StmtList).Stmts; len(s) > 0 {
		p.printNodes(s)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtAltFor(n node.Node) {
	nn := n.(*stmt.AltFor)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "for")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Init)
	p.printMeta(nn, meta.ForInitSemicolonToken)
	io.WriteString(p.w, ";")
	p.joinPrint(",", nn.Cond)
	p.printMeta(nn, meta.ForCondSemicolonToken)
	io.WriteString(p.w, ";")
	p.joinPrint(",", nn.Loop)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.ColonToken)
	io.WriteString(p.w, ":")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)

	p.printMeta(nn, meta.EndforToken)
	io.WriteString(p.w, "endfor")
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtAltForeach(n node.Node) {
	nn := n.(*stmt.AltForeach)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "foreach")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Expr)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.printMeta(nn, meta.AsToken)
	io.WriteString(p.w, "as")

	if nn.Key != nil {
		if len((*nn.Key.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Key)
		p.printMeta(nn, meta.DoubleArrowToken)
		io.WriteString(p.w, "=>")
	} else {
		if len((*nn.Variable.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Variable)

	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.ColonToken)
	io.WriteString(p.w, ":")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)

	p.printMeta(nn, meta.EndforeachToken)
	io.WriteString(p.w, "endforeach")
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtAltIf(n node.Node) {
	nn := n.(*stmt.AltIf)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "if")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.ColonToken)
	io.WriteString(p.w, ":")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)

	for _, elseif := range nn.ElseIf {
		p.Print(elseif)
	}

	if nn.Else != nil {
		p.Print(nn.Else)
	}

	p.printMeta(nn, meta.EndifToken)
	io.WriteString(p.w, "endif")
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtAltSwitch(n node.Node) {
	nn := n.(*stmt.AltSwitch)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "switch")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.ColonToken)
	io.WriteString(p.w, ":")

	p.printMeta(nn.CaseList, meta.CaseSeparatorToken)
	s := nn.CaseList.Cases
	p.printNodes(s)

	p.printMeta(nn, meta.EndswitchToken)
	io.WriteString(p.w, "endswitch")
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtAltWhile(n node.Node) {
	nn := n.(*stmt.AltWhile)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "while")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.ColonToken)
	io.WriteString(p.w, ":")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)

	p.printMeta(nn, meta.EndwhileToken)
	io.WriteString(p.w, "endwhile")
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtBreak(n node.Node) {
	nn := n.(*stmt.Break)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "break")
	if nn.Expr != nil {
		if len((*nn.Expr.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Expr)
	}

	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtCase(n node.Node) {
	nn := n.(*stmt.Case)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "case")
	if len((*nn.Cond.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Cond)
	r := p.printMeta(nn, meta.CaseSeparatorToken)
	if !r {
		io.WriteString(p.w, ":")
	}

	if len(nn.Stmts) > 0 {
		p.printNodes(nn.Stmts)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtCatch(n node.Node) {
	nn := n.(*stmt.Catch)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "catch")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint("|", nn.Types)
	p.Print(nn.Variable)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtClassMethod(n node.Node) {
	nn := n.(*stmt.ClassMethod)
	p.printMeta(nn, meta.NodeStart)

	if nn.Modifiers != nil {
		for k, m := range nn.Modifiers {
			if k > 0 && len(m.GetMeta().FindBy(meta.AndFilter(meta.TokenNameFilter(meta.NodeStart), meta.TypeFilter(meta.WhiteSpaceType)))) == 0 {
				io.WriteString(p.w, " ")
			}
			p.Print(m)
		}

		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}
	p.printMeta(nn, meta.FunctionToken)
	io.WriteString(p.w, "function")

	if nn.ReturnsRef {
		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.printMeta(nn, meta.AmpersandToken)
		io.WriteString(p.w, "&")
	} else {
		if len((*nn.MethodName.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.MethodName)
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	if nn.ReturnType != nil {
		p.printMeta(nn.ReturnType, meta.ColonToken)
		io.WriteString(p.w, ":")
		p.Print(nn.ReturnType)
	}

	p.Print(nn.Stmt)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtClass(n node.Node) {
	nn := n.(*stmt.Class)
	p.printMeta(nn, meta.NodeStart)

	if nn.Modifiers != nil {
		for k, m := range nn.Modifiers {
			if k > 0 && len(m.GetMeta().FindBy(meta.AndFilter(meta.TokenNameFilter(meta.NodeStart), meta.TypeFilter(meta.WhiteSpaceType)))) == 0 {
				io.WriteString(p.w, " ")
			}
			p.Print(m)
		}

		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}
	p.printMeta(nn, meta.ClassToken)
	io.WriteString(p.w, "class")

	if nn.ClassName != nil {
		if len((*nn.ClassName.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.ClassName)
	}

	if nn.ArgumentList != nil {
		p.printMeta(nn.ArgumentList, meta.OpenParenthesisToken)
		io.WriteString(p.w, "(")
		p.joinPrint(",", nn.ArgumentList.Arguments)
		p.printMeta(nn.ArgumentList, meta.CommaToken)
		p.printMeta(nn.ArgumentList, meta.CloseParenthesisToken)
		io.WriteString(p.w, ")")
	}

	if nn.Extends != nil {
		p.printMeta(nn.Extends, meta.NodeStart)
		if len((*nn.Extends.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.printMeta(nn.Extends, meta.ExtendsToken)
		io.WriteString(p.w, "extends")
		if len((*nn.Extends.ClassName.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Extends.ClassName)
	}

	if nn.Implements != nil {
		p.printMeta(nn.Implements, meta.NodeStart)
		if len((*nn.Implements.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.printMeta(nn.Implements, meta.ImplementsToken)
		io.WriteString(p.w, "implements")
		if len((*nn.Implements.InterfaceNames[0].GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.joinPrint(",", nn.Implements.InterfaceNames)
	}

	p.printMeta(nn, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtClassConstList(n node.Node) {
	nn := n.(*stmt.ClassConstList)
	p.printMeta(nn, meta.NodeStart)

	if nn.Modifiers != nil {
		for k, m := range nn.Modifiers {
			if k > 0 && len(m.GetMeta().FindBy(meta.AndFilter(meta.TokenNameFilter(meta.NodeStart), meta.TypeFilter(meta.WhiteSpaceType)))) == 0 {
				io.WriteString(p.w, " ")
			}
			p.Print(m)
		}

		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}
	p.printMeta(nn, meta.ConstToken)
	io.WriteString(p.w, "const")

	if len((*nn.Consts[0].GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Consts)

	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtConstList(n node.Node) {
	nn := n.(*stmt.ConstList)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "const")

	if len((*nn.Consts[0].GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Consts)

	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtConstant(n node.Node) {
	nn := n.(*stmt.Constant)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.ConstantName)
	p.printMeta(nn, meta.EqualToken)
	io.WriteString(p.w, "=")
	p.Print(nn.Expr)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtContinue(n node.Node) {
	nn := n.(*stmt.Continue)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "continue")

	if nn.Expr != nil {
		if len((*nn.Expr.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Expr)
	}

	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtDeclare(n node.Node) {
	nn := n.(*stmt.Declare)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "declare")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Consts)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	if nn.Alt {
		p.printMeta(nn, meta.ColonToken)
		io.WriteString(p.w, ":")

		s := nn.Stmt.(*stmt.StmtList)
		p.printNodes(s.Stmts)

		p.printMeta(nn, meta.EnddeclareToken)
		io.WriteString(p.w, "enddeclare")
		p.printMeta(nn, meta.SemiColonToken)
		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, ";")
		}

	} else {
		p.Print(nn.Stmt)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtDefault(n node.Node) {
	nn := n.(*stmt.Default)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "default")
	r := p.printMeta(nn, meta.CaseSeparatorToken)
	if !r {
		io.WriteString(p.w, ":")
	}

	if len(nn.Stmts) > 0 {
		p.printNodes(nn.Stmts)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtDo(n node.Node) {
	nn := n.(*stmt.Do)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "do")

	if _, ok := nn.Stmt.(*stmt.StmtList); !ok {
		if len((*nn.Stmt.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Stmt)

	p.printMeta(nn, meta.WhileToken)
	io.WriteString(p.w, "while")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtEcho(n node.Node) {
	nn := n.(*stmt.Echo)

	if p.s == HtmlState {
		if len((*n.GetMeta())) == 0 {
			io.WriteString(p.w, "<?=")
		}

		p.SetState(PhpState)
	} else {
		if len((*n.GetMeta())) == 0 {
			io.WriteString(p.w, "echo")
		}
		if len((*nn.Exprs[0].GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}

	p.printMeta(nn, meta.NodeStart)

	p.joinPrint(",", nn.Exprs)
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtElseif(n node.Node) {
	nn := n.(*stmt.ElseIf)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "elseif")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtElse(n node.Node) {
	nn := n.(*stmt.Else)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "else")

	if _, ok := nn.Stmt.(*stmt.StmtList); !ok {
		if len((*nn.Stmt.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Stmt)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtExpression(n node.Node) {
	nn := n.(*stmt.Expression)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Expr)

	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtFinally(n node.Node) {
	nn := n.(*stmt.Finally)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "finally")
	p.printMeta(nn, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtFor(n node.Node) {
	nn := n.(*stmt.For)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "for")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Init)
	p.printMeta(nn, meta.ForInitSemicolonToken)
	io.WriteString(p.w, ";")
	p.joinPrint(",", nn.Cond)
	p.printMeta(nn, meta.ForCondSemicolonToken)
	io.WriteString(p.w, ";")
	p.joinPrint(",", nn.Loop)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtForeach(n node.Node) {
	nn := n.(*stmt.Foreach)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "foreach")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Expr)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.printMeta(nn, meta.AsToken)
	io.WriteString(p.w, "as")

	if nn.Key != nil {
		if len((*nn.Key.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Key)
		p.printMeta(nn, meta.DoubleArrowToken)
		io.WriteString(p.w, "=>")
	} else {
		if len((*nn.Variable.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Variable)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtFunction(n node.Node) {
	nn := n.(*stmt.Function)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "function")

	if nn.ReturnsRef {
		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.printMeta(nn, meta.AmpersandToken)
		io.WriteString(p.w, "&")
	} else {
		if len((*nn.FunctionName.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.FunctionName)

	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	if nn.ReturnType != nil {
		p.printMeta(nn.ReturnType, meta.ColonToken)
		io.WriteString(p.w, ":")
		p.Print(nn.ReturnType)
	}

	p.printMeta(nn, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtGlobal(n node.Node) {
	nn := n.(*stmt.Global)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "global")
	p.joinPrint(",", nn.Vars)
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtGoto(n node.Node) {
	nn := n.(*stmt.Goto)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "goto")
	if len((*nn.Label.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Label)
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtGroupUse(n node.Node) {
	nn := n.(*stmt.GroupUse)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "use")

	if nn.UseType != nil {
		if len((*nn.UseType.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.UseType)
	}

	if len((*nn.Prefix.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Prefix)
	p.printMeta(nn, meta.NsSeparatorToken)
	io.WriteString(p.w, "\\")
	p.printMeta(nn, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.joinPrint(",", nn.UseList)
	p.printMeta(nn, meta.CommaToken)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtHaltCompiler(n node.Node) {
	nn := n.(*stmt.HaltCompiler)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "__halt_compiler")
	p.printMeta(n, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.printMeta(n, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtIf(n node.Node) {
	nn := n.(*stmt.If)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "if")
	p.printMeta(n, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printMeta(n, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	if nn.ElseIf != nil {
		p.printNodes(nn.ElseIf)
	}

	if nn.Else != nil {
		p.Print(nn.Else)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtInlineHTML(n node.Node) {
	nn := n.(*stmt.InlineHtml)
	p.printMeta(nn, meta.NodeStart)

	if p.s == PhpState && len(n.GetMeta().FindBy(meta.ValueFilter("?>"))) == 0 {
		io.WriteString(p.w, "?>")
	}
	p.SetState(HtmlState)

	io.WriteString(p.w, nn.Value)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtInterface(n node.Node) {
	nn := n.(*stmt.Interface)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "interface")

	if len((*nn.InterfaceName.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}

	p.Print(nn.InterfaceName)

	if nn.Extends != nil {
		p.printMeta(nn.Extends, meta.NodeStart)
		if len((*nn.Extends.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.printMeta(nn.Extends, meta.ExtendsToken)
		io.WriteString(p.w, "extends")
		if len((*nn.Extends.InterfaceNames[0].GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.joinPrint(",", nn.Extends.InterfaceNames)
	}

	p.printMeta(n, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printMeta(n, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtLabel(n node.Node) {
	nn := n.(*stmt.Label)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.LabelName)
	p.printMeta(n, meta.ColonToken)
	io.WriteString(p.w, ":")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtNamespace(n node.Node) {
	nn := n.(*stmt.Namespace)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "namespace")

	if nn.NamespaceName != nil {
		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.NamespaceName)
	}

	if nn.Stmts != nil {
		p.printMeta(n, meta.OpenCurlyBracesToken)
		io.WriteString(p.w, "{")
		p.printNodes(nn.Stmts)
		p.printMeta(n, meta.CloseCurlyBracesToken)
		io.WriteString(p.w, "}")
	} else {
		p.printMeta(nn, meta.SemiColonToken)
		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, ";")
		}
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtNop(n node.Node) {
	p.printMeta(n, meta.NodeStart)
	if len((*n.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}
	p.printMeta(n, meta.NodeEnd)
}

func (p *Printer) printStmtPropertyList(n node.Node) {
	nn := n.(*stmt.PropertyList)
	p.printMeta(nn, meta.NodeStart)

	for k, m := range nn.Modifiers {
		if k > 0 && len(m.GetMeta().FindBy(meta.AndFilter(meta.TokenNameFilter(meta.NodeStart), meta.TypeFilter(meta.WhiteSpaceType)))) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(m)
	}

	if len((*nn.Properties[0].GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}

	p.joinPrint(",", nn.Properties)
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtProperty(n node.Node) {
	nn := n.(*stmt.Property)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Variable)

	if nn.Expr != nil {
		p.printMeta(n, meta.EqualToken)
		io.WriteString(p.w, "=")
		p.Print(nn.Expr)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtReturn(n node.Node) {
	nn := n.(*stmt.Return)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "return")
	if len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtStaticVar(n node.Node) {
	nn := n.(*stmt.StaticVar)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Variable)

	if nn.Expr != nil {
		p.printMeta(nn, meta.EqualToken)
		io.WriteString(p.w, "=")
		p.Print(nn.Expr)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtStatic(n node.Node) {
	nn := n.(*stmt.Static)
	p.printMeta(nn, meta.NodeStart)
	io.WriteString(p.w, "static")
	p.joinPrint(",", nn.Vars)
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}
	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtStmtList(n node.Node) {
	nn := n.(*stmt.StmtList)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtSwitch(n node.Node) {
	nn := n.(*stmt.Switch)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "switch")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.printMeta(nn.CaseList, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.printMeta(nn.CaseList, meta.CaseSeparatorToken)
	p.printNodes(nn.CaseList.Cases)
	p.printMeta(nn.CaseList, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtThrow(n node.Node) {
	nn := n.(*stmt.Throw)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "throw")
	if len((*nn.Expr.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtTraitAdaptationList(n node.Node) {
	nn := n.(*stmt.TraitAdaptationList)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "{")
	p.printNodes(nn.Adaptations)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtTraitMethodRef(n node.Node) {
	nn := n.(*stmt.TraitMethodRef)
	p.printMeta(nn, meta.NodeStart)

	if nn.Trait != nil {
		p.Print(nn.Trait)
		p.printMeta(nn, meta.PaamayimNekudotayimToken)
		io.WriteString(p.w, "::")
	}
	p.Print(nn.Method)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtTraitUseAlias(n node.Node) {
	nn := n.(*stmt.TraitUseAlias)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Ref)
	p.printMeta(nn, meta.AsToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "as")

	if nn.Modifier != nil {
		if len((*nn.Modifier.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Modifier)
	}

	if nn.Alias != nil {
		if len((*nn.Alias.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Alias)
	}

	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtTraitUsePrecedence(n node.Node) {
	nn := n.(*stmt.TraitUsePrecedence)
	p.printMeta(nn, meta.NodeStart)

	p.Print(nn.Ref)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.printMeta(nn, meta.InsteadofToken)
	io.WriteString(p.w, "insteadof")
	if len((*nn.Insteadof[0].GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Insteadof)

	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtTraitUse(n node.Node) {
	nn := n.(*stmt.TraitUse)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "use")
	if len((*nn.Traits[0].GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Traits)

	p.Print(nn.TraitAdaptationList)

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtTrait(n node.Node) {
	nn := n.(*stmt.Trait)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "trait")
	if len((*nn.TraitName.GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.TraitName)

	p.printMeta(nn, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtTry(n node.Node) {
	nn := n.(*stmt.Try)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "try")
	p.printMeta(nn, meta.OpenCurlyBracesToken)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printMeta(nn, meta.CloseCurlyBracesToken)
	io.WriteString(p.w, "}")

	if nn.Catches != nil {
		p.printNodes(nn.Catches)
	}

	if nn.Finally != nil {
		p.Print(nn.Finally)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtUnset(n node.Node) {
	nn := n.(*stmt.Unset)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "unset")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Vars)
	p.printMeta(nn, meta.CommaToken)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtUseList(n node.Node) {
	nn := n.(*stmt.UseList)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "use")

	if nn.UseType != nil {
		if len((*nn.UseType.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.UseType)
	}

	if len((*nn.Uses[0].GetMeta())) == 0 {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Uses)
	p.printMeta(nn, meta.SemiColonToken)
	if len((*nn.GetMeta())) == 0 {
		io.WriteString(p.w, ";")
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtUse(n node.Node) {
	nn := n.(*stmt.Use)
	p.printMeta(nn, meta.NodeStart)

	if nn.UseType != nil {
		p.Print(nn.UseType)
		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Use)

	if nn.Alias != nil {
		if len((*nn.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.printMeta(nn, meta.AsToken)
		io.WriteString(p.w, "as")
		if len((*nn.Alias.GetMeta())) == 0 {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Alias)
	}

	p.printMeta(nn, meta.NodeEnd)
}

func (p *Printer) printStmtWhile(n node.Node) {
	nn := n.(*stmt.While)
	p.printMeta(nn, meta.NodeStart)

	io.WriteString(p.w, "while")
	p.printMeta(nn, meta.OpenParenthesisToken)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printMeta(nn, meta.CloseParenthesisToken)
	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	p.printMeta(nn, meta.NodeEnd)
}
