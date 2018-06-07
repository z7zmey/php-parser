package printer

import (
	"io"
	"strings"

	"github.com/z7zmey/php-parser/node/stmt"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/expr/cast"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
)

type Printer struct {
	w           io.Writer
	indentStr   string
	indentDepth int
}

// NewPrinter -  Constructor for Printer
func NewPrinter(w io.Writer, indentStr string) *Printer {
	return &Printer{
		w:           w,
		indentStr:   indentStr,
		indentDepth: 0,
	}
}

func (p *Printer) Print(n node.Node) {
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

func (p *Printer) printIndent() {
	for i := 0; i < p.indentDepth; i++ {
		io.WriteString(p.w, p.indentStr)
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
		p.printReference(n)
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
	case *expr.Die:
		p.printExprDie(n)
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
	v := n.(*node.Root)

	if len(v.Stmts) > 0 {
		firstStmt := v.Stmts[0]
		v.Stmts = v.Stmts[1:]

		switch fs := firstStmt.(type) {
		case *stmt.InlineHtml:
			io.WriteString(p.w, fs.Value)
			io.WriteString(p.w, "<?php\n")
		default:
			io.WriteString(p.w, "<?php\n")
			p.printIndent()
			p.Print(fs)
			io.WriteString(p.w, "\n")
		}
	}
	p.indentDepth--
	p.printNodes(v.Stmts)
	io.WriteString(p.w, "\n")
}

func (p *Printer) printNodeIdentifier(n node.Node) {
	v := n.(*node.Identifier).Value
	io.WriteString(p.w, v)
}

func (p *Printer) printNodeParameter(n node.Node) {
	nn := n.(*node.Parameter)

	if nn.VariableType != nil {
		p.Print(nn.VariableType)
		io.WriteString(p.w, " ")
	}

	if nn.ByRef {
		io.WriteString(p.w, "&")
	}

	if nn.Variadic {
		io.WriteString(p.w, "...")
	}

	p.Print(nn.Variable)

	if nn.DefaultValue != nil {
		io.WriteString(p.w, " = ")
		p.Print(nn.DefaultValue)
	}
}

func (p *Printer) printNodeNullable(n node.Node) {
	nn := n.(*node.Nullable)

	io.WriteString(p.w, "?")
	p.Print(nn.Expr)
}

func (p *Printer) printNodeArgument(n node.Node) {
	nn := n.(*node.Argument)

	if nn.IsReference {
		io.WriteString(p.w, "&")
	}

	if nn.Variadic {
		io.WriteString(p.w, "...")
	}

	p.Print(nn.Expr)
}

// name

func (p *Printer) printNameNamePart(n node.Node) {
	v := n.(*name.NamePart).Value
	io.WriteString(p.w, v)
}

func (p *Printer) printNameName(n node.Node) {
	nn := n.(*name.Name)

	for k, part := range nn.Parts {
		if k > 0 {
			io.WriteString(p.w, "\\")
		}

		p.Print(part)
	}
}

func (p *Printer) printNameFullyQualified(n node.Node) {
	nn := n.(*name.FullyQualified)

	for _, part := range nn.Parts {
		io.WriteString(p.w, "\\")
		p.Print(part)
	}
}

func (p *Printer) printNameRelative(n node.Node) {
	nn := n.(*name.Relative)

	io.WriteString(p.w, "namespace")
	for _, part := range nn.Parts {
		io.WriteString(p.w, "\\")
		p.Print(part)
	}
}

// scalar

func (p *Printer) printScalarLNumber(n node.Node) {
	v := n.(*scalar.Lnumber).Value
	io.WriteString(p.w, v)
}

func (p *Printer) printScalarDNumber(n node.Node) {
	v := n.(*scalar.Dnumber).Value
	io.WriteString(p.w, v)
}

func (p *Printer) printScalarString(n node.Node) {
	v := n.(*scalar.String).Value

	io.WriteString(p.w, v)
}

func (p *Printer) printScalarEncapsedStringPart(n node.Node) {
	v := n.(*scalar.EncapsedStringPart).Value
	io.WriteString(p.w, v)
}

func (p *Printer) printScalarEncapsed(n node.Node) {
	io.WriteString(p.w, "\"")

	for _, nn := range n.(*scalar.Encapsed).Parts {
		p.Print(nn)
	}

	io.WriteString(p.w, "\"")
}

func (p *Printer) printScalarHeredoc(n node.Node) {
	nn := n.(*scalar.Heredoc)

	io.WriteString(p.w, "<<<")
	io.WriteString(p.w, nn.Label)
	io.WriteString(p.w, "\n")

	for _, nn := range nn.Parts {
		p.Print(nn)
	}

	io.WriteString(p.w, strings.Trim(nn.Label, "\"'"))
}

func (p *Printer) printScalarMagicConstant(n node.Node) {
	v := n.(*scalar.MagicConstant).Value
	io.WriteString(p.w, v)
}

// Assign

func (p *Printer) printAssign(n node.Node) {
	nn := n.(*assign.Assign)
	p.Print(nn.Variable)
	io.WriteString(p.w, " = ")
	p.Print(nn.Expression)
}

func (p *Printer) printReference(n node.Node) {
	nn := n.(*assign.Reference)
	p.Print(nn.Variable)
	io.WriteString(p.w, " =& ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignBitwiseAnd(n node.Node) {
	nn := n.(*assign.BitwiseAnd)
	p.Print(nn.Variable)
	io.WriteString(p.w, " &= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignBitwiseOr(n node.Node) {
	nn := n.(*assign.BitwiseOr)
	p.Print(nn.Variable)
	io.WriteString(p.w, " |= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignBitwiseXor(n node.Node) {
	nn := n.(*assign.BitwiseXor)
	p.Print(nn.Variable)
	io.WriteString(p.w, " ^= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignConcat(n node.Node) {
	nn := n.(*assign.Concat)
	p.Print(nn.Variable)
	io.WriteString(p.w, " .= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignDiv(n node.Node) {
	nn := n.(*assign.Div)
	p.Print(nn.Variable)
	io.WriteString(p.w, " /= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignMinus(n node.Node) {
	nn := n.(*assign.Minus)
	p.Print(nn.Variable)
	io.WriteString(p.w, " -= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignMod(n node.Node) {
	nn := n.(*assign.Mod)
	p.Print(nn.Variable)
	io.WriteString(p.w, " %= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignMul(n node.Node) {
	nn := n.(*assign.Mul)
	p.Print(nn.Variable)
	io.WriteString(p.w, " *= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignPlus(n node.Node) {
	nn := n.(*assign.Plus)
	p.Print(nn.Variable)
	io.WriteString(p.w, " += ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignPow(n node.Node) {
	nn := n.(*assign.Pow)
	p.Print(nn.Variable)
	io.WriteString(p.w, " **= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignShiftLeft(n node.Node) {
	nn := n.(*assign.ShiftLeft)
	p.Print(nn.Variable)
	io.WriteString(p.w, " <<= ")
	p.Print(nn.Expression)
}

func (p *Printer) printAssignShiftRight(n node.Node) {
	nn := n.(*assign.ShiftRight)
	p.Print(nn.Variable)
	io.WriteString(p.w, " >>= ")
	p.Print(nn.Expression)
}

// binary

func (p *Printer) printBinaryBitwiseAnd(n node.Node) {
	nn := n.(*binary.BitwiseAnd)

	p.Print(nn.Left)
	io.WriteString(p.w, " & ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryBitwiseOr(n node.Node) {
	nn := n.(*binary.BitwiseOr)

	p.Print(nn.Left)
	io.WriteString(p.w, " | ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryBitwiseXor(n node.Node) {
	nn := n.(*binary.BitwiseXor)

	p.Print(nn.Left)
	io.WriteString(p.w, " ^ ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryBooleanAnd(n node.Node) {
	nn := n.(*binary.BooleanAnd)

	p.Print(nn.Left)
	io.WriteString(p.w, " && ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryBooleanOr(n node.Node) {
	nn := n.(*binary.BooleanOr)

	p.Print(nn.Left)
	io.WriteString(p.w, " || ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryCoalesce(n node.Node) {
	nn := n.(*binary.Coalesce)

	p.Print(nn.Left)
	io.WriteString(p.w, " ?? ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryConcat(n node.Node) {
	nn := n.(*binary.Concat)

	p.Print(nn.Left)
	io.WriteString(p.w, " . ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryDiv(n node.Node) {
	nn := n.(*binary.Div)

	p.Print(nn.Left)
	io.WriteString(p.w, " / ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryEqual(n node.Node) {
	nn := n.(*binary.Equal)

	p.Print(nn.Left)
	io.WriteString(p.w, " == ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryGreaterOrEqual(n node.Node) {
	nn := n.(*binary.GreaterOrEqual)

	p.Print(nn.Left)
	io.WriteString(p.w, " >= ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryGreater(n node.Node) {
	nn := n.(*binary.Greater)

	p.Print(nn.Left)
	io.WriteString(p.w, " > ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryIdentical(n node.Node) {
	nn := n.(*binary.Identical)

	p.Print(nn.Left)
	io.WriteString(p.w, " === ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryLogicalAnd(n node.Node) {
	nn := n.(*binary.LogicalAnd)

	p.Print(nn.Left)
	io.WriteString(p.w, " and ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryLogicalOr(n node.Node) {
	nn := n.(*binary.LogicalOr)

	p.Print(nn.Left)
	io.WriteString(p.w, " or ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryLogicalXor(n node.Node) {
	nn := n.(*binary.LogicalXor)

	p.Print(nn.Left)
	io.WriteString(p.w, " xor ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryMinus(n node.Node) {
	nn := n.(*binary.Minus)

	p.Print(nn.Left)
	io.WriteString(p.w, " - ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryMod(n node.Node) {
	nn := n.(*binary.Mod)

	p.Print(nn.Left)
	io.WriteString(p.w, " % ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryMul(n node.Node) {
	nn := n.(*binary.Mul)

	p.Print(nn.Left)
	io.WriteString(p.w, " * ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryNotEqual(n node.Node) {
	nn := n.(*binary.NotEqual)

	p.Print(nn.Left)
	io.WriteString(p.w, " != ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryNotIdentical(n node.Node) {
	nn := n.(*binary.NotIdentical)

	p.Print(nn.Left)
	io.WriteString(p.w, " !== ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryPlus(n node.Node) {
	nn := n.(*binary.Plus)

	p.Print(nn.Left)
	io.WriteString(p.w, " + ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryPow(n node.Node) {
	nn := n.(*binary.Pow)

	p.Print(nn.Left)
	io.WriteString(p.w, " ** ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryShiftLeft(n node.Node) {
	nn := n.(*binary.ShiftLeft)

	p.Print(nn.Left)
	io.WriteString(p.w, " << ")
	p.Print(nn.Right)
}

func (p *Printer) printBinaryShiftRight(n node.Node) {
	nn := n.(*binary.ShiftRight)

	p.Print(nn.Left)
	io.WriteString(p.w, " >> ")
	p.Print(nn.Right)
}

func (p *Printer) printBinarySmallerOrEqual(n node.Node) {
	nn := n.(*binary.SmallerOrEqual)

	p.Print(nn.Left)
	io.WriteString(p.w, " <= ")
	p.Print(nn.Right)
}

func (p *Printer) printBinarySmaller(n node.Node) {
	nn := n.(*binary.Smaller)

	p.Print(nn.Left)
	io.WriteString(p.w, " < ")
	p.Print(nn.Right)
}

func (p *Printer) printBinarySpaceship(n node.Node) {
	nn := n.(*binary.Spaceship)

	p.Print(nn.Left)
	io.WriteString(p.w, " <=> ")
	p.Print(nn.Right)
}

// cast

func (p *Printer) printArray(n node.Node) {
	nn := n.(*cast.Array)

	io.WriteString(p.w, "(array)")
	p.Print(nn.Expr)
}

func (p *Printer) printBool(n node.Node) {
	nn := n.(*cast.Bool)

	io.WriteString(p.w, "(bool)")
	p.Print(nn.Expr)
}

func (p *Printer) printDouble(n node.Node) {
	nn := n.(*cast.Double)

	io.WriteString(p.w, "(float)")
	p.Print(nn.Expr)
}

func (p *Printer) printInt(n node.Node) {
	nn := n.(*cast.Int)

	io.WriteString(p.w, "(int)")
	p.Print(nn.Expr)
}

func (p *Printer) printObject(n node.Node) {
	nn := n.(*cast.Object)

	io.WriteString(p.w, "(object)")
	p.Print(nn.Expr)
}

func (p *Printer) printString(n node.Node) {
	nn := n.(*cast.String)

	io.WriteString(p.w, "(string)")
	p.Print(nn.Expr)
}

func (p *Printer) printUnset(n node.Node) {
	nn := n.(*cast.Unset)

	io.WriteString(p.w, "(unset)")
	p.Print(nn.Expr)
}

// expr

func (p *Printer) printExprArrayDimFetch(n node.Node) {
	nn := n.(*expr.ArrayDimFetch)
	p.Print(nn.Variable)
	io.WriteString(p.w, "[")
	p.Print(nn.Dim)
	io.WriteString(p.w, "]")
}

func (p *Printer) printExprArrayItem(n node.Node) {
	nn := n.(*expr.ArrayItem)

	if nn.Key != nil {
		p.Print(nn.Key)
		io.WriteString(p.w, " => ")
	}

	p.Print(nn.Val)
}

func (p *Printer) printExprArray(n node.Node) {
	nn := n.(*expr.Array)

	io.WriteString(p.w, "array(")
	p.joinPrint(", ", nn.Items)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprBitwiseNot(n node.Node) {
	nn := n.(*expr.BitwiseNot)
	io.WriteString(p.w, "~")
	p.Print(nn.Expr)
}

func (p *Printer) printExprBooleanNot(n node.Node) {
	nn := n.(*expr.BooleanNot)
	io.WriteString(p.w, "!")
	p.Print(nn.Expr)
}

func (p *Printer) printExprClassConstFetch(n node.Node) {
	nn := n.(*expr.ClassConstFetch)

	p.Print(nn.Class)
	io.WriteString(p.w, "::")
	io.WriteString(p.w, nn.ConstantName.(*node.Identifier).Value)
}

func (p *Printer) printExprClone(n node.Node) {
	nn := n.(*expr.Clone)

	io.WriteString(p.w, "clone ")
	p.Print(nn.Expr)
}

func (p *Printer) printExprClosureUse(n node.Node) {
	nn := n.(*expr.ClosureUse)

	io.WriteString(p.w, "use (")
	p.joinPrint(", ", nn.Uses)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprClosure(n node.Node) {
	nn := n.(*expr.Closure)

	if nn.Static {
		io.WriteString(p.w, "static ")
	}

	io.WriteString(p.w, "function ")

	if nn.ReturnsRef {
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

func (p *Printer) printExprConstFetch(n node.Node) {
	nn := n.(*expr.ConstFetch)

	p.Print(nn.Constant)
}

func (p *Printer) printExprDie(n node.Node) {
	nn := n.(*expr.Die)

	io.WriteString(p.w, "die(")
	p.Print(nn.Expr)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprEmpty(n node.Node) {
	nn := n.(*expr.Empty)

	io.WriteString(p.w, "empty(")
	p.Print(nn.Expr)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprErrorSuppress(n node.Node) {
	nn := n.(*expr.ErrorSuppress)

	io.WriteString(p.w, "@")
	p.Print(nn.Expr)
}

func (p *Printer) printExprEval(n node.Node) {
	nn := n.(*expr.Eval)

	io.WriteString(p.w, "eval(")
	p.Print(nn.Expr)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprExit(n node.Node) {
	nn := n.(*expr.Exit)

	io.WriteString(p.w, "exit(")
	p.Print(nn.Expr)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprFunctionCall(n node.Node) {
	nn := n.(*expr.FunctionCall)

	p.Print(nn.Function)
	io.WriteString(p.w, "(")
	p.joinPrint(", ", nn.ArgumentList.Arguments)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprInclude(n node.Node) {
	nn := n.(*expr.Include)

	io.WriteString(p.w, "include ")
	p.Print(nn.Expr)
}

func (p *Printer) printExprIncludeOnce(n node.Node) {
	nn := n.(*expr.IncludeOnce)

	io.WriteString(p.w, "include_once ")
	p.Print(nn.Expr)
}

func (p *Printer) printExprInstanceOf(n node.Node) {
	nn := n.(*expr.InstanceOf)

	p.Print(nn.Expr)
	io.WriteString(p.w, " instanceof ")
	p.Print(nn.Class)
}

func (p *Printer) printExprIsset(n node.Node) {
	nn := n.(*expr.Isset)

	io.WriteString(p.w, "isset(")
	p.joinPrint(", ", nn.Variables)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprList(n node.Node) {
	nn := n.(*expr.List)

	io.WriteString(p.w, "list(")
	p.joinPrint(", ", nn.Items)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprMethodCall(n node.Node) {
	nn := n.(*expr.MethodCall)

	p.Print(nn.Variable)
	io.WriteString(p.w, "->")
	p.Print(nn.Method)
	io.WriteString(p.w, "(")
	p.joinPrint(", ", nn.ArgumentList.Arguments)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprNew(n node.Node) {
	nn := n.(*expr.New)

	io.WriteString(p.w, "new ")
	p.Print(nn.Class)

	if nn.ArgumentList != nil {
		io.WriteString(p.w, "(")
		p.joinPrint(", ", nn.ArgumentList.Arguments)
		io.WriteString(p.w, ")")
	}
}

func (p *Printer) printExprPostDec(n node.Node) {
	nn := n.(*expr.PostDec)

	p.Print(nn.Variable)
	io.WriteString(p.w, "--")
}

func (p *Printer) printExprPostInc(n node.Node) {
	nn := n.(*expr.PostInc)

	p.Print(nn.Variable)
	io.WriteString(p.w, "++")
}

func (p *Printer) printExprPreDec(n node.Node) {
	nn := n.(*expr.PreDec)

	io.WriteString(p.w, "--")
	p.Print(nn.Variable)
}

func (p *Printer) printExprPreInc(n node.Node) {
	nn := n.(*expr.PreInc)

	io.WriteString(p.w, "++")
	p.Print(nn.Variable)
}

func (p *Printer) printExprPrint(n node.Node) {
	nn := n.(*expr.Print)

	io.WriteString(p.w, "print(")
	p.Print(nn.Expr)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprPropertyFetch(n node.Node) {
	nn := n.(*expr.PropertyFetch)

	p.Print(nn.Variable)
	io.WriteString(p.w, "->")
	p.Print(nn.Property)
}

func (p *Printer) printExprReference(n node.Node) {
	nn := n.(*expr.Reference)

	io.WriteString(p.w, "&")
	p.Print(nn.Variable)
}

func (p *Printer) printExprRequire(n node.Node) {
	nn := n.(*expr.Require)

	io.WriteString(p.w, "require ")
	p.Print(nn.Expr)
}

func (p *Printer) printExprRequireOnce(n node.Node) {
	nn := n.(*expr.RequireOnce)

	io.WriteString(p.w, "require_once ")
	p.Print(nn.Expr)
}

func (p *Printer) printExprShellExec(n node.Node) {
	nn := n.(*expr.ShellExec)

	io.WriteString(p.w, "`")
	for _, part := range nn.Parts {
		p.Print(part)
	}
	io.WriteString(p.w, "`")
}

func (p *Printer) printExprShortArray(n node.Node) {
	nn := n.(*expr.ShortArray)

	io.WriteString(p.w, "[")
	p.joinPrint(", ", nn.Items)
	io.WriteString(p.w, "]")
}

func (p *Printer) printExprShortList(n node.Node) {
	nn := n.(*expr.ShortList)

	io.WriteString(p.w, "[")
	p.joinPrint(", ", nn.Items)
	io.WriteString(p.w, "]")
}

func (p *Printer) printExprStaticCall(n node.Node) {
	nn := n.(*expr.StaticCall)

	p.Print(nn.Class)
	io.WriteString(p.w, "::")
	p.Print(nn.Call)
	io.WriteString(p.w, "(")
	p.joinPrint(", ", nn.ArgumentList.Arguments)
	io.WriteString(p.w, ")")
}

func (p *Printer) printExprStaticPropertyFetch(n node.Node) {
	nn := n.(*expr.StaticPropertyFetch)

	p.Print(nn.Class)
	io.WriteString(p.w, "::")
	p.Print(nn.Property)
}

func (p *Printer) printExprTernary(n node.Node) {
	nn := n.(*expr.Ternary)

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

func (p *Printer) printExprUnaryMinus(n node.Node) {
	nn := n.(*expr.UnaryMinus)

	io.WriteString(p.w, "-")
	p.Print(nn.Expr)
}

func (p *Printer) printExprUnaryPlus(n node.Node) {
	nn := n.(*expr.UnaryPlus)

	io.WriteString(p.w, "+")
	p.Print(nn.Expr)
}

func (p *Printer) printExprVariable(n node.Node) {
	io.WriteString(p.w, "$")
	p.Print(n.(*expr.Variable).VarName)
}

func (p *Printer) printExprYieldFrom(n node.Node) {
	nn := n.(*expr.YieldFrom)

	io.WriteString(p.w, "yield from ")
	p.Print(nn.Expr)
}

func (p *Printer) printExprYield(n node.Node) {
	nn := n.(*expr.Yield)

	io.WriteString(p.w, "yield ")

	if nn.Key != nil {
		p.Print(nn.Key)
		io.WriteString(p.w, " => ")
	}

	p.Print(nn.Value)
}

// smtm

func (p *Printer) printStmtAltElseIf(n node.Node) {
	nn := n.(*stmt.AltElseIf)

	io.WriteString(p.w, "elseif (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ") :")

	if s := nn.Stmt.(*stmt.StmtList).Stmts; len(s) > 0 {
		io.WriteString(p.w, "\n")
		p.printNodes(s)
	}
}

func (p *Printer) printStmtAltElse(n node.Node) {
	nn := n.(*stmt.AltElse)

	io.WriteString(p.w, "else :")

	if s := nn.Stmt.(*stmt.StmtList).Stmts; len(s) > 0 {
		io.WriteString(p.w, "\n")
		p.printNodes(s)
	}
}

func (p *Printer) printStmtAltFor(n node.Node) {
	nn := n.(*stmt.AltFor)

	io.WriteString(p.w, "for (")
	p.joinPrint(", ", nn.Init)
	io.WriteString(p.w, "; ")
	p.joinPrint(", ", nn.Cond)
	io.WriteString(p.w, "; ")
	p.joinPrint(", ", nn.Loop)
	io.WriteString(p.w, ") :\n")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()

	io.WriteString(p.w, "endfor;")
}

func (p *Printer) printStmtAltForeach(n node.Node) {
	nn := n.(*stmt.AltForeach)

	io.WriteString(p.w, "foreach (")
	p.Print(nn.Expr)
	io.WriteString(p.w, " as ")

	if nn.Key != nil {
		p.Print(nn.Key)
		io.WriteString(p.w, " => ")
	}

	p.Print(nn.Variable)

	io.WriteString(p.w, ") :\n")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "endforeach;")
}

func (p *Printer) printStmtAltIf(n node.Node) {
	nn := n.(*stmt.AltIf)

	io.WriteString(p.w, "if (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ") :\n")

	s := nn.Stmt.(*stmt.StmtList)
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

func (p *Printer) printStmtAltSwitch(n node.Node) {
	nn := n.(*stmt.AltSwitch)

	io.WriteString(p.w, "switch (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ") :\n")

	s := nn.CaseList.Cases
	p.printNodes(s)

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "endswitch;")
}

func (p *Printer) printStmtAltWhile(n node.Node) {
	nn := n.(*stmt.AltWhile)

	io.WriteString(p.w, "while (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ") :\n")

	s := nn.Stmt.(*stmt.StmtList)
	p.printNodes(s.Stmts)

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "endwhile;")
}

func (p *Printer) printStmtBreak(n node.Node) {
	nn := n.(*stmt.Break)

	io.WriteString(p.w, "break")
	if nn.Expr != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.Expr)
	}

	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtCase(n node.Node) {
	nn := n.(*stmt.Case)

	io.WriteString(p.w, "case ")
	p.Print(nn.Cond)
	io.WriteString(p.w, ":")

	if len(nn.Stmts) > 0 {
		io.WriteString(p.w, "\n")
		p.printNodes(nn.Stmts)
	}
}

func (p *Printer) printStmtCatch(n node.Node) {
	nn := n.(*stmt.Catch)

	io.WriteString(p.w, "catch (")
	p.joinPrint(" | ", nn.Types)
	io.WriteString(p.w, " ")
	p.Print(nn.Variable)
	io.WriteString(p.w, ") {\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *Printer) printStmtClassMethod(n node.Node) {
	nn := n.(*stmt.ClassMethod)

	if nn.Modifiers != nil {
		p.joinPrint(" ", nn.Modifiers)
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "function ")

	if nn.ReturnsRef {
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
	case *stmt.StmtList:
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

func (p *Printer) printStmtClass(n node.Node) {
	nn := n.(*stmt.Class)

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
		p.joinPrint(", ", nn.ArgumentList.Arguments)
		io.WriteString(p.w, ")")
	}

	if nn.Extends != nil {
		io.WriteString(p.w, " extends ")
		p.Print(nn.Extends.ClassName)
	}

	if nn.Implements != nil {
		io.WriteString(p.w, " implements ")
		p.joinPrint(", ", nn.Implements.InterfaceNames)
	}

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "{\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *Printer) printStmtClassConstList(n node.Node) {
	nn := n.(*stmt.ClassConstList)

	if nn.Modifiers != nil {
		p.joinPrint(" ", nn.Modifiers)
		io.WriteString(p.w, " ")
	}
	io.WriteString(p.w, "const ")

	p.joinPrint(", ", nn.Consts)

	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtConstant(n node.Node) {
	nn := n.(*stmt.Constant)

	p.Print(nn.ConstantName)
	io.WriteString(p.w, " = ")
	p.Print(nn.Expr)
}

func (p *Printer) printStmtContinue(n node.Node) {
	nn := n.(*stmt.Continue)

	io.WriteString(p.w, "continue")
	if nn.Expr != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.Expr)
	}

	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtDeclare(n node.Node) {
	nn := n.(*stmt.Declare)

	io.WriteString(p.w, "declare(")
	p.joinPrint(", ", nn.Consts)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		p.Print(s)
		break
	case *stmt.StmtList:
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

func (p *Printer) printStmtDefault(n node.Node) {
	nn := n.(*stmt.Default)
	io.WriteString(p.w, "default:")

	if len(nn.Stmts) > 0 {
		io.WriteString(p.w, "\n")
		p.printNodes(nn.Stmts)
	}
}

func (p *Printer) printStmtDo(n node.Node) {
	nn := n.(*stmt.Do)
	io.WriteString(p.w, "do")

	switch s := nn.Stmt.(type) {
	case *stmt.StmtList:
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

func (p *Printer) printStmtEcho(n node.Node) {
	nn := n.(*stmt.Echo)
	io.WriteString(p.w, "echo ")
	p.joinPrint(", ", nn.Exprs)
	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtElseif(n node.Node) {
	nn := n.(*stmt.ElseIf)

	io.WriteString(p.w, "elseif (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		p.Print(s)
		break
	case *stmt.StmtList:
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

func (p *Printer) printStmtElse(n node.Node) {
	nn := n.(*stmt.Else)

	io.WriteString(p.w, "else")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		p.Print(s)
		break
	case *stmt.StmtList:
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

func (p *Printer) printStmtExpression(n node.Node) {
	nn := n.(*stmt.Expression)

	p.Print(nn.Expr)

	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtFinally(n node.Node) {
	nn := n.(*stmt.Finally)

	io.WriteString(p.w, "finally {\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *Printer) printStmtFor(n node.Node) {
	nn := n.(*stmt.For)

	io.WriteString(p.w, "for (")
	p.joinPrint(", ", nn.Init)
	io.WriteString(p.w, "; ")
	p.joinPrint(", ", nn.Cond)
	io.WriteString(p.w, "; ")
	p.joinPrint(", ", nn.Loop)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		p.Print(s)
		break
	case *stmt.StmtList:
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

func (p *Printer) printStmtForeach(n node.Node) {
	nn := n.(*stmt.Foreach)

	io.WriteString(p.w, "foreach (")
	p.Print(nn.Expr)
	io.WriteString(p.w, " as ")

	if nn.Key != nil {
		p.Print(nn.Key)
		io.WriteString(p.w, " => ")
	}

	p.Print(nn.Variable)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		p.Print(s)
		break
	case *stmt.StmtList:
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

func (p *Printer) printStmtFunction(n node.Node) {
	nn := n.(*stmt.Function)

	io.WriteString(p.w, "function ")

	if nn.ReturnsRef {
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

func (p *Printer) printStmtGlobal(n node.Node) {
	nn := n.(*stmt.Global)

	io.WriteString(p.w, "global ")
	p.joinPrint(", ", nn.Vars)
	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtGoto(n node.Node) {
	nn := n.(*stmt.Goto)

	io.WriteString(p.w, "goto ")
	p.Print(nn.Label)
	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtGroupUse(n node.Node) {
	nn := n.(*stmt.GroupUse)

	io.WriteString(p.w, "use ")

	if nn.UseType != nil {
		p.Print(nn.UseType)
		io.WriteString(p.w, " ")
	}

	p.Print(nn.Prefix)
	io.WriteString(p.w, "\\{")
	p.joinPrint(", ", nn.UseList)
	io.WriteString(p.w, "};")
}

func (p *Printer) printStmtHaltCompiler(n node.Node) {
	io.WriteString(p.w, "__halt_compiler();")
}

func (p *Printer) printStmtIf(n node.Node) {
	nn := n.(*stmt.If)

	io.WriteString(p.w, "if (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		p.Print(s)
		break
	case *stmt.StmtList:
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

func (p *Printer) printStmtInlineHTML(n node.Node) {
	nn := n.(*stmt.InlineHtml)

	io.WriteString(p.w, "?>")
	io.WriteString(p.w, nn.Value)
	io.WriteString(p.w, "<?php")
}

func (p *Printer) printStmtInterface(n node.Node) {
	nn := n.(*stmt.Interface)

	io.WriteString(p.w, "interface")

	if nn.InterfaceName != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.InterfaceName)
	}

	if nn.Extends != nil {
		io.WriteString(p.w, " extends ")
		p.joinPrint(", ", nn.Extends.InterfaceNames)
	}

	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "{\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *Printer) printStmtLabel(n node.Node) {
	nn := n.(*stmt.Label)

	p.Print(nn.LabelName)
	io.WriteString(p.w, ":")
}

func (p *Printer) printStmtNamespace(n node.Node) {
	nn := n.(*stmt.Namespace)

	io.WriteString(p.w, "namespace")

	if nn.NamespaceName != nil {
		io.WriteString(p.w, " ")
		p.Print(nn.NamespaceName)
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

func (p *Printer) printStmtNop(n node.Node) {
	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtPropertyList(n node.Node) {
	nn := n.(*stmt.PropertyList)

	p.joinPrint(" ", nn.Modifiers)
	io.WriteString(p.w, " ")
	p.joinPrint(", ", nn.Properties)
	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtProperty(n node.Node) {
	nn := n.(*stmt.Property)

	p.Print(nn.Variable)

	if nn.Expr != nil {
		io.WriteString(p.w, " = ")
		p.Print(nn.Expr)
	}
}

func (p *Printer) printStmtReturn(n node.Node) {
	nn := n.(*stmt.Return)

	io.WriteString(p.w, "return ")
	p.Print(nn.Expr)
	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtStaticVar(n node.Node) {
	nn := n.(*stmt.StaticVar)
	p.Print(nn.Variable)

	if nn.Expr != nil {
		io.WriteString(p.w, " = ")
		p.Print(nn.Expr)
	}
}

func (p *Printer) printStmtStatic(n node.Node) {
	nn := n.(*stmt.Static)

	io.WriteString(p.w, "static ")
	p.joinPrint(", ", nn.Vars)
	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtStmtList(n node.Node) {
	nn := n.(*stmt.StmtList)

	io.WriteString(p.w, "{\n")
	p.printNodes(nn.Stmts)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *Printer) printStmtSwitch(n node.Node) {
	nn := n.(*stmt.Switch)

	io.WriteString(p.w, "switch (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ")")

	io.WriteString(p.w, " {\n")
	p.printNodes(nn.CaseList.Cases)
	io.WriteString(p.w, "\n")
	p.printIndent()
	io.WriteString(p.w, "}")
}

func (p *Printer) printStmtThrow(n node.Node) {
	nn := n.(*stmt.Throw)

	io.WriteString(p.w, "throw ")
	p.Print(nn.Expr)
	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtTraitMethodRef(n node.Node) {
	nn := n.(*stmt.TraitMethodRef)

	p.Print(nn.Trait)
	io.WriteString(p.w, "::")
	p.Print(nn.Method)
}

func (p *Printer) printStmtTraitUseAlias(n node.Node) {
	nn := n.(*stmt.TraitUseAlias)

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

func (p *Printer) printStmtTraitUsePrecedence(n node.Node) {
	nn := n.(*stmt.TraitUsePrecedence)

	p.Print(nn.Ref)
	io.WriteString(p.w, " insteadof ")
	p.joinPrint(", ", nn.Insteadof)

	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtTraitUse(n node.Node) {
	nn := n.(*stmt.TraitUse)

	io.WriteString(p.w, "use ")
	p.joinPrint(", ", nn.Traits)

	if nn.TraitAdaptationList != nil {
		adaptations := nn.TraitAdaptationList.Adaptations
		io.WriteString(p.w, " {\n")
		p.printNodes(adaptations)
		io.WriteString(p.w, "\n")
		p.printIndent()
		io.WriteString(p.w, "}")
	} else {
		io.WriteString(p.w, ";")
	}
}

func (p *Printer) printStmtTrait(n node.Node) {
	nn := n.(*stmt.Trait)

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

func (p *Printer) printStmtTry(n node.Node) {
	nn := n.(*stmt.Try)

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

func (p *Printer) printStmtUnset(n node.Node) {
	nn := n.(*stmt.Unset)

	io.WriteString(p.w, "unset(")
	p.joinPrint(", ", nn.Vars)
	io.WriteString(p.w, ");")
}

func (p *Printer) printStmtUseList(n node.Node) {
	nn := n.(*stmt.UseList)

	io.WriteString(p.w, "use ")

	if nn.UseType != nil {
		p.Print(nn.UseType)
		io.WriteString(p.w, " ")
	}

	p.joinPrint(", ", nn.Uses)
	io.WriteString(p.w, ";")
}

func (p *Printer) printStmtUse(n node.Node) {
	nn := n.(*stmt.Use)

	if nn.UseType != nil {
		p.Print(nn.UseType)
		io.WriteString(p.w, " ")
	}

	p.Print(nn.Use)

	if nn.Alias != nil {
		io.WriteString(p.w, " as ")
		p.Print(nn.Alias)
	}
}

func (p *Printer) printStmtWhile(n node.Node) {
	nn := n.(*stmt.While)

	io.WriteString(p.w, "while (")
	p.Print(nn.Cond)
	io.WriteString(p.w, ")")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		p.Print(s)
		break
	case *stmt.StmtList:
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
