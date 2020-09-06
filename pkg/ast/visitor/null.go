package visitor

import (
	"github.com/z7zmey/php-parser/pkg/ast"
)

type Null struct {
}

func (v *Null) Enter(_ string, _ bool) {
	// do nothing
}
func (v *Null) Leave(_ string, _ bool) {
	// do nothing
}

func (v *Null) EnterNode(_ ast.Vertex) bool {
	return true
}

func (v *Null) LeaveNode(_ ast.Vertex) {
	// do nothing
}

func (v *Null) Root(_ *ast.Root) {
	// do nothing
}

func (v *Null) Nullable(_ *ast.Nullable) {
	// do nothing
}

func (v *Null) Reference(_ *ast.Reference) {
	// do nothing
}

func (v *Null) Variadic(_ *ast.Variadic) {
	// do nothing
}

func (v *Null) Parameter(_ *ast.Parameter) {
	// do nothing
}

func (v *Null) Identifier(_ *ast.Identifier) {
	// do nothing
}

func (v *Null) ArgumentList(_ *ast.ArgumentList) {
	// do nothing
}

func (v *Null) Argument(_ *ast.Argument) {
	// do nothing
}

func (v *Null) StmtAltForeach(_ *ast.StmtAltForeach) {
	// do nothing
}

func (v *Null) StmtBreak(_ *ast.StmtBreak) {
	// do nothing
}

func (v *Null) StmtCase(_ *ast.StmtCase) {
	// do nothing
}

func (v *Null) StmtCatch(_ *ast.StmtCatch) {
	// do nothing
}

func (v *Null) StmtClass(_ *ast.StmtClass) {
	// do nothing
}

func (v *Null) StmtClassConstList(_ *ast.StmtClassConstList) {
	// do nothing
}

func (v *Null) StmtClassExtends(_ *ast.StmtClassExtends) {
	// do nothing
}

func (v *Null) StmtClassImplements(_ *ast.StmtClassImplements) {
	// do nothing
}

func (v *Null) StmtClassMethod(_ *ast.StmtClassMethod) {
	// do nothing
}

func (v *Null) StmtConstList(_ *ast.StmtConstList) {
	// do nothing
}

func (v *Null) StmtConstant(_ *ast.StmtConstant) {
	// do nothing
}

func (v *Null) StmtContinue(_ *ast.StmtContinue) {
	// do nothing
}

func (v *Null) StmtDeclare(_ *ast.StmtDeclare) {
	// do nothing
}

func (v *Null) StmtDefault(_ *ast.StmtDefault) {
	// do nothing
}

func (v *Null) StmtDo(_ *ast.StmtDo) {
	// do nothing
}

func (v *Null) StmtEcho(_ *ast.StmtEcho) {
	// do nothing
}

func (v *Null) StmtElse(_ *ast.StmtElse) {
	// do nothing
}

func (v *Null) StmtElseIf(_ *ast.StmtElseIf) {
	// do nothing
}

func (v *Null) StmtExpression(_ *ast.StmtExpression) {
	// do nothing
}

func (v *Null) StmtFinally(_ *ast.StmtFinally) {
	// do nothing
}

func (v *Null) StmtFor(_ *ast.StmtFor) {
	// do nothing
}

func (v *Null) StmtForeach(_ *ast.StmtForeach) {
	// do nothing
}

func (v *Null) StmtFunction(_ *ast.StmtFunction) {
	// do nothing
}

func (v *Null) StmtGlobal(_ *ast.StmtGlobal) {
	// do nothing
}

func (v *Null) StmtGoto(_ *ast.StmtGoto) {
	// do nothing
}

func (v *Null) StmtHaltCompiler(_ *ast.StmtHaltCompiler) {
	// do nothing
}

func (v *Null) StmtIf(_ *ast.StmtIf) {
	// do nothing
}

func (v *Null) StmtInlineHtml(_ *ast.StmtInlineHtml) {
	// do nothing
}

func (v *Null) StmtInterface(_ *ast.StmtInterface) {
	// do nothing
}

func (v *Null) StmtInterfaceExtends(_ *ast.StmtInterfaceExtends) {
	// do nothing
}

func (v *Null) StmtLabel(_ *ast.StmtLabel) {
	// do nothing
}

func (v *Null) StmtNamespace(_ *ast.StmtNamespace) {
	// do nothing
}

func (v *Null) StmtNop(_ *ast.StmtNop) {
	// do nothing
}

func (v *Null) StmtProperty(_ *ast.StmtProperty) {
	// do nothing
}

func (v *Null) StmtPropertyList(_ *ast.StmtPropertyList) {
	// do nothing
}

func (v *Null) StmtReturn(_ *ast.StmtReturn) {
	// do nothing
}

func (v *Null) StmtStatic(_ *ast.StmtStatic) {
	// do nothing
}

func (v *Null) StmtStaticVar(_ *ast.StmtStaticVar) {
	// do nothing
}

func (v *Null) StmtStmtList(_ *ast.StmtStmtList) {
	// do nothing
}

func (v *Null) StmtSwitch(_ *ast.StmtSwitch) {
	// do nothing
}

func (v *Null) StmtThrow(_ *ast.StmtThrow) {
	// do nothing
}

func (v *Null) StmtTrait(_ *ast.StmtTrait) {
	// do nothing
}

func (v *Null) StmtTraitAdaptationList(_ *ast.StmtTraitAdaptationList) {
	// do nothing
}

func (v *Null) StmtTraitMethodRef(_ *ast.StmtTraitMethodRef) {
	// do nothing
}

func (v *Null) StmtTraitUse(_ *ast.StmtTraitUse) {
	// do nothing
}

func (v *Null) StmtTraitUseAlias(_ *ast.StmtTraitUseAlias) {
	// do nothing
}

func (v *Null) StmtTraitUsePrecedence(_ *ast.StmtTraitUsePrecedence) {
	// do nothing
}

func (v *Null) StmtTry(_ *ast.StmtTry) {
	// do nothing
}

func (v *Null) StmtUnset(_ *ast.StmtUnset) {
	// do nothing
}

func (v *Null) StmtUse(_ *ast.StmtUse) {
	// do nothing
}

func (v *Null) StmtGroupUse(_ *ast.StmtGroupUse) {
	// do nothing
}

func (v *Null) StmtUseDeclaration(_ *ast.StmtUseDeclaration) {
	// do nothing
}

func (v *Null) StmtWhile(_ *ast.StmtWhile) {
	// do nothing
}

func (v *Null) ExprArray(_ *ast.ExprArray) {
	// do nothing
}

func (v *Null) ExprArrayDimFetch(_ *ast.ExprArrayDimFetch) {
	// do nothing
}

func (v *Null) ExprArrayItem(_ *ast.ExprArrayItem) {
	// do nothing
}

func (v *Null) ExprArrowFunction(_ *ast.ExprArrowFunction) {
	// do nothing
}

func (v *Null) ExprBitwiseNot(_ *ast.ExprBitwiseNot) {
	// do nothing
}

func (v *Null) ExprBooleanNot(_ *ast.ExprBooleanNot) {
	// do nothing
}

func (v *Null) ExprClassConstFetch(_ *ast.ExprClassConstFetch) {
	// do nothing
}

func (v *Null) ExprClone(_ *ast.ExprClone) {
	// do nothing
}

func (v *Null) ExprClosure(_ *ast.ExprClosure) {
	// do nothing
}

func (v *Null) ExprClosureUse(_ *ast.ExprClosureUse) {
	// do nothing
}

func (v *Null) ExprConstFetch(_ *ast.ExprConstFetch) {
	// do nothing
}

func (v *Null) ExprEmpty(_ *ast.ExprEmpty) {
	// do nothing
}

func (v *Null) ExprErrorSuppress(_ *ast.ExprErrorSuppress) {
	// do nothing
}

func (v *Null) ExprEval(_ *ast.ExprEval) {
	// do nothing
}

func (v *Null) ExprExit(_ *ast.ExprExit) {
	// do nothing
}

func (v *Null) ExprFunctionCall(_ *ast.ExprFunctionCall) {
	// do nothing
}

func (v *Null) ExprInclude(_ *ast.ExprInclude) {
	// do nothing
}

func (v *Null) ExprIncludeOnce(_ *ast.ExprIncludeOnce) {
	// do nothing
}

func (v *Null) ExprInstanceOf(_ *ast.ExprInstanceOf) {
	// do nothing
}

func (v *Null) ExprIsset(_ *ast.ExprIsset) {
	// do nothing
}

func (v *Null) ExprList(_ *ast.ExprList) {
	// do nothing
}

func (v *Null) ExprMethodCall(_ *ast.ExprMethodCall) {
	// do nothing
}

func (v *Null) ExprNew(_ *ast.ExprNew) {
	// do nothing
}

func (v *Null) ExprPostDec(_ *ast.ExprPostDec) {
	// do nothing
}

func (v *Null) ExprPostInc(_ *ast.ExprPostInc) {
	// do nothing
}

func (v *Null) ExprPreDec(_ *ast.ExprPreDec) {
	// do nothing
}

func (v *Null) ExprPreInc(_ *ast.ExprPreInc) {
	// do nothing
}

func (v *Null) ExprPrint(_ *ast.ExprPrint) {
	// do nothing
}

func (v *Null) ExprPropertyFetch(_ *ast.ExprPropertyFetch) {
	// do nothing
}

func (v *Null) ExprReference(_ *ast.ExprReference) {
	// do nothing
}

func (v *Null) ExprRequire(_ *ast.ExprRequire) {
	// do nothing
}

func (v *Null) ExprRequireOnce(_ *ast.ExprRequireOnce) {
	// do nothing
}

func (v *Null) ExprShellExec(_ *ast.ExprShellExec) {
	// do nothing
}

func (v *Null) ExprShortArray(_ *ast.ExprShortArray) {
	// do nothing
}

func (v *Null) ExprShortList(_ *ast.ExprShortList) {
	// do nothing
}

func (v *Null) ExprStaticCall(_ *ast.ExprStaticCall) {
	// do nothing
}

func (v *Null) ExprStaticPropertyFetch(_ *ast.ExprStaticPropertyFetch) {
	// do nothing
}

func (v *Null) ExprTernary(_ *ast.ExprTernary) {
	// do nothing
}

func (v *Null) ExprUnaryMinus(_ *ast.ExprUnaryMinus) {
	// do nothing
}

func (v *Null) ExprUnaryPlus(_ *ast.ExprUnaryPlus) {
	// do nothing
}

func (v *Null) ExprVariable(_ *ast.ExprVariable) {
	// do nothing
}

func (v *Null) ExprYield(_ *ast.ExprYield) {
	// do nothing
}

func (v *Null) ExprYieldFrom(_ *ast.ExprYieldFrom) {
	// do nothing
}

func (v *Null) ExprAssign(_ *ast.ExprAssign) {
	// do nothing
}

func (v *Null) ExprAssignReference(_ *ast.ExprAssignReference) {
	// do nothing
}

func (v *Null) ExprAssignBitwiseAnd(_ *ast.ExprAssignBitwiseAnd) {
	// do nothing
}

func (v *Null) ExprAssignBitwiseOr(_ *ast.ExprAssignBitwiseOr) {
	// do nothing
}

func (v *Null) ExprAssignBitwiseXor(_ *ast.ExprAssignBitwiseXor) {
	// do nothing
}

func (v *Null) ExprAssignCoalesce(_ *ast.ExprAssignCoalesce) {
	// do nothing
}

func (v *Null) ExprAssignConcat(_ *ast.ExprAssignConcat) {
	// do nothing
}

func (v *Null) ExprAssignDiv(_ *ast.ExprAssignDiv) {
	// do nothing
}

func (v *Null) ExprAssignMinus(_ *ast.ExprAssignMinus) {
	// do nothing
}

func (v *Null) ExprAssignMod(_ *ast.ExprAssignMod) {
	// do nothing
}

func (v *Null) ExprAssignMul(_ *ast.ExprAssignMul) {
	// do nothing
}

func (v *Null) ExprAssignPlus(_ *ast.ExprAssignPlus) {
	// do nothing
}

func (v *Null) ExprAssignPow(_ *ast.ExprAssignPow) {
	// do nothing
}

func (v *Null) ExprAssignShiftLeft(_ *ast.ExprAssignShiftLeft) {
	// do nothing
}

func (v *Null) ExprAssignShiftRight(_ *ast.ExprAssignShiftRight) {
	// do nothing
}

func (v *Null) ExprBinaryBitwiseAnd(_ *ast.ExprBinaryBitwiseAnd) {
	// do nothing
}

func (v *Null) ExprBinaryBitwiseOr(_ *ast.ExprBinaryBitwiseOr) {
	// do nothing
}

func (v *Null) ExprBinaryBitwiseXor(_ *ast.ExprBinaryBitwiseXor) {
	// do nothing
}

func (v *Null) ExprBinaryBooleanAnd(_ *ast.ExprBinaryBooleanAnd) {
	// do nothing
}

func (v *Null) ExprBinaryBooleanOr(_ *ast.ExprBinaryBooleanOr) {
	// do nothing
}

func (v *Null) ExprBinaryCoalesce(_ *ast.ExprBinaryCoalesce) {
	// do nothing
}

func (v *Null) ExprBinaryConcat(_ *ast.ExprBinaryConcat) {
	// do nothing
}

func (v *Null) ExprBinaryDiv(_ *ast.ExprBinaryDiv) {
	// do nothing
}

func (v *Null) ExprBinaryEqual(_ *ast.ExprBinaryEqual) {
	// do nothing
}

func (v *Null) ExprBinaryGreater(_ *ast.ExprBinaryGreater) {
	// do nothing
}

func (v *Null) ExprBinaryGreaterOrEqual(_ *ast.ExprBinaryGreaterOrEqual) {
	// do nothing
}

func (v *Null) ExprBinaryIdentical(_ *ast.ExprBinaryIdentical) {
	// do nothing
}

func (v *Null) ExprBinaryLogicalAnd(_ *ast.ExprBinaryLogicalAnd) {
	// do nothing
}

func (v *Null) ExprBinaryLogicalOr(_ *ast.ExprBinaryLogicalOr) {
	// do nothing
}

func (v *Null) ExprBinaryLogicalXor(_ *ast.ExprBinaryLogicalXor) {
	// do nothing
}

func (v *Null) ExprBinaryMinus(_ *ast.ExprBinaryMinus) {
	// do nothing
}

func (v *Null) ExprBinaryMod(_ *ast.ExprBinaryMod) {
	// do nothing
}

func (v *Null) ExprBinaryMul(_ *ast.ExprBinaryMul) {
	// do nothing
}

func (v *Null) ExprBinaryNotEqual(_ *ast.ExprBinaryNotEqual) {
	// do nothing
}

func (v *Null) ExprBinaryNotIdentical(_ *ast.ExprBinaryNotIdentical) {
	// do nothing
}

func (v *Null) ExprBinaryPlus(_ *ast.ExprBinaryPlus) {
	// do nothing
}

func (v *Null) ExprBinaryPow(_ *ast.ExprBinaryPow) {
	// do nothing
}

func (v *Null) ExprBinaryShiftLeft(_ *ast.ExprBinaryShiftLeft) {
	// do nothing
}

func (v *Null) ExprBinaryShiftRight(_ *ast.ExprBinaryShiftRight) {
	// do nothing
}

func (v *Null) ExprBinarySmaller(_ *ast.ExprBinarySmaller) {
	// do nothing
}

func (v *Null) ExprBinarySmallerOrEqual(_ *ast.ExprBinarySmallerOrEqual) {
	// do nothing
}

func (v *Null) ExprBinarySpaceship(_ *ast.ExprBinarySpaceship) {
	// do nothing
}

func (v *Null) ExprCastArray(_ *ast.ExprCastArray) {
	// do nothing
}

func (v *Null) ExprCastBool(_ *ast.ExprCastBool) {
	// do nothing
}

func (v *Null) ExprCastDouble(_ *ast.ExprCastDouble) {
	// do nothing
}

func (v *Null) ExprCastInt(_ *ast.ExprCastInt) {
	// do nothing
}

func (v *Null) ExprCastObject(_ *ast.ExprCastObject) {
	// do nothing
}

func (v *Null) ExprCastString(_ *ast.ExprCastString) {
	// do nothing
}

func (v *Null) ExprCastUnset(_ *ast.ExprCastUnset) {
	// do nothing
}

func (v *Null) ScalarDnumber(_ *ast.ScalarDnumber) {
	// do nothing
}

func (v *Null) ScalarEncapsed(_ *ast.ScalarEncapsed) {
	// do nothing
}

func (v *Null) ScalarEncapsedStringPart(_ *ast.ScalarEncapsedStringPart) {
	// do nothing
}

func (v *Null) ScalarHeredoc(_ *ast.ScalarHeredoc) {
	// do nothing
}

func (v *Null) ScalarLnumber(_ *ast.ScalarLnumber) {
	// do nothing
}

func (v *Null) ScalarMagicConstant(_ *ast.ScalarMagicConstant) {
	// do nothing
}

func (v *Null) ScalarString(_ *ast.ScalarString) {
	// do nothing
}

func (v *Null) NameName(_ *ast.NameName) {
	// do nothing
}

func (v *Null) NameFullyQualified(_ *ast.NameFullyQualified) {
	// do nothing
}

func (v *Null) NameRelative(_ *ast.NameRelative) {
	// do nothing
}

func (v *Null) NameNamePart(_ *ast.NameNamePart) {
	// do nothing
}

func (v *Null) ParserAs(_ *ast.ParserAs) {
	// do nothing
}

func (v *Null) ParserNsSeparator(_ *ast.ParserNsSeparator) {
	// do nothing
}

func (v *Null) ParserBrackets(_ *ast.ParserBrackets) {
	// do nothing
}
