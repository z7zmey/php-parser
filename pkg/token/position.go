package token

type Position int

//go:generate stringer -type=Position -output ./position_string.go
const (
	Start Position = iota
	End
	SemiColon
	AltEnd
	Ampersand
	Name
	Key
	Var
	ReturnType
	CaseSeparator
	LexicalVars
	Params
	Ref
	Cast
	Expr
	InitExpr
	CondExpr
	IncExpr
	True
	Cond

	Namespace
	Static
	Use
	For
	Foreach
	Declare
	Label
	Finally
	List
	Default
	Function
	Alias
	Equal
	Array
	Isset
	Echo
	Try
	Catch
	Unset

	Stmts
	VarList
	ConstList
	NameList
	ParamList
	ModifierList
	ArrayPairList
	CaseListStart
	CaseListEnd
	PropertyList
	ParameterList
	AdaptationList
	LexicalVarList

	CloseParenthesisToken
)

type Collection map[Position][]*Token

func (c Collection) IsEmpty() bool {
	for _, v := range c {
		if len(v) > 0 {
			return false
		}
	}
	return true
}
