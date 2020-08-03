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
	Class
	Use
	While
	For
	Switch
	Foreach
	Declare
	Label
	Finally
	List
	Default
	If
	ElseIf
	Else
	Function
	Alias
	Equal
	Exit
	Array
	Isset
	Empty
	Eval
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
	ArgumentList
	PropertyList
	ParameterList
	AdaptationList
	LexicalVarList

	OpenParenthesisToken
	CloseParenthesisToken
)

type Collection map[Position][]Token

func (c Collection) IsEmpty() bool {
	for _, v := range c {
		if len(v) > 0 {
			return false
		}
	}
	return true
}
