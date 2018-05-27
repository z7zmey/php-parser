package comment

// TokenName is used to specify a comment position
type TokenName int

const (
	UnknownToken TokenName = iota
	IncludeToken
	IncludeOnceToken
	ExitToken
	IfToken
	LnumberToken
	DnumberToken
	StringToken
	StringVarnameToken
	VariableToken
	NumStringToken
	InlineHTMLToken
	EncapsedAndWhitespaceToken
	ConstantEncapsedStringToken
	EchoToken
	DoToken
	WhileToken
	EndwhileToken
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
	ClassCToken
	TraitCToken
	MethodCToken
	FuncCToken
	LineToken
	FileToken
	StartHeredocToken
	DollarOpenCurlyBracesToken
	CurlyOpenToken
	PaamayimNekudotayimToken
	NamespaceToken
	NsCToken
	DirToken
	NsSeparatorToken
	EllipsisToken
	EvalToken
	RequireToken
	RequireOnceToken
	LogicalOrToken
	LogicalXorToken
	LogicalAndToken
	InstanceofToken
	NewToken
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
	IdentifierToken
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
	DollarToken           // '$'
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
