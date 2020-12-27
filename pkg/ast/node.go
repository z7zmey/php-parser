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

func (n *Root) Accept(v NodeVisitor) {
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

func (n *Nullable) Accept(v NodeVisitor) {
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

func (n *Parameter) Accept(v NodeVisitor) {
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

func (n *Identifier) Accept(v NodeVisitor) {
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

func (n *Argument) Accept(v NodeVisitor) {
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

func (n *ScalarDnumber) Accept(v NodeVisitor) {
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

func (n *ScalarEncapsed) Accept(v NodeVisitor) {
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

func (n *ScalarEncapsedStringPart) Accept(v NodeVisitor) {
	v.ScalarEncapsedStringPart(n)
}

func (n *ScalarEncapsedStringPart) GetPosition() *position.Position {
	return n.Position
}

// ScalarHeredoc node
type ScalarHeredoc struct {
	Position        *position.Position
	OpenHeredocTkn  *token.Token
	Parts           []Vertex
	CloseHeredocTkn *token.Token
}

func (n *ScalarHeredoc) Accept(v NodeVisitor) {
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

func (n *ScalarLnumber) Accept(v NodeVisitor) {
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

func (n *ScalarMagicConstant) Accept(v NodeVisitor) {
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

func (n *ScalarString) Accept(v NodeVisitor) {
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

func (n *StmtBreak) Accept(v NodeVisitor) {
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

func (n *StmtCase) Accept(v NodeVisitor) {
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

func (n *StmtCatch) Accept(v NodeVisitor) {
	v.StmtCatch(n)
}

func (n *StmtCatch) GetPosition() *position.Position {
	return n.Position
}

// StmtClass node
type StmtClass struct {
	Position             *position.Position
	Modifiers            []Vertex
	ClassTkn             *token.Token
	ClassName            Vertex
	OpenParenthesisTkn   *token.Token
	Arguments            []Vertex
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
	Extends              Vertex
	Implements           Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtClass) Accept(v NodeVisitor) {
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

func (n *StmtClassConstList) Accept(v NodeVisitor) {
	v.StmtClassConstList(n)
}

func (n *StmtClassConstList) GetPosition() *position.Position {
	return n.Position
}

// StmtClassExtends node
type StmtClassExtends struct {
	Position  *position.Position
	ExtendTkn *token.Token
	ClassName Vertex
}

func (n *StmtClassExtends) Accept(v NodeVisitor) {
	v.StmtClassExtends(n)
}

func (n *StmtClassExtends) GetPosition() *position.Position {
	return n.Position
}

// StmtClassImplements node
type StmtClassImplements struct {
	Position       *position.Position
	ImplementsTkn  *token.Token
	InterfaceNames []Vertex
	SeparatorTkns  []*token.Token
}

func (n *StmtClassImplements) Accept(v NodeVisitor) {
	v.StmtClassImplements(n)
}

func (n *StmtClassImplements) GetPosition() *position.Position {
	return n.Position
}

// StmtClassMethod node
type StmtClassMethod struct {
	Position            *position.Position
	Modifiers           []Vertex
	FunctionTkn         *token.Token
	AmpersandTkn        *token.Token
	MethodName          Vertex
	OpenParenthesisTkn  *token.Token
	Params              []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	ReturnType          Vertex
	Stmt                Vertex
}

func (n *StmtClassMethod) Accept(v NodeVisitor) {
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

func (n *StmtConstList) Accept(v NodeVisitor) {
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

func (n *StmtConstant) Accept(v NodeVisitor) {
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

func (n *StmtContinue) Accept(v NodeVisitor) {
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

func (n *StmtDeclare) Accept(v NodeVisitor) {
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

func (n *StmtDefault) Accept(v NodeVisitor) {
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

func (n *StmtDo) Accept(v NodeVisitor) {
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

func (n *StmtEcho) Accept(v NodeVisitor) {
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

func (n *StmtElse) Accept(v NodeVisitor) {
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

func (n *StmtElseIf) Accept(v NodeVisitor) {
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

func (n *StmtExpression) Accept(v NodeVisitor) {
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

func (n *StmtFinally) Accept(v NodeVisitor) {
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

func (n *StmtFor) Accept(v NodeVisitor) {
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

func (n *StmtForeach) Accept(v NodeVisitor) {
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
	FunctionName         Vertex
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

func (n *StmtFunction) Accept(v NodeVisitor) {
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

func (n *StmtGlobal) Accept(v NodeVisitor) {
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

func (n *StmtGoto) Accept(v NodeVisitor) {
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

func (n *StmtHaltCompiler) Accept(v NodeVisitor) {
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

func (n *StmtIf) Accept(v NodeVisitor) {
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

func (n *StmtInlineHtml) Accept(v NodeVisitor) {
	v.StmtInlineHtml(n)
}

func (n *StmtInlineHtml) GetPosition() *position.Position {
	return n.Position
}

// StmtInterface node
type StmtInterface struct {
	Position             *position.Position
	InterfaceTkn         *token.Token
	InterfaceName        Vertex
	Extends              Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtInterface) Accept(v NodeVisitor) {
	v.StmtInterface(n)
}

func (n *StmtInterface) GetPosition() *position.Position {
	return n.Position
}

// StmtInterfaceExtends node
type StmtInterfaceExtends struct {
	Position       *position.Position
	ExtendsTkn     *token.Token
	InterfaceNames []Vertex
	SeparatorTkns  []*token.Token
}

func (n *StmtInterfaceExtends) Accept(v NodeVisitor) {
	v.StmtInterfaceExtends(n)
}

func (n *StmtInterfaceExtends) GetPosition() *position.Position {
	return n.Position
}

// StmtLabel node
type StmtLabel struct {
	Position  *position.Position
	LabelName Vertex
	ColonTkn  *token.Token
}

func (n *StmtLabel) Accept(v NodeVisitor) {
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

func (n *StmtNamespace) Accept(v NodeVisitor) {
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

func (n *StmtNop) Accept(v NodeVisitor) {
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

func (n *StmtProperty) Accept(v NodeVisitor) {
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
	Properties    []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtPropertyList) Accept(v NodeVisitor) {
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

func (n *StmtReturn) Accept(v NodeVisitor) {
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

func (n *StmtStatic) Accept(v NodeVisitor) {
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

func (n *StmtStaticVar) Accept(v NodeVisitor) {
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

func (n *StmtStmtList) Accept(v NodeVisitor) {
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
	CaseList             []Vertex
	CloseCurlyBracketTkn *token.Token
	EndSwitchTkn         *token.Token
	SemiColonTkn         *token.Token
}

func (n *StmtSwitch) Accept(v NodeVisitor) {
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

func (n *StmtThrow) Accept(v NodeVisitor) {
	v.StmtThrow(n)
}

func (n *StmtThrow) GetPosition() *position.Position {
	return n.Position
}

// StmtTrait node
type StmtTrait struct {
	Position             *position.Position
	TraitTkn             *token.Token
	TraitName            Vertex
	Extends              Vertex
	Implements           Vertex
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *StmtTrait) Accept(v NodeVisitor) {
	v.StmtTrait(n)
}

func (n *StmtTrait) GetPosition() *position.Position {
	return n.Position
}

// StmtTraitMethodRef node
type StmtTraitMethodRef struct {
	Position       *position.Position
	Trait          Vertex
	DoubleColonTkn *token.Token
	Method         Vertex
}

func (n *StmtTraitMethodRef) Accept(v NodeVisitor) {
	v.StmtTraitMethodRef(n)
}

func (n *StmtTraitMethodRef) GetPosition() *position.Position {
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

func (n *StmtTraitUse) Accept(v NodeVisitor) {
	v.StmtTraitUse(n)
}

func (n *StmtTraitUse) GetPosition() *position.Position {
	return n.Position
}

// StmtTraitUseAlias node
type StmtTraitUseAlias struct {
	Position     *position.Position
	Ref          Vertex
	AsTkn        *token.Token
	Modifier     Vertex
	Alias        Vertex
	SemiColonTkn *token.Token
}

func (n *StmtTraitUseAlias) Accept(v NodeVisitor) {
	v.StmtTraitUseAlias(n)
}

func (n *StmtTraitUseAlias) GetPosition() *position.Position {
	return n.Position
}

// StmtTraitUsePrecedence node
type StmtTraitUsePrecedence struct {
	Position      *position.Position
	Ref           Vertex
	InsteadofTkn  *token.Token
	Insteadof     []Vertex
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

func (n *StmtTraitUsePrecedence) Accept(v NodeVisitor) {
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

func (n *StmtTry) Accept(v NodeVisitor) {
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

func (n *StmtUnset) Accept(v NodeVisitor) {
	v.StmtUnset(n)
}

func (n *StmtUnset) GetPosition() *position.Position {
	return n.Position
}

// StmtUse node
type StmtUse struct {
	Position        *position.Position
	UseTkn          *token.Token
	Type            Vertex
	UseDeclarations []Vertex
	SeparatorTkns   []*token.Token
	SemiColonTkn    *token.Token
}

func (n *StmtUse) Accept(v NodeVisitor) {
	v.StmtUse(n)
}

func (n *StmtUse) GetPosition() *position.Position {
	return n.Position
}

// StmtGroupUse node
type StmtGroupUse struct {
	Position              *position.Position
	UseTkn                *token.Token
	Type                  Vertex
	LeadingNsSeparatorTkn *token.Token
	Prefix                Vertex
	NsSeparatorTkn        *token.Token
	OpenCurlyBracketTkn   *token.Token
	UseDeclarations       []Vertex
	SeparatorTkns         []*token.Token
	CloseCurlyBracketTkn  *token.Token
	SemiColonTkn          *token.Token
}

func (n *StmtGroupUse) Accept(v NodeVisitor) {
	v.StmtGroupUse(n)
}

func (n *StmtGroupUse) GetPosition() *position.Position {
	return n.Position
}

// StmtUseDeclaration node
type StmtUseDeclaration struct {
	Position       *position.Position
	Type           Vertex
	NsSeparatorTkn *token.Token
	Use            Vertex
	AsTkn          *token.Token
	Alias          Vertex
}

func (n *StmtUseDeclaration) Accept(v NodeVisitor) {
	v.StmtUseDeclaration(n)
}

func (n *StmtUseDeclaration) GetPosition() *position.Position {
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

func (n *StmtWhile) Accept(v NodeVisitor) {
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

func (n *ExprArray) Accept(v NodeVisitor) {
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

func (n *ExprArrayDimFetch) Accept(v NodeVisitor) {
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

func (n *ExprArrayItem) Accept(v NodeVisitor) {
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

func (n *ExprArrowFunction) Accept(v NodeVisitor) {
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

func (n *ExprBitwiseNot) Accept(v NodeVisitor) {
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

func (n *ExprBooleanNot) Accept(v NodeVisitor) {
	v.ExprBooleanNot(n)
}

func (n *ExprBooleanNot) GetPosition() *position.Position {
	return n.Position
}

// ExprClassConstFetch node
type ExprClassConstFetch struct {
	Position       *position.Position
	Class          Vertex
	DoubleColonTkn *token.Token
	ConstantName   Vertex
}

func (n *ExprClassConstFetch) Accept(v NodeVisitor) {
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

func (n *ExprClone) Accept(v NodeVisitor) {
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
	Use                    []Vertex
	UseSeparatorTkns       []*token.Token
	UseCloseParenthesisTkn *token.Token
	ColonTkn               *token.Token
	ReturnType             Vertex
	OpenCurlyBracketTkn    *token.Token
	Stmts                  []Vertex
	CloseCurlyBracketTkn   *token.Token
}

func (n *ExprClosure) Accept(v NodeVisitor) {
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

func (n *ExprClosureUse) Accept(v NodeVisitor) {
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

func (n *ExprConstFetch) Accept(v NodeVisitor) {
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

func (n *ExprEmpty) Accept(v NodeVisitor) {
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

func (n *ExprErrorSuppress) Accept(v NodeVisitor) {
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

func (n *ExprEval) Accept(v NodeVisitor) {
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

func (n *ExprExit) Accept(v NodeVisitor) {
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
	Arguments           []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ExprFunctionCall) Accept(v NodeVisitor) {
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

func (n *ExprInclude) Accept(v NodeVisitor) {
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

func (n *ExprIncludeOnce) Accept(v NodeVisitor) {
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

func (n *ExprInstanceOf) Accept(v NodeVisitor) {
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

func (n *ExprIsset) Accept(v NodeVisitor) {
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

func (n *ExprList) Accept(v NodeVisitor) {
	v.ExprList(n)
}

func (n *ExprList) GetPosition() *position.Position {
	return n.Position
}

// ExprMethodCall node
type ExprMethodCall struct {
	Position            *position.Position
	Var                 Vertex
	ObjectOperatorTkn   *token.Token
	Method              Vertex
	OpenParenthesisTkn  *token.Token
	Arguments           []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ExprMethodCall) Accept(v NodeVisitor) {
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
	Arguments           []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ExprNew) Accept(v NodeVisitor) {
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

func (n *ExprPostDec) Accept(v NodeVisitor) {
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

func (n *ExprPostInc) Accept(v NodeVisitor) {
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

func (n *ExprPreDec) Accept(v NodeVisitor) {
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

func (n *ExprPreInc) Accept(v NodeVisitor) {
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

func (n *ExprPrint) Accept(v NodeVisitor) {
	v.ExprPrint(n)
}

func (n *ExprPrint) GetPosition() *position.Position {
	return n.Position
}

// ExprPropertyFetch node
type ExprPropertyFetch struct {
	Position          *position.Position
	Var               Vertex
	ObjectOperatorTkn *token.Token
	Property          Vertex
}

func (n *ExprPropertyFetch) Accept(v NodeVisitor) {
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

func (n *ExprRequire) Accept(v NodeVisitor) {
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

func (n *ExprRequireOnce) Accept(v NodeVisitor) {
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

func (n *ExprShellExec) Accept(v NodeVisitor) {
	v.ExprShellExec(n)
}

func (n *ExprShellExec) GetPosition() *position.Position {
	return n.Position
}

// ExprStaticCall node
type ExprStaticCall struct {
	Position            *position.Position
	Class               Vertex
	DoubleColonTkn      *token.Token
	Call                Vertex
	OpenParenthesisTkn  *token.Token
	Arguments           []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ExprStaticCall) Accept(v NodeVisitor) {
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
	Property       Vertex
}

func (n *ExprStaticPropertyFetch) Accept(v NodeVisitor) {
	v.ExprStaticPropertyFetch(n)
}

func (n *ExprStaticPropertyFetch) GetPosition() *position.Position {
	return n.Position
}

// ExprTernary node
type ExprTernary struct {
	Position    *position.Position
	Condition   Vertex
	QuestionTkn *token.Token
	IfTrue      Vertex
	ColonTkn    *token.Token
	IfFalse     Vertex
}

func (n *ExprTernary) Accept(v NodeVisitor) {
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

func (n *ExprUnaryMinus) Accept(v NodeVisitor) {
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

func (n *ExprUnaryPlus) Accept(v NodeVisitor) {
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
	VarName              Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *ExprVariable) Accept(v NodeVisitor) {
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
	Value          Vertex
}

func (n *ExprYield) Accept(v NodeVisitor) {
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

func (n *ExprYieldFrom) Accept(v NodeVisitor) {
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

func (n *ExprCastArray) Accept(v NodeVisitor) {
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

func (n *ExprCastBool) Accept(v NodeVisitor) {
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

func (n *ExprCastDouble) Accept(v NodeVisitor) {
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

func (n *ExprCastInt) Accept(v NodeVisitor) {
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

func (n *ExprCastObject) Accept(v NodeVisitor) {
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

func (n *ExprCastString) Accept(v NodeVisitor) {
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

func (n *ExprCastUnset) Accept(v NodeVisitor) {
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

func (n *ExprAssign) Accept(v NodeVisitor) {
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

func (n *ExprAssignReference) Accept(v NodeVisitor) {
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

func (n *ExprAssignBitwiseAnd) Accept(v NodeVisitor) {
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

func (n *ExprAssignBitwiseOr) Accept(v NodeVisitor) {
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

func (n *ExprAssignBitwiseXor) Accept(v NodeVisitor) {
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

func (n *ExprAssignCoalesce) Accept(v NodeVisitor) {
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

func (n *ExprAssignConcat) Accept(v NodeVisitor) {
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

func (n *ExprAssignDiv) Accept(v NodeVisitor) {
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

func (n *ExprAssignMinus) Accept(v NodeVisitor) {
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

func (n *ExprAssignMod) Accept(v NodeVisitor) {
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

func (n *ExprAssignMul) Accept(v NodeVisitor) {
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

func (n *ExprAssignPlus) Accept(v NodeVisitor) {
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

func (n *ExprAssignPow) Accept(v NodeVisitor) {
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

func (n *ExprAssignShiftLeft) Accept(v NodeVisitor) {
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

func (n *ExprAssignShiftRight) Accept(v NodeVisitor) {
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

func (n *ExprBinaryBitwiseAnd) Accept(v NodeVisitor) {
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

func (n *ExprBinaryBitwiseOr) Accept(v NodeVisitor) {
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

func (n *ExprBinaryBitwiseXor) Accept(v NodeVisitor) {
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

func (n *ExprBinaryBooleanAnd) Accept(v NodeVisitor) {
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

func (n *ExprBinaryBooleanOr) Accept(v NodeVisitor) {
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

func (n *ExprBinaryCoalesce) Accept(v NodeVisitor) {
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

func (n *ExprBinaryConcat) Accept(v NodeVisitor) {
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

func (n *ExprBinaryDiv) Accept(v NodeVisitor) {
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

func (n *ExprBinaryEqual) Accept(v NodeVisitor) {
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

func (n *ExprBinaryGreater) Accept(v NodeVisitor) {
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

func (n *ExprBinaryGreaterOrEqual) Accept(v NodeVisitor) {
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

func (n *ExprBinaryIdentical) Accept(v NodeVisitor) {
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

func (n *ExprBinaryLogicalAnd) Accept(v NodeVisitor) {
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

func (n *ExprBinaryLogicalOr) Accept(v NodeVisitor) {
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

func (n *ExprBinaryLogicalXor) Accept(v NodeVisitor) {
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

func (n *ExprBinaryMinus) Accept(v NodeVisitor) {
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

func (n *ExprBinaryMod) Accept(v NodeVisitor) {
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

func (n *ExprBinaryMul) Accept(v NodeVisitor) {
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

func (n *ExprBinaryNotEqual) Accept(v NodeVisitor) {
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

func (n *ExprBinaryNotIdentical) Accept(v NodeVisitor) {
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

func (n *ExprBinaryPlus) Accept(v NodeVisitor) {
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

func (n *ExprBinaryPow) Accept(v NodeVisitor) {
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

func (n *ExprBinaryShiftLeft) Accept(v NodeVisitor) {
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

func (n *ExprBinaryShiftRight) Accept(v NodeVisitor) {
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

func (n *ExprBinarySmaller) Accept(v NodeVisitor) {
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

func (n *ExprBinarySmallerOrEqual) Accept(v NodeVisitor) {
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

func (n *ExprBinarySpaceship) Accept(v NodeVisitor) {
	v.ExprBinarySpaceship(n)
}

func (n *ExprBinarySpaceship) GetPosition() *position.Position {
	return n.Position
}

type NameName struct {
	Position      *position.Position
	Parts         []Vertex
	SeparatorTkns []*token.Token
}

func (n *NameName) Accept(v NodeVisitor) {
	v.NameName(n)
}

func (n *NameName) GetPosition() *position.Position {
	return n.Position
}

type NameFullyQualified struct {
	Position       *position.Position
	NsSeparatorTkn *token.Token
	Parts          []Vertex
	SeparatorTkns  []*token.Token
}

func (n *NameFullyQualified) Accept(v NodeVisitor) {
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

func (n *NameRelative) Accept(v NodeVisitor) {
	v.NameRelative(n)
}

func (n *NameRelative) GetPosition() *position.Position {
	return n.Position
}

type NameNamePart struct {
	Position  *position.Position
	StringTkn *token.Token
	Value     []byte
}

func (n *NameNamePart) Accept(v NodeVisitor) {
	v.NameNamePart(n)
}

func (n *NameNamePart) GetPosition() *position.Position {
	return n.Position
}

// TODO: move to private section

type ParserBrackets struct {
	Position        *position.Position
	OpenBracketTkn  *token.Token
	Child           Vertex
	CloseBracketTkn *token.Token
}

func (n *ParserBrackets) Accept(v NodeVisitor) {
	v.ParserBrackets(n)
}

func (n *ParserBrackets) GetPosition() *position.Position {
	return n.Position
}

type ParserSeparatedList struct {
	Position      *position.Position
	Items         []Vertex
	SeparatorTkns []*token.Token
}

func (n *ParserSeparatedList) Accept(v NodeVisitor) {
	// do nothing
}

func (n *ParserSeparatedList) GetPosition() *position.Position {
	return n.Position
}

// TraitAdaptationList node
type TraitAdaptationList struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Adaptations          []Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *TraitAdaptationList) Accept(v NodeVisitor) {
	// do nothing
}

func (n *TraitAdaptationList) GetPosition() *position.Position {
	return n.Position
}

// ArgumentList node
type ArgumentList struct {
	Position            *position.Position
	OpenParenthesisTkn  *token.Token
	Arguments           []Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ArgumentList) Accept(v NodeVisitor) {
	// do nothing
}

func (n *ArgumentList) GetPosition() *position.Position {
	return n.Position
}

type ReturnType struct {
	Position *position.Position
	ColonTkn *token.Token
	Type     Vertex
}

func (n *ReturnType) Accept(v NodeVisitor) {
	// do nothing
}

func (n *ReturnType) GetPosition() *position.Position {
	return n.Position
}
