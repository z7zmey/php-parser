package printer

import (
	"io"
	"strings"

	"github.com/z7zmey/php-parser/freefloating"

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
	if p.s == HtmlState && !isInlineHtml && !isRoot {
		if n.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, "<?php ")
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

func (p *Printer) printFreeFloating(n node.Node, pos freefloating.Position) {
	if n == nil {
		return
	}

	for _, m := range (*n.GetFreeFloating())[pos] {
		io.WriteString(p.w, m.Value)
	}
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
	case *assign.Coalesce:
		p.printAssignCoalesce(n)
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
	case *expr.ArrowFunction:
		p.printExprArrowFunction(n)
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
	p.printFreeFloating(nn, freefloating.Start)
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printNodeIdentifier(n node.Node) {
	nn := n.(*node.Identifier)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, nn.Value)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printNodeParameter(n node.Node) {
	nn := n.(*node.Parameter)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.VariableType != nil {
		p.Print(nn.VariableType)
	}
	p.printFreeFloating(nn, freefloating.OptionalType)

	if nn.ByRef {
		io.WriteString(p.w, "&")
	}
	p.printFreeFloating(nn, freefloating.Ampersand)

	if nn.Variadic {
		io.WriteString(p.w, "...")
	}
	p.printFreeFloating(nn, freefloating.Variadic)

	p.Print(nn.Variable)

	if nn.DefaultValue != nil {
		p.printFreeFloating(nn, freefloating.Var)
		io.WriteString(p.w, "=")
		p.Print(nn.DefaultValue)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printNodeNullable(n node.Node) {
	nn := n.(*node.Nullable)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "?")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printNodeArgument(n node.Node) {
	nn := n.(*node.Argument)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.IsReference {
		io.WriteString(p.w, "&")
	}
	p.printFreeFloating(nn, freefloating.Ampersand)

	if nn.Variadic {
		io.WriteString(p.w, "...")
	}
	p.printFreeFloating(nn, freefloating.Variadic)

	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

// name

func (p *Printer) printNameNamePart(n node.Node) {
	nn := n.(*name.NamePart)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, nn.Value)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printNameName(n node.Node) {
	nn := n.(*name.Name)
	p.printFreeFloating(nn, freefloating.Start)

	p.joinPrint("\\", nn.Parts)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printNameFullyQualified(n node.Node) {
	nn := n.(*name.FullyQualified)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "\\")
	p.joinPrint("\\", nn.Parts)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printNameRelative(n node.Node) {
	nn := n.(*name.Relative)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "namespace")
	p.printFreeFloating(nn, freefloating.Namespace)

	for _, part := range nn.Parts {
		io.WriteString(p.w, "\\")
		p.Print(part)
	}

	p.printFreeFloating(nn, freefloating.End)
}

// scalar

func (p *Printer) printScalarLNumber(n node.Node) {
	nn := n.(*scalar.Lnumber)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, nn.Value)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printScalarDNumber(n node.Node) {
	nn := n.(*scalar.Dnumber)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, nn.Value)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printScalarString(n node.Node) {
	nn := n.(*scalar.String)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, nn.Value)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printScalarEncapsedStringPart(n node.Node) {
	nn := n.(*scalar.EncapsedStringPart)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, nn.Value)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printScalarEncapsed(n node.Node) {
	nn := n.(*scalar.Encapsed)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "\"")
	for _, part := range nn.Parts {
		switch part.(type) {
		case *expr.ArrayDimFetch:
			s := (*part.GetFreeFloating())[freefloating.Start]
			if len(s) > 0 && s[0].Value == "${" {
				p.printExprArrayDimFetchWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		case *expr.Variable:
			s := (*part.GetFreeFloating())[freefloating.Start]
			if len(s) > 0 && s[0].Value == "${" {
				p.printExprVariableWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		default:
			p.Print(part)
		}
	}
	io.WriteString(p.w, "\"")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printScalarHeredoc(n node.Node) {
	nn := n.(*scalar.Heredoc)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, nn.Label)

	for _, part := range nn.Parts {
		switch part.(type) {
		case *expr.ArrayDimFetch:
			s := (*part.GetFreeFloating())[freefloating.Start]
			if len(s) > 0 && s[0].Value == "${" {
				p.printExprArrayDimFetchWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		case *expr.Variable:
			s := (*part.GetFreeFloating())[freefloating.Start]
			if len(s) > 0 && s[0].Value == "${" {
				p.printExprVariableWithoutLeadingDollar(part)
			} else {
				p.Print(part)
			}
		default:
			p.Print(part)
		}
	}

	io.WriteString(p.w, strings.Trim(nn.Label, "<\"'\n"))

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printScalarMagicConstant(n node.Node) {
	nn := n.(*scalar.MagicConstant)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, nn.Value)
	p.printFreeFloating(nn, freefloating.End)
}

// Assign

func (p *Printer) printAssign(n node.Node) {
	nn := n.(*assign.Assign)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignReference(n node.Node) {
	nn := n.(*assign.Reference)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "=")
	p.printFreeFloating(nn, freefloating.Equal)
	io.WriteString(p.w, "&")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignBitwiseAnd(n node.Node) {
	nn := n.(*assign.BitwiseAnd)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "&")
	io.WriteString(p.w, "=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignBitwiseOr(n node.Node) {
	nn := n.(*assign.BitwiseOr)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "|=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignBitwiseXor(n node.Node) {
	nn := n.(*assign.BitwiseXor)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "^=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignCoalesce(n node.Node) {
	nn := n.(*assign.Coalesce)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "??=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignConcat(n node.Node) {
	nn := n.(*assign.Concat)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, ".=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignDiv(n node.Node) {
	nn := n.(*assign.Div)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "/=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignMinus(n node.Node) {
	nn := n.(*assign.Minus)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "-=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignMod(n node.Node) {
	nn := n.(*assign.Mod)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "%=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignMul(n node.Node) {
	nn := n.(*assign.Mul)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "*=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignPlus(n node.Node) {
	nn := n.(*assign.Plus)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "+=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignPow(n node.Node) {
	nn := n.(*assign.Pow)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "**=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignShiftLeft(n node.Node) {
	nn := n.(*assign.ShiftLeft)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "<<=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printAssignShiftRight(n node.Node) {
	nn := n.(*assign.ShiftRight)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, ">>=")
	p.Print(nn.Expression)

	p.printFreeFloating(nn, freefloating.End)
}

// binary

func (p *Printer) printBinaryBitwiseAnd(n node.Node) {
	nn := n.(*binary.BitwiseAnd)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "&")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryBitwiseOr(n node.Node) {
	nn := n.(*binary.BitwiseOr)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "|")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryBitwiseXor(n node.Node) {
	nn := n.(*binary.BitwiseXor)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "^")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryBooleanAnd(n node.Node) {
	nn := n.(*binary.BooleanAnd)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "&&")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryBooleanOr(n node.Node) {
	nn := n.(*binary.BooleanOr)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "||")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryCoalesce(n node.Node) {
	nn := n.(*binary.Coalesce)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "??")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryConcat(n node.Node) {
	nn := n.(*binary.Concat)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ".")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryDiv(n node.Node) {
	nn := n.(*binary.Div)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "/")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryEqual(n node.Node) {
	nn := n.(*binary.Equal)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "==")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryGreaterOrEqual(n node.Node) {
	nn := n.(*binary.GreaterOrEqual)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ">=")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryGreater(n node.Node) {
	nn := n.(*binary.Greater)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ">")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryIdentical(n node.Node) {
	nn := n.(*binary.Identical)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "===")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryLogicalAnd(n node.Node) {
	nn := n.(*binary.LogicalAnd)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "and")
	if nn.Right.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryLogicalOr(n node.Node) {
	nn := n.(*binary.LogicalOr)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "or")
	if nn.Right.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryLogicalXor(n node.Node) {
	nn := n.(*binary.LogicalXor)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "xor")
	if nn.Right.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryMinus(n node.Node) {
	nn := n.(*binary.Minus)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "-")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryMod(n node.Node) {
	nn := n.(*binary.Mod)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "%")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryMul(n node.Node) {
	nn := n.(*binary.Mul)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "*")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryNotEqual(n node.Node) {
	nn := n.(*binary.NotEqual)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	p.printFreeFloating(nn, freefloating.Equal)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "!=")
	}
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryNotIdentical(n node.Node) {
	nn := n.(*binary.NotIdentical)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "!==")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryPlus(n node.Node) {
	nn := n.(*binary.Plus)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "+")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryPow(n node.Node) {
	nn := n.(*binary.Pow)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "**")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryShiftLeft(n node.Node) {
	nn := n.(*binary.ShiftLeft)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "<<")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinaryShiftRight(n node.Node) {
	nn := n.(*binary.ShiftRight)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ">>")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinarySmallerOrEqual(n node.Node) {
	nn := n.(*binary.SmallerOrEqual)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "<=")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinarySmaller(n node.Node) {
	nn := n.(*binary.Smaller)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "<")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBinarySpaceship(n node.Node) {
	nn := n.(*binary.Spaceship)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Left)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, "<=>")
	p.Print(nn.Right)

	p.printFreeFloating(nn, freefloating.End)
}

// cast

func (p *Printer) printArray(n node.Node) {
	nn := n.(*cast.Array)
	p.printFreeFloating(nn, freefloating.Start)

	p.printFreeFloating(nn, freefloating.Cast)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "(array)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printBool(n node.Node) {
	nn := n.(*cast.Bool)
	p.printFreeFloating(nn, freefloating.Start)

	p.printFreeFloating(nn, freefloating.Cast)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "(boolean)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printDouble(n node.Node) {
	nn := n.(*cast.Double)
	p.printFreeFloating(nn, freefloating.Start)

	p.printFreeFloating(nn, freefloating.Cast)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "(float)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printInt(n node.Node) {
	nn := n.(*cast.Int)
	p.printFreeFloating(nn, freefloating.Start)

	p.printFreeFloating(nn, freefloating.Cast)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "(integer)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printObject(n node.Node) {
	nn := n.(*cast.Object)
	p.printFreeFloating(nn, freefloating.Start)

	p.printFreeFloating(nn, freefloating.Cast)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "(object)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printString(n node.Node) {
	nn := n.(*cast.String)
	p.printFreeFloating(nn, freefloating.Start)

	p.printFreeFloating(nn, freefloating.Cast)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "(string)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printUnset(n node.Node) {
	nn := n.(*cast.Unset)
	p.printFreeFloating(nn, freefloating.Start)

	p.printFreeFloating(nn, freefloating.Cast)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "(unset)")
	}

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

// expr

func (p *Printer) printExprArrayDimFetch(n node.Node) {
	nn := n.(*expr.ArrayDimFetch)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "[")
	}
	p.Print(nn.Dim)
	p.printFreeFloating(nn, freefloating.Expr)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "]")
	}
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprArrayDimFetchWithoutLeadingDollar(n node.Node) {
	nn := n.(*expr.ArrayDimFetch)
	p.printFreeFloating(nn, freefloating.Start)
	p.printExprVariableWithoutLeadingDollar(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "[")
	}
	p.Print(nn.Dim)
	p.printFreeFloating(nn, freefloating.Expr)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "]")
	}
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprArrayItem(n node.Node) {
	nn := n.(*expr.ArrayItem)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.Unpack {
		io.WriteString(p.w, "...")
	}

	if nn.Key != nil {
		p.Print(nn.Key)
		p.printFreeFloating(nn, freefloating.Expr)
		io.WriteString(p.w, "=>")
	}

	p.Print(nn.Val)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprArray(n node.Node) {
	nn := n.(*expr.Array)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "array")
	p.printFreeFloating(nn, freefloating.Array)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, freefloating.ArrayPairList)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprArrowFunction(n node.Node) {
	nn := n.(*expr.ArrowFunction)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.Static {
		io.WriteString(p.w, "static")
	}
	p.printFreeFloating(nn, freefloating.Static)
	if nn.Static && n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "fn")
	p.printFreeFloating(nn, freefloating.Function)

	if nn.ReturnsRef {
		io.WriteString(p.w, "&")
	}
	p.printFreeFloating(nn, freefloating.Ampersand)

	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, freefloating.ParameterList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Params)

	if nn.ReturnType != nil {
		io.WriteString(p.w, ":")
		p.Print(nn.ReturnType)
	}
	p.printFreeFloating(nn, freefloating.ReturnType)

	io.WriteString(p.w, "=>")

	p.printNode(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprBitwiseNot(n node.Node) {
	nn := n.(*expr.BitwiseNot)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "~")
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprBooleanNot(n node.Node) {
	nn := n.(*expr.BooleanNot)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "!")
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprClassConstFetch(n node.Node) {
	nn := n.(*expr.ClassConstFetch)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Class)
	p.printFreeFloating(nn, freefloating.Name)
	io.WriteString(p.w, "::")
	p.Print(nn.ConstantName)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprClone(n node.Node) {
	nn := n.(*expr.Clone)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "clone")
	if nn.Expr.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprClosureUse(n node.Node) {
	nn := n.(*expr.ClosureUse)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "use")
	p.printFreeFloating(nn, freefloating.Use)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Uses)
	p.printFreeFloating(nn, freefloating.LexicalVarList)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprClosure(n node.Node) {
	nn := n.(*expr.Closure)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.Static {
		io.WriteString(p.w, "static")
	}
	p.printFreeFloating(nn, freefloating.Static)
	if nn.Static && n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "function")
	p.printFreeFloating(nn, freefloating.Function)

	if nn.ReturnsRef {
		io.WriteString(p.w, "&")
	}
	p.printFreeFloating(nn, freefloating.Ampersand)

	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, freefloating.ParameterList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Params)

	if nn.ClosureUse != nil {
		p.Print(nn.ClosureUse)
	}
	p.printFreeFloating(nn, freefloating.LexicalVars)

	if nn.ReturnType != nil {
		io.WriteString(p.w, ":")
		p.Print(nn.ReturnType)
	}
	p.printFreeFloating(nn, freefloating.ReturnType)

	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprConstFetch(n node.Node) {
	nn := n.(*expr.ConstFetch)
	p.printFreeFloating(nn, freefloating.Start)
	p.Print(nn.Constant)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprEmpty(n node.Node) {
	nn := n.(*expr.Empty)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "empty")
	p.printFreeFloating(nn, freefloating.Empty)
	io.WriteString(p.w, "(")
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprErrorSuppress(n node.Node) {
	nn := n.(*expr.ErrorSuppress)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "@")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprEval(n node.Node) {
	nn := n.(*expr.Eval)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "eval")
	p.printFreeFloating(nn, freefloating.Eval)
	io.WriteString(p.w, "(")
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprExit(n node.Node) {
	nn := n.(*expr.Exit)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.Die {
		io.WriteString(p.w, "die")
	} else {
		io.WriteString(p.w, "exit")
	}
	p.printFreeFloating(nn, freefloating.Exit)

	if nn.Expr != nil && nn.Expr.GetFreeFloating().IsEmpty() && nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprFunctionCall(n node.Node) {
	nn := n.(*expr.FunctionCall)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Function)

	p.printFreeFloating(nn.ArgumentList, freefloating.Start)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.ArgumentList.Arguments)
	p.printFreeFloating(nn.ArgumentList, freefloating.ArgumentList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn.ArgumentList, freefloating.End)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprInclude(n node.Node) {
	nn := n.(*expr.Include)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "include")
	if nn.Expr.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprIncludeOnce(n node.Node) {
	nn := n.(*expr.IncludeOnce)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "include_once")
	if nn.Expr.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprInstanceOf(n node.Node) {
	nn := n.(*expr.InstanceOf)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.Expr)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "instanceof")

	if nn.Class.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Class)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprIsset(n node.Node) {
	nn := n.(*expr.Isset)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "isset")
	p.printFreeFloating(nn, freefloating.Isset)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Variables)
	p.printFreeFloating(nn, freefloating.VarList)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprList(n node.Node) {
	nn := n.(*expr.List)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "list")
	p.printFreeFloating(nn, freefloating.List)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, freefloating.ArrayPairList)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprMethodCall(n node.Node) {
	nn := n.(*expr.MethodCall)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "->")
	p.Print(nn.Method)

	p.printFreeFloating(nn.ArgumentList, freefloating.Start)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.ArgumentList.Arguments)
	p.printFreeFloating(nn.ArgumentList, freefloating.ArgumentList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn.ArgumentList, freefloating.End)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprNew(n node.Node) {
	nn := n.(*expr.New)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "new")
	if nn.Class.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Class)

	if nn.ArgumentList != nil {
		p.printFreeFloating(nn.ArgumentList, freefloating.Start)
		io.WriteString(p.w, "(")
		p.joinPrint(",", nn.ArgumentList.Arguments)
		p.printFreeFloating(nn.ArgumentList, freefloating.ArgumentList)
		io.WriteString(p.w, ")")
		p.printFreeFloating(nn.ArgumentList, freefloating.End)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprPostDec(n node.Node) {
	nn := n.(*expr.PostDec)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "--")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprPostInc(n node.Node) {
	nn := n.(*expr.PostInc)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "++")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprPreDec(n node.Node) {
	nn := n.(*expr.PreDec)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "--")
	p.Print(nn.Variable)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprPreInc(n node.Node) {
	nn := n.(*expr.PreInc)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "++")
	p.Print(nn.Variable)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprPrint(n node.Node) {
	nn := n.(*expr.Print)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "print")
	if nn.Expr.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprPropertyFetch(n node.Node) {
	nn := n.(*expr.PropertyFetch)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, "->")
	p.Print(nn.Property)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprReference(n node.Node) {
	nn := n.(*expr.Reference)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "&")
	p.Print(nn.Variable)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprRequire(n node.Node) {
	nn := n.(*expr.Require)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "require")
	if nn.Expr.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprRequireOnce(n node.Node) {
	nn := n.(*expr.RequireOnce)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "require_once")
	if nn.Expr.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprShellExec(n node.Node) {
	nn := n.(*expr.ShellExec)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "`")
	p.joinPrint("", nn.Parts)
	io.WriteString(p.w, "`")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprShortArray(n node.Node) {
	nn := n.(*expr.ShortArray)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "[")
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, freefloating.ArrayPairList)
	io.WriteString(p.w, "]")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprShortList(n node.Node) {
	nn := n.(*expr.ShortList)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "[")
	p.joinPrint(",", nn.Items)
	p.printFreeFloating(nn, freefloating.ArrayPairList)
	io.WriteString(p.w, "]")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprStaticCall(n node.Node) {
	nn := n.(*expr.StaticCall)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Class)
	p.printFreeFloating(nn, freefloating.Name)
	io.WriteString(p.w, "::")
	p.Print(nn.Call)

	p.printFreeFloating(nn.ArgumentList, freefloating.Start)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.ArgumentList.Arguments)
	p.printFreeFloating(nn.ArgumentList, freefloating.ArgumentList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn.ArgumentList, freefloating.End)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprStaticPropertyFetch(n node.Node) {
	nn := n.(*expr.StaticPropertyFetch)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Class)
	p.printFreeFloating(nn, freefloating.Name)
	io.WriteString(p.w, "::")
	p.Print(nn.Property)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprTernary(n node.Node) {
	nn := n.(*expr.Ternary)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Condition)
	p.printFreeFloating(nn, freefloating.Cond)
	io.WriteString(p.w, "?")

	if nn.IfTrue != nil {
		p.Print(nn.IfTrue)
	}
	p.printFreeFloating(nn, freefloating.True)

	io.WriteString(p.w, ":")
	p.Print(nn.IfFalse)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprUnaryMinus(n node.Node) {
	nn := n.(*expr.UnaryMinus)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "-")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprUnaryPlus(n node.Node) {
	nn := n.(*expr.UnaryPlus)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "+")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprVariable(n node.Node) {
	nn := n.(*expr.Variable)
	p.printFreeFloating(nn, freefloating.Start)

	p.printFreeFloating(nn, freefloating.Dollar)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "$")
	}

	p.Print(nn.VarName)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprVariableWithoutLeadingDollar(n node.Node) {
	nn := n.(*expr.Variable)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.VarName)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprYieldFrom(n node.Node) {
	nn := n.(*expr.YieldFrom)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "yield from")
	if nn.Expr.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printExprYield(n node.Node) {
	nn := n.(*expr.Yield)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "yield")

	if nn.Key != nil {
		if nn.Key.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Key)
		p.printFreeFloating(nn, freefloating.Expr)
		io.WriteString(p.w, "=>")
	} else {
		if nn.Value.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Value)

	p.printFreeFloating(nn, freefloating.End)
}

// smtm

func (p *Printer) printStmtAltElseIf(n node.Node) {
	nn := n.(*stmt.AltElseIf)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "elseif")
	p.printFreeFloating(nn, freefloating.ElseIf)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Cond)
	io.WriteString(p.w, ":")

	if s := nn.Stmt.(*stmt.StmtList).Stmts; len(s) > 0 {
		p.printNodes(s)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtAltElse(n node.Node) {
	nn := n.(*stmt.AltElse)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "else")
	p.printFreeFloating(nn, freefloating.Else)
	io.WriteString(p.w, ":")

	if s := nn.Stmt.(*stmt.StmtList).Stmts; len(s) > 0 {
		p.printNodes(s)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtAltFor(n node.Node) {
	nn := n.(*stmt.AltFor)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "for")
	p.printFreeFloating(nn, freefloating.For)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Init)
	p.printFreeFloating(nn, freefloating.InitExpr)
	io.WriteString(p.w, ";")
	p.joinPrint(",", nn.Cond)
	p.printFreeFloating(nn, freefloating.CondExpr)
	io.WriteString(p.w, ";")
	p.joinPrint(",", nn.Loop)
	p.printFreeFloating(nn, freefloating.IncExpr)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Cond)
	io.WriteString(p.w, ":")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)

	io.WriteString(p.w, "endfor")
	p.printFreeFloating(nn, freefloating.AltEnd)
	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtAltForeach(n node.Node) {
	nn := n.(*stmt.AltForeach)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "foreach")
	p.printFreeFloating(nn, freefloating.Foreach)
	io.WriteString(p.w, "(")
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.Expr)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "as")

	if nn.Key != nil {
		if nn.Key.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Key)
		p.printFreeFloating(nn, freefloating.Key)
		io.WriteString(p.w, "=>")
	} else {
		if nn.Variable.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)

	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Cond)

	io.WriteString(p.w, ":")
	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)

	io.WriteString(p.w, "endforeach")
	p.printFreeFloating(nn, freefloating.AltEnd)
	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtAltIf(n node.Node) {
	nn := n.(*stmt.AltIf)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "if")
	p.printFreeFloating(nn, freefloating.If)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Cond)
	io.WriteString(p.w, ":")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)

	for _, elseif := range nn.ElseIf {
		p.Print(elseif)
	}

	if nn.Else != nil {
		p.Print(nn.Else)
	}

	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "endif")
	p.printFreeFloating(nn, freefloating.AltEnd)
	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtAltSwitch(n node.Node) {
	nn := n.(*stmt.AltSwitch)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "switch")
	p.printFreeFloating(nn, freefloating.Switch)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Cond)
	io.WriteString(p.w, ":")

	p.printFreeFloating(nn.CaseList, freefloating.Start)
	p.printFreeFloating(nn.CaseList, freefloating.CaseListStart)
	p.printNodes(nn.CaseList.Cases)
	p.printFreeFloating(nn.CaseList, freefloating.CaseListEnd)
	p.printFreeFloating(nn.CaseList, freefloating.End)

	io.WriteString(p.w, "endswitch")
	p.printFreeFloating(nn, freefloating.AltEnd)
	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtAltWhile(n node.Node) {
	nn := n.(*stmt.AltWhile)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "while")
	p.printFreeFloating(nn, freefloating.While)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Cond)
	io.WriteString(p.w, ":")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)

	io.WriteString(p.w, "endwhile")
	p.printFreeFloating(nn, freefloating.AltEnd)
	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtBreak(n node.Node) {
	nn := n.(*stmt.Break)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "break")
	if nn.Expr != nil {
		if nn.Expr.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Expr)
	}
	p.printFreeFloating(nn, freefloating.Expr)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtCase(n node.Node) {
	nn := n.(*stmt.Case)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "case")
	if nn.Cond.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Cond)
	p.printFreeFloating(nn, freefloating.Expr)
	p.printFreeFloating(nn, freefloating.CaseSeparator)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ":")
	}

	if len(nn.Stmts) > 0 {
		p.printNodes(nn.Stmts)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtCatch(n node.Node) {
	nn := n.(*stmt.Catch)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "catch")
	p.printFreeFloating(nn, freefloating.Catch)
	io.WriteString(p.w, "(")
	p.joinPrint("|", nn.Types)
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Cond)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtClassMethod(n node.Node) {
	nn := n.(*stmt.ClassMethod)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.Modifiers != nil {
		for k, m := range nn.Modifiers {
			if k > 0 && m.GetFreeFloating().IsEmpty() {
				io.WriteString(p.w, " ")
			}
			p.Print(m)
		}

		if nn.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}
	p.printFreeFloating(nn, freefloating.ModifierList)
	io.WriteString(p.w, "function")
	p.printFreeFloating(nn, freefloating.Function)

	if nn.ReturnsRef {
		if nn.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "&")
		p.printFreeFloating(nn, freefloating.Ampersand)
	} else {
		if nn.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.MethodName)
	p.printFreeFloating(nn, freefloating.Name)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, freefloating.ParameterList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Params)

	if nn.ReturnType != nil {
		io.WriteString(p.w, ":")
		p.Print(nn.ReturnType)
	}

	p.Print(nn.Stmt)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtClass(n node.Node) {
	nn := n.(*stmt.Class)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.Modifiers != nil {
		for k, m := range nn.Modifiers {
			if k > 0 && m.GetFreeFloating().IsEmpty() {
				io.WriteString(p.w, " ")
			}
			p.Print(m)
		}

		if nn.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}
	p.printFreeFloating(nn, freefloating.ModifierList)
	io.WriteString(p.w, "class")
	p.printFreeFloating(nn, freefloating.Class)

	if nn.ClassName != nil {
		if nn.ClassName.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.ClassName)
	}

	if nn.ArgumentList != nil {
		p.printFreeFloating(nn.ArgumentList, freefloating.Start)
		io.WriteString(p.w, "(")
		p.joinPrint(",", nn.ArgumentList.Arguments)
		p.printFreeFloating(nn.ArgumentList, freefloating.ArgumentList)
		io.WriteString(p.w, ")")
		p.printFreeFloating(nn.ArgumentList, freefloating.End)
	}

	if nn.Extends != nil {
		p.printFreeFloating(nn.Extends, freefloating.Start)
		if nn.Extends.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "extends")
		if nn.Extends.ClassName.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Extends.ClassName)
	}

	if nn.Implements != nil {
		p.printFreeFloating(nn.Implements, freefloating.Start)
		if nn.Implements.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "implements")
		if nn.Implements.InterfaceNames[0].GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.joinPrint(",", nn.Implements.InterfaceNames)
	}

	p.printFreeFloating(nn, freefloating.Name)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtClassConstList(n node.Node) {
	nn := n.(*stmt.ClassConstList)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.Modifiers != nil {
		for k, m := range nn.Modifiers {
			if k > 0 && m.GetFreeFloating().IsEmpty() {
				io.WriteString(p.w, " ")
			}
			p.Print(m)
		}

		if nn.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}
	p.printFreeFloating(nn, freefloating.ModifierList)
	io.WriteString(p.w, "const")

	if nn.Consts[0].GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Consts)
	p.printFreeFloating(nn, freefloating.ConstList)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtConstList(n node.Node) {
	nn := n.(*stmt.ConstList)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "const")

	if nn.Consts[0].GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Consts)
	p.printFreeFloating(nn, freefloating.Stmts)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtConstant(n node.Node) {
	nn := n.(*stmt.Constant)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.ConstantName)
	p.printFreeFloating(nn, freefloating.Name)
	io.WriteString(p.w, "=")
	p.Print(nn.Expr)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtContinue(n node.Node) {
	nn := n.(*stmt.Continue)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "continue")

	if nn.Expr != nil {
		if nn.Expr.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Expr)
	}
	p.printFreeFloating(nn, freefloating.Expr)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtDeclare(n node.Node) {
	nn := n.(*stmt.Declare)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "declare")
	p.printFreeFloating(nn, freefloating.Declare)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Consts)
	p.printFreeFloating(nn, freefloating.ConstList)
	io.WriteString(p.w, ")")

	if nn.Alt {
		p.printFreeFloating(nn, freefloating.Cond)
		io.WriteString(p.w, ":")

		s := nn.Stmt.(*stmt.StmtList)
		p.printNodes(s.Stmts)
		p.printFreeFloating(nn, freefloating.Stmts)

		io.WriteString(p.w, "enddeclare")
		p.printFreeFloating(nn, freefloating.AltEnd)

		p.printFreeFloating(nn, freefloating.SemiColon)
		if nn.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, ";")
		}
	} else {
		p.Print(nn.Stmt)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtDefault(n node.Node) {
	nn := n.(*stmt.Default)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "default")
	p.printFreeFloating(nn, freefloating.Default)
	p.printFreeFloating(nn, freefloating.CaseSeparator)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ":")
	}

	if len(nn.Stmts) > 0 {
		p.printNodes(nn.Stmts)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtDo(n node.Node) {
	nn := n.(*stmt.Do)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "do")

	if _, ok := nn.Stmt.(*stmt.StmtList); !ok {
		if nn.Stmt.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Stmt)
	p.printFreeFloating(nn, freefloating.Stmts)

	io.WriteString(p.w, "while")
	p.printFreeFloating(nn, freefloating.While)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Cond)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtEcho(n node.Node) {
	nn := n.(*stmt.Echo)

	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "echo")
	}
	if nn.Exprs[0].GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}

	p.printFreeFloating(nn, freefloating.Start)
	p.printFreeFloating(nn, freefloating.Echo)

	p.joinPrint(",", nn.Exprs)
	p.printFreeFloating(nn, freefloating.Expr)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtElseif(n node.Node) {
	nn := n.(*stmt.ElseIf)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "elseif")
	p.printFreeFloating(nn, freefloating.ElseIf)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtElse(n node.Node) {
	nn := n.(*stmt.Else)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "else")

	if _, ok := nn.Stmt.(*stmt.StmtList); !ok {
		if nn.Stmt.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.Stmt)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtExpression(n node.Node) {
	nn := n.(*stmt.Expression)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.Expr)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtFinally(n node.Node) {
	nn := n.(*stmt.Finally)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "finally")
	p.printFreeFloating(nn, freefloating.Finally)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtFor(n node.Node) {
	nn := n.(*stmt.For)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "for")
	p.printFreeFloating(nn, freefloating.For)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Init)
	p.printFreeFloating(nn, freefloating.InitExpr)
	io.WriteString(p.w, ";")
	p.joinPrint(",", nn.Cond)
	p.printFreeFloating(nn, freefloating.CondExpr)
	io.WriteString(p.w, ";")
	p.joinPrint(",", nn.Loop)
	p.printFreeFloating(nn, freefloating.IncExpr)
	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtForeach(n node.Node) {
	nn := n.(*stmt.Foreach)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "foreach")
	p.printFreeFloating(nn, freefloating.Foreach)
	io.WriteString(p.w, "(")

	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.Expr)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "as")

	if nn.Key != nil {
		if nn.Key.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Key)
		p.printFreeFloating(nn, freefloating.Key)
		io.WriteString(p.w, "=>")
	} else {
		if nn.Variable.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}
	p.Print(nn.Variable)
	p.printFreeFloating(nn, freefloating.Var)

	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtFunction(n node.Node) {
	nn := n.(*stmt.Function)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "function")
	p.printFreeFloating(nn, freefloating.Function)

	if nn.ReturnsRef {
		if nn.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "&")
	} else {
		if nn.FunctionName.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.Print(nn.FunctionName)
	p.printFreeFloating(nn, freefloating.Name)

	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Params)
	p.printFreeFloating(nn, freefloating.ParamList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.Params)

	if nn.ReturnType != nil {
		io.WriteString(p.w, ":")
		p.Print(nn.ReturnType)
	}
	p.printFreeFloating(nn, freefloating.ReturnType)

	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtGlobal(n node.Node) {
	nn := n.(*stmt.Global)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "global")
	p.joinPrint(",", nn.Vars)
	p.printFreeFloating(nn, freefloating.VarList)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtGoto(n node.Node) {
	nn := n.(*stmt.Goto)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "goto")
	if nn.Label.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Label)
	p.printFreeFloating(nn, freefloating.Label)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtGroupUse(n node.Node) {
	nn := n.(*stmt.GroupUse)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "use")
	p.printFreeFloating(nn, freefloating.Use)

	if nn.UseType != nil {
		if nn.UseType.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.UseType)
	}

	if nn.Prefix.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Prefix)
	io.WriteString(p.w, "\\")
	p.printFreeFloating(nn, freefloating.Slash)

	io.WriteString(p.w, "{")
	p.joinPrint(",", nn.UseList)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")
	p.printFreeFloating(nn, freefloating.UseDeclarationList)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtHaltCompiler(n node.Node) {
	nn := n.(*stmt.HaltCompiler)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "__halt_compiler")
	p.printFreeFloating(nn, freefloating.HaltCompiller)
	io.WriteString(p.w, "(")
	p.printFreeFloating(nn, freefloating.OpenParenthesisToken)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.CloseParenthesisToken)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtIf(n node.Node) {
	nn := n.(*stmt.If)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "if")
	p.printFreeFloating(n, freefloating.If)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printFreeFloating(n, freefloating.Expr)
	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	if nn.ElseIf != nil {
		p.printNodes(nn.ElseIf)
	}

	if nn.Else != nil {
		p.Print(nn.Else)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtInlineHTML(n node.Node) {
	nn := n.(*stmt.InlineHtml)
	p.printFreeFloating(nn, freefloating.Start)

	if p.s == PhpState && nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, "?>")
	}
	p.SetState(HtmlState)

	io.WriteString(p.w, nn.Value)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtInterface(n node.Node) {
	nn := n.(*stmt.Interface)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "interface")

	if nn.InterfaceName.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}

	p.Print(nn.InterfaceName)

	if nn.Extends != nil {
		p.printFreeFloating(nn.Extends, freefloating.Start)
		if nn.Extends.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "extends")
		if nn.Extends.InterfaceNames[0].GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.joinPrint(",", nn.Extends.InterfaceNames)
	}

	p.printFreeFloating(nn, freefloating.Name)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtLabel(n node.Node) {
	nn := n.(*stmt.Label)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.LabelName)
	p.printFreeFloating(nn, freefloating.Label)

	io.WriteString(p.w, ":")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtNamespace(n node.Node) {
	nn := n.(*stmt.Namespace)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "namespace")

	if nn.NamespaceName != nil {
		if nn.NamespaceName.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.NamespaceName)
	}

	if nn.Stmts != nil {
		p.printFreeFloating(nn, freefloating.Namespace)
		io.WriteString(p.w, "{")
		p.printNodes(nn.Stmts)
		p.printFreeFloating(nn, freefloating.Stmts)
		io.WriteString(p.w, "}")
	} else {
		p.printFreeFloating(nn, freefloating.SemiColon)
		if nn.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, ";")
		}
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtNop(n node.Node) {
	p.printFreeFloating(n, freefloating.Start)
	p.printFreeFloating(n, freefloating.SemiColon)
	if n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}
	p.printFreeFloating(n, freefloating.End)
}

func (p *Printer) printStmtPropertyList(n node.Node) {
	nn := n.(*stmt.PropertyList)
	p.printFreeFloating(nn, freefloating.Start)

	for k, m := range nn.Modifiers {
		if k > 0 && m.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(m)
	}

	if nn.Type != nil && nn.Type.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}

	p.Print(nn.Type)

	if nn.Properties[0].GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}

	p.joinPrint(",", nn.Properties)
	p.printFreeFloating(n, freefloating.PropertyList)

	p.printFreeFloating(n, freefloating.SemiColon)
	if n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtProperty(n node.Node) {
	nn := n.(*stmt.Property)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Variable)

	if nn.Expr != nil {
		p.printFreeFloating(nn, freefloating.Var)
		io.WriteString(p.w, "=")
		p.Print(nn.Expr)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtReturn(n node.Node) {
	nn := n.(*stmt.Return)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "return")
	if nn.Expr != nil && nn.Expr.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.Expr)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtStaticVar(n node.Node) {
	nn := n.(*stmt.StaticVar)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Variable)

	if nn.Expr != nil {
		p.printFreeFloating(nn, freefloating.Var)
		io.WriteString(p.w, "=")
		p.Print(nn.Expr)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtStatic(n node.Node) {
	nn := n.(*stmt.Static)
	p.printFreeFloating(nn, freefloating.Start)
	io.WriteString(p.w, "static")

	p.joinPrint(",", nn.Vars)
	p.printFreeFloating(nn, freefloating.VarList)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtStmtList(n node.Node) {
	nn := n.(*stmt.StmtList)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtSwitch(n node.Node) {
	nn := n.(*stmt.Switch)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "switch")
	p.printFreeFloating(nn, freefloating.Switch)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")

	p.printFreeFloating(nn.CaseList, freefloating.Start)
	io.WriteString(p.w, "{")
	p.printFreeFloating(nn.CaseList, freefloating.CaseListStart)
	p.printNodes(nn.CaseList.Cases)
	p.printFreeFloating(nn.CaseList, freefloating.CaseListEnd)
	io.WriteString(p.w, "}")
	p.printFreeFloating(nn.CaseList, freefloating.End)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtThrow(n node.Node) {
	nn := n.(*stmt.Throw)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "throw")
	if nn.Expr.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.Expr)
	p.printFreeFloating(nn, freefloating.Expr)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtTraitAdaptationList(n node.Node) {
	nn := n.(*stmt.TraitAdaptationList)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "{")
	p.printNodes(nn.Adaptations)
	p.printFreeFloating(nn, freefloating.AdaptationList)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtTraitMethodRef(n node.Node) {
	nn := n.(*stmt.TraitMethodRef)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.Trait != nil {
		p.Print(nn.Trait)
		p.printFreeFloating(nn, freefloating.Name)
		io.WriteString(p.w, "::")
	}

	p.Print(nn.Method)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtTraitUseAlias(n node.Node) {
	nn := n.(*stmt.TraitUseAlias)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Ref)
	p.printFreeFloating(nn, freefloating.Ref)

	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "as")

	if nn.Modifier != nil {
		if nn.Modifier.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Modifier)
	}

	if nn.Alias != nil {
		if nn.Alias.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Alias)
	}
	p.printFreeFloating(nn, freefloating.Alias)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtTraitUsePrecedence(n node.Node) {
	nn := n.(*stmt.TraitUsePrecedence)
	p.printFreeFloating(nn, freefloating.Start)

	p.Print(nn.Ref)
	p.printFreeFloating(nn, freefloating.Ref)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}

	io.WriteString(p.w, "insteadof")
	if nn.Insteadof[0].GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Insteadof)
	p.printFreeFloating(nn, freefloating.NameList)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtTraitUse(n node.Node) {
	nn := n.(*stmt.TraitUse)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "use")
	if nn.Traits[0].GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Traits)

	p.Print(nn.TraitAdaptationList)

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtTrait(n node.Node) {
	nn := n.(*stmt.Trait)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "trait")
	if nn.TraitName.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.Print(nn.TraitName)

	p.printFreeFloating(nn, freefloating.Name)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtTry(n node.Node) {
	nn := n.(*stmt.Try)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "try")
	p.printFreeFloating(nn, freefloating.Try)
	io.WriteString(p.w, "{")
	p.printNodes(nn.Stmts)
	p.printFreeFloating(nn, freefloating.Stmts)
	io.WriteString(p.w, "}")

	if nn.Catches != nil {
		p.printNodes(nn.Catches)
	}

	if nn.Finally != nil {
		p.Print(nn.Finally)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtUnset(n node.Node) {
	nn := n.(*stmt.Unset)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "unset")
	p.printFreeFloating(nn, freefloating.Unset)
	io.WriteString(p.w, "(")
	p.joinPrint(",", nn.Vars)
	p.printFreeFloating(nn, freefloating.VarList)
	io.WriteString(p.w, ")")
	p.printFreeFloating(nn, freefloating.CloseParenthesisToken)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if n.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtUseList(n node.Node) {
	nn := n.(*stmt.UseList)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "use")

	if nn.UseType != nil {
		if nn.UseType.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.UseType)
	}

	if nn.Uses[0].GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, " ")
	}
	p.joinPrint(",", nn.Uses)
	p.printFreeFloating(nn, freefloating.UseDeclarationList)

	p.printFreeFloating(nn, freefloating.SemiColon)
	if nn.GetFreeFloating().IsEmpty() {
		io.WriteString(p.w, ";")
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtUse(n node.Node) {
	nn := n.(*stmt.Use)
	p.printFreeFloating(nn, freefloating.Start)

	if nn.UseType != nil {
		p.Print(nn.UseType)
		if nn.UseType.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
	}

	p.printFreeFloating(nn, freefloating.Slash)

	p.Print(nn.Use)

	if nn.Alias != nil {
		if nn.Alias.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		io.WriteString(p.w, "as")
		if nn.Alias.GetFreeFloating().IsEmpty() {
			io.WriteString(p.w, " ")
		}
		p.Print(nn.Alias)
	}

	p.printFreeFloating(nn, freefloating.End)
}

func (p *Printer) printStmtWhile(n node.Node) {
	nn := n.(*stmt.While)
	p.printFreeFloating(nn, freefloating.Start)

	io.WriteString(p.w, "while")
	p.printFreeFloating(nn, freefloating.While)
	io.WriteString(p.w, "(")
	p.Print(nn.Cond)
	p.printFreeFloating(nn, freefloating.Expr)
	io.WriteString(p.w, ")")

	p.Print(nn.Stmt)

	p.printFreeFloating(nn, freefloating.End)
}
