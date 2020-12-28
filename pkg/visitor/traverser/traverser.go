package traverser

import (
	"github.com/z7zmey/php-parser/pkg/ast"
)

type Traverser struct {
	v ast.Visitor
}

func NewTraverser(v ast.Visitor) *Traverser {
	return &Traverser{
		v: v,
	}
}

func (t *Traverser) Traverse(n ast.Vertex) {
	if n != nil {
		n.Accept(t)
	}
}

func (t *Traverser) Root(n *ast.Root) {
	n.Accept(t.v)

	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) Nullable(n *ast.Nullable) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) Parameter(n *ast.Parameter) {
	n.Accept(t.v)

	t.Traverse(n.Type)
	t.Traverse(n.Var)
	t.Traverse(n.DefaultValue)
}

func (t *Traverser) Identifier(n *ast.Identifier) {
	n.Accept(t.v)
}

func (t *Traverser) Argument(n *ast.Argument) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) StmtBreak(n *ast.StmtBreak) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) StmtCase(n *ast.StmtCase) {
	n.Accept(t.v)

	t.Traverse(n.Cond)
	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtCatch(n *ast.StmtCatch) {
	n.Accept(t.v)

	for _, nn := range n.Types {
		nn.Accept(t)
	}
	t.Traverse(n.Var)
	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtClass(n *ast.StmtClass) {
	n.Accept(t.v)

	for _, nn := range n.Modifiers {
		nn.Accept(t)
	}
	t.Traverse(n.Name)
	for _, nn := range n.Args {
		nn.Accept(t)
	}
	t.Traverse(n.Extends)
	for _, nn := range n.Implements {
		nn.Accept(t)
	}
	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtClassConstList(n *ast.StmtClassConstList) {
	n.Accept(t.v)

	for _, nn := range n.Modifiers {
		nn.Accept(t)
	}
	for _, nn := range n.Consts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtClassMethod(n *ast.StmtClassMethod) {
	n.Accept(t.v)

	for _, nn := range n.Modifiers {
		nn.Accept(t)
	}
	t.Traverse(n.Name)
	for _, nn := range n.Params {
		nn.Accept(t)
	}
	t.Traverse(n.ReturnType)
	t.Traverse(n.Stmt)
}

func (t *Traverser) StmtConstList(n *ast.StmtConstList) {
	n.Accept(t.v)

	for _, nn := range n.Consts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtConstant(n *ast.StmtConstant) {
	n.Accept(t.v)

	t.Traverse(n.Name)
	t.Traverse(n.Expr)
}

func (t *Traverser) StmtContinue(n *ast.StmtContinue) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) StmtDeclare(n *ast.StmtDeclare) {
	n.Accept(t.v)

	for _, nn := range n.Consts {
		nn.Accept(t)
	}
	t.Traverse(n.Stmt)
}

func (t *Traverser) StmtDefault(n *ast.StmtDefault) {
	n.Accept(t.v)

	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtDo(n *ast.StmtDo) {
	n.Accept(t.v)

	t.Traverse(n.Stmt)
	t.Traverse(n.Cond)
}

func (t *Traverser) StmtEcho(n *ast.StmtEcho) {
	n.Accept(t.v)

	for _, nn := range n.Exprs {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtElse(n *ast.StmtElse) {
	n.Accept(t.v)

	t.Traverse(n.Stmt)
}

func (t *Traverser) StmtElseIf(n *ast.StmtElseIf) {
	n.Accept(t.v)

	t.Traverse(n.Cond)
	t.Traverse(n.Stmt)
}

func (t *Traverser) StmtExpression(n *ast.StmtExpression) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) StmtFinally(n *ast.StmtFinally) {
	n.Accept(t.v)

	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtFor(n *ast.StmtFor) {
	n.Accept(t.v)

	for _, nn := range n.Init {
		nn.Accept(t)
	}
	for _, nn := range n.Cond {
		nn.Accept(t)
	}
	for _, nn := range n.Loop {
		nn.Accept(t)
	}
	t.Traverse(n.Stmt)
}

func (t *Traverser) StmtForeach(n *ast.StmtForeach) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
	t.Traverse(n.Key)
	t.Traverse(n.Var)
	t.Traverse(n.Stmt)
}

func (t *Traverser) StmtFunction(n *ast.StmtFunction) {
	n.Accept(t.v)

	t.Traverse(n.Name)
	for _, nn := range n.Params {
		nn.Accept(t)
	}
	t.Traverse(n.ReturnType)
	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtGlobal(n *ast.StmtGlobal) {
	n.Accept(t.v)

	for _, nn := range n.Vars {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtGoto(n *ast.StmtGoto) {
	n.Accept(t.v)

	t.Traverse(n.Label)
}

func (t *Traverser) StmtHaltCompiler(n *ast.StmtHaltCompiler) {
	n.Accept(t.v)
}

func (t *Traverser) StmtIf(n *ast.StmtIf) {
	n.Accept(t.v)

	t.Traverse(n.Cond)
	t.Traverse(n.Stmt)
	for _, nn := range n.ElseIf {
		nn.Accept(t)
	}
	t.Traverse(n.Else)
}

func (t *Traverser) StmtInlineHtml(n *ast.StmtInlineHtml) {
	n.Accept(t.v)
}

func (t *Traverser) StmtInterface(n *ast.StmtInterface) {
	n.Accept(t.v)

	t.Traverse(n.Name)
	for _, nn := range n.Extends {
		nn.Accept(t)
	}
	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtLabel(n *ast.StmtLabel) {
	n.Accept(t.v)

	t.Traverse(n.Name)
}

func (t *Traverser) StmtNamespace(n *ast.StmtNamespace) {
	n.Accept(t.v)

	t.Traverse(n.Name)
	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtNop(n *ast.StmtNop) {
	n.Accept(t.v)
}

func (t *Traverser) StmtProperty(n *ast.StmtProperty) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) StmtPropertyList(n *ast.StmtPropertyList) {
	n.Accept(t.v)

	for _, nn := range n.Modifiers {
		nn.Accept(t)
	}
	t.Traverse(n.Type)
	for _, nn := range n.Props {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtReturn(n *ast.StmtReturn) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) StmtStatic(n *ast.StmtStatic) {
	n.Accept(t.v)

	for _, nn := range n.Vars {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtStaticVar(n *ast.StmtStaticVar) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) StmtStmtList(n *ast.StmtStmtList) {
	n.Accept(t.v)

	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtSwitch(n *ast.StmtSwitch) {
	n.Accept(t.v)

	t.Traverse(n.Cond)
	for _, nn := range n.Cases {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtThrow(n *ast.StmtThrow) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) StmtTrait(n *ast.StmtTrait) {
	n.Accept(t.v)

	t.Traverse(n.Name)
	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtTraitUse(n *ast.StmtTraitUse) {
	n.Accept(t.v)

	for _, nn := range n.Traits {
		nn.Accept(t)
	}
	for _, nn := range n.Adaptations {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtTraitUseAlias(n *ast.StmtTraitUseAlias) {
	n.Accept(t.v)

	t.Traverse(n.Trait)
	t.Traverse(n.Method)
	t.Traverse(n.Modifier)
	t.Traverse(n.Alias)
}

func (t *Traverser) StmtTraitUsePrecedence(n *ast.StmtTraitUsePrecedence) {
	n.Accept(t.v)

	t.Traverse(n.Trait)
	t.Traverse(n.Method)
	for _, nn := range n.Insteadof {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtTry(n *ast.StmtTry) {
	n.Accept(t.v)

	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
	for _, nn := range n.Catches {
		nn.Accept(t)
	}
	t.Traverse(n.Finally)
}

func (t *Traverser) StmtUnset(n *ast.StmtUnset) {
	n.Accept(t.v)

	for _, nn := range n.Vars {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtUse(n *ast.StmtUseList) {
	n.Accept(t.v)

	t.Traverse(n.Type)
	for _, nn := range n.Uses {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtGroupUse(n *ast.StmtGroupUseList) {
	n.Accept(t.v)

	t.Traverse(n.Type)
	t.Traverse(n.Prefix)
	for _, nn := range n.Uses {
		nn.Accept(t)
	}
}

func (t *Traverser) StmtUseDeclaration(n *ast.StmtUse) {
	n.Accept(t.v)

	t.Traverse(n.Type)
	t.Traverse(n.Use)
	t.Traverse(n.Alias)
}

func (t *Traverser) StmtWhile(n *ast.StmtWhile) {
	n.Accept(t.v)

	t.Traverse(n.Cond)
	t.Traverse(n.Stmt)
}

func (t *Traverser) ExprArray(n *ast.ExprArray) {
	n.Accept(t.v)

	for _, nn := range n.Items {
		nn.Accept(t)
	}
}

func (t *Traverser) ExprArrayDimFetch(n *ast.ExprArrayDimFetch) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Dim)
}

func (t *Traverser) ExprArrayItem(n *ast.ExprArrayItem) {
	n.Accept(t.v)

	t.Traverse(n.Key)
	t.Traverse(n.Val)
}

func (t *Traverser) ExprArrowFunction(n *ast.ExprArrowFunction) {
	n.Accept(t.v)

	for _, nn := range n.Params {
		nn.Accept(t)
	}
	t.Traverse(n.ReturnType)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprBitwiseNot(n *ast.ExprBitwiseNot) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprBooleanNot(n *ast.ExprBooleanNot) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprBrackets(n *ast.ExprBrackets) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	n.Accept(t.v)

	t.Traverse(n.Class)
	t.Traverse(n.Const)
}

func (t *Traverser) ExprClone(n *ast.ExprClone) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprClosure(n *ast.ExprClosure) {
	n.Accept(t.v)

	for _, nn := range n.Params {
		nn.Accept(t)
	}
	for _, nn := range n.Uses {
		nn.Accept(t)
	}
	t.Traverse(n.ReturnType)
	for _, nn := range n.Stmts {
		nn.Accept(t)
	}
}

func (t *Traverser) ExprClosureUse(n *ast.ExprClosureUse) {
	n.Accept(t.v)

	t.Traverse(n.Var)
}

func (t *Traverser) ExprConstFetch(n *ast.ExprConstFetch) {
	n.Accept(t.v)

	t.Traverse(n.Const)
}

func (t *Traverser) ExprEmpty(n *ast.ExprEmpty) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprErrorSuppress(n *ast.ExprErrorSuppress) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprEval(n *ast.ExprEval) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprExit(n *ast.ExprExit) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprFunctionCall(n *ast.ExprFunctionCall) {
	n.Accept(t.v)

	t.Traverse(n.Function)
	for _, nn := range n.Args {
		nn.Accept(t)
	}
}

func (t *Traverser) ExprInclude(n *ast.ExprInclude) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprIncludeOnce(n *ast.ExprIncludeOnce) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprInstanceOf(n *ast.ExprInstanceOf) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
	t.Traverse(n.Class)
}

func (t *Traverser) ExprIsset(n *ast.ExprIsset) {
	n.Accept(t.v)

	for _, nn := range n.Vars {
		nn.Accept(t)
	}
}

func (t *Traverser) ExprList(n *ast.ExprList) {
	n.Accept(t.v)

	for _, nn := range n.Items {
		nn.Accept(t)
	}
}

func (t *Traverser) ExprMethodCall(n *ast.ExprMethodCall) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Method)
	for _, nn := range n.Args {
		nn.Accept(t)
	}
}

func (t *Traverser) ExprNew(n *ast.ExprNew) {
	n.Accept(t.v)

	t.Traverse(n.Class)
	for _, nn := range n.Args {
		nn.Accept(t)
	}
}

func (t *Traverser) ExprPostDec(n *ast.ExprPostDec) {
	n.Accept(t.v)

	t.Traverse(n.Var)
}

func (t *Traverser) ExprPostInc(n *ast.ExprPostInc) {
	n.Accept(t.v)

	t.Traverse(n.Var)
}

func (t *Traverser) ExprPreDec(n *ast.ExprPreDec) {
	n.Accept(t.v)

	t.Traverse(n.Var)
}

func (t *Traverser) ExprPreInc(n *ast.ExprPreInc) {
	n.Accept(t.v)

	t.Traverse(n.Var)
}

func (t *Traverser) ExprPrint(n *ast.ExprPrint) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprPropertyFetch(n *ast.ExprPropertyFetch) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Prop)
}

func (t *Traverser) ExprRequire(n *ast.ExprRequire) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprRequireOnce(n *ast.ExprRequireOnce) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprShellExec(n *ast.ExprShellExec) {
	n.Accept(t.v)

	for _, nn := range n.Parts {
		nn.Accept(t)
	}
}

func (t *Traverser) ExprStaticCall(n *ast.ExprStaticCall) {
	n.Accept(t.v)

	t.Traverse(n.Class)
	t.Traverse(n.Call)
	for _, nn := range n.Args {
		nn.Accept(t)
	}
}

func (t *Traverser) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch) {
	n.Accept(t.v)

	t.Traverse(n.Class)
	t.Traverse(n.Prop)
}

func (t *Traverser) ExprTernary(n *ast.ExprTernary) {
	n.Accept(t.v)

	t.Traverse(n.Cond)
	t.Traverse(n.IfTrue)
	t.Traverse(n.IfFalse)
}

func (t *Traverser) ExprUnaryMinus(n *ast.ExprUnaryMinus) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprUnaryPlus(n *ast.ExprUnaryPlus) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprVariable(n *ast.ExprVariable) {
	n.Accept(t.v)

	t.Traverse(n.Name)
}

func (t *Traverser) ExprYield(n *ast.ExprYield) {
	n.Accept(t.v)

	t.Traverse(n.Key)
	t.Traverse(n.Val)
}

func (t *Traverser) ExprYieldFrom(n *ast.ExprYieldFrom) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssign(n *ast.ExprAssign) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignReference(n *ast.ExprAssignReference) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignBitwiseAnd(n *ast.ExprAssignBitwiseAnd) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignBitwiseOr(n *ast.ExprAssignBitwiseOr) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignBitwiseXor(n *ast.ExprAssignBitwiseXor) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignCoalesce(n *ast.ExprAssignCoalesce) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignConcat(n *ast.ExprAssignConcat) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignDiv(n *ast.ExprAssignDiv) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignMinus(n *ast.ExprAssignMinus) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignMod(n *ast.ExprAssignMod) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignMul(n *ast.ExprAssignMul) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignPlus(n *ast.ExprAssignPlus) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignPow(n *ast.ExprAssignPow) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignShiftLeft(n *ast.ExprAssignShiftLeft) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprAssignShiftRight(n *ast.ExprAssignShiftRight) {
	n.Accept(t.v)

	t.Traverse(n.Var)
	t.Traverse(n.Expr)
}

func (t *Traverser) ExprBinaryBitwiseAnd(n *ast.ExprBinaryBitwiseAnd) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryBitwiseOr(n *ast.ExprBinaryBitwiseOr) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryBitwiseXor(n *ast.ExprBinaryBitwiseXor) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryBooleanAnd(n *ast.ExprBinaryBooleanAnd) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryBooleanOr(n *ast.ExprBinaryBooleanOr) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryCoalesce(n *ast.ExprBinaryCoalesce) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryConcat(n *ast.ExprBinaryConcat) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryDiv(n *ast.ExprBinaryDiv) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryEqual(n *ast.ExprBinaryEqual) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryGreater(n *ast.ExprBinaryGreater) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryGreaterOrEqual(n *ast.ExprBinaryGreaterOrEqual) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryIdentical(n *ast.ExprBinaryIdentical) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryLogicalAnd(n *ast.ExprBinaryLogicalAnd) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryLogicalOr(n *ast.ExprBinaryLogicalOr) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryLogicalXor(n *ast.ExprBinaryLogicalXor) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryMinus(n *ast.ExprBinaryMinus) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryMod(n *ast.ExprBinaryMod) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryMul(n *ast.ExprBinaryMul) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryNotEqual(n *ast.ExprBinaryNotEqual) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryNotIdentical(n *ast.ExprBinaryNotIdentical) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryPlus(n *ast.ExprBinaryPlus) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryPow(n *ast.ExprBinaryPow) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryShiftLeft(n *ast.ExprBinaryShiftLeft) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinaryShiftRight(n *ast.ExprBinaryShiftRight) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinarySmaller(n *ast.ExprBinarySmaller) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinarySmallerOrEqual(n *ast.ExprBinarySmallerOrEqual) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprBinarySpaceship(n *ast.ExprBinarySpaceship) {
	n.Accept(t.v)

	t.Traverse(n.Left)
	t.Traverse(n.Right)
}

func (t *Traverser) ExprCastArray(n *ast.ExprCastArray) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprCastBool(n *ast.ExprCastBool) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprCastDouble(n *ast.ExprCastDouble) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprCastInt(n *ast.ExprCastInt) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprCastObject(n *ast.ExprCastObject) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprCastString(n *ast.ExprCastString) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ExprCastUnset(n *ast.ExprCastUnset) {
	n.Accept(t.v)

	t.Traverse(n.Expr)
}

func (t *Traverser) ScalarDnumber(n *ast.ScalarDnumber) {
	n.Accept(t.v)
}

func (t *Traverser) ScalarEncapsed(n *ast.ScalarEncapsed) {
	n.Accept(t.v)

	for _, nn := range n.Parts {
		nn.Accept(t)
	}
}

func (t *Traverser) ScalarEncapsedStringPart(n *ast.ScalarEncapsedStringPart) {
	n.Accept(t.v)
}

func (t *Traverser) ScalarEncapsedStringVar(n *ast.ScalarEncapsedStringVar) {
	n.Accept(t.v)

	t.Traverse(n.Name)
	t.Traverse(n.Dim)
}

func (t *Traverser) ScalarEncapsedStringBrackets(n *ast.ScalarEncapsedStringBrackets) {
	n.Accept(t.v)

	t.Traverse(n.Var)
}

func (t *Traverser) ScalarHeredoc(n *ast.ScalarHeredoc) {
	n.Accept(t.v)

	for _, nn := range n.Parts {
		nn.Accept(t)
	}
}

func (t *Traverser) ScalarLnumber(n *ast.ScalarLnumber) {
	n.Accept(t.v)
}

func (t *Traverser) ScalarMagicConstant(n *ast.ScalarMagicConstant) {
	n.Accept(t.v)
}

func (t *Traverser) ScalarString(n *ast.ScalarString) {
	n.Accept(t.v)
}

func (t *Traverser) NameName(n *ast.Name) {
	n.Accept(t.v)

	for _, nn := range n.Parts {
		nn.Accept(t)
	}
}

func (t *Traverser) NameFullyQualified(n *ast.NameFullyQualified) {
	n.Accept(t.v)

	for _, nn := range n.Parts {
		nn.Accept(t)
	}
}

func (t *Traverser) NameRelative(n *ast.NameRelative) {
	n.Accept(t.v)

	for _, nn := range n.Parts {
		nn.Accept(t)
	}
}

func (t *Traverser) NameNamePart(n *ast.NamePart) {
	n.Accept(t.v)
}
