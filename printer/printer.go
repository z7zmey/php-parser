package printer

import (
	"io"

	"github.com/z7zmey/php-parser/node/stmt"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/expr/cast"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/scalar"
)

func Print(o io.Writer, n node.Node) {
	fn := getPrintFuncByNode(n)
	fn(o, n)
}

func joinPrint(glue string, o io.Writer, nn []node.Node) {
	for k, n := range nn {
		if k > 0 {
			io.WriteString(o, glue)
		}

		Print(o, n)
	}
}

func printNodes(o io.Writer, nn []node.Node) {
	// TODO: handle indentations
	for _, n := range nn {
		Print(o, n)
	}
}

func getPrintFuncByNode(n node.Node) func(o io.Writer, n node.Node) {
	switch n.(type) {

	// node

	case *node.Identifier:
		return printNodeIdentifier
	case *node.Parameter:
		return printNodeParameter
	case *node.Nullable:
		return printNodeNullable
	case *node.Argument:
		return printNodeArgument

	// name

	case *name.NamePart:
		return printNameNamePart
	case *name.Name:
		return printNameName
	case *name.FullyQualified:
		return printNameFullyQualified
	case *name.Relative:
		return printNameRelative

	// scalar

	case *scalar.Lnumber:
		return printScalarLNumber
	case *scalar.Dnumber:
		return printScalarDNumber
	case *scalar.String:
		return printScalarString
	case *scalar.EncapsedStringPart:
		return printScalarEncapsedStringPart
	case *scalar.Encapsed:
		return printScalarEncapsed
	case *scalar.MagicConstant:
		return printScalarMagicConstant

	// assign

	case *assign.Assign:
		return printAssign
	case *assign.AssignRef:
		return printAssignRef
	case *assign.BitwiseAnd:
		return printAssignBitwiseAnd
	case *assign.BitwiseOr:
		return printAssignBitwiseOr
	case *assign.BitwiseXor:
		return printAssignBitwiseXor
	case *assign.Concat:
		return printAssignConcat
	case *assign.Div:
		return printAssignDiv
	case *assign.Minus:
		return printAssignMinus
	case *assign.Mod:
		return printAssignMod
	case *assign.Mul:
		return printAssignMul
	case *assign.Plus:
		return printAssignPlus
	case *assign.Pow:
		return printAssignPow
	case *assign.ShiftLeft:
		return printAssignShiftLeft
	case *assign.ShiftRight:
		return printAssignShiftRight

	// binary

	case *binary.BitwiseAnd:
		return printBinaryBitwiseAnd
	case *binary.BitwiseOr:
		return printBinaryBitwiseOr
	case *binary.BitwiseXor:
		return printBinaryBitwiseXor
	case *binary.BooleanAnd:
		return printBinaryBooleanAnd
	case *binary.BooleanOr:
		return printBinaryBooleanOr
	case *binary.Coalesce:
		return printBinaryCoalesce
	case *binary.Concat:
		return printBinaryConcat
	case *binary.Div:
		return printBinaryDiv
	case *binary.Equal:
		return printBinaryEqual
	case *binary.GreaterOrEqual:
		return printBinaryGreaterOrEqual
	case *binary.Greater:
		return printBinaryGreater
	case *binary.Identical:
		return printBinaryIdentical
	case *binary.LogicalAnd:
		return printBinaryLogicalAnd
	case *binary.LogicalOr:
		return printBinaryLogicalOr
	case *binary.LogicalXor:
		return printBinaryLogicalXor
	case *binary.Minus:
		return printBinaryMinus
	case *binary.Mod:
		return printBinaryMod
	case *binary.Mul:
		return printBinaryMul
	case *binary.NotEqual:
		return printBinaryNotEqual
	case *binary.NotIdentical:
		return printBinaryNotIdentical
	case *binary.Plus:
		return printBinaryPlus
	case *binary.Pow:
		return printBinaryPow
	case *binary.ShiftLeft:
		return printBinaryShiftLeft
	case *binary.ShiftRight:
		return printBinaryShiftRight
	case *binary.SmallerOrEqual:
		return printBinarySmallerOrEqual
	case *binary.Smaller:
		return printBinarySmaller
	case *binary.Spaceship:
		return printBinarySpaceship

	// cast

	case *cast.CastArray:
		return printCastArray
	case *cast.CastBool:
		return printCastBool
	case *cast.CastDouble:
		return printCastDouble
	case *cast.CastInt:
		return printCastInt
	case *cast.CastObject:
		return printCastObject
	case *cast.CastString:
		return printCastString
	case *cast.CastUnset:
		return printCastUnset

	// expr

	case *expr.ArrayDimFetch:
		return printExprArrayDimFetch
	case *expr.ArrayItem:
		return printExprArrayItem
	case *expr.Array:
		return printExprArray
	case *expr.BitwiseNot:
		return printExprBitwiseNot
	case *expr.BooleanNot:
		return printExprBooleanNot
	case *expr.ClassConstFetch:
		return printExprClassConstFetch
	case *expr.Clone:
		return printExprClone
	case *expr.ClosureUse:
		return printExprClosureUse
	case *expr.Closure:
		return printExprClosure
	case *expr.ConstFetch:
		return printExprConstFetch
	case *expr.Die:
		return printExprDie
	case *expr.Empty:
		return printExprEmpty
	case *expr.ErrorSuppress:
		return printExprErrorSuppress
	case *expr.Eval:
		return printExprEval
	case *expr.Exit:
		return printExprExit
	case *expr.FunctionCall:
		return printExprFunctionCall
	case *expr.Include:
		return printExprInclude
	case *expr.IncludeOnce:
		return printExprIncludeOnce
	case *expr.InstanceOf:
		return printExprInstanceOf
	case *expr.Isset:
		return printExprIsset
	case *expr.List:
		return printExprList
	case *expr.MethodCall:
		return printExprMethodCall
	case *expr.New:
		return printExprNew
	case *expr.PostDec:
		return printExprPostDec
	case *expr.PostInc:
		return printExprPostInc
	case *expr.PreDec:
		return printExprPreDec
	case *expr.PreInc:
		return printExprPreInc
	case *expr.Print:
		return printExprPrint
	case *expr.PropertyFetch:
		return printExprPropertyFetch
	case *expr.Require:
		return printExprRequire
	case *expr.RequireOnce:
		return printExprRequireOnce
	case *expr.ShellExec:
		return printExprShellExec
	case *expr.ShortArray:
		return printExprShortArray
	case *expr.ShortList:
		return printExprShortList
	case *expr.StaticCall:
		return printExprStaticCall
	case *expr.StaticPropertyFetch:
		return printExprStaticPropertyFetch
	case *expr.Ternary:
		return printExprTernary
	case *expr.UnaryMinus:
		return printExprUnaryMinus
	case *expr.UnaryPlus:
		return printExprUnaryPlus
	case *expr.Variable:
		return printExprVariable
	case *expr.YieldFrom:
		return printExprYieldFrom
	case *expr.Yield:
		return printExprYield

	// stmt

	case *stmt.AltElseIf:
		return printStmtAltElseIf
	case *stmt.AltElse:
		return printStmtAltElse
	case *stmt.AltFor:
		return printStmtAltFor
	case *stmt.AltForeach:
		return printStmtAltForeach
	case *stmt.AltIf:
		return printStmtAltIf
	case *stmt.AltSwitch:
		return printStmtAltSwitch
	case *stmt.AltWhile:
		return printStmtAltWhile
	case *stmt.Break:
		return printStmtBreak
	case *stmt.Case:
		return printStmtCase
	case *stmt.Catch:
		return printStmtCatch
	case *stmt.ClassMethod:
		return printStmtClassMethod
	case *stmt.Class:
		return printStmtClass
	case *stmt.ClassConstList:
		return printStmtClassConstList
	case *stmt.Constant:
		return printStmtConstant
	case *stmt.Continue:
		return printStmtContinue
	case *stmt.Declare:
		return printStmtDeclare
	case *stmt.Default:
		return printStmtDefault
	case *stmt.Do:
		return printStmtDo
	case *stmt.Echo:
		return printStmtEcho
	case *stmt.ElseIf:
		return printStmtElseif
	case *stmt.Else:
		return printStmtElse
	case *stmt.Expression:
		return printStmtExpression
	case *stmt.Finally:
		return printStmtFinally

	case *stmt.StmtList:
		return printStmtStmtList
	case *stmt.Nop:
		return printStmtNop
	}

	panic("printer is missing for the node")
}

// node

func printNodeIdentifier(o io.Writer, n node.Node) {
	v := n.(*node.Identifier).Value
	io.WriteString(o, v)
}

func printNodeParameter(o io.Writer, n node.Node) {
	nn := n.(*node.Parameter)

	if nn.VariableType != nil {
		Print(o, nn.VariableType)
		io.WriteString(o, " ")
	}

	if nn.ByRef {
		io.WriteString(o, "&")
	}

	if nn.Variadic {
		io.WriteString(o, "...")
	}

	Print(o, nn.Variable)

	if nn.DefaultValue != nil {
		io.WriteString(o, " = ")
		Print(o, nn.DefaultValue)
	}
}

func printNodeNullable(o io.Writer, n node.Node) {
	nn := n.(*node.Nullable)

	io.WriteString(o, "?")
	Print(o, nn.Expr)
}

func printNodeArgument(o io.Writer, n node.Node) {
	nn := n.(*node.Argument)

	if nn.IsReference {
		io.WriteString(o, "&")
	}

	if nn.Variadic {
		io.WriteString(o, "...")
	}

	Print(o, nn.Expr)
}

// name

func printNameNamePart(o io.Writer, n node.Node) {
	v := n.(*name.NamePart).Value
	io.WriteString(o, v)
}

func printNameName(o io.Writer, n node.Node) {
	nn := n.(*name.Name)

	for k, part := range nn.Parts {
		if k > 0 {
			io.WriteString(o, "\\")
		}

		Print(o, part)
	}
}

func printNameFullyQualified(o io.Writer, n node.Node) {
	nn := n.(*name.FullyQualified)

	for _, part := range nn.Parts {
		io.WriteString(o, "\\")
		Print(o, part)
	}
}

func printNameRelative(o io.Writer, n node.Node) {
	nn := n.(*name.Relative)

	io.WriteString(o, "namespace")
	for _, part := range nn.Parts {
		io.WriteString(o, "\\")
		Print(o, part)
	}
}

// scalar

func printScalarLNumber(o io.Writer, n node.Node) {
	v := n.(*scalar.Lnumber).Value
	io.WriteString(o, v)
}

func printScalarDNumber(o io.Writer, n node.Node) {
	v := n.(*scalar.Dnumber).Value
	io.WriteString(o, v)
}

func printScalarString(o io.Writer, n node.Node) {
	v := n.(*scalar.String).Value

	io.WriteString(o, "'")
	io.WriteString(o, v)
	io.WriteString(o, "'")
}

func printScalarEncapsedStringPart(o io.Writer, n node.Node) {
	v := n.(*scalar.EncapsedStringPart).Value
	io.WriteString(o, v)
}

func printScalarEncapsed(o io.Writer, n node.Node) {
	io.WriteString(o, "\"")

	for _, nn := range n.(*scalar.Encapsed).Parts {
		Print(o, nn)
	}

	io.WriteString(o, "\"")
}

func printScalarMagicConstant(o io.Writer, n node.Node) {
	v := n.(*scalar.MagicConstant).Value
	io.WriteString(o, v)
}

// Assign

func printAssign(o io.Writer, n node.Node) {
	nn := n.(*assign.Assign)
	Print(o, nn.Variable)
	io.WriteString(o, " = ")
	Print(o, nn.Expression)
}

func printAssignRef(o io.Writer, n node.Node) {
	nn := n.(*assign.AssignRef)
	Print(o, nn.Variable)
	io.WriteString(o, " =& ")
	Print(o, nn.Expression)
}

func printAssignBitwiseAnd(o io.Writer, n node.Node) {
	nn := n.(*assign.BitwiseAnd)
	Print(o, nn.Variable)
	io.WriteString(o, " &= ")
	Print(o, nn.Expression)
}

func printAssignBitwiseOr(o io.Writer, n node.Node) {
	nn := n.(*assign.BitwiseOr)
	Print(o, nn.Variable)
	io.WriteString(o, " |= ")
	Print(o, nn.Expression)
}

func printAssignBitwiseXor(o io.Writer, n node.Node) {
	nn := n.(*assign.BitwiseXor)
	Print(o, nn.Variable)
	io.WriteString(o, " ^= ")
	Print(o, nn.Expression)
}

func printAssignConcat(o io.Writer, n node.Node) {
	nn := n.(*assign.Concat)
	Print(o, nn.Variable)
	io.WriteString(o, " .= ")
	Print(o, nn.Expression)
}

func printAssignDiv(o io.Writer, n node.Node) {
	nn := n.(*assign.Div)
	Print(o, nn.Variable)
	io.WriteString(o, " /= ")
	Print(o, nn.Expression)
}

func printAssignMinus(o io.Writer, n node.Node) {
	nn := n.(*assign.Minus)
	Print(o, nn.Variable)
	io.WriteString(o, " -= ")
	Print(o, nn.Expression)
}

func printAssignMod(o io.Writer, n node.Node) {
	nn := n.(*assign.Mod)
	Print(o, nn.Variable)
	io.WriteString(o, " %= ")
	Print(o, nn.Expression)
}

func printAssignMul(o io.Writer, n node.Node) {
	nn := n.(*assign.Mul)
	Print(o, nn.Variable)
	io.WriteString(o, " *= ")
	Print(o, nn.Expression)
}

func printAssignPlus(o io.Writer, n node.Node) {
	nn := n.(*assign.Plus)
	Print(o, nn.Variable)
	io.WriteString(o, " += ")
	Print(o, nn.Expression)
}

func printAssignPow(o io.Writer, n node.Node) {
	nn := n.(*assign.Pow)
	Print(o, nn.Variable)
	io.WriteString(o, " **= ")
	Print(o, nn.Expression)
}

func printAssignShiftLeft(o io.Writer, n node.Node) {
	nn := n.(*assign.ShiftLeft)
	Print(o, nn.Variable)
	io.WriteString(o, " <<= ")
	Print(o, nn.Expression)
}

func printAssignShiftRight(o io.Writer, n node.Node) {
	nn := n.(*assign.ShiftRight)
	Print(o, nn.Variable)
	io.WriteString(o, " >>= ")
	Print(o, nn.Expression)
}

// binary

func printBinaryBitwiseAnd(o io.Writer, n node.Node) {
	nn := n.(*binary.BitwiseAnd)

	Print(o, nn.Left)
	io.WriteString(o, " & ")
	Print(o, nn.Right)
}

func printBinaryBitwiseOr(o io.Writer, n node.Node) {
	nn := n.(*binary.BitwiseOr)

	Print(o, nn.Left)
	io.WriteString(o, " | ")
	Print(o, nn.Right)
}

func printBinaryBitwiseXor(o io.Writer, n node.Node) {
	nn := n.(*binary.BitwiseXor)

	Print(o, nn.Left)
	io.WriteString(o, " ^ ")
	Print(o, nn.Right)
}

func printBinaryBooleanAnd(o io.Writer, n node.Node) {
	nn := n.(*binary.BooleanAnd)

	Print(o, nn.Left)
	io.WriteString(o, " && ")
	Print(o, nn.Right)
}

func printBinaryBooleanOr(o io.Writer, n node.Node) {
	nn := n.(*binary.BooleanOr)

	Print(o, nn.Left)
	io.WriteString(o, " || ")
	Print(o, nn.Right)
}

func printBinaryCoalesce(o io.Writer, n node.Node) {
	nn := n.(*binary.Coalesce)

	Print(o, nn.Left)
	io.WriteString(o, " ?? ")
	Print(o, nn.Right)
}

func printBinaryConcat(o io.Writer, n node.Node) {
	nn := n.(*binary.Concat)

	Print(o, nn.Left)
	io.WriteString(o, " . ")
	Print(o, nn.Right)
}

func printBinaryDiv(o io.Writer, n node.Node) {
	nn := n.(*binary.Div)

	Print(o, nn.Left)
	io.WriteString(o, " / ")
	Print(o, nn.Right)
}

func printBinaryEqual(o io.Writer, n node.Node) {
	nn := n.(*binary.Equal)

	Print(o, nn.Left)
	io.WriteString(o, " == ")
	Print(o, nn.Right)
}

func printBinaryGreaterOrEqual(o io.Writer, n node.Node) {
	nn := n.(*binary.GreaterOrEqual)

	Print(o, nn.Left)
	io.WriteString(o, " >= ")
	Print(o, nn.Right)
}

func printBinaryGreater(o io.Writer, n node.Node) {
	nn := n.(*binary.Greater)

	Print(o, nn.Left)
	io.WriteString(o, " > ")
	Print(o, nn.Right)
}

func printBinaryIdentical(o io.Writer, n node.Node) {
	nn := n.(*binary.Identical)

	Print(o, nn.Left)
	io.WriteString(o, " === ")
	Print(o, nn.Right)
}

func printBinaryLogicalAnd(o io.Writer, n node.Node) {
	nn := n.(*binary.LogicalAnd)

	Print(o, nn.Left)
	io.WriteString(o, " and ")
	Print(o, nn.Right)
}

func printBinaryLogicalOr(o io.Writer, n node.Node) {
	nn := n.(*binary.LogicalOr)

	Print(o, nn.Left)
	io.WriteString(o, " or ")
	Print(o, nn.Right)
}

func printBinaryLogicalXor(o io.Writer, n node.Node) {
	nn := n.(*binary.LogicalXor)

	Print(o, nn.Left)
	io.WriteString(o, " xor ")
	Print(o, nn.Right)
}

func printBinaryMinus(o io.Writer, n node.Node) {
	nn := n.(*binary.Minus)

	Print(o, nn.Left)
	io.WriteString(o, " - ")
	Print(o, nn.Right)
}

func printBinaryMod(o io.Writer, n node.Node) {
	nn := n.(*binary.Mod)

	Print(o, nn.Left)
	io.WriteString(o, " % ")
	Print(o, nn.Right)
}

func printBinaryMul(o io.Writer, n node.Node) {
	nn := n.(*binary.Mul)

	Print(o, nn.Left)
	io.WriteString(o, " * ")
	Print(o, nn.Right)
}

func printBinaryNotEqual(o io.Writer, n node.Node) {
	nn := n.(*binary.NotEqual)

	Print(o, nn.Left)
	io.WriteString(o, " != ")
	Print(o, nn.Right)
}

func printBinaryNotIdentical(o io.Writer, n node.Node) {
	nn := n.(*binary.NotIdentical)

	Print(o, nn.Left)
	io.WriteString(o, " !== ")
	Print(o, nn.Right)
}

func printBinaryPlus(o io.Writer, n node.Node) {
	nn := n.(*binary.Plus)

	Print(o, nn.Left)
	io.WriteString(o, " + ")
	Print(o, nn.Right)
}

func printBinaryPow(o io.Writer, n node.Node) {
	nn := n.(*binary.Pow)

	Print(o, nn.Left)
	io.WriteString(o, " ** ")
	Print(o, nn.Right)
}

func printBinaryShiftLeft(o io.Writer, n node.Node) {
	nn := n.(*binary.ShiftLeft)

	Print(o, nn.Left)
	io.WriteString(o, " << ")
	Print(o, nn.Right)
}

func printBinaryShiftRight(o io.Writer, n node.Node) {
	nn := n.(*binary.ShiftRight)

	Print(o, nn.Left)
	io.WriteString(o, " >> ")
	Print(o, nn.Right)
}

func printBinarySmallerOrEqual(o io.Writer, n node.Node) {
	nn := n.(*binary.SmallerOrEqual)

	Print(o, nn.Left)
	io.WriteString(o, " <= ")
	Print(o, nn.Right)
}

func printBinarySmaller(o io.Writer, n node.Node) {
	nn := n.(*binary.Smaller)

	Print(o, nn.Left)
	io.WriteString(o, " < ")
	Print(o, nn.Right)
}

func printBinarySpaceship(o io.Writer, n node.Node) {
	nn := n.(*binary.Spaceship)

	Print(o, nn.Left)
	io.WriteString(o, " <=> ")
	Print(o, nn.Right)
}

// cast

func printCastArray(o io.Writer, n node.Node) {
	nn := n.(*cast.CastArray)

	io.WriteString(o, "(array)")
	Print(o, nn.Expr)
}

func printCastBool(o io.Writer, n node.Node) {
	nn := n.(*cast.CastBool)

	io.WriteString(o, "(bool)")
	Print(o, nn.Expr)
}

func printCastDouble(o io.Writer, n node.Node) {
	nn := n.(*cast.CastDouble)

	io.WriteString(o, "(float)")
	Print(o, nn.Expr)
}

func printCastInt(o io.Writer, n node.Node) {
	nn := n.(*cast.CastInt)

	io.WriteString(o, "(int)")
	Print(o, nn.Expr)
}

func printCastObject(o io.Writer, n node.Node) {
	nn := n.(*cast.CastObject)

	io.WriteString(o, "(object)")
	Print(o, nn.Expr)
}

func printCastString(o io.Writer, n node.Node) {
	nn := n.(*cast.CastString)

	io.WriteString(o, "(string)")
	Print(o, nn.Expr)
}

func printCastUnset(o io.Writer, n node.Node) {
	nn := n.(*cast.CastUnset)

	io.WriteString(o, "(unset)")
	Print(o, nn.Expr)
}

// expr

func printExprArrayDimFetch(o io.Writer, n node.Node) {
	nn := n.(*expr.ArrayDimFetch)
	Print(o, nn.Variable)
	io.WriteString(o, "[")
	Print(o, nn.Dim)
	io.WriteString(o, "]")
}

func printExprArrayItem(o io.Writer, n node.Node) {
	nn := n.(*expr.ArrayItem)

	if nn.Key != nil {
		Print(o, nn.Key)
		io.WriteString(o, " => ")
	}

	if nn.ByRef {
		io.WriteString(o, "&")
	}

	Print(o, nn.Val)
}

func printExprArray(o io.Writer, n node.Node) {
	nn := n.(*expr.Array)

	io.WriteString(o, "array(")
	joinPrint(", ", o, nn.Items)
	io.WriteString(o, ")")
}

func printExprBitwiseNot(o io.Writer, n node.Node) {
	nn := n.(*expr.BitwiseNot)
	io.WriteString(o, "~")
	Print(o, nn.Expr)
}

func printExprBooleanNot(o io.Writer, n node.Node) {
	nn := n.(*expr.BooleanNot)
	io.WriteString(o, "!")
	Print(o, nn.Expr)
}

func printExprClassConstFetch(o io.Writer, n node.Node) {
	nn := n.(*expr.ClassConstFetch)

	Print(o, nn.Class)
	io.WriteString(o, "::")
	io.WriteString(o, nn.ConstantName.(*node.Identifier).Value)
}

func printExprClone(o io.Writer, n node.Node) {
	nn := n.(*expr.Clone)

	io.WriteString(o, "clone ")
	Print(o, nn.Expr)
}

func printExprClosureUse(o io.Writer, n node.Node) {
	nn := n.(*expr.ClosureUse)

	if nn.ByRef {
		io.WriteString(o, "&")
	}

	Print(o, nn.Variable)
}

func printExprClosure(o io.Writer, n node.Node) {
	nn := n.(*expr.Closure)

	if nn.Static {
		io.WriteString(o, "static ")
	}

	io.WriteString(o, "function ")

	if nn.ReturnsRef {
		io.WriteString(o, "&")
	}

	io.WriteString(o, "(")
	joinPrint(", ", o, nn.Params)
	io.WriteString(o, ")")

	if nn.Uses != nil {
		io.WriteString(o, " use (")
		joinPrint(", ", o, nn.Uses)
		io.WriteString(o, ")")
	}

	if nn.ReturnType != nil {
		io.WriteString(o, ": ")
		Print(o, nn.ReturnType)
	}

	io.WriteString(o, " {")
	for _, stmt := range nn.Stmts {
		Print(o, stmt)
	}
	io.WriteString(o, "}")
}

func printExprConstFetch(o io.Writer, n node.Node) {
	nn := n.(*expr.ConstFetch)

	Print(o, nn.Constant)
}

func printExprDie(o io.Writer, n node.Node) {
	nn := n.(*expr.Die)

	io.WriteString(o, "die(")
	Print(o, nn.Expr)
	io.WriteString(o, ")")
}

func printExprEmpty(o io.Writer, n node.Node) {
	nn := n.(*expr.Empty)

	io.WriteString(o, "empty(")
	Print(o, nn.Expr)
	io.WriteString(o, ")")
}

func printExprErrorSuppress(o io.Writer, n node.Node) {
	nn := n.(*expr.ErrorSuppress)

	io.WriteString(o, "@")
	Print(o, nn.Expr)
}

func printExprEval(o io.Writer, n node.Node) {
	nn := n.(*expr.Eval)

	io.WriteString(o, "eval(")
	Print(o, nn.Expr)
	io.WriteString(o, ")")
}

func printExprExit(o io.Writer, n node.Node) {
	nn := n.(*expr.Exit)

	io.WriteString(o, "exit(")
	Print(o, nn.Expr)
	io.WriteString(o, ")")
}

func printExprFunctionCall(o io.Writer, n node.Node) {
	nn := n.(*expr.FunctionCall)

	Print(o, nn.Function)
	io.WriteString(o, "(")
	joinPrint(", ", o, nn.Arguments)
	io.WriteString(o, ")")
}

func printExprInclude(o io.Writer, n node.Node) {
	nn := n.(*expr.Include)

	io.WriteString(o, "include ")
	Print(o, nn.Expr)
}

func printExprIncludeOnce(o io.Writer, n node.Node) {
	nn := n.(*expr.IncludeOnce)

	io.WriteString(o, "include_once ")
	Print(o, nn.Expr)
}

func printExprInstanceOf(o io.Writer, n node.Node) {
	nn := n.(*expr.InstanceOf)

	Print(o, nn.Expr)
	io.WriteString(o, " instanceof ")
	Print(o, nn.Class)
}

func printExprIsset(o io.Writer, n node.Node) {
	nn := n.(*expr.Isset)

	io.WriteString(o, "isset(")
	joinPrint(", ", o, nn.Variables)
	io.WriteString(o, ")")
}

func printExprList(o io.Writer, n node.Node) {
	nn := n.(*expr.List)

	io.WriteString(o, "list(")
	joinPrint(", ", o, nn.Items)
	io.WriteString(o, ")")
}

func printExprMethodCall(o io.Writer, n node.Node) {
	nn := n.(*expr.MethodCall)

	Print(o, nn.Variable)
	io.WriteString(o, "->")
	Print(o, nn.Method)
	io.WriteString(o, "(")
	joinPrint(", ", o, nn.Arguments)
	io.WriteString(o, ")")
}

func printExprNew(o io.Writer, n node.Node) {
	nn := n.(*expr.New)

	io.WriteString(o, "new ")
	Print(o, nn.Class)

	if nn.Arguments != nil {
		io.WriteString(o, "(")
		joinPrint(", ", o, nn.Arguments)
		io.WriteString(o, ")")
	}
}

func printExprPostDec(o io.Writer, n node.Node) {
	nn := n.(*expr.PostDec)

	Print(o, nn.Variable)
	io.WriteString(o, "--")
}

func printExprPostInc(o io.Writer, n node.Node) {
	nn := n.(*expr.PostInc)

	Print(o, nn.Variable)
	io.WriteString(o, "++")
}

func printExprPreDec(o io.Writer, n node.Node) {
	nn := n.(*expr.PreDec)

	io.WriteString(o, "--")
	Print(o, nn.Variable)
}

func printExprPreInc(o io.Writer, n node.Node) {
	nn := n.(*expr.PreInc)

	io.WriteString(o, "++")
	Print(o, nn.Variable)
}

func printExprPrint(o io.Writer, n node.Node) {
	nn := n.(*expr.Print)

	io.WriteString(o, "print(")
	Print(o, nn.Expr)
	io.WriteString(o, ")")
}

func printExprPropertyFetch(o io.Writer, n node.Node) {
	nn := n.(*expr.PropertyFetch)

	Print(o, nn.Variable)
	io.WriteString(o, "->")
	Print(o, nn.Property)
}

func printExprRequire(o io.Writer, n node.Node) {
	nn := n.(*expr.Require)

	io.WriteString(o, "require ")
	Print(o, nn.Expr)
}

func printExprRequireOnce(o io.Writer, n node.Node) {
	nn := n.(*expr.RequireOnce)

	io.WriteString(o, "require_once ")
	Print(o, nn.Expr)
}

func printExprShellExec(o io.Writer, n node.Node) {
	nn := n.(*expr.ShellExec)

	io.WriteString(o, "`")
	for _, part := range nn.Parts {
		Print(o, part)
	}
	io.WriteString(o, "`")
}

func printExprShortArray(o io.Writer, n node.Node) {
	nn := n.(*expr.ShortArray)

	io.WriteString(o, "[")
	joinPrint(", ", o, nn.Items)
	io.WriteString(o, "]")
}

func printExprShortList(o io.Writer, n node.Node) {
	nn := n.(*expr.ShortList)

	io.WriteString(o, "[")
	joinPrint(", ", o, nn.Items)
	io.WriteString(o, "]")
}

func printExprStaticCall(o io.Writer, n node.Node) {
	nn := n.(*expr.StaticCall)

	Print(o, nn.Class)
	io.WriteString(o, "::")
	Print(o, nn.Call)
	io.WriteString(o, "(")
	joinPrint(", ", o, nn.Arguments)
	io.WriteString(o, ")")
}

func printExprStaticPropertyFetch(o io.Writer, n node.Node) {
	nn := n.(*expr.StaticPropertyFetch)

	Print(o, nn.Class)
	io.WriteString(o, "::")
	Print(o, nn.Property)
}

func printExprTernary(o io.Writer, n node.Node) {
	nn := n.(*expr.Ternary)

	Print(o, nn.Condition)
	io.WriteString(o, " ?")

	if nn.IfFalse != nil {
		io.WriteString(o, " ")
		Print(o, nn.IfFalse)
		io.WriteString(o, " ")
	}

	io.WriteString(o, ": ")
	Print(o, nn.IfTrue)
}

func printExprUnaryMinus(o io.Writer, n node.Node) {
	nn := n.(*expr.UnaryMinus)

	io.WriteString(o, "-")
	Print(o, nn.Expr)
}

func printExprUnaryPlus(o io.Writer, n node.Node) {
	nn := n.(*expr.UnaryPlus)

	io.WriteString(o, "+")
	Print(o, nn.Expr)
}

func printExprVariable(o io.Writer, n node.Node) {
	io.WriteString(o, "$")
	Print(o, n.(*expr.Variable).VarName)
}

func printExprYieldFrom(o io.Writer, n node.Node) {
	nn := n.(*expr.YieldFrom)

	io.WriteString(o, "yield from ")
	Print(o, nn.Expr)
}

func printExprYield(o io.Writer, n node.Node) {
	nn := n.(*expr.Yield)

	io.WriteString(o, "yield ")

	if nn.Key != nil {
		Print(o, nn.Key)
		io.WriteString(o, " => ")
	}

	Print(o, nn.Value)
}

// smtm

func printStmtAltElseIf(o io.Writer, n node.Node) {
	nn := n.(*stmt.AltElseIf)

	io.WriteString(o, "elseif (")
	Print(o, nn.Cond)
	io.WriteString(o, ") :\n")

	s := nn.Stmt.(*stmt.StmtList)
	printNodes(o, s.Stmts)
}

func printStmtAltElse(o io.Writer, n node.Node) {
	nn := n.(*stmt.AltElse)

	io.WriteString(o, "else :\n")

	s := nn.Stmt.(*stmt.StmtList)
	printNodes(o, s.Stmts)
}

func printStmtAltFor(o io.Writer, n node.Node) {
	nn := n.(*stmt.AltFor)

	io.WriteString(o, "for (")
	joinPrint(", ", o, nn.Init)
	io.WriteString(o, "; ")
	joinPrint(", ", o, nn.Cond)
	io.WriteString(o, "; ")
	joinPrint(", ", o, nn.Loop)
	io.WriteString(o, ") :\n")

	s := nn.Stmt.(*stmt.StmtList)
	printNodes(o, s.Stmts)

	io.WriteString(o, "endfor;\n")
}

func printStmtAltForeach(o io.Writer, n node.Node) {
	nn := n.(*stmt.AltForeach)

	io.WriteString(o, "foreach (")
	Print(o, nn.Expr)
	io.WriteString(o, " as ")

	if nn.Key != nil {
		Print(o, nn.Key)
		io.WriteString(o, " => ")
	}

	if nn.ByRef {
		io.WriteString(o, "&")
	}

	Print(o, nn.Variable)

	io.WriteString(o, ") :\n")

	s := nn.Stmt.(*stmt.StmtList)
	printNodes(o, s.Stmts)

	io.WriteString(o, "endforeach;\n")
}

func printStmtAltIf(o io.Writer, n node.Node) {
	nn := n.(*stmt.AltIf)

	io.WriteString(o, "if (")
	Print(o, nn.Cond)
	io.WriteString(o, ") :\n")

	s := nn.Stmt.(*stmt.StmtList)
	printNodes(o, s.Stmts)

	for _, elseif := range nn.ElseIf {
		Print(o, elseif)
	}

	if nn.Else != nil {
		Print(o, nn.Else)
	}

	io.WriteString(o, "endif;\n")
}

func printStmtAltSwitch(o io.Writer, n node.Node) {
	nn := n.(*stmt.AltSwitch)

	io.WriteString(o, "switch (")
	Print(o, nn.Cond)
	io.WriteString(o, ") :\n")

	s := nn.Cases
	printNodes(o, s)

	io.WriteString(o, "endswitch;\n")
}

func printStmtAltWhile(o io.Writer, n node.Node) {
	nn := n.(*stmt.AltWhile)

	io.WriteString(o, "while (")
	Print(o, nn.Cond)
	io.WriteString(o, ") :\n")

	s := nn.Stmt.(*stmt.StmtList)
	printNodes(o, s.Stmts)

	io.WriteString(o, "endwhile;\n")
}

func printStmtBreak(o io.Writer, n node.Node) {
	nn := n.(*stmt.Break)

	io.WriteString(o, "break")
	if nn.Expr != nil {
		io.WriteString(o, " ")
		Print(o, nn.Expr)
	}

	io.WriteString(o, ";\n")
}

func printStmtCase(o io.Writer, n node.Node) {
	nn := n.(*stmt.Case)

	io.WriteString(o, "case ")
	Print(o, nn.Cond)
	io.WriteString(o, ":\n")
	printNodes(o, nn.Stmts)
}

func printStmtCatch(o io.Writer, n node.Node) {
	nn := n.(*stmt.Catch)

	io.WriteString(o, "catch (")
	joinPrint(" | ", o, nn.Types)
	io.WriteString(o, " ")
	Print(o, nn.Variable)
	io.WriteString(o, ") {\n")
	printNodes(o, nn.Stmts)
	io.WriteString(o, "}\n")
}

func printStmtClassMethod(o io.Writer, n node.Node) {
	nn := n.(*stmt.ClassMethod)

	if nn.Modifiers != nil {
		joinPrint(" ", o, nn.Modifiers)
		io.WriteString(o, " ")
	}
	io.WriteString(o, "function ")

	if nn.ReturnsRef {
		io.WriteString(o, "&")
	}

	Print(o, nn.MethodName)
	io.WriteString(o, "(")
	joinPrint(", ", o, nn.Params)
	io.WriteString(o, ")")

	if nn.ReturnType != nil {
		io.WriteString(o, ": ")
		Print(o, nn.ReturnType)
	}

	io.WriteString(o, "\n{\n") // TODO: handle indentation
	printNodes(o, nn.Stmts)
	io.WriteString(o, "}\n") // TODO: handle indentation
}

func printStmtClass(o io.Writer, n node.Node) {
	nn := n.(*stmt.Class)

	if nn.Modifiers != nil {
		joinPrint(" ", o, nn.Modifiers)
		io.WriteString(o, " ")
	}
	io.WriteString(o, "class")

	if nn.ClassName != nil {
		io.WriteString(o, " ")
		Print(o, nn.ClassName)
	}

	if nn.Args != nil {
		io.WriteString(o, "(")
		joinPrint(", ", o, nn.Args)
		io.WriteString(o, ")")
	}

	if nn.Extends != nil {
		io.WriteString(o, " extends ")
		Print(o, nn.Extends)
	}

	if nn.Implements != nil {
		io.WriteString(o, " implements ")
		joinPrint(", ", o, nn.Implements)
	}

	io.WriteString(o, "\n{\n") // TODO: handle indentation
	printNodes(o, nn.Stmts)
	io.WriteString(o, "}\n") // TODO: handle indentation
}

func printStmtClassConstList(o io.Writer, n node.Node) {
	nn := n.(*stmt.ClassConstList)

	if nn.Modifiers != nil {
		joinPrint(" ", o, nn.Modifiers)
		io.WriteString(o, " ")
	}
	io.WriteString(o, "const ")

	joinPrint(", ", o, nn.Consts)

	io.WriteString(o, ";\n")
}

func printStmtConstant(o io.Writer, n node.Node) {
	nn := n.(*stmt.Constant)

	Print(o, nn.ConstantName)
	io.WriteString(o, " = ")
	Print(o, nn.Expr)
}

func printStmtContinue(o io.Writer, n node.Node) {
	nn := n.(*stmt.Continue)

	io.WriteString(o, "continue")
	if nn.Expr != nil {
		io.WriteString(o, " ")
		Print(o, nn.Expr)
	}

	io.WriteString(o, ";\n")
}

func printStmtDeclare(o io.Writer, n node.Node) {
	nn := n.(*stmt.Declare)

	io.WriteString(o, "declare(")
	joinPrint(", ", o, nn.Consts)
	io.WriteString(o, ")")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		Print(o, s)
		break
	case *stmt.StmtList:
		io.WriteString(o, " {\n")
		printNodes(o, s.Stmts)
		io.WriteString(o, "}\n")
	default:
		io.WriteString(o, "\n")
		Print(o, s)
	}
}

func printStmtDefault(o io.Writer, n node.Node) {
	nn := n.(*stmt.Default)
	io.WriteString(o, "default:\n")
	printNodes(o, nn.Stmts)
}

func printStmtDo(o io.Writer, n node.Node) {
	nn := n.(*stmt.Do)
	io.WriteString(o, "do")

	switch s := nn.Stmt.(type) {
	case *stmt.StmtList:
		io.WriteString(o, " {\n")
		printNodes(o, s.Stmts)
		io.WriteString(o, "} ")
	default:
		io.WriteString(o, "\n")
		Print(o, s)
	}

	io.WriteString(o, "while (")
	Print(o, nn.Cond)
	io.WriteString(o, ");\n")
}

func printStmtEcho(o io.Writer, n node.Node) {
	nn := n.(*stmt.Echo)
	io.WriteString(o, "echo ")
	joinPrint(", ", o, nn.Exprs)
	io.WriteString(o, ";\n")
}

func printStmtElseif(o io.Writer, n node.Node) {
	nn := n.(*stmt.ElseIf)

	io.WriteString(o, "elseif (")
	Print(o, nn.Cond)
	io.WriteString(o, ")")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		Print(o, s)
		break
	case *stmt.StmtList:
		io.WriteString(o, " {\n")
		printNodes(o, s.Stmts)
		io.WriteString(o, "}\n")
	default:
		io.WriteString(o, "\n")
		Print(o, s)
	}
}

func printStmtElse(o io.Writer, n node.Node) {
	nn := n.(*stmt.Else)

	io.WriteString(o, "else")

	switch s := nn.Stmt.(type) {
	case *stmt.Nop:
		Print(o, s)
		break
	case *stmt.StmtList:
		io.WriteString(o, " {\n")
		printNodes(o, s.Stmts)
		io.WriteString(o, "}\n")
	default:
		io.WriteString(o, "\n")
		Print(o, s)
	}
}

func printStmtExpression(o io.Writer, n node.Node) {
	nn := n.(*stmt.Expression)

	Print(o, nn.Expr)

	io.WriteString(o, ";\n")
}

func printStmtFinally(o io.Writer, n node.Node) {
	nn := n.(*stmt.Finally)

	io.WriteString(o, "finally {\n")
	printNodes(o, nn.Stmts)
	io.WriteString(o, "}\n")
}

func printStmtStmtList(o io.Writer, n node.Node) {
	nn := n.(*stmt.StmtList)

	printNodes(o, nn.Stmts)
}

func printStmtNop(o io.Writer, n node.Node) {
	io.WriteString(o, ";\n")
}
