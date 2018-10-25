package meta

// TokenName is used to specify a comment position
type TokenName int

//go:generate stringer -type=TokenName -output ./tokenName_string.go
const (
	NodeStart TokenName = iota
	NodeEnd
	MagicConstantToken
	IncludeToken
	IncludeOnceToken
	ExitToken
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
	ForToken
	EndforToken
	ForeachToken
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
	UseToken
	InsteadofToken
	GlobalToken
	VarToken
	UnsetToken
	IssetToken
	EmptyToken
	ClassToken
	TraitToken
	InterfaceToken
	ExtendsToken
	ImplementsToken
	DoubleArrowToken
	ListToken
	ArrayToken
	CallableToken
	StartHeredocToken
	CurlyOpenToken
	PaamayimNekudotayimToken
	UseLeadingNsSeparatorToken
	NsSeparatorToken
	EllipsisToken
	EvalToken
	RequireToken
	RequireOnceToken
	LogicalOrToken
	LogicalXorToken
	LogicalAndToken
	InstanceofToken
	NewAnchor
	CloneToken
	ElseifToken
	ElseToken
	EndifToken
	PrintToken
	YieldToken
	StaticToken
	AbstractToken
	FinalToken
	PrivateToken
	ProtectedToken
	PublicToken
	IncToken
	DecToken
	YieldFromToken
	ObjectOperatorToken
	IntCastToken
	DoubleCastToken
	StringCastToken
	ArrayCastToken
	ObjectCastToken
	BoolCastToken
	UnsetCastToken
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
	DoubleQuoteToken      // '"'
	BackquoteToken        // '`'
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
	ExclamationMarkToken  // '!'
	TildeToken            // '~'
	AtToken               // '@'
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
