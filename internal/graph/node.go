package graph

import (
	"encoding/json"

	"github.com/z7zmey/php-parser/pkg/ast"
)

var EdgeTypeNode = NewEdgeType("node")

type NodeID uint32

type Node struct {
	Type  ast.NodeType
	Flag  ast.NodeFlag
	Group ast.NodeGroup

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

// Create saves new Node in store
func (b *NodeStorage) Create(n Node) NodeID {
	*b = append(*b, n)
	return NodeID(len(*b))
}

// Save modified Node
func (b NodeStorage) Save(id NodeID, n Node) {
	b[id-1] = n
}

// Get returns Node by NodeID
func (b NodeStorage) Get(id NodeID) Node {
	return b[id-1]
}
