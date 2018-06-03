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

var TokenNames = map[TokenName]string{
	UnknownToken:                "UnknownToken",
	IncludeToken:                "IncludeToken",
	IncludeOnceToken:            "IncludeOnceToken",
	ExitToken:                   "ExitToken",
	IfToken:                     "IfToken",
	LnumberToken:                "LnumberToken",
	DnumberToken:                "DnumberToken",
	StringToken:                 "StringToken",
	StringVarnameToken:          "StringVarnameToken",
	VariableToken:               "VariableToken",
	NumStringToken:              "NumStringToken",
	InlineHTMLToken:             "InlineHTMLToken",
	EncapsedAndWhitespaceToken:  "EncapsedAndWhitespaceToken",
	ConstantEncapsedStringToken: "ConstantEncapsedStringToken",
	EchoToken:                   "EchoToken",
	DoToken:                     "DoToken",
	WhileToken:                  "WhileToken",
	EndwhileToken:               "EndwhileToken",
	ForInitSemicolonToken:       "ForInitSemicolonToken",
	ForCondSemicolonToken:       "ForCondSemicolonToken",
	ForToken:                    "ForToken",
	EndforToken:                 "EndforToken",
	ForeachToken:                "ForeachToken",
	EndforeachToken:             "EndforeachToken",
	DeclareToken:                "DeclareToken",
	EnddeclareToken:             "EnddeclareToken",
	AsToken:                     "AsToken",
	SwitchToken:                 "SwitchToken",
	EndswitchToken:              "EndswitchToken",
	CaseToken:                   "CaseToken",
	DefaultToken:                "DefaultToken",
	BreakToken:                  "BreakToken",
	ContinueToken:               "ContinueToken",
	GotoToken:                   "GotoToken",
	FunctionToken:               "FunctionToken",
	ConstToken:                  "ConstToken",
	ReturnToken:                 "ReturnToken",
	TryToken:                    "TryToken",
	CatchToken:                  "CatchToken",
	FinallyToken:                "FinallyToken",
	ThrowToken:                  "ThrowToken",
	UseToken:                    "UseToken",
	InsteadofToken:              "InsteadofToken",
	GlobalToken:                 "GlobalToken",
	VarToken:                    "VarToken",
	UnsetToken:                  "UnsetToken",
	IssetToken:                  "IssetToken",
	EmptyToken:                  "EmptyToken",
	ClassToken:                  "ClassToken",
	TraitToken:                  "TraitToken",
	InterfaceToken:              "InterfaceToken",
	ExtendsToken:                "ExtendsToken",
	ImplementsToken:             "ImplementsToken",
	DoubleArrowToken:            "DoubleArrowToken",
	ListToken:                   "ListToken",
	ArrayToken:                  "ArrayToken",
	CallableToken:               "CallableToken",
	ClassCToken:                 "ClassCToken",
	TraitCToken:                 "TraitCToken",
	MethodCToken:                "MethodCToken",
	FuncCToken:                  "FuncCToken",
	LineToken:                   "LineToken",
	FileToken:                   "FileToken",
	StartHeredocToken:           "StartHeredocToken",
	DollarOpenCurlyBracesToken:  "DollarOpenCurlyBracesToken",
	CurlyOpenToken:              "CurlyOpenToken",
	PaamayimNekudotayimToken:    "PaamayimNekudotayimToken",
	NamespaceToken:              "NamespaceToken",
	NsCToken:                    "NsCToken",
	DirToken:                    "DirToken",
	NsSeparatorToken:            "NsSeparatorToken",
	EllipsisToken:               "EllipsisToken",
	EvalToken:                   "EvalToken",
	RequireToken:                "RequireToken",
	RequireOnceToken:            "RequireOnceToken",
	LogicalOrToken:              "LogicalOrToken",
	LogicalXorToken:             "LogicalXorToken",
	LogicalAndToken:             "LogicalAndToken",
	InstanceofToken:             "InstanceofToken",
	NewToken:                    "NewToken",
	CloneToken:                  "CloneToken",
	ElseifToken:                 "ElseifToken",
	ElseToken:                   "ElseToken",
	EndifToken:                  "EndifToken",
	PrintToken:                  "PrintToken",
	YieldToken:                  "YieldToken",
	StaticToken:                 "StaticToken",
	AbstractToken:               "AbstractToken",
	FinalToken:                  "FinalToken",
	PrivateToken:                "PrivateToken",
	ProtectedToken:              "ProtectedToken",
	PublicToken:                 "PublicToken",
	IncToken:                    "IncToken",
	DecToken:                    "DecToken",
	YieldFromToken:              "YieldFromToken",
	ObjectOperatorToken:         "ObjectOperatorToken",
	IntCastToken:                "IntCastToken",
	DoubleCastToken:             "DoubleCastToken",
	StringCastToken:             "StringCastToken",
	ArrayCastToken:              "ArrayCastToken",
	ObjectCastToken:             "ObjectCastToken",
	BoolCastToken:               "BoolCastToken",
	UnsetCastToken:              "UnsetCastToken",
	CoalesceToken:               "CoalesceToken",
	SpaceshipToken:              "SpaceshipToken",
	PlusEqualToken:              "PlusEqualToken",
	MinusEqualToken:             "MinusEqualToken",
	MulEqualToken:               "MulEqualToken",
	PowEqualToken:               "PowEqualToken",
	DivEqualToken:               "DivEqualToken",
	ConcatEqualToken:            "ConcatEqualToken",
	ModEqualToken:               "ModEqualToken",
	AndEqualToken:               "AndEqualToken",
	OrEqualToken:                "OrEqualToken",
	XorEqualToken:               "XorEqualToken",
	SlEqualToken:                "SlEqualToken",
	SrEqualToken:                "SrEqualToken",
	BooleanOrToken:              "BooleanOrToken",
	BooleanAndToken:             "BooleanAndToken",
	PowToken:                    "PowToken",
	SlToken:                     "SlToken",
	SrToken:                     "SrToken",
	IsIdenticalToken:            "IsIdenticalToken",
	IsNotIdenticalToken:         "IsNotIdenticalToken",
	IsEqualToken:                "IsEqualToken",
	IsNotEqualToken:             "IsNotEqualToken",
	IsSmallerOrEqualToken:       "IsSmallerOrEqualToken",
	IsGreaterOrEqualToken:       "IsGreaterOrEqualToken",
	HaltCompilerToken:           "HaltCompilerToken",
	IdentifierToken:             "IdentifierToken",
	CaseSeparatorToken:          "CaseSeparatorToken",
	DoubleQuoteToken:            "DoubleQuoteToken",
	BackquoteToken:              "BackquoteToken",
	OpenCurlyBracesToken:        "OpenCurlyBracesToken",
	CloseCurlyBracesToken:       "CloseCurlyBracesToken",
	SemiColonToken:              "SemiColonToken",
	ColonToken:                  "ColonToken",
	OpenParenthesisToken:        "OpenParenthesisToken",
	CloseParenthesisToken:       "CloseParenthesisToken",
	OpenSquareBracket:           "OpenSquareBracket",
	CloseSquareBracket:          "CloseSquareBracket",
	QuestionMarkToken:           "QuestionMarkToken",
	AmpersandToken:              "AmpersandToken",
	MinusToken:                  "MinusToken",
	PlusToken:                   "PlusToken",
	ExclamationMarkToken:        "ExclamationMarkToken",
	TildeToken:                  "TildeToken",
	AtToken:                     "AtToken",
	DollarToken:                 "DollarToken",
	CommaToken:                  "CommaToken",
	VerticalBarToken:            "VerticalBarToken",
	EqualToken:                  "EqualToken",
	CaretToken:                  "CaretToken",
	AsteriskToken:               "AsteriskToken",
	SlashToken:                  "SlashToken",
	PercentToken:                "PercentToken",
	LessToken:                   "LessToken",
	GreaterToken:                "GreaterToken",
	DotToken:                    "DotToken",
}
