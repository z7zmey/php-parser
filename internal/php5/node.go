package php5

import (
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
)

type ParserBrackets struct {
	Position        *position.Position
	OpenBracketTkn  *token.Token
	Child           ast.Vertex
	CloseBracketTkn *token.Token
}

func (n *ParserBrackets) Accept(v ast.Visitor) {
	// do nothing
}

func (n *ParserBrackets) GetPosition() *position.Position {
	return n.Position
}

type ParserSeparatedList struct {
	Position      *position.Position
	Items         []ast.Vertex
	SeparatorTkns []*token.Token
}

func (n *ParserSeparatedList) Accept(v ast.Visitor) {
	// do nothing
}

func (n *ParserSeparatedList) GetPosition() *position.Position {
	return n.Position
}

// TraitAdaptationList node
type TraitAdaptationList struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Adaptations          []ast.Vertex
	CloseCurlyBracketTkn *token.Token
}

func (n *TraitAdaptationList) Accept(v ast.Visitor) {
	// do nothing
}

func (n *TraitAdaptationList) GetPosition() *position.Position {
	return n.Position
}

// ArgumentList node
type ArgumentList struct {
	Position            *position.Position
	OpenParenthesisTkn  *token.Token
	Arguments           []ast.Vertex
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

func (n *ArgumentList) Accept(v ast.Visitor) {
	// do nothing
}

func (n *ArgumentList) GetPosition() *position.Position {
	return n.Position
}

// TraitMethodRef node
type TraitMethodRef struct {
	Position       *position.Position
	Trait          ast.Vertex
	DoubleColonTkn *token.Token
	Method         ast.Vertex
}

func (n *TraitMethodRef) Accept(v ast.Visitor) {
	// do nothing
}

func (n *TraitMethodRef) GetPosition() *position.Position {
	return n.Position
}
