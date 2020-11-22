package ast

import (
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
)

type Node struct {
	StartTokens []token.Token
	EndTokens   []token.Token
	Tokens      token.Collection
	Position    *position.Position
}

func (n *Node) GetNode() *Node {
	return n
}

func (n *Node) GetPosition() *position.Position {
	return n.Position
}

// Root node
type Root struct {
	Node
	Stmts []Vertex
}

func (n *Root) Accept(v NodeVisitor) {
	v.Root(n)
}

// Nullable node
type Nullable struct {
	Node
	QuestionTkn *token.Token
	Expr        Vertex
}

func (n *Nullable) Accept(v NodeVisitor) {
	v.Nullable(n)
}

// Reference node
type Reference struct {
	Node
	AmpersandTkn *token.Token
	Var          Vertex
}

func (n *Reference) Accept(v NodeVisitor) {
	v.Reference(n)
}

// Variadic node
type Variadic struct {
	Node
	VariadicTkn *token.Token
	Var         Vertex
}

func (n *Variadic) Accept(v NodeVisitor) {
	v.Variadic(n)
}

// Parameter node
type Parameter struct {
	Node
	Type         Vertex
	Var          Vertex
	EqualTkn     *token.Token
	DefaultValue Vertex
}

func (n *Parameter) Accept(v NodeVisitor) {
	v.Parameter(n)
}

// Identifier node
type Identifier struct {
	Node
	IdentifierTkn *token.Token
	Value         []byte
}

func (n *Identifier) Accept(v NodeVisitor) {
	v.Identifier(n)
}

// ArgumentList node
type ArgumentList struct {
	Node
	OpenParenthesisTkn  *token.Token
	Arguments           []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ArgumentList) Accept(v NodeVisitor) {
	v.ArgumentList(n)
}

// Argument node
type Argument struct {
	Node
	VariadicTkn  *token.Token
	AmpersandTkn *token.Token
	Expr         Vertex
}

func (n *Argument) Accept(v NodeVisitor) {
	v.Argument(n)
}

// ScalarDnumber node
type ScalarDnumber struct {
	Node
	NumberTkn *token.Token
	Value     []byte
}

func (n *ScalarDnumber) Accept(v NodeVisitor) {
	v.ScalarDnumber(n)
}

// ScalarEncapsed node
type ScalarEncapsed struct {
	Node
	OpenQoteTkn  *token.Token
	Parts        []Vertex
	CloseQoteTkn *token.Token
}

func (n *ScalarEncapsed) Accept(v NodeVisitor) {
	v.ScalarEncapsed(n)
}

// ScalarEncapsedStringPart node
type ScalarEncapsedStringPart struct {
	Node
	EncapsedStrTkn *token.Token
	Value          []byte
}

func (n *ScalarEncapsedStringPart) Accept(v NodeVisitor) {
	v.ScalarEncapsedStringPart(n)
}

// ScalarHeredoc node
type ScalarHeredoc struct {
	Node
	OpenHeredocTkn  *token.Token
	Parts           []Vertex
	CloseHeredocTkn *token.Token
}

func (n *ScalarHeredoc) Accept(v NodeVisitor) {
	v.ScalarHeredoc(n)
}

// ScalarLnumber node
type ScalarLnumber struct {
	Node
	NumberTkn *token.Token
	Value     []byte
}

func (n *ScalarLnumber) Accept(v NodeVisitor) {
	v.ScalarLnumber(n)
}

// ScalarMagicConstant node
type ScalarMagicConstant struct {
	Node
	MagicConstTkn *token.Token
	Value         []byte
}

func (n *ScalarMagicConstant) Accept(v NodeVisitor) {
	v.ScalarMagicConstant(n)
}

// ScalarString node
type ScalarString struct {
	Node
	MinusTkn  *token.Token
	StringTkn *token.Token
	Value     []byte
}

func (n *ScalarString) Accept(v NodeVisitor) {
	v.ScalarString(n)
}

// StmtBreak node
type StmtBreak struct {
	Node
	BreakTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (n *StmtBreak) Accept(v NodeVisitor) {
	v.StmtBreak(n)
}

// StmtCase node
type StmtCase struct {
	Node
	CaseTkn          *token.Token
	Cond             Vertex
	CaseSeparatorTkn *token.Token
	Stmts            []Vertex
}

func (n *StmtCase) Accept(v NodeVisitor) {
	v.StmtCase(n)
}

// StmtCatch node
type StmtCatch struct {
	Node
	CatchTkn             *token.Token
	OpenParenthesisTkn   *token.Token
	Types                []Vertex
	SeparatorTkns        []*token.Token
	Var                  Vertex
	CloseParenthesisTkn  *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtCatch) Accept(v NodeVisitor) {
	v.StmtCatch(n)
}

// StmtClass node
type StmtClass struct {
	Node
	Modifiers         []Vertex
	ClassTkn          *token.Token
	ClassName         Vertex
	ArgumentList      Vertex
	Extends           *StmtClassExtends
	Implements        *StmtClassImplements
	OpenCurlyBracket  *token.Token
	Stmts             []Vertex
	CloseCurlyBracket *token.Token
}

func (n *StmtClass) Accept(v NodeVisitor) {
	v.StmtClass(n)
}

// StmtClassConstList node
type StmtClassConstList struct {
	Node
	Modifiers    []Vertex
	ConstTkn     *token.Token
	Consts       []Vertex
	SemiColonTkn *token.Token
}

func (n *StmtClassConstList) Accept(v NodeVisitor) {
	v.StmtClassConstList(n)
}

// StmtClassExtends node
type StmtClassExtends struct {
	Node
	ClassName Vertex
}

func (n *StmtClassExtends) Accept(v NodeVisitor) {
	v.StmtClassExtends(n)
}

// StmtClassImplements node
type StmtClassImplements struct {
	Node
	InterfaceNames []Vertex
}

func (n *StmtClassImplements) Accept(v NodeVisitor) {
	v.StmtClassImplements(n)
}

// StmtClassMethod node
type StmtClassMethod struct {
	Node
	ReturnsRef bool
	MethodName Vertex
	Modifiers  []Vertex
	Params     []Vertex
	ReturnType Vertex
	Stmt       Vertex
}

func (n *StmtClassMethod) Accept(v NodeVisitor) {
	v.StmtClassMethod(n)
}

// StmtConstList node
type StmtConstList struct {
	Node
	ConstTkn      *token.Token
	Consts        []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtConstList) Accept(v NodeVisitor) {
	v.StmtConstList(n)
}

// StmtConstant node
type StmtConstant struct {
	Node
	Name     Vertex
	EqualTkn *token.Token
	Expr     Vertex
	CommaTkn *token.Token
}

func (n *StmtConstant) Accept(v NodeVisitor) {
	v.StmtConstant(n)
}

// StmtContinue node
type StmtContinue struct {
	Node
	ContinueTkn  *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (n *StmtContinue) Accept(v NodeVisitor) {
	v.StmtContinue(n)
}

// StmtDeclare node
type StmtDeclare struct {
	Node
	Alt                 bool
	DeclareTkn          *token.Token
	OpenParenthesisTkn  *token.Token
	Consts              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndDeclareTkn       *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtDeclare) Accept(v NodeVisitor) {
	v.StmtDeclare(n)
}

// StmtDefault node
type StmtDefault struct {
	Node
	DefaultTkn       *token.Token
	CaseSeparatorTkn *token.Token
	Stmts            []Vertex
}

func (n *StmtDefault) Accept(v NodeVisitor) {
	v.StmtDefault(n)
}

// StmtDo node
type StmtDo struct {
	Node
	DoTkn               *token.Token
	Stmt                Vertex
	WhileTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtDo) Accept(v NodeVisitor) {
	v.StmtDo(n)
}

// StmtEcho node
type StmtEcho struct {
	Node
	EchoTkn       *token.Token
	Exprs         []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtEcho) Accept(v NodeVisitor) {
	v.StmtEcho(n)
}

// StmtElse node
type StmtElse struct {
	Node
	Alt      bool
	ElseTkn  *token.Token
	ColonTkn *token.Token
	Stmt     Vertex
}

func (n *StmtElse) Accept(v NodeVisitor) {
	v.StmtElse(n)
}

// StmtElseIf node
type StmtElseIf struct {
	Node
	Alt                 bool
	ElseIfTkn           *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
}

func (n *StmtElseIf) Accept(v NodeVisitor) {
	v.StmtElseIf(n)
}

// StmtExpression node
type StmtExpression struct {
	Node
	Expr Vertex
}

func (n *StmtExpression) Accept(v NodeVisitor) {
	v.StmtExpression(n)
}

// StmtFinally node
type StmtFinally struct {
	Node
	FinallyTkn           *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtFinally) Accept(v NodeVisitor) {
	v.StmtFinally(n)
}

// StmtFor node
type StmtFor struct {
	Node
	Alt                 bool
	ForTkn              *token.Token
	OpenParenthesisTkn  *token.Token
	Init                []Vertex
	InitSemiColonTkn    *token.Token
	Cond                []Vertex
	CondSemiColonTkn    *token.Token
	Loop                []Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndForTkn           *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtFor) Accept(v NodeVisitor) {
	v.StmtFor(n)
}

// StmtForeach node
type StmtForeach struct {
	Node
	Alt                 bool
	ForeachTkn          *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	AsTkn               *token.Token
	Key                 Vertex
	DoubleArrowTkn      *token.Token
	Var                 Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndForeachTkn       *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtForeach) Accept(v NodeVisitor) {
	v.StmtForeach(n)
}

// StmtFunction node
type StmtFunction struct {
	Node
	ReturnsRef   bool
	FunctionName Vertex
	Params       []Vertex
	ReturnType   Vertex
	Stmts        []Vertex
}

func (n *StmtFunction) Accept(v NodeVisitor) {
	v.StmtFunction(n)
}

// StmtGlobal node
type StmtGlobal struct {
	Node
	GlobalTkn     *token.Token
	Vars          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtGlobal) Accept(v NodeVisitor) {
	v.StmtGlobal(n)
}

// StmtGoto node
type StmtGoto struct {
	Node
	GotoTkn      *token.Token
	Label        Vertex
	SemiColonTkn *token.Token
}

func (n *StmtGoto) Accept(v NodeVisitor) {
	v.StmtGoto(n)
}

// StmtHaltCompiler node
type StmtHaltCompiler struct {
	Node
	HaltCompilerTkn     *token.Token
	OpenParenthesisTkn  *token.Token
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtHaltCompiler) Accept(v NodeVisitor) {
	v.StmtHaltCompiler(n)
}

// StmtIf node
type StmtIf struct {
	Node
	Alt                 bool
	IfTkn               *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	ElseIf              []Vertex
	Else                Vertex
	EndIfTkn            *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtIf) Accept(v NodeVisitor) {
	v.StmtIf(n)
}

// StmtInlineHtml node
type StmtInlineHtml struct {
	Node
	InlineHtmlTkn *token.Token
	Value         []byte
}

func (n *StmtInlineHtml) Accept(v NodeVisitor) {
	v.StmtInlineHtml(n)
}

// StmtInterface node
type StmtInterface struct {
	Node
	InterfaceName Vertex
	Extends       *StmtInterfaceExtends
	Stmts         []Vertex
}

func (n *StmtInterface) Accept(v NodeVisitor) {
	v.StmtInterface(n)
}

// StmtInterfaceExtends node
type StmtInterfaceExtends struct {
	Node
	InterfaceNames []Vertex
}

func (n *StmtInterfaceExtends) Accept(v NodeVisitor) {
	v.StmtInterfaceExtends(n)
}

// StmtLabel node
type StmtLabel struct {
	Node
	LabelName Vertex
	ColonTkn  *token.Token
}

func (n *StmtLabel) Accept(v NodeVisitor) {
	v.StmtLabel(n)
}

// StmtNamespace node
type StmtNamespace struct {
	Node
	NsTkn             *token.Token
	Name              Vertex
	OpenCurlyBracket  *token.Token
	Stmts             []Vertex
	CloseCurlyBracket *token.Token
	SemiColonTkn      *token.Token
}

func (n *StmtNamespace) Accept(v NodeVisitor) {
	v.StmtNamespace(n)
}

// StmtNop node
type StmtNop struct {
	Node
	SemiColonTkn *token.Token
}

func (n *StmtNop) Accept(v NodeVisitor) {
	v.StmtNop(n)
}

// StmtProperty node
type StmtProperty struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *StmtProperty) Accept(v NodeVisitor) {
	v.StmtProperty(n)
}

// StmtPropertyList node
type StmtPropertyList struct {
	Node
	Modifiers  []Vertex
	Type       Vertex
	Properties []Vertex
}

func (n *StmtPropertyList) Accept(v NodeVisitor) {
	v.StmtPropertyList(n)
}

// StmtReturn node
type StmtReturn struct {
	Node
	ReturnTkn    *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (n *StmtReturn) Accept(v NodeVisitor) {
	v.StmtReturn(n)
}

// StmtStatic node
type StmtStatic struct {
	Node
	StaticTkn     *token.Token
	Vars          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtStatic) Accept(v NodeVisitor) {
	v.StmtStatic(n)
}

// StmtStaticVar node
type StmtStaticVar struct {
	Node
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *StmtStaticVar) Accept(v NodeVisitor) {
	v.StmtStaticVar(n)
}

// StmtStmtList node
type StmtStmtList struct {
	Node
	OpenCurlyBracket  *token.Token
	Stmts             []Vertex
	CloseCurlyBracket *token.Token
}

func (n *StmtStmtList) Accept(v NodeVisitor) {
	v.StmtStmtList(n)
}

// StmtSwitch node
type StmtSwitch struct {
	Node
	Alt                  bool
	SwitchTkn            *token.Token
	OpenParenthesisTkn   *token.Token
	Cond                 Vertex
	CloseParenthesisTkn  *token.Token
	ColonTkn             *token.Token
	OpenCurlyBracketTkn  *token.Token
	CaseSeparatorTkn     *token.Token
	CaseList             []Vertex
	CloseCurlyBracketTkn *token.Token
	EndSwitchTkn         *token.Token
	SemiColonTkn         *token.Token
}

func (n *StmtSwitch) Accept(v NodeVisitor) {
	v.StmtSwitch(n)
}

// StmtThrow node
type StmtThrow struct {
	Node
	ThrowTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (n *StmtThrow) Accept(v NodeVisitor) {
	v.StmtThrow(n)
}

// StmtTrait node
type StmtTrait struct {
	Node
	TraitTkn          *token.Token
	TraitName         Vertex
	Extends           *StmtClassExtends
	Implements        *StmtClassImplements
	OpenCurlyBracket  *token.Token
	Stmts             []Vertex
	CloseCurlyBracket *token.Token
}

func (n *StmtTrait) Accept(v NodeVisitor) {
	v.StmtTrait(n)
}

// StmtTraitAdaptationList node
type StmtTraitAdaptationList struct {
	Node
	Adaptations []Vertex
}

func (n *StmtTraitAdaptationList) Accept(v NodeVisitor) {
	v.StmtTraitAdaptationList(n)
}

// StmtTraitMethodRef node
type StmtTraitMethodRef struct {
	Node
	Trait  Vertex
	Method Vertex
}

func (n *StmtTraitMethodRef) Accept(v NodeVisitor) {
	v.StmtTraitMethodRef(n)
}

// StmtTraitUse node
type StmtTraitUse struct {
	Node
	Traits              []Vertex
	TraitAdaptationList Vertex
}

func (n *StmtTraitUse) Accept(v NodeVisitor) {
	v.StmtTraitUse(n)
}

// StmtTraitUseAlias node
type StmtTraitUseAlias struct {
	Node
	Ref      Vertex
	Modifier Vertex
	Alias    Vertex
}

func (n *StmtTraitUseAlias) Accept(v NodeVisitor) {
	v.StmtTraitUseAlias(n)
}

// StmtTraitUsePrecedence node
type StmtTraitUsePrecedence struct {
	Node
	Ref       Vertex
	Insteadof []Vertex
}

func (n *StmtTraitUsePrecedence) Accept(v NodeVisitor) {
	v.StmtTraitUsePrecedence(n)
}

// StmtTry node
type StmtTry struct {
	Node
	TryTkn            *token.Token
	OpenCurlyBracket  *token.Token
	Stmts             []Vertex
	CloseCurlyBracket *token.Token
	Catches           []Vertex
	Finally           Vertex
}

func (n *StmtTry) Accept(v NodeVisitor) {
	v.StmtTry(n)
}

// StmtUnset node
type StmtUnset struct {
	Node
	UnsetTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Vars                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtUnset) Accept(v NodeVisitor) {
	v.StmtUnset(n)
}

// StmtUse node
type StmtUse struct {
	Node
	UseTkn          *token.Token
	Type            Vertex
	UseDeclarations []Vertex
	SemiColonTkn    *token.Token
}

func (n *StmtUse) Accept(v NodeVisitor) {
	v.StmtUse(n)
}

// StmtGroupUse node
type StmtGroupUse struct {
	Node
	UseTkn                *token.Token
	Type                  Vertex
	LeadingNsSeparatorTkn *token.Token
	Prefix                Vertex
	NsSeparatorTkn        *token.Token
	OpenCurlyBracketTkn   *token.Token
	UseDeclarations       []Vertex
	CloseCurlyBracketTkn  *token.Token
	SemiColonTkn          *token.Token
}

func (n *StmtGroupUse) Accept(v NodeVisitor) {
	v.StmtGroupUse(n)
}

// StmtUseDeclaration node
type StmtUseDeclaration struct {
	Node
	Type           Vertex
	NsSeparatorTkn *token.Token
	Use            Vertex
	AsTkn          *token.Token
	Alias          Vertex
	CommaTkn       *token.Token
}

func (n *StmtUseDeclaration) Accept(v NodeVisitor) {
	v.StmtUseDeclaration(n)
}

// StmtWhile node
type StmtWhile struct {
	Node
	Alt                 bool
	WhileTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndWhileTkn         *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtWhile) Accept(v NodeVisitor) {
	v.StmtWhile(n)
}

// ExprArray node
type ExprArray struct {
	Node
	Items []Vertex
}

func (n *ExprArray) Accept(v NodeVisitor) {
	v.ExprArray(n)
}

// ExprArrayDimFetch node
type ExprArrayDimFetch struct {
	Node
	Var Vertex
	Dim Vertex
}

func (n *ExprArrayDimFetch) Accept(v NodeVisitor) {
	v.ExprArrayDimFetch(n)
}

// ExprArrayItem node
type ExprArrayItem struct {
	Node
	Unpack bool
	Key    Vertex
	Val    Vertex
}

func (n *ExprArrayItem) Accept(v NodeVisitor) {
	v.ExprArrayItem(n)
}

// ExprArrowFunction node
type ExprArrowFunction struct {
	Node
	ReturnsRef bool
	Static     bool
	Params     []Vertex
	ReturnType Vertex
	Expr       Vertex
}

func (n *ExprArrowFunction) Accept(v NodeVisitor) {
	v.ExprArrowFunction(n)
}

// ExprBitwiseNot node
type ExprBitwiseNot struct {
	Node
	Expr Vertex
}

func (n *ExprBitwiseNot) Accept(v NodeVisitor) {
	v.ExprBitwiseNot(n)
}

// ExprBooleanNot node
type ExprBooleanNot struct {
	Node
	Expr Vertex
}

func (n *ExprBooleanNot) Accept(v NodeVisitor) {
	v.ExprBooleanNot(n)
}

// ExprClassConstFetch node
type ExprClassConstFetch struct {
	Node
	Class        Vertex
	ConstantName Vertex
}

func (n *ExprClassConstFetch) Accept(v NodeVisitor) {
	v.ExprClassConstFetch(n)
}

// ExprClone node
type ExprClone struct {
	Node
	Expr Vertex
}

func (n *ExprClone) Accept(v NodeVisitor) {
	v.ExprClone(n)
}

// ExprClosure node
type ExprClosure struct {
	Node
	ReturnsRef bool
	Static     bool
	Params     []Vertex
	ClosureUse *ExprClosureUse
	ReturnType Vertex
	Stmts      []Vertex
}

func (n *ExprClosure) Accept(v NodeVisitor) {
	v.ExprClosure(n)
}

// ExprClosureUse node
type ExprClosureUse struct {
	Node
	Uses []Vertex
}

func (n *ExprClosureUse) Accept(v NodeVisitor) {
	v.ExprClosureUse(n)
}

// ExprConstFetch node
type ExprConstFetch struct {
	Node
	Const Vertex
}

func (n *ExprConstFetch) Accept(v NodeVisitor) {
	v.ExprConstFetch(n)
}

// ExprEmpty node
type ExprEmpty struct {
	Node
	Expr Vertex
}

func (n *ExprEmpty) Accept(v NodeVisitor) {
	v.ExprEmpty(n)
}

// ExprErrorSuppress node
type ExprErrorSuppress struct {
	Node
	Expr Vertex
}

func (n *ExprErrorSuppress) Accept(v NodeVisitor) {
	v.ExprErrorSuppress(n)
}

// ExprEval node
type ExprEval struct {
	Node
	Expr Vertex
}

func (n *ExprEval) Accept(v NodeVisitor) {
	v.ExprEval(n)
}

// ExprExit node
type ExprExit struct {
	Node
	Die  bool
	Expr Vertex
}

func (n *ExprExit) Accept(v NodeVisitor) {
	v.ExprExit(n)
}

// ExprFunctionCall node
type ExprFunctionCall struct {
	Node
	Function     Vertex
	ArgumentList *ArgumentList
}

func (n *ExprFunctionCall) Accept(v NodeVisitor) {
	v.ExprFunctionCall(n)
}

// ExprInclude node
type ExprInclude struct {
	Node
	Expr Vertex
}

func (n *ExprInclude) Accept(v NodeVisitor) {
	v.ExprInclude(n)
}

// ExprIncludeOnce node
type ExprIncludeOnce struct {
	Node
	Expr Vertex
}

func (n *ExprIncludeOnce) Accept(v NodeVisitor) {
	v.ExprIncludeOnce(n)
}

// ExprInstanceOf node
type ExprInstanceOf struct {
	Node
	Expr  Vertex
	Class Vertex
}

func (n *ExprInstanceOf) Accept(v NodeVisitor) {
	v.ExprInstanceOf(n)
}

// ExprIsset node
type ExprIsset struct {
	Node
	Vars []Vertex
}

func (n *ExprIsset) Accept(v NodeVisitor) {
	v.ExprIsset(n)
}

// ExprList node
type ExprList struct {
	Node
	Items []Vertex
}

func (n *ExprList) Accept(v NodeVisitor) {
	v.ExprList(n)
}

// ExprMethodCall node
type ExprMethodCall struct {
	Node
	Var          Vertex
	Method       Vertex
	ArgumentList *ArgumentList
}

func (n *ExprMethodCall) Accept(v NodeVisitor) {
	v.ExprMethodCall(n)
}

// ExprNew node
type ExprNew struct {
	Node
	Class        Vertex
	ArgumentList *ArgumentList
}

func (n *ExprNew) Accept(v NodeVisitor) {
	v.ExprNew(n)
}

// ExprPostDec node
type ExprPostDec struct {
	Node
	Var Vertex
}

func (n *ExprPostDec) Accept(v NodeVisitor) {
	v.ExprPostDec(n)
}

// ExprPostInc node
type ExprPostInc struct {
	Node
	Var Vertex
}

func (n *ExprPostInc) Accept(v NodeVisitor) {
	v.ExprPostInc(n)
}

// ExprPreDec node
type ExprPreDec struct {
	Node
	Var Vertex
}

func (n *ExprPreDec) Accept(v NodeVisitor) {
	v.ExprPreDec(n)
}

// ExprPreInc node
type ExprPreInc struct {
	Node
	Var Vertex
}

func (n *ExprPreInc) Accept(v NodeVisitor) {
	v.ExprPreInc(n)
}

// ExprPrint node
type ExprPrint struct {
	Node
	Expr Vertex
}

func (n *ExprPrint) Accept(v NodeVisitor) {
	v.ExprPrint(n)
}

// ExprPropertyFetch node
type ExprPropertyFetch struct {
	Node
	Var      Vertex
	Property Vertex
}

func (n *ExprPropertyFetch) Accept(v NodeVisitor) {
	v.ExprPropertyFetch(n)
}

// ExprReference node
type ExprReference struct {
	Node
	Var Vertex
}

func (n *ExprReference) Accept(v NodeVisitor) {
	v.ExprReference(n)
}

// ExprRequire node
type ExprRequire struct {
	Node
	Expr Vertex
}

func (n *ExprRequire) Accept(v NodeVisitor) {
	v.ExprRequire(n)
}

// ExprRequireOnce node
type ExprRequireOnce struct {
	Node
	Expr Vertex
}

func (n *ExprRequireOnce) Accept(v NodeVisitor) {
	v.ExprRequireOnce(n)
}

// ExprShellExec node
type ExprShellExec struct {
	Node
	Parts []Vertex
}

func (n *ExprShellExec) Accept(v NodeVisitor) {
	v.ExprShellExec(n)
}

// ExprShortArray node
type ExprShortArray struct {
	Node
	Items []Vertex
}

func (n *ExprShortArray) Accept(v NodeVisitor) {
	v.ExprShortArray(n)
}

// ExprShortList node
type ExprShortList struct {
	Node
	Items []Vertex
}

func (n *ExprShortList) Accept(v NodeVisitor) {
	v.ExprShortList(n)
}

// ExprStaticCall node
type ExprStaticCall struct {
	Node
	Class        Vertex
	Call         Vertex
	ArgumentList *ArgumentList
}

func (n *ExprStaticCall) Accept(v NodeVisitor) {
	v.ExprStaticCall(n)
}

// ExprStaticPropertyFetch node
type ExprStaticPropertyFetch struct {
	Node
	Class    Vertex
	Property Vertex
}

func (n *ExprStaticPropertyFetch) Accept(v NodeVisitor) {
	v.ExprStaticPropertyFetch(n)
}

// ExprTernary node
type ExprTernary struct {
	Node
	Condition Vertex
	IfTrue    Vertex
	IfFalse   Vertex
}

func (n *ExprTernary) Accept(v NodeVisitor) {
	v.ExprTernary(n)
}

// ExprUnaryMinus node
type ExprUnaryMinus struct {
	Node
	Expr Vertex
}

func (n *ExprUnaryMinus) Accept(v NodeVisitor) {
	v.ExprUnaryMinus(n)
}

// ExprUnaryPlus node
type ExprUnaryPlus struct {
	Node
	Expr Vertex
}

func (n *ExprUnaryPlus) Accept(v NodeVisitor) {
	v.ExprUnaryPlus(n)
}

// ExprVariable node
type ExprVariable struct {
	Node
	VarName Vertex
}

func (n *ExprVariable) Accept(v NodeVisitor) {
	v.ExprVariable(n)
}

// ExprYield node
type ExprYield struct {
	Node
	Key   Vertex
	Value Vertex
}

func (n *ExprYield) Accept(v NodeVisitor) {
	v.ExprYield(n)
}

// ExprYieldFrom node
type ExprYieldFrom struct {
	Node
	Expr Vertex
}

func (n *ExprYieldFrom) Accept(v NodeVisitor) {
	v.ExprYieldFrom(n)
}

// ExprCastArray node
type ExprCastArray struct {
	Node
	Expr Vertex
}

func (n *ExprCastArray) Accept(v NodeVisitor) {
	v.ExprCastArray(n)
}

// ExprCastBool node
type ExprCastBool struct {
	Node
	Expr Vertex
}

func (n *ExprCastBool) Accept(v NodeVisitor) {
	v.ExprCastBool(n)
}

// ExprCastDouble node
type ExprCastDouble struct {
	Node
	Expr Vertex
}

func (n *ExprCastDouble) Accept(v NodeVisitor) {
	v.ExprCastDouble(n)
}

// ExprCastInt node
type ExprCastInt struct {
	Node
	Expr Vertex
}

func (n *ExprCastInt) Accept(v NodeVisitor) {
	v.ExprCastInt(n)
}

// ExprCastObject node
type ExprCastObject struct {
	Node
	Expr Vertex
}

func (n *ExprCastObject) Accept(v NodeVisitor) {
	v.ExprCastObject(n)
}

// ExprCastString node
type ExprCastString struct {
	Node
	Expr Vertex
}

func (n *ExprCastString) Accept(v NodeVisitor) {
	v.ExprCastString(n)
}

// ExprCastUnset node
type ExprCastUnset struct {
	Node
	Expr Vertex
}

func (n *ExprCastUnset) Accept(v NodeVisitor) {
	v.ExprCastUnset(n)
}

// ExprAssign node
type ExprAssign struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssign) Accept(v NodeVisitor) {
	v.ExprAssign(n)
}

// ExprAssignReference node
type ExprAssignReference struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignReference) Accept(v NodeVisitor) {
	v.ExprAssignReference(n)
}

// ExprAssignBitwiseAnd node
type ExprAssignBitwiseAnd struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignBitwiseAnd) Accept(v NodeVisitor) {
	v.ExprAssignBitwiseAnd(n)
}

// ExprAssignBitwiseOr node
type ExprAssignBitwiseOr struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignBitwiseOr) Accept(v NodeVisitor) {
	v.ExprAssignBitwiseOr(n)
}

// ExprAssignBitwiseXor node
type ExprAssignBitwiseXor struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignBitwiseXor) Accept(v NodeVisitor) {
	v.ExprAssignBitwiseXor(n)
}

// ExprAssignCoalesce node
type ExprAssignCoalesce struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignCoalesce) Accept(v NodeVisitor) {
	v.ExprAssignCoalesce(n)
}

// ExprAssignConcat node
type ExprAssignConcat struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignConcat) Accept(v NodeVisitor) {
	v.ExprAssignConcat(n)
}

// ExprAssignDiv node
type ExprAssignDiv struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignDiv) Accept(v NodeVisitor) {
	v.ExprAssignDiv(n)
}

// ExprAssignMinus node
type ExprAssignMinus struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignMinus) Accept(v NodeVisitor) {
	v.ExprAssignMinus(n)
}

// ExprAssignMod node
type ExprAssignMod struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignMod) Accept(v NodeVisitor) {
	v.ExprAssignMod(n)
}

// ExprAssignMul node
type ExprAssignMul struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignMul) Accept(v NodeVisitor) {
	v.ExprAssignMul(n)
}

// ExprAssignPlus node
type ExprAssignPlus struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignPlus) Accept(v NodeVisitor) {
	v.ExprAssignPlus(n)
}

// ExprAssignPow node
type ExprAssignPow struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignPow) Accept(v NodeVisitor) {
	v.ExprAssignPow(n)
}

// ExprAssignShiftLeft node
type ExprAssignShiftLeft struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignShiftLeft) Accept(v NodeVisitor) {
	v.ExprAssignShiftLeft(n)
}

// ExprAssignShiftRight node
type ExprAssignShiftRight struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssignShiftRight) Accept(v NodeVisitor) {
	v.ExprAssignShiftRight(n)
}

// ExprBinaryBitwiseAnd node
type ExprBinaryBitwiseAnd struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryBitwiseAnd) Accept(v NodeVisitor) {
	v.ExprBinaryBitwiseAnd(n)
}

// ExprBinaryBitwiseOr node
type ExprBinaryBitwiseOr struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryBitwiseOr) Accept(v NodeVisitor) {
	v.ExprBinaryBitwiseOr(n)
}

// ExprBinaryBitwiseXor node
type ExprBinaryBitwiseXor struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryBitwiseXor) Accept(v NodeVisitor) {
	v.ExprBinaryBitwiseXor(n)
}

// ExprBinaryBooleanAnd node
type ExprBinaryBooleanAnd struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryBooleanAnd) Accept(v NodeVisitor) {
	v.ExprBinaryBooleanAnd(n)
}

// ExprBinaryBooleanOr node
type ExprBinaryBooleanOr struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryBooleanOr) Accept(v NodeVisitor) {
	v.ExprBinaryBooleanOr(n)
}

// ExprBinaryCoalesce node
type ExprBinaryCoalesce struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryCoalesce) Accept(v NodeVisitor) {
	v.ExprBinaryCoalesce(n)
}

// ExprBinaryConcat node
type ExprBinaryConcat struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryConcat) Accept(v NodeVisitor) {
	v.ExprBinaryConcat(n)
}

// ExprBinaryDiv node
type ExprBinaryDiv struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryDiv) Accept(v NodeVisitor) {
	v.ExprBinaryDiv(n)
}

// ExprBinaryEqual node
type ExprBinaryEqual struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryEqual) Accept(v NodeVisitor) {
	v.ExprBinaryEqual(n)
}

// ExprBinaryGreater node
type ExprBinaryGreater struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryGreater) Accept(v NodeVisitor) {
	v.ExprBinaryGreater(n)
}

// ExprBinaryGreaterOrEqual node
type ExprBinaryGreaterOrEqual struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryGreaterOrEqual) Accept(v NodeVisitor) {
	v.ExprBinaryGreaterOrEqual(n)
}

// ExprBinaryIdentical node
type ExprBinaryIdentical struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryIdentical) Accept(v NodeVisitor) {
	v.ExprBinaryIdentical(n)
}

// ExprBinaryLogicalAnd node
type ExprBinaryLogicalAnd struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryLogicalAnd) Accept(v NodeVisitor) {
	v.ExprBinaryLogicalAnd(n)
}

// ExprBinaryLogicalOr node
type ExprBinaryLogicalOr struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryLogicalOr) Accept(v NodeVisitor) {
	v.ExprBinaryLogicalOr(n)
}

// ExprBinaryLogicalXor node
type ExprBinaryLogicalXor struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryLogicalXor) Accept(v NodeVisitor) {
	v.ExprBinaryLogicalXor(n)
}

// ExprBinaryMinus node
type ExprBinaryMinus struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryMinus) Accept(v NodeVisitor) {
	v.ExprBinaryMinus(n)
}

// ExprBinaryMod node
type ExprBinaryMod struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryMod) Accept(v NodeVisitor) {
	v.ExprBinaryMod(n)
}

// ExprBinaryMul node
type ExprBinaryMul struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryMul) Accept(v NodeVisitor) {
	v.ExprBinaryMul(n)
}

// ExprBinaryNotEqual node
type ExprBinaryNotEqual struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryNotEqual) Accept(v NodeVisitor) {
	v.ExprBinaryNotEqual(n)
}

// ExprBinaryNotIdentical node
type ExprBinaryNotIdentical struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryNotIdentical) Accept(v NodeVisitor) {
	v.ExprBinaryNotIdentical(n)
}

// ExprBinaryPlus node
type ExprBinaryPlus struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryPlus) Accept(v NodeVisitor) {
	v.ExprBinaryPlus(n)
}

// ExprBinaryPow node
type ExprBinaryPow struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryPow) Accept(v NodeVisitor) {
	v.ExprBinaryPow(n)
}

// ExprBinaryShiftLeft node
type ExprBinaryShiftLeft struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryShiftLeft) Accept(v NodeVisitor) {
	v.ExprBinaryShiftLeft(n)
}

// ExprBinaryShiftRight node
type ExprBinaryShiftRight struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinaryShiftRight) Accept(v NodeVisitor) {
	v.ExprBinaryShiftRight(n)
}

// ExprBinarySmaller node
type ExprBinarySmaller struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinarySmaller) Accept(v NodeVisitor) {
	v.ExprBinarySmaller(n)
}

// ExprBinarySmallerOrEqual node
type ExprBinarySmallerOrEqual struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinarySmallerOrEqual) Accept(v NodeVisitor) {
	v.ExprBinarySmallerOrEqual(n)
}

// ExprBinarySpaceship node
type ExprBinarySpaceship struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinarySpaceship) Accept(v NodeVisitor) {
	v.ExprBinarySpaceship(n)
}

type NameName struct {
	Node
	Parts            []Vertex
	ListSeparatorTkn *token.Token
}

func (n *NameName) Accept(v NodeVisitor) {
	v.NameName(n)
}

type NameFullyQualified struct {
	Node
	NsSeparatorTkn   *token.Token
	Parts            []Vertex
	ListSeparatorTkn *token.Token
}

func (n *NameFullyQualified) Accept(v NodeVisitor) {
	v.NameFullyQualified(n)
}

type NameRelative struct {
	Node
	NsTkn            *token.Token
	NsSeparatorTkn   *token.Token
	Parts            []Vertex
	ListSeparatorTkn *token.Token
}

func (n *NameRelative) Accept(v NodeVisitor) {
	v.NameRelative(n)
}

type NameNamePart struct {
	Node
	NsSeparatorTkn *token.Token
	StringTkn      *token.Token
	Value          []byte
}

func (n *NameNamePart) Accept(v NodeVisitor) {
	v.NameNamePart(n)
}

type ParserBrackets struct {
	Node
	OpenBracketTkn  *token.Token
	Child           Vertex
	CloseBracketTkn *token.Token
}

func (n *ParserBrackets) Accept(v NodeVisitor) {
	v.ParserBrackets(n)
}

type ParserSeparatedList struct {
	Node
	Items         []Vertex
	SeparatorTkns []*token.Token
}

func (n *ParserSeparatedList) Accept(v NodeVisitor) {
	v.ParserSeparatedList(n)
}
