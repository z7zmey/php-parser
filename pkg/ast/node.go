package ast

import (
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
)

// Root node
type Root struct {
	Position *position.Position
	Stmts    []Vertex
	EndTkn   *token.Token
}

func (n *Root) Accept(v Visitor) {
	v.Root(n)
}

func (n *Root) GetPosition() *position.Position {
	return n.Position
}

// Nullable node
type Nullable struct {
	Position    *position.Position
	QuestionTkn *token.Token
	Expr        Vertex
}

func (n *Nullable) Accept(v Visitor) {
	v.Nullable(n)
}

func (n *Nullable) GetPosition() *position.Position {
	return n.Position
}

// Parameter node
type Parameter struct {
	Position     *position.Position
	Type         Vertex
	AmpersandTkn *token.Token
	VariadicTkn  *token.Token
	Var          Vertex
	EqualTkn     *token.Token
	DefaultValue Vertex
}

func (n *Parameter) Accept(v Visitor) {
	v.Parameter(n)
}

func (n *Parameter) GetPosition() *position.Position {
	return n.Position
}

// Identifier node
type Identifier struct {
	Position      *position.Position
	IdentifierTkn *token.Token
	Value         []byte
}

func (n *Identifier) Accept(v Visitor) {
	v.Identifier(n)
}

func (n *Identifier) GetPosition() *position.Position {
	return n.Position
}

// Argument node
type Argument struct {
	Position     *position.Position
	VariadicTkn  *token.Token
	AmpersandTkn *token.Token
	Expr         Vertex
}

func (n *Argument) Accept(v Visitor) {
	v.Argument(n)
}

func (n *Argument) GetPosition() *position.Position {
	return n.Position
}

// ScalarDnumber node
type ScalarDnumber struct {
	Position  *position.Position
	NumberTkn *token.Token
	Value     []byte
}

func (n *ScalarDnumber) Accept(v Visitor) {
	v.ScalarDnumber(n)
}

func (n *ScalarDnumber) GetPosition() *position.Position {
	return n.Position
}

// ScalarEncapsed node
type ScalarEncapsed struct {
	Position      *position.Position
	OpenQuoteTkn  *token.Token
	Parts         []Vertex
	CloseQuoteTkn *token.Token
}

func (n *ScalarEncapsed) Accept(v Visitor) {
	v.ScalarEncapsed(n)
}

func (n *ScalarEncapsed) GetPosition() *position.Position {
	return n.Position
}

// ScalarEncapsedStringPart node
type ScalarEncapsedStringPart struct {
	Position       *position.Position
	EncapsedStrTkn *token.Token
	Value          []byte
}

func (n *ScalarEncapsedStringPart) Accept(v Visitor) {
	v.ScalarEncapsedStringPart(n)
}

func (n *ScalarEncapsedStringPart) GetPosition() *position.Position {
	return n.Position
}

// ScalarEncapsedStringVar node
type ScalarEncapsedStringVar struct {
	Position                  *position.Position
	DollarOpenCurlyBracketTkn *token.Token
	Name                      Vertex
	OpenSquareBracketTkn      *token.Token
	Dim                       Vertex
	CloseSquareBracketTkn     *token.Token
	CloseCurlyBracketTkn      *token.Token
}

func (n *ScalarEncapsedStringVar) Accept(v Visitor) {
	v.ScalarEncapsedStringVar(n)
}

func (n *ScalarEncapsedStringVar) GetPosition() *position.Position {
	return n.Position
}

// ScalarEncapsedStringVar node
type ScalarEncapsedStringBrackets struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Var                  Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *ScalarEncapsedStringBrackets) Accept(v Visitor) {
	v.ScalarEncapsedStringBrackets(n)
}

func (n *ScalarEncapsedStringBrackets) GetPosition() *position.Position {
	return n.Position
}

// ScalarHeredoc node
type ScalarHeredoc struct {
	Position        *position.Position
	OpenHeredocTkn  *token.Token
	Parts           []Vertex
	CloseHeredocTkn *token.Token
}

func (n *ScalarHeredoc) Accept(v Visitor) {
	v.ScalarHeredoc(n)
}

func (n *ScalarHeredoc) GetPosition() *position.Position {
	return n.Position
}

// ScalarLnumber node
type ScalarLnumber struct {
	Position  *position.Position
	NumberTkn *token.Token
	Value     []byte
}

func (n *ScalarLnumber) Accept(v Visitor) {
	v.ScalarLnumber(n)
}

func (n *ScalarLnumber) GetPosition() *position.Position {
	return n.Position
}

// ScalarMagicConstant node
type ScalarMagicConstant struct {
	Position      *position.Position
	MagicConstTkn *token.Token
	Value         []byte
}

func (n *ScalarMagicConstant) Accept(v Visitor) {
	v.ScalarMagicConstant(n)
}

func (n *ScalarMagicConstant) GetPosition() *position.Position {
	return n.Position
}

// ScalarString node
type ScalarString struct {
	Position  *position.Position
	MinusTkn  *token.Token
	StringTkn *token.Token
	Value     []byte
}

func (n *ScalarString) Accept(v Visitor) {
	v.ScalarString(n)
}

func (n *ScalarString) GetPosition() *position.Position {
	return n.Position
}

// StmtBreak node
type StmtBreak struct {
	Position     *position.Position
	BreakTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (n *StmtBreak) Accept(v Visitor) {
	v.StmtBreak(n)
}

func (n *StmtBreak) GetPosition() *position.Position {
	return n.Position
}

// StmtCase node
type StmtCase struct {
	Position         *position.Position
	CaseTkn          *token.Token
	Cond             Vertex
	CaseSeparatorTkn *token.Token
	Stmts            []Vertex
}

func (n *StmtCase) Accept(v Visitor) {
	v.StmtCase(n)
}

func (n *StmtCase) GetPosition() *position.Position {
	return n.Position
}

// StmtCatch node
type StmtCatch struct {
	Position             *position.Position
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

func (n *StmtCatch) Accept(v Visitor) {
	v.StmtCatch(n)
}

func (n *StmtCatch) GetPosition() *position.Position {
	return n.Position
}

// StmtClass node
type StmtClass struct {
	Position                *position.Position
	Modifiers               []Vertex
	ClassTkn                *token.Token
	Name                    Vertex
	OpenParenthesisTkn      *token.Token
	Args                    []Vertex
	SeparatorTkns           []*token.Token
	CloseParenthesisTkn     *token.Token
	ExtendsTkn              *token.Token
	Extends                 Vertex
	ImplementsTkn           *token.Token
	Implements              []Vertex
	ImplementsSeparatorTkns []*token.Token
	OpenCurlyBracketTkn     *token.Token
	Stmts                   []Vertex
	CloseCurlyBracketTkn    *token.Token
}

func (n *StmtClass) Accept(v Visitor) {
	v.StmtClass(n)
}

func (n *StmtClass) GetPosition() *position.Position {
	return n.Position
}

// StmtClassConstList node
type StmtClassConstList struct {
	Position      *position.Position
	Modifiers     []Vertex
	ConstTkn      *token.Token
	Consts        []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtClassConstList) Accept(v Visitor) {
	v.StmtClassConstList(n)
}

func (n *StmtClassConstList) GetPosition() *position.Position {
	return n.Position
}

// StmtClassMethod node
type StmtClassMethod struct {
	Position            *position.Position
	Modifiers           []Vertex
	FunctionTkn         *token.Token
	AmpersandTkn        *token.Token
	Name                Vertex
	OpenParenthesisTkn  *token.Token
	Params              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	ReturnType          Vertex
	Stmt                Vertex
}

func (n *StmtClassMethod) Accept(v Visitor) {
	v.StmtClassMethod(n)
}

func (n *StmtClassMethod) GetPosition() *position.Position {
	return n.Position
}

// StmtConstList node
type StmtConstList struct {
	Position      *position.Position
	ConstTkn      *token.Token
	Consts        []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtConstList) Accept(v Visitor) {
	v.StmtConstList(n)
}

func (n *StmtConstList) GetPosition() *position.Position {
	return n.Position
}

// StmtConstant node
type StmtConstant struct {
	Position *position.Position
	Name     Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *StmtConstant) Accept(v Visitor) {
	v.StmtConstant(n)
}

func (n *StmtConstant) GetPosition() *position.Position {
	return n.Position
}

// StmtContinue node
type StmtContinue struct {
	Position     *position.Position
	ContinueTkn  *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (n *StmtContinue) Accept(v Visitor) {
	v.StmtContinue(n)
}

func (n *StmtContinue) GetPosition() *position.Position {
	return n.Position
}

// StmtDeclare node
type StmtDeclare struct {
	Position            *position.Position
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

func (n *StmtDeclare) Accept(v Visitor) {
	v.StmtDeclare(n)
}

func (n *StmtDeclare) GetPosition() *position.Position {
	return n.Position
}

// StmtDefault node
type StmtDefault struct {
	Position         *position.Position
	DefaultTkn       *token.Token
	CaseSeparatorTkn *token.Token
	Stmts            []Vertex
}

func (n *StmtDefault) Accept(v Visitor) {
	v.StmtDefault(n)
}

func (n *StmtDefault) GetPosition() *position.Position {
	return n.Position
}

// StmtDo node
type StmtDo struct {
	Position            *position.Position
	DoTkn               *token.Token
	Stmt                Vertex
	WhileTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtDo) Accept(v Visitor) {
	v.StmtDo(n)
}

func (n *StmtDo) GetPosition() *position.Position {
	return n.Position
}

// StmtEcho node
type StmtEcho struct {
	Position      *position.Position
	EchoTkn       *token.Token
	Exprs         []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtEcho) Accept(v Visitor) {
	v.StmtEcho(n)
}

func (n *StmtEcho) GetPosition() *position.Position {
	return n.Position
}

// StmtElse node
type StmtElse struct {
	Position *position.Position
	ElseTkn  *token.Token
	ColonTkn *token.Token
	Stmt     Vertex
}

func (n *StmtElse) Accept(v Visitor) {
	v.StmtElse(n)
}

func (n *StmtElse) GetPosition() *position.Position {
	return n.Position
}

// StmtElseIf node
type StmtElseIf struct {
	Position            *position.Position
	ElseIfTkn           *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
}

func (n *StmtElseIf) Accept(v Visitor) {
	v.StmtElseIf(n)
}

func (n *StmtElseIf) GetPosition() *position.Position {
	return n.Position
}

// StmtExpression node
type StmtExpression struct {
	Position     *position.Position
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (n *StmtExpression) Accept(v Visitor) {
	v.StmtExpression(n)
}

func (n *StmtExpression) GetPosition() *position.Position {
	return n.Position
}

// StmtFinally node
type StmtFinally struct {
	Position             *position.Position
	FinallyTkn           *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtFinally) Accept(v Visitor) {
	v.StmtFinally(n)
}

func (n *StmtFinally) GetPosition() *position.Position {
	return n.Position
}

// StmtFor node
type StmtFor struct {
	Position            *position.Position
	ForTkn              *token.Token
	OpenParenthesisTkn  *token.Token
	Init                []Vertex
	InitSeparatorTkns   []*token.Token
	InitSemiColonTkn    *token.Token
	Cond                []Vertex
	CondSeparatorTkns   []*token.Token
	CondSemiColonTkn    *token.Token
	Loop                []Vertex
	LoopSeparatorTkns   []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndForTkn           *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtFor) Accept(v Visitor) {
	v.StmtFor(n)
}

func (n *StmtFor) GetPosition() *position.Position {
	return n.Position
}

// StmtForeach node
type StmtForeach struct {
	Position            *position.Position
	ForeachTkn          *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	AsTkn               *token.Token
	Key                 Vertex
	DoubleArrowTkn      *token.Token
	AmpersandTkn        *token.Token
	Var                 Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndForeachTkn       *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtForeach) Accept(v Visitor) {
	v.StmtForeach(n)
}

func (n *StmtForeach) GetPosition() *position.Position {
	return n.Position
}

// StmtFunction node
type StmtFunction struct {
	Position             *position.Position
	FunctionTkn          *token.Token
	AmpersandTkn         *token.Token
	Name                 Vertex
	OpenParenthesisTkn   *token.Token
	Params               []Vertex
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
	ColonTkn             *token.Token
	ReturnType           Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtFunction) Accept(v Visitor) {
	v.StmtFunction(n)
}

func (n *StmtFunction) GetPosition() *position.Position {
	return n.Position
}

// StmtGlobal node
type StmtGlobal struct {
	Position      *position.Position
	GlobalTkn     *token.Token
	Vars          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtGlobal) Accept(v Visitor) {
	v.StmtGlobal(n)
}

func (n *StmtGlobal) GetPosition() *position.Position {
	return n.Position
}

// StmtGoto node
type StmtGoto struct {
	Position     *position.Position
	GotoTkn      *token.Token
	Label        Vertex
	SemiColonTkn *token.Token
}

func (n *StmtGoto) Accept(v Visitor) {
	v.StmtGoto(n)
}

func (n *StmtGoto) GetPosition() *position.Position {
	return n.Position
}

// StmtHaltCompiler node
type StmtHaltCompiler struct {
	Position            *position.Position
	HaltCompilerTkn     *token.Token
	OpenParenthesisTkn  *token.Token
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtHaltCompiler) Accept(v Visitor) {
	v.StmtHaltCompiler(n)
}

func (n *StmtHaltCompiler) GetPosition() *position.Position {
	return n.Position
}

// StmtIf node
type StmtIf struct {
	Position            *position.Position
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

func (n *StmtIf) Accept(v Visitor) {
	v.StmtIf(n)
}

func (n *StmtIf) GetPosition() *position.Position {
	return n.Position
}

// StmtInlineHtml node
type StmtInlineHtml struct {
	Position      *position.Position
	InlineHtmlTkn *token.Token
	Value         []byte
}

func (n *StmtInlineHtml) Accept(v Visitor) {
	v.StmtInlineHtml(n)
}

func (n *StmtInlineHtml) GetPosition() *position.Position {
	return n.Position
}

// StmtInterface node
type StmtInterface struct {
	Position             *position.Position
	InterfaceTkn         *token.Token
	Name                 Vertex
	ExtendsTkn           *token.Token
	Extends              []Vertex
	ExtendsSeparatorTkns []*token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtInterface) Accept(v Visitor) {
	v.StmtInterface(n)
}

func (n *StmtInterface) GetPosition() *position.Position {
	return n.Position
}

// StmtLabel node
type StmtLabel struct {
	Position *position.Position
	Name     Vertex
	ColonTkn *token.Token
}

func (n *StmtLabel) Accept(v Visitor) {
	v.StmtLabel(n)
}

func (n *StmtLabel) GetPosition() *position.Position {
	return n.Position
}

// StmtNamespace node
type StmtNamespace struct {
	Position             *position.Position
	NsTkn                *token.Token
	Name                 Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
	SemiColonTkn         *token.Token
}

func (n *StmtNamespace) Accept(v Visitor) {
	v.StmtNamespace(n)
}

func (n *StmtNamespace) GetPosition() *position.Position {
	return n.Position
}

// StmtNop node
type StmtNop struct {
	Position     *position.Position
	SemiColonTkn *token.Token
}

func (n *StmtNop) Accept(v Visitor) {
	v.StmtNop(n)
}

func (n *StmtNop) GetPosition() *position.Position {
	return n.Position
}

// StmtProperty node
type StmtProperty struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *StmtProperty) Accept(v Visitor) {
	v.StmtProperty(n)
}

func (n *StmtProperty) GetPosition() *position.Position {
	return n.Position
}

// StmtPropertyList node
type StmtPropertyList struct {
	Position      *position.Position
	Modifiers     []Vertex
	Type          Vertex
	Props         []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtPropertyList) Accept(v Visitor) {
	v.StmtPropertyList(n)
}

func (n *StmtPropertyList) GetPosition() *position.Position {
	return n.Position
}

// StmtReturn node
type StmtReturn struct {
	Position     *position.Position
	ReturnTkn    *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (n *StmtReturn) Accept(v Visitor) {
	v.StmtReturn(n)
}

func (n *StmtReturn) GetPosition() *position.Position {
	return n.Position
}

// StmtStatic node
type StmtStatic struct {
	Position      *position.Position
	StaticTkn     *token.Token
	Vars          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtStatic) Accept(v Visitor) {
	v.StmtStatic(n)
}

func (n *StmtStatic) GetPosition() *position.Position {
	return n.Position
}

// StmtStaticVar node
type StmtStaticVar struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *StmtStaticVar) Accept(v Visitor) {
	v.StmtStaticVar(n)
}

func (n *StmtStaticVar) GetPosition() *position.Position {
	return n.Position
}

// StmtStmtList node
type StmtStmtList struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtStmtList) Accept(v Visitor) {
	v.StmtStmtList(n)
}

func (n *StmtStmtList) GetPosition() *position.Position {
	return n.Position
}

// StmtSwitch node
type StmtSwitch struct {
	Position             *position.Position
	SwitchTkn            *token.Token
	OpenParenthesisTkn   *token.Token
	Cond                 Vertex
	CloseParenthesisTkn  *token.Token
	ColonTkn             *token.Token
	OpenCurlyBracketTkn  *token.Token
	CaseSeparatorTkn     *token.Token
	Cases                []Vertex
	CloseCurlyBracketTkn *token.Token
	EndSwitchTkn         *token.Token
	SemiColonTkn         *token.Token
}

func (n *StmtSwitch) Accept(v Visitor) {
	v.StmtSwitch(n)
}

func (n *StmtSwitch) GetPosition() *position.Position {
	return n.Position
}

// StmtThrow node
type StmtThrow struct {
	Position     *position.Position
	ThrowTkn     *token.Token
	Expr         Vertex
	SemiColonTkn *token.Token
}

func (n *StmtThrow) Accept(v Visitor) {
	v.StmtThrow(n)
}

func (n *StmtThrow) GetPosition() *position.Position {
	return n.Position
}

// StmtTrait node
type StmtTrait struct {
	Position             *position.Position
	TraitTkn             *token.Token
	Name                 Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtTrait) Accept(v Visitor) {
	v.StmtTrait(n)
}

func (n *StmtTrait) GetPosition() *position.Position {
	return n.Position
}

// StmtTraitUse node
type StmtTraitUse struct {
	Position             *position.Position
	UseTkn               *token.Token
	Traits               []Vertex
	SeparatorTkns        []*token.Token
	OpenCurlyBracketTkn  *token.Token
	Adaptations          []Vertex
	CloseCurlyBracketTkn *token.Token
	SemiColonTkn         *token.Token
}

func (n *StmtTraitUse) Accept(v Visitor) {
	v.StmtTraitUse(n)
}

func (n *StmtTraitUse) GetPosition() *position.Position {
	return n.Position
}

// StmtTraitUseAlias node
type StmtTraitUseAlias struct {
	Position       *position.Position
	Trait          Vertex
	DoubleColonTkn *token.Token
	Method         Vertex
	AsTkn          *token.Token
	Modifier       Vertex
	Alias          Vertex
	SemiColonTkn   *token.Token
}

func (n *StmtTraitUseAlias) Accept(v Visitor) {
	v.StmtTraitUseAlias(n)
}

func (n *StmtTraitUseAlias) GetPosition() *position.Position {
	return n.Position
}

// StmtTraitUsePrecedence node
type StmtTraitUsePrecedence struct {
	Position       *position.Position
	Trait          Vertex
	DoubleColonTkn *token.Token
	Method         Vertex
	InsteadofTkn   *token.Token
	Insteadof      []Vertex
	SeparatorTkns  []*token.Token
	SemiColonTkn   *token.Token
}

func (n *StmtTraitUsePrecedence) Accept(v Visitor) {
	v.StmtTraitUsePrecedence(n)
}

func (n *StmtTraitUsePrecedence) GetPosition() *position.Position {
	return n.Position
}

// StmtTry node
type StmtTry struct {
	Position             *position.Position
	TryTkn               *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
	Catches              []Vertex
	Finally              Vertex
}

func (n *StmtTry) Accept(v Visitor) {
	v.StmtTry(n)
}

func (n *StmtTry) GetPosition() *position.Position {
	return n.Position
}

// StmtUnset node
type StmtUnset struct {
	Position            *position.Position
	UnsetTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Vars                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtUnset) Accept(v Visitor) {
	v.StmtUnset(n)
}

func (n *StmtUnset) GetPosition() *position.Position {
	return n.Position
}

// StmtUseList node
type StmtUseList struct {
	Position      *position.Position
	UseTkn        *token.Token
	Type          Vertex
	Uses          []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtUseList) Accept(v Visitor) {
	v.StmtUse(n)
}

func (n *StmtUseList) GetPosition() *position.Position {
	return n.Position
}

// StmtGroupUseList node
type StmtGroupUseList struct {
	Position              *position.Position
	UseTkn                *token.Token
	Type                  Vertex
	LeadingNsSeparatorTkn *token.Token
	Prefix                Vertex
	NsSeparatorTkn        *token.Token
	OpenCurlyBracketTkn   *token.Token
	Uses                  []Vertex
	SeparatorTkns         []*token.Token
	CloseCurlyBracketTkn  *token.Token
	SemiColonTkn          *token.Token
}

func (n *StmtGroupUseList) Accept(v Visitor) {
	v.StmtGroupUse(n)
}

func (n *StmtGroupUseList) GetPosition() *position.Position {
	return n.Position
}

// StmtUse node
type StmtUse struct {
	Position       *position.Position
	Type           Vertex
	NsSeparatorTkn *token.Token
	Use            Vertex
	AsTkn          *token.Token
	Alias          Vertex
}

func (n *StmtUse) Accept(v Visitor) {
	v.StmtUseDeclaration(n)
}

func (n *StmtUse) GetPosition() *position.Position {
	return n.Position
}

// StmtWhile node
type StmtWhile struct {
	Position            *position.Position
	WhileTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Vertex
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Vertex
	EndWhileTkn         *token.Token
	SemiColonTkn        *token.Token
}

func (n *StmtWhile) Accept(v Visitor) {
	v.StmtWhile(n)
}

func (n *StmtWhile) GetPosition() *position.Position {
	return n.Position
}

// ExprArray node
type ExprArray struct {
	Position        *position.Position
	ArrayTkn        *token.Token
	OpenBracketTkn  *token.Token
	Items           []Vertex
	SeparatorTkns   []*token.Token
	CloseBracketTkn *token.Token
}

func (n *ExprArray) Accept(v Visitor) {
	v.ExprArray(n)
}

func (n *ExprArray) GetPosition() *position.Position {
	return n.Position
}

// ExprArrayDimFetch node
type ExprArrayDimFetch struct {
	Position        *position.Position
	Var             Vertex
	OpenBracketTkn  *token.Token
	Dim             Vertex
	CloseBracketTkn *token.Token
}

func (n *ExprArrayDimFetch) Accept(v Visitor) {
	v.ExprArrayDimFetch(n)
}

func (n *ExprArrayDimFetch) GetPosition() *position.Position {
	return n.Position
}

// ExprArrayItem node
type ExprArrayItem struct {
	Position       *position.Position
	EllipsisTkn    *token.Token
	Key            Vertex
	DoubleArrowTkn *token.Token
	AmpersandTkn   *token.Token
	Val            Vertex
}

func (n *ExprArrayItem) Accept(v Visitor) {
	v.ExprArrayItem(n)
}

func (n *ExprArrayItem) GetPosition() *position.Position {
	return n.Position
}

// ExprArrowFunction node
type ExprArrowFunction struct {
	Position            *position.Position
	StaticTkn           *token.Token
	FnTkn               *token.Token
	AmpersandTkn        *token.Token
	OpenParenthesisTkn  *token.Token
	Params              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	ReturnType          Vertex
	DoubleArrowTkn      *token.Token
	Expr                Vertex
}

func (n *ExprArrowFunction) Accept(v Visitor) {
	v.ExprArrowFunction(n)
}

func (n *ExprArrowFunction) GetPosition() *position.Position {
	return n.Position
}

// ExprBitwiseNot node
type ExprBitwiseNot struct {
	Position *position.Position
	TildaTkn *token.Token
	Expr     Vertex
}

func (n *ExprBitwiseNot) Accept(v Visitor) {
	v.ExprBitwiseNot(n)
}

func (n *ExprBitwiseNot) GetPosition() *position.Position {
	return n.Position
}

// ExprBooleanNot node
type ExprBooleanNot struct {
	Position       *position.Position
	ExclamationTkn *token.Token
	Expr           Vertex
}

func (n *ExprBooleanNot) Accept(v Visitor) {
	v.ExprBooleanNot(n)
}

func (n *ExprBooleanNot) GetPosition() *position.Position {
	return n.Position
}

type ExprBrackets struct {
	Position            *position.Position
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

func (n *ExprBrackets) Accept(v Visitor) {
	v.ExprBrackets(n)
}

func (n *ExprBrackets) GetPosition() *position.Position {
	return n.Position
}

// ExprClassConstFetch node
type ExprClassConstFetch struct {
	Position       *position.Position
	Class          Vertex
	DoubleColonTkn *token.Token
	Const          Vertex
}

func (n *ExprClassConstFetch) Accept(v Visitor) {
	v.ExprClassConstFetch(n)
}

func (n *ExprClassConstFetch) GetPosition() *position.Position {
	return n.Position
}

// ExprClone node
type ExprClone struct {
	Position *position.Position
	CloneTkn *token.Token
	Expr     Vertex
}

func (n *ExprClone) Accept(v Visitor) {
	v.ExprClone(n)
}

func (n *ExprClone) GetPosition() *position.Position {
	return n.Position
}

// ExprClosure node
type ExprClosure struct {
	Position               *position.Position
	StaticTkn              *token.Token
	FunctionTkn            *token.Token
	AmpersandTkn           *token.Token
	OpenParenthesisTkn     *token.Token
	Params                 []Vertex
	SeparatorTkns          []*token.Token
	CloseParenthesisTkn    *token.Token
	UseTkn                 *token.Token
	UseOpenParenthesisTkn  *token.Token
	Uses                   []Vertex
	UseSeparatorTkns       []*token.Token
	UseCloseParenthesisTkn *token.Token
	ColonTkn               *token.Token
	ReturnType             Vertex
	OpenCurlyBracketTkn    *token.Token
	Stmts                  []Vertex
	CloseCurlyBracketTkn   *token.Token
}

func (n *ExprClosure) Accept(v Visitor) {
	v.ExprClosure(n)
}

func (n *ExprClosure) GetPosition() *position.Position {
	return n.Position
}

// ExprClosureUse node
type ExprClosureUse struct {
	Position     *position.Position
	AmpersandTkn *token.Token
	Var          Vertex
}

func (n *ExprClosureUse) Accept(v Visitor) {
	v.ExprClosureUse(n)
}

func (n *ExprClosureUse) GetPosition() *position.Position {
	return n.Position
}

// ExprConstFetch node
type ExprConstFetch struct {
	Position *position.Position
	Const    Vertex
}

func (n *ExprConstFetch) Accept(v Visitor) {
	v.ExprConstFetch(n)
}

func (n *ExprConstFetch) GetPosition() *position.Position {
	return n.Position
}

// ExprEmpty node
type ExprEmpty struct {
	Position            *position.Position
	EmptyTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

func (n *ExprEmpty) Accept(v Visitor) {
	v.ExprEmpty(n)
}

func (n *ExprEmpty) GetPosition() *position.Position {
	return n.Position
}

// ExprErrorSuppress node
type ExprErrorSuppress struct {
	Position *position.Position
	AtTkn    *token.Token
	Expr     Vertex
}

func (n *ExprErrorSuppress) Accept(v Visitor) {
	v.ExprErrorSuppress(n)
}

func (n *ExprErrorSuppress) GetPosition() *position.Position {
	return n.Position
}

// ExprEval node
type ExprEval struct {
	Position            *position.Position
	EvalTkn             *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

func (n *ExprEval) Accept(v Visitor) {
	v.ExprEval(n)
}

func (n *ExprEval) GetPosition() *position.Position {
	return n.Position
}

// ExprExit node
type ExprExit struct {
	Position            *position.Position
	ExitTkn             *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Vertex
	CloseParenthesisTkn *token.Token
}

func (n *ExprExit) Accept(v Visitor) {
	v.ExprExit(n)
}

func (n *ExprExit) GetPosition() *position.Position {
	return n.Position
}

// ExprFunctionCall node
type ExprFunctionCall struct {
	Position            *position.Position
	Function            Vertex
	OpenParenthesisTkn  *token.Token
	Args                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ExprFunctionCall) Accept(v Visitor) {
	v.ExprFunctionCall(n)
}

func (n *ExprFunctionCall) GetPosition() *position.Position {
	return n.Position
}

// ExprInclude node
type ExprInclude struct {
	Position   *position.Position
	IncludeTkn *token.Token
	Expr       Vertex
}

func (n *ExprInclude) Accept(v Visitor) {
	v.ExprInclude(n)
}

func (n *ExprInclude) GetPosition() *position.Position {
	return n.Position
}

// ExprIncludeOnce node
type ExprIncludeOnce struct {
	Position       *position.Position
	IncludeOnceTkn *token.Token
	Expr           Vertex
}

func (n *ExprIncludeOnce) Accept(v Visitor) {
	v.ExprIncludeOnce(n)
}

func (n *ExprIncludeOnce) GetPosition() *position.Position {
	return n.Position
}

// ExprInstanceOf node
type ExprInstanceOf struct {
	Position      *position.Position
	Expr          Vertex
	InstanceOfTkn *token.Token
	Class         Vertex
}

func (n *ExprInstanceOf) Accept(v Visitor) {
	v.ExprInstanceOf(n)
}

func (n *ExprInstanceOf) GetPosition() *position.Position {
	return n.Position
}

// ExprIsset node
type ExprIsset struct {
	Position            *position.Position
	IssetTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Vars                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ExprIsset) Accept(v Visitor) {
	v.ExprIsset(n)
}

func (n *ExprIsset) GetPosition() *position.Position {
	return n.Position
}

// ExprList node
type ExprList struct {
	Position        *position.Position
	ListTkn         *token.Token
	OpenBracketTkn  *token.Token
	Items           []Vertex
	SeparatorTkns   []*token.Token
	CloseBracketTkn *token.Token
}

func (n *ExprList) Accept(v Visitor) {
	v.ExprList(n)
}

func (n *ExprList) GetPosition() *position.Position {
	return n.Position
}

// ExprMethodCall node
type ExprMethodCall struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Method               Vertex
	CloseCurlyBracketTkn *token.Token
	OpenParenthesisTkn   *token.Token
	Args                 []Vertex
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
}

func (n *ExprMethodCall) Accept(v Visitor) {
	v.ExprMethodCall(n)
}

func (n *ExprMethodCall) GetPosition() *position.Position {
	return n.Position
}

// ExprNew node
type ExprNew struct {
	Position            *position.Position
	NewTkn              *token.Token
	Class               Vertex
	OpenParenthesisTkn  *token.Token
	Args                []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ExprNew) Accept(v Visitor) {
	v.ExprNew(n)
}

func (n *ExprNew) GetPosition() *position.Position {
	return n.Position
}

// ExprPostDec node
type ExprPostDec struct {
	Position *position.Position
	Var      Vertex
	DecTkn   *token.Token
}

func (n *ExprPostDec) Accept(v Visitor) {
	v.ExprPostDec(n)
}

func (n *ExprPostDec) GetPosition() *position.Position {
	return n.Position
}

// ExprPostInc node
type ExprPostInc struct {
	Position *position.Position
	Var      Vertex
	IncTkn   *token.Token
}

func (n *ExprPostInc) Accept(v Visitor) {
	v.ExprPostInc(n)
}

func (n *ExprPostInc) GetPosition() *position.Position {
	return n.Position
}

// ExprPreDec node
type ExprPreDec struct {
	Position *position.Position
	DecTkn   *token.Token
	Var      Vertex
}

func (n *ExprPreDec) Accept(v Visitor) {
	v.ExprPreDec(n)
}

func (n *ExprPreDec) GetPosition() *position.Position {
	return n.Position
}

// ExprPreInc node
type ExprPreInc struct {
	Position *position.Position
	IncTkn   *token.Token
	Var      Vertex
}

func (n *ExprPreInc) Accept(v Visitor) {
	v.ExprPreInc(n)
}

func (n *ExprPreInc) GetPosition() *position.Position {
	return n.Position
}

// ExprPrint node
type ExprPrint struct {
	Position *position.Position
	PrintTkn *token.Token
	Expr     Vertex
}

func (n *ExprPrint) Accept(v Visitor) {
	v.ExprPrint(n)
}

func (n *ExprPrint) GetPosition() *position.Position {
	return n.Position
}

// ExprPropertyFetch node
type ExprPropertyFetch struct {
	Position             *position.Position
	Var                  Vertex
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Prop                 Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *ExprPropertyFetch) Accept(v Visitor) {
	v.ExprPropertyFetch(n)
}

func (n *ExprPropertyFetch) GetPosition() *position.Position {
	return n.Position
}

// ExprRequire node
type ExprRequire struct {
	Position   *position.Position
	RequireTkn *token.Token
	Expr       Vertex
}

func (n *ExprRequire) Accept(v Visitor) {
	v.ExprRequire(n)
}

func (n *ExprRequire) GetPosition() *position.Position {
	return n.Position
}

// ExprRequireOnce node
type ExprRequireOnce struct {
	Position       *position.Position
	RequireOnceTkn *token.Token
	Expr           Vertex
}

func (n *ExprRequireOnce) Accept(v Visitor) {
	v.ExprRequireOnce(n)
}

func (n *ExprRequireOnce) GetPosition() *position.Position {
	return n.Position
}

// ExprShellExec node
type ExprShellExec struct {
	Position         *position.Position
	OpenBacktickTkn  *token.Token
	Parts            []Vertex
	CloseBacktickTkn *token.Token
}

func (n *ExprShellExec) Accept(v Visitor) {
	v.ExprShellExec(n)
}

func (n *ExprShellExec) GetPosition() *position.Position {
	return n.Position
}

// ExprStaticCall node
type ExprStaticCall struct {
	Position             *position.Position
	Class                Vertex
	DoubleColonTkn       *token.Token
	OpenCurlyBracketTkn  *token.Token
	Call                 Vertex
	CloseCurlyBracketTkn *token.Token
	OpenParenthesisTkn   *token.Token
	Args                 []Vertex
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
}

func (n *ExprStaticCall) Accept(v Visitor) {
	v.ExprStaticCall(n)
}

func (n *ExprStaticCall) GetPosition() *position.Position {
	return n.Position
}

// ExprStaticPropertyFetch node
type ExprStaticPropertyFetch struct {
	Position       *position.Position
	Class          Vertex
	DoubleColonTkn *token.Token
	Prop           Vertex
}

func (n *ExprStaticPropertyFetch) Accept(v Visitor) {
	v.ExprStaticPropertyFetch(n)
}

func (n *ExprStaticPropertyFetch) GetPosition() *position.Position {
	return n.Position
}

// ExprTernary node
type ExprTernary struct {
	Position    *position.Position
	Cond        Vertex
	QuestionTkn *token.Token
	IfTrue      Vertex
	ColonTkn    *token.Token
	IfFalse     Vertex
}

func (n *ExprTernary) Accept(v Visitor) {
	v.ExprTernary(n)
}

func (n *ExprTernary) GetPosition() *position.Position {
	return n.Position
}

// ExprUnaryMinus node
type ExprUnaryMinus struct {
	Position *position.Position
	MinusTkn *token.Token
	Expr     Vertex
}

func (n *ExprUnaryMinus) Accept(v Visitor) {
	v.ExprUnaryMinus(n)
}

func (n *ExprUnaryMinus) GetPosition() *position.Position {
	return n.Position
}

// ExprUnaryPlus node
type ExprUnaryPlus struct {
	Position *position.Position
	PlusTkn  *token.Token
	Expr     Vertex
}

func (n *ExprUnaryPlus) Accept(v Visitor) {
	v.ExprUnaryPlus(n)
}

func (n *ExprUnaryPlus) GetPosition() *position.Position {
	return n.Position
}

// ExprVariable node
type ExprVariable struct {
	Position             *position.Position
	DollarTkn            *token.Token
	OpenCurlyBracketTkn  *token.Token
	Name                 Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *ExprVariable) Accept(v Visitor) {
	v.ExprVariable(n)
}

func (n *ExprVariable) GetPosition() *position.Position {
	return n.Position
}

// ExprYield node
type ExprYield struct {
	Position       *position.Position
	YieldTkn       *token.Token
	Key            Vertex
	DoubleArrowTkn *token.Token
	Val            Vertex
}

func (n *ExprYield) Accept(v Visitor) {
	v.ExprYield(n)
}

func (n *ExprYield) GetPosition() *position.Position {
	return n.Position
}

// ExprYieldFrom node
type ExprYieldFrom struct {
	Position     *position.Position
	YieldFromTkn *token.Token
	Expr         Vertex
}

func (n *ExprYieldFrom) Accept(v Visitor) {
	v.ExprYieldFrom(n)
}

func (n *ExprYieldFrom) GetPosition() *position.Position {
	return n.Position
}

// ExprCastArray node
type ExprCastArray struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (n *ExprCastArray) Accept(v Visitor) {
	v.ExprCastArray(n)
}

func (n *ExprCastArray) GetPosition() *position.Position {
	return n.Position
}

// ExprCastBool node
type ExprCastBool struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (n *ExprCastBool) Accept(v Visitor) {
	v.ExprCastBool(n)
}

func (n *ExprCastBool) GetPosition() *position.Position {
	return n.Position
}

// ExprCastDouble node
type ExprCastDouble struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (n *ExprCastDouble) Accept(v Visitor) {
	v.ExprCastDouble(n)
}

func (n *ExprCastDouble) GetPosition() *position.Position {
	return n.Position
}

// ExprCastInt node
type ExprCastInt struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (n *ExprCastInt) Accept(v Visitor) {
	v.ExprCastInt(n)
}

func (n *ExprCastInt) GetPosition() *position.Position {
	return n.Position
}

// ExprCastObject node
type ExprCastObject struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (n *ExprCastObject) Accept(v Visitor) {
	v.ExprCastObject(n)
}

func (n *ExprCastObject) GetPosition() *position.Position {
	return n.Position
}

// ExprCastString node
type ExprCastString struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (n *ExprCastString) Accept(v Visitor) {
	v.ExprCastString(n)
}

func (n *ExprCastString) GetPosition() *position.Position {
	return n.Position
}

// ExprCastUnset node
type ExprCastUnset struct {
	Position *position.Position
	CastTkn  *token.Token
	Expr     Vertex
}

func (n *ExprCastUnset) Accept(v Visitor) {
	v.ExprCastUnset(n)
}

func (n *ExprCastUnset) GetPosition() *position.Position {
	return n.Position
}

// ExprAssign node
type ExprAssign struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssign) Accept(v Visitor) {
	v.ExprAssign(n)
}

func (n *ExprAssign) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignReference node
type ExprAssignReference struct {
	Position     *position.Position
	Var          Vertex
	EqualTkn     *token.Token
	AmpersandTkn *token.Token
	Expr         Vertex
}

func (n *ExprAssignReference) Accept(v Visitor) {
	v.ExprAssignReference(n)
}

func (n *ExprAssignReference) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignBitwiseAnd node
type ExprAssignBitwiseAnd struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignBitwiseAnd) Accept(v Visitor) {
	v.ExprAssignBitwiseAnd(n)
}

func (n *ExprAssignBitwiseAnd) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignBitwiseOr node
type ExprAssignBitwiseOr struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignBitwiseOr) Accept(v Visitor) {
	v.ExprAssignBitwiseOr(n)
}

func (n *ExprAssignBitwiseOr) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignBitwiseXor node
type ExprAssignBitwiseXor struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignBitwiseXor) Accept(v Visitor) {
	v.ExprAssignBitwiseXor(n)
}

func (n *ExprAssignBitwiseXor) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignCoalesce node
type ExprAssignCoalesce struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignCoalesce) Accept(v Visitor) {
	v.ExprAssignCoalesce(n)
}

func (n *ExprAssignCoalesce) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignConcat node
type ExprAssignConcat struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignConcat) Accept(v Visitor) {
	v.ExprAssignConcat(n)
}

func (n *ExprAssignConcat) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignDiv node
type ExprAssignDiv struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignDiv) Accept(v Visitor) {
	v.ExprAssignDiv(n)
}

func (n *ExprAssignDiv) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignMinus node
type ExprAssignMinus struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignMinus) Accept(v Visitor) {
	v.ExprAssignMinus(n)
}

func (n *ExprAssignMinus) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignMod node
type ExprAssignMod struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignMod) Accept(v Visitor) {
	v.ExprAssignMod(n)
}

func (n *ExprAssignMod) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignMul node
type ExprAssignMul struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignMul) Accept(v Visitor) {
	v.ExprAssignMul(n)
}

func (n *ExprAssignMul) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignPlus node
type ExprAssignPlus struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignPlus) Accept(v Visitor) {
	v.ExprAssignPlus(n)
}

func (n *ExprAssignPlus) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignPow node
type ExprAssignPow struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignPow) Accept(v Visitor) {
	v.ExprAssignPow(n)
}

func (n *ExprAssignPow) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignShiftLeft node
type ExprAssignShiftLeft struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignShiftLeft) Accept(v Visitor) {
	v.ExprAssignShiftLeft(n)
}

func (n *ExprAssignShiftLeft) GetPosition() *position.Position {
	return n.Position
}

// ExprAssignShiftRight node
type ExprAssignShiftRight struct {
	Position *position.Position
	Var      Vertex
	EqualTkn *token.Token
	Expr     Vertex
}

func (n *ExprAssignShiftRight) Accept(v Visitor) {
	v.ExprAssignShiftRight(n)
}

func (n *ExprAssignShiftRight) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryBitwiseAnd node
type ExprBinaryBitwiseAnd struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryBitwiseAnd) Accept(v Visitor) {
	v.ExprBinaryBitwiseAnd(n)
}

func (n *ExprBinaryBitwiseAnd) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryBitwiseOr node
type ExprBinaryBitwiseOr struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryBitwiseOr) Accept(v Visitor) {
	v.ExprBinaryBitwiseOr(n)
}

func (n *ExprBinaryBitwiseOr) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryBitwiseXor node
type ExprBinaryBitwiseXor struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryBitwiseXor) Accept(v Visitor) {
	v.ExprBinaryBitwiseXor(n)
}

func (n *ExprBinaryBitwiseXor) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryBooleanAnd node
type ExprBinaryBooleanAnd struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryBooleanAnd) Accept(v Visitor) {
	v.ExprBinaryBooleanAnd(n)
}

func (n *ExprBinaryBooleanAnd) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryBooleanOr node
type ExprBinaryBooleanOr struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryBooleanOr) Accept(v Visitor) {
	v.ExprBinaryBooleanOr(n)
}

func (n *ExprBinaryBooleanOr) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryCoalesce node
type ExprBinaryCoalesce struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryCoalesce) Accept(v Visitor) {
	v.ExprBinaryCoalesce(n)
}

func (n *ExprBinaryCoalesce) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryConcat node
type ExprBinaryConcat struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryConcat) Accept(v Visitor) {
	v.ExprBinaryConcat(n)
}

func (n *ExprBinaryConcat) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryDiv node
type ExprBinaryDiv struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryDiv) Accept(v Visitor) {
	v.ExprBinaryDiv(n)
}

func (n *ExprBinaryDiv) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryEqual node
type ExprBinaryEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryEqual) Accept(v Visitor) {
	v.ExprBinaryEqual(n)
}

func (n *ExprBinaryEqual) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryGreater node
type ExprBinaryGreater struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryGreater) Accept(v Visitor) {
	v.ExprBinaryGreater(n)
}

func (n *ExprBinaryGreater) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryGreaterOrEqual node
type ExprBinaryGreaterOrEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryGreaterOrEqual) Accept(v Visitor) {
	v.ExprBinaryGreaterOrEqual(n)
}

func (n *ExprBinaryGreaterOrEqual) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryIdentical node
type ExprBinaryIdentical struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryIdentical) Accept(v Visitor) {
	v.ExprBinaryIdentical(n)
}

func (n *ExprBinaryIdentical) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryLogicalAnd node
type ExprBinaryLogicalAnd struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryLogicalAnd) Accept(v Visitor) {
	v.ExprBinaryLogicalAnd(n)
}

func (n *ExprBinaryLogicalAnd) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryLogicalOr node
type ExprBinaryLogicalOr struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryLogicalOr) Accept(v Visitor) {
	v.ExprBinaryLogicalOr(n)
}

func (n *ExprBinaryLogicalOr) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryLogicalXor node
type ExprBinaryLogicalXor struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryLogicalXor) Accept(v Visitor) {
	v.ExprBinaryLogicalXor(n)
}

func (n *ExprBinaryLogicalXor) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryMinus node
type ExprBinaryMinus struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryMinus) Accept(v Visitor) {
	v.ExprBinaryMinus(n)
}

func (n *ExprBinaryMinus) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryMod node
type ExprBinaryMod struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryMod) Accept(v Visitor) {
	v.ExprBinaryMod(n)
}

func (n *ExprBinaryMod) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryMul node
type ExprBinaryMul struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryMul) Accept(v Visitor) {
	v.ExprBinaryMul(n)
}

func (n *ExprBinaryMul) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryNotEqual node
type ExprBinaryNotEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryNotEqual) Accept(v Visitor) {
	v.ExprBinaryNotEqual(n)
}

func (n *ExprBinaryNotEqual) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryNotIdentical node
type ExprBinaryNotIdentical struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryNotIdentical) Accept(v Visitor) {
	v.ExprBinaryNotIdentical(n)
}

func (n *ExprBinaryNotIdentical) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryPlus node
type ExprBinaryPlus struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryPlus) Accept(v Visitor) {
	v.ExprBinaryPlus(n)
}

func (n *ExprBinaryPlus) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryPow node
type ExprBinaryPow struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryPow) Accept(v Visitor) {
	v.ExprBinaryPow(n)
}

func (n *ExprBinaryPow) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryShiftLeft node
type ExprBinaryShiftLeft struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryShiftLeft) Accept(v Visitor) {
	v.ExprBinaryShiftLeft(n)
}

func (n *ExprBinaryShiftLeft) GetPosition() *position.Position {
	return n.Position
}

// ExprBinaryShiftRight node
type ExprBinaryShiftRight struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinaryShiftRight) Accept(v Visitor) {
	v.ExprBinaryShiftRight(n)
}

func (n *ExprBinaryShiftRight) GetPosition() *position.Position {
	return n.Position
}

// ExprBinarySmaller node
type ExprBinarySmaller struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinarySmaller) Accept(v Visitor) {
	v.ExprBinarySmaller(n)
}

func (n *ExprBinarySmaller) GetPosition() *position.Position {
	return n.Position
}

// ExprBinarySmallerOrEqual node
type ExprBinarySmallerOrEqual struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinarySmallerOrEqual) Accept(v Visitor) {
	v.ExprBinarySmallerOrEqual(n)
}

func (n *ExprBinarySmallerOrEqual) GetPosition() *position.Position {
	return n.Position
}

// ExprBinarySpaceship node
type ExprBinarySpaceship struct {
	Position *position.Position
	Left     Vertex
	OpTkn    *token.Token
	Right    Vertex
}

func (n *ExprBinarySpaceship) Accept(v Visitor) {
	v.ExprBinarySpaceship(n)
}

func (n *ExprBinarySpaceship) GetPosition() *position.Position {
	return n.Position
}

type Name struct {
	Position      *position.Position
	Parts         []Vertex
	SeparatorTkns []*token.Token
}

func (n *Name) Accept(v Visitor) {
	v.NameName(n)
}

func (n *Name) GetPosition() *position.Position {
	return n.Position
}

type NameFullyQualified struct {
	Position       *position.Position
	NsSeparatorTkn *token.Token
	Parts          []Vertex
	SeparatorTkns  []*token.Token
}

func (n *NameFullyQualified) Accept(v Visitor) {
	v.NameFullyQualified(n)
}

func (n *NameFullyQualified) GetPosition() *position.Position {
	return n.Position
}

type NameRelative struct {
	Position       *position.Position
	NsTkn          *token.Token
	NsSeparatorTkn *token.Token
	Parts          []Vertex
	SeparatorTkns  []*token.Token
}

func (n *NameRelative) Accept(v Visitor) {
	v.NameRelative(n)
}

func (n *NameRelative) GetPosition() *position.Position {
	return n.Position
}

type NamePart struct {
	Position  *position.Position
	StringTkn *token.Token
	Value     []byte
}

func (n *NamePart) Accept(v Visitor) {
	v.NameNamePart(n)
}

func (n *NamePart) GetPosition() *position.Position {
	return n.Position
}
