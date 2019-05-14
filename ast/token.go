package ast

type TokenGroup uint8

//go:generate stringer -type=TokenGroup -trimprefix=TokenGroup -output ./tokengroup_string.go
const (
	TokenGroupStart TokenGroup = iota
	TokenGroupEnd
	TokenGroupSlash
	TokenGroupColon
	TokenGroupSemiColon
	TokenGroupAltEnd
	TokenGroupDollar
	TokenGroupAmpersand
	TokenGroupName
	TokenGroupPrefix
	TokenGroupKey
	TokenGroupVar
	TokenGroupUseType
	TokenGroupReturnType
	TokenGroupOptionalType
	TokenGroupCaseSeparator
	TokenGroupLexicalVars
	TokenGroupParams
	TokenGroupRef
	TokenGroupCast
	TokenGroupExpr
	TokenGroupInitExpr
	TokenGroupCondExpr
	TokenGroupIncExpr
	TokenGroupTrue
	TokenGroupCond

	TokenGroupHaltCompiller
	TokenGroupNamespace
	TokenGroupStatic
	TokenGroupClass
	TokenGroupUse
	TokenGroupWhile
	TokenGroupFor
	TokenGroupSwitch
	TokenGroupBreak
	TokenGroupForeach
	TokenGroupDeclare
	TokenGroupLabel
	TokenGroupFinally
	TokenGroupList
	TokenGroupDefault
	TokenGroupIf
	TokenGroupElseIf
	TokenGroupElse
	TokenGroupVariadic
	TokenGroupFunction
	TokenGroupAlias
	TokenGroupAs
	TokenGroupEqual
	TokenGroupExit
	TokenGroupArray
	TokenGroupIsset
	TokenGroupEmpty
	TokenGroupEval
	TokenGroupEcho
	TokenGroupTry
	TokenGroupCatch
	TokenGroupUnset

	TokenGroupStmts
	TokenGroupVarList
	TokenGroupConstList
	TokenGroupNameList
	TokenGroupParamList
	TokenGroupModifierList
	TokenGroupArrayPairList
	TokenGroupCaseListStart
	TokenGroupCaseListEnd
	TokenGroupArgumentList
	TokenGroupPropertyList
	TokenGroupParameterList
	TokenGroupAdaptationList
	TokenGroupLexicalVarList
	TokenGroupUseDeclarationList

	TokenGroupOpenParenthesisToken
	TokenGroupCloseParenthesisToken
)
