package linear

import "github.com/z7zmey/php-parser/ast"

type TokenID int32

type Token struct {
	Type  ast.TokenType
	Group ast.TokenGroup
	Next  TokenID
	Pos   PositionID
}

type TokenStorage struct {
	buf []Token
}

// NewTokenStorage creates new TokenStorage
func NewTokenStorage(buf []Token) *TokenStorage {
	return &TokenStorage{buf}
}

// Reset storage
func (b *TokenStorage) Reset() {
	b.buf = b.buf[:0]
}

// Create saves new Node in store
func (b *TokenStorage) Create(s Token) TokenID {
	b.buf = append(b.buf, s)
	return TokenID(len(b.buf))
}

// Save modified Node
func (b *TokenStorage) Save(id TokenID, s Token) {
	b.buf[id-1] = s
}

// Get returns Node by NodeID
func (b TokenStorage) Get(id TokenID) Token {
	return b.buf[id-1]
}

// GetAll returns all Nodes
func (b TokenStorage) GetAll() []Token {
	return b.buf
}
