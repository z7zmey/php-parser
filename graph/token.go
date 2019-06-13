package graph

import (
	"encoding/json"

	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/scanner"
)

var EdgeTypeToken = NewEdgeType("token")

type TokenID int32

type Token struct {
	Type  scanner.TokenType
	Group ast.TokenGroup
	Next  TokenID
	Pos   PositionID
}

type token struct {
	Type  string
	Group string
	Next  TokenID
	Pos   PositionID
}

func (t Token) MarshalJSON() ([]byte, error) {
	out := token{
		Type:  t.Type.String(),
		Group: t.Group.String(),
		Next:  t.Next,
		Pos:   t.Pos,
	}

	return json.Marshal(out)
}

type TokenStorage []Token

// Create saves new Node in store
func (b *TokenStorage) Create(s Token) TokenID {
	*b = append(*b, s)
	return TokenID(len(*b))
}

// Save modified Node
func (b TokenStorage) Save(id TokenID, s Token) {
	b[id-1] = s
}

// Get returns Node by NodeID
func (b TokenStorage) Get(id TokenID) Token {
	return b[id-1]
}
