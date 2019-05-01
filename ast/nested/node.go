package nested

import (
	"encoding/json"

	"github.com/z7zmey/php-parser/ast"
)

type Node struct {
	Type     ast.NodeType
	Flags    ast.NodeFlag
	Position ast.Position
	Children map[ast.EdgeType][]Node
	Value    string
}

type node struct {
	Type     string            `json:"type"`
	Flags    []string          `json:"flags"`
	Value    string            `json:"value"`
	Position ast.Position      `json:"position"`
	Children map[string][]Node `json:"children"`
}

func (n Node) MarshalJSON() ([]byte, error) {
	children := map[string][]Node{}

	for k, v := range n.Children {
		children[k.String()] = v
	}

	out := node{
		Type:     n.Type.String(),
		Flags:    n.Flags.GetFlagNames(),
		Value:    n.Value,
		Position: n.Position,
		Children: children,
	}

	return json.Marshal(out)
}
