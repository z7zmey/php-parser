package linear

import "github.com/z7zmey/php-parser/ast"

type NodeID uint32

type Node struct {
	Type ast.NodeType
	Flag ast.NodeFlag

	Parent NodeID
	Child  NodeID
	Next   NodeID

	Key ast.EdgeType

	Pos PositionID
}

// NodeStorage store nodes
type NodeStorage struct {
	buf []Node
}

// NewNodeStorage creates new NodeStorage
func NewNodeStorage(buf []Node) *NodeStorage {
	return &NodeStorage{buf}
}

// Reset storage
func (b *NodeStorage) Reset() {
	b.buf = b.buf[:0]
}

// Create saves new Node in store
func (b *NodeStorage) Create(n Node) NodeID {
	b.buf = append(b.buf, n)
	return NodeID(len(b.buf))
}

// Save modified Node
func (b *NodeStorage) Save(id NodeID, n Node) {
	b.buf[id-1] = n
}

// Get returns Node by NodeID
func (b NodeStorage) Get(id NodeID) Node {
	return b.buf[id-1]
}

// GetAll returns all Nodes
func (b NodeStorage) GetAll() []Node {
	return b.buf
}
