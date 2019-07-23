package graph

import (
	"encoding/json"

	"github.com/z7zmey/php-parser/pkg/ast"
)

var EdgeTypeNode = NewEdgeType("node")

type NodeID uint32

type Node struct {
	ast.SimpleNode

	Edges EdgeList
}

type node struct {
	Type  string
	Group string

	Edges EdgeList
}

func (n Node) MarshalJSON() ([]byte, error) {
	out := node{
		Type:  n.Type.String(),
		Group: n.Group.String(),
		Edges: n.Edges,
	}

	return json.Marshal(out)
}

// NodeStorage store nodes
type NodeStorage []Node

// Put saves new Node in store
func (b *NodeStorage) Put(n Node) NodeID {
	*b = append(*b, n)
	return NodeID(len(*b))
}

// Get returns Node by NodeID
func (b NodeStorage) Get(id NodeID) Node {
	return b[id-1]
}
