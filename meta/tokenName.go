package meta

// TokenName is used to specify a comment position
type TokenName int

//go:generate stringer -type=TokenName -output ./tokenName_string.go
const (
	NodeStart TokenName = iota
	NodeEnd
	IfToken
	StringVarnameToken
	NumStringToken
	InlineHTMLToken
	EchoToken
	DoToken
	WhileToken
	EndwhileToken
	ForInitSemicolonToken
	ForCondSemicolonToken
	EndforToken
	EndforeachToken
	DeclareToken
	EnddeclareToken
	AsToken
	SwitchToken
	EndswitchToken
	CaseToken
	DefaultToken
	BreakToken
	ContinueToken
	GotoToken
	FunctionToken
	ConstToken
	ReturnToken
	TryToken
	CatchToken
	FinallyToken
	ThrowToken
	InsteadofToken
	GlobalToken
	VarToken
	UnsetToken
	ClassToken
	TraitToken
	InterfaceToken
	ExtendsToken
	ImplementsToken
	DoubleArrowToken
	CallableToken
	CurlyOpenToken
	PaamayimNekudotayimToken
	UseLeadingNsSeparatorToken
	NsSeparatorToken
	EllipsisToken
	LogicalOrToken
	LogicalXorToken
	LogicalAndToken
	InstanceofToken
	EndifToken
	AbstractToken
	FinalToken
	PrivateToken
	ProtectedToken
	PublicToken
	IncToken
	DecToken
	ObjectOperatorToken
	CoalesceToken
	SpaceshipToken
	PlusEqualToken
	MinusEqualToken
	MulEqualToken
	PowEqualToken
	DivEqualToken
	ConcatEqualToken
	ModEqualToken
	AndEqualToken
	OrEqualToken
	XorEqualToken
	SlEqualToken
	SrEqualToken
	BooleanOrToken
	BooleanAndToken
	PowToken
	SlToken
	SrToken
	IsIdenticalToken
	IsNotIdenticalToken
	IsEqualToken
	IsNotEqualToken
	IsSmallerOrEqualToken
	IsGreaterOrEqualToken
	HaltCompilerToken
	CaseSeparatorToken    // ';' or ':'
	OpenCurlyBracesToken  // '{'
	CloseCurlyBracesToken // '}'
	SemiColonToken        // ';'
	ColonToken            // ':'
	OpenParenthesisToken  // '('
	CloseParenthesisToken // ')'
	OpenSquareBracket     // '['
	CloseSquareBracket    // ']'
	QuestionMarkToken     // '?'
	AmpersandToken        // '&'
	MinusToken            // '-'
	PlusToken             // '+'
	CommaToken            // ','
	VerticalBarToken      // '|'
	EqualToken            // '='
	CaretToken            // '^'
	AsteriskToken         // '*'
	SlashToken            // '/'
	PercentToken          // '%'
	LessToken             // '<'
	GreaterToken          // '>'
	DotToken              // '.'
)
