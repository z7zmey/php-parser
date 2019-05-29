package graph

import "github.com/z7zmey/php-parser/ast"

var EdgeTypePosition = NewEdgeType("position")

type PositionID uint32

// PositionStorage stores positions
type PositionStorage struct {
	buf []ast.Position
}

// NewPositionStorage creates new position buffer
func NewPositionStorage(buf []ast.Position) *PositionStorage {
	return &PositionStorage{buf}
}

func (b *PositionStorage) Reset() {
	b.buf = b.buf[:0]
}

// Create saves new Position in store
func (b *PositionStorage) Create(n ast.Position) PositionID {
	b.buf = append(b.buf, n)
	return PositionID(len(b.buf))
}

// Save modified Position
func (b *PositionStorage) Save(id PositionID, n ast.Position) {
	b.buf[id-1] = n
}

// Get returns position by PositionID
func (b PositionStorage) Get(id PositionID) ast.Position {
	return b.buf[id-1]
}

// GetAll returns all Positions
func (b PositionStorage) GetAll() []ast.Position {
	return b.buf
}
