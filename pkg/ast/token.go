package ast

import (
	"encoding/json"

	"github.com/z7zmey/php-parser/internal/scanner"
)

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

type Token struct {
	Type     scanner.TokenType
	Group    TokenGroup
	Position Position
	Value    string
}

func (t Token) MarshalJSON() ([]byte, error) {
	out := struct {
		Type     string   `json:"type"`
		Value    string   `json:"value"`
		Position Position `json:"position"`
	}{
		Type:     t.Type.String(),
		Value:    t.Value,
		Position: t.Position,
	}

	return json.Marshal(out)
}
