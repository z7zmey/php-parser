package ast

type PositionID uint32

// Position stores range of bytes and lines in a file
// ps - positionStart, pe - positionEnd, ls - lineStart, le - lineEnd
type Position struct {
	PS, PE, LS, LE int
}

// PositionStorage stores positions
type PositionStorage struct {
	buf []Position
}

// NewPositionStorage creates new position buffer
func NewPositionStorage(buf []Position) *PositionStorage {
	return &PositionStorage{buf}
}

func (b *PositionStorage) Reset() {
	b.buf = b.buf[:0]
}

// Create saves new Position in store
func (b *PositionStorage) Create(n Position) PositionID {
	b.buf = append(b.buf, n)
	return PositionID(len(b.buf))
}

// Save modified Position
func (b *PositionStorage) Save(id PositionID, n Position) {
	b.buf[id-1] = n
}

// Get returns position by PositionID
func (b PositionStorage) Get(id PositionID) Position {
	return b.buf[id-1]
}

// GetAll returns all Positions
func (b PositionStorage) GetAll() []Position {
	return b.buf
}
