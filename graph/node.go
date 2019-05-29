package graph

import (
	"encoding/json"

	"github.com/z7zmey/php-parser/ast"
)

var EdgeTypeNode = NewEdgeType("node")

type NodeID uint32

type Node struct {
	Type  ast.NodeType
	Flag  ast.NodeFlag
	Group ast.NodeGroup

	Edge EdgeID
}

type node struct {
	Type  string
	Group string

	Edge EdgeID
}

func (n Node) MarshalJSON() ([]byte, error) {
	out := node{
		Type:  n.Type.String(),
		Group: n.Group.String(),
		Edge:  n.Edge,
	}

	return json.Marshal(out)
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
