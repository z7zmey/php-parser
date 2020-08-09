package visitor

import (
	"github.com/z7zmey/php-parser/pkg/ast"
)

type FilterParserNodes struct {
	Null
}

func (v *FilterParserNodes) EnterNode(n ast.Vertex) bool {
	n.Accept(v)
	return true
}

func (v *FilterParserNodes) StmtGroupUseList(n *ast.StmtGroupUseList) {
	if nn, ok := n.Prefix.(*ast.ParserNsSeparator); ok {
		n.Prefix = nn.Child
	}

	if nn, ok := n.UseList.(*ast.ParserNsSeparator); ok {
		n.UseList = nn.Child
	}

	if nn, ok := n.UseList.(*ast.ParserBrackets); ok {
		n.UseList = nn.Child
	}
}

func (v *FilterParserNodes) StmtUseList(n *ast.StmtUseList) {
	for k, v := range n.UseDeclarations {
		if nn, ok := v.(*ast.ParserNsSeparator); ok {
			n.UseDeclarations[k] = nn.Child
		}
	}
}

func (v *FilterParserNodes) StmtUseDeclaration(n *ast.StmtUseDeclaration) {
	if nn, ok := n.Alias.(*ast.ParserAs); ok {
		n.Alias = nn.Child
	}
}

func (v *FilterParserNodes) StmtAltIf(n *ast.StmtAltIf) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}

	if nn, ok := n.Stmt.(*ast.ParserBrackets); ok {
		n.Stmt = nn.Child
	}
}

func (v *FilterParserNodes) StmtAltElseIf(n *ast.StmtAltElseIf) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}

	if nn, ok := n.Stmt.(*ast.ParserBrackets); ok {
		n.Stmt = nn.Child
	}
}

func (v *FilterParserNodes) StmtAltElse(n *ast.StmtAltElse) {
	if nn, ok := n.Stmt.(*ast.ParserBrackets); ok {
		n.Stmt = nn.Child
	}
}

func (v *FilterParserNodes) StmtIf(n *ast.StmtIf) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtElseIf(n *ast.StmtElseIf) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtWhile(n *ast.StmtWhile) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtAltWhile(n *ast.StmtAltWhile) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtDo(n *ast.StmtDo) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtSwitch(n *ast.StmtSwitch) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtAltSwitch(n *ast.StmtAltSwitch) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprExit(n *ast.ExprExit) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtContinue(n *ast.StmtContinue) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtBreak(n *ast.StmtBreak) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprClone(n *ast.ExprClone) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprPrint(n *ast.ExprPrint) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtExpression(n *ast.StmtExpression) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtEcho(n *ast.StmtEcho) {
	for k, v := range n.Exprs {
		for {
			if nn, ok := v.(*ast.ParserBrackets); ok {
				v = nn.Child
			} else {
				break
			}
		}

		n.Exprs[k] = v
	}
}

func (v *FilterParserNodes) ExprIsset(n *ast.ExprIsset) {
	for k, v := range n.Vars {
		for {
			if nn, ok := v.(*ast.ParserBrackets); ok {
				v = nn.Child
			} else {
				break
			}
		}

		n.Vars[k] = v
	}
}

func (v *FilterParserNodes) StmtReturn(n *ast.StmtReturn) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtForeach(n *ast.StmtForeach) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtAltForeach(n *ast.StmtAltForeach) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprYield(n *ast.ExprYield) {
	for {
		if nn, ok := n.Key.(*ast.ParserBrackets); ok {
			n.Key = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Value.(*ast.ParserBrackets); ok {
			n.Value = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) Argument(n *ast.Argument) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtThrow(n *ast.StmtThrow) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtCase(n *ast.StmtCase) {
	for {
		if nn, ok := n.Cond.(*ast.ParserBrackets); ok {
			n.Cond = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprVariable(n *ast.ExprVariable) {
	for {
		if nn, ok := n.VarName.(*ast.ParserBrackets); ok {
			n.VarName = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) StmtFor(n *ast.StmtFor) {
	for k, v := range n.Init {
		for {
			if nn, ok := v.(*ast.ParserBrackets); ok {
				v = nn.Child
			} else {
				break
			}
		}

		n.Init[k] = v
	}

	for k, v := range n.Cond {
		for {
			if nn, ok := v.(*ast.ParserBrackets); ok {
				v = nn.Child
			} else {
				break
			}
		}

		n.Cond[k] = v
	}

	for k, v := range n.Loop {
		for {
			if nn, ok := v.(*ast.ParserBrackets); ok {
				v = nn.Child
			} else {
				break
			}
		}

		n.Loop[k] = v
	}
}

func (v *FilterParserNodes) ExprAssign(n *ast.ExprAssign) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignBitwiseAnd(n *ast.ExprAssignBitwiseAnd) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignBitwiseOr(n *ast.ExprAssignBitwiseOr) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignBitwiseXor(n *ast.ExprAssignBitwiseXor) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignCoalesce(n *ast.ExprAssignCoalesce) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignConcat(n *ast.ExprAssignConcat) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignDiv(n *ast.ExprAssignDiv) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignMinus(n *ast.ExprAssignMinus) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignMod(n *ast.ExprAssignMod) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignMul(n *ast.ExprAssignMul) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignPlus(n *ast.ExprAssignPlus) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignPow(n *ast.ExprAssignPow) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignShiftLeft(n *ast.ExprAssignShiftLeft) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprAssignShiftRight(n *ast.ExprAssignShiftRight) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}
func (v *FilterParserNodes) ExprBinaryBitwiseAnd(n *ast.ExprBinaryBitwiseAnd) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryBitwiseOr(n *ast.ExprBinaryBitwiseOr) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryBitwiseXor(n *ast.ExprBinaryBitwiseXor) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryBooleanAnd(n *ast.ExprBinaryBooleanAnd) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryBooleanOr(n *ast.ExprBinaryBooleanOr) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryCoalesce(n *ast.ExprBinaryCoalesce) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryConcat(n *ast.ExprBinaryConcat) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryDiv(n *ast.ExprBinaryDiv) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryEqual(n *ast.ExprBinaryEqual) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryGreater(n *ast.ExprBinaryGreater) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryGreaterOrEqual(n *ast.ExprBinaryGreaterOrEqual) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryIdentical(n *ast.ExprBinaryIdentical) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryLogicalAnd(n *ast.ExprBinaryLogicalAnd) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryLogicalOr(n *ast.ExprBinaryLogicalOr) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryLogicalXor(n *ast.ExprBinaryLogicalXor) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryMinus(n *ast.ExprBinaryMinus) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryMod(n *ast.ExprBinaryMod) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryMul(n *ast.ExprBinaryMul) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryNotEqual(n *ast.ExprBinaryNotEqual) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryNotIdentical(n *ast.ExprBinaryNotIdentical) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryPlus(n *ast.ExprBinaryPlus) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryPow(n *ast.ExprBinaryPow) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryShiftLeft(n *ast.ExprBinaryShiftLeft) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinaryShiftRight(n *ast.ExprBinaryShiftRight) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinarySmaller(n *ast.ExprBinarySmaller) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinarySmallerOrEqual(n *ast.ExprBinarySmallerOrEqual) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBinarySpaceship(n *ast.ExprBinarySpaceship) {
	for {
		if nn, ok := n.Left.(*ast.ParserBrackets); ok {
			n.Left = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Right.(*ast.ParserBrackets); ok {
			n.Right = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprUnaryMinus(n *ast.ExprUnaryMinus) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprUnaryPlus(n *ast.ExprUnaryPlus) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBooleanNot(n *ast.ExprBooleanNot) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprBitwiseNot(n *ast.ExprBitwiseNot) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprInstanceOf(n *ast.ExprInstanceOf) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprTernary(n *ast.ExprTernary) {
	for {
		if nn, ok := n.Condition.(*ast.ParserBrackets); ok {
			n.Condition = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.IfTrue.(*ast.ParserBrackets); ok {
			n.IfTrue = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.IfFalse.(*ast.ParserBrackets); ok {
			n.IfFalse = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprCastArray(n *ast.ExprCastArray) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprCastBool(n *ast.ExprCastBool) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprCastDouble(n *ast.ExprCastDouble) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprCastInt(n *ast.ExprCastInt) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprCastObject(n *ast.ExprCastObject) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprCastString(n *ast.ExprCastString) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprCastUnset(n *ast.ExprCastUnset) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprErrorSuppress(n *ast.ExprErrorSuppress) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprArrayDimFetch(n *ast.ExprArrayDimFetch) {
	for {
		if nn, ok := n.Dim.(*ast.ParserBrackets); ok {
			n.Dim = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprArrayItem(n *ast.ExprArrayItem) {
	for {
		if nn, ok := n.Key.(*ast.ParserBrackets); ok {
			n.Key = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Val.(*ast.ParserBrackets); ok {
			n.Val = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprEmpty(n *ast.ExprEmpty) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprInclude(n *ast.ExprInclude) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprIncludeOnce(n *ast.ExprIncludeOnce) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprEval(n *ast.ExprEval) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprRequire(n *ast.ExprRequire) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprRequireOnce(n *ast.ExprRequireOnce) {
	for {
		if nn, ok := n.Expr.(*ast.ParserBrackets); ok {
			n.Expr = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprPropertyFetch(n *ast.ExprPropertyFetch) {
	for {
		if nn, ok := n.Var.(*ast.ParserBrackets); ok {
			n.Var = nn.Child
		} else {
			break
		}
	}

	for {
		if nn, ok := n.Property.(*ast.ParserBrackets); ok {
			n.Property = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprFunctionCall(n *ast.ExprFunctionCall) {
	for {
		if nn, ok := n.Function.(*ast.ParserBrackets); ok {
			n.Function = nn.Child
		} else {
			break
		}
	}
}

func (v *FilterParserNodes) ExprStaticCall(n *ast.ExprStaticCall) {
	for {
		if nn, ok := n.Call.(*ast.ParserBrackets); ok {
			n.Call = nn.Child
		} else {
			break
		}
	}
}
