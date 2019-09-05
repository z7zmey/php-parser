package stxtree

import (
	"github.com/z7zmey/php-parser/internal/graph"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/traverser"
)

const (
	NodeTypeNode uint = iota
	NodeTypePosition
	NodeTypeToken
)

type Graph struct {
	graph.Graph

	FileData  []byte
	Nodes     []ast.SimpleNode
	Positions []ast.Position
	Tokens    []ast.Token

	RootNode graph.NodeID
}

func (g *Graph) Reset() {
	g.FileData = g.FileData[:0]
	g.Nodes = g.Nodes[:0]
	g.Positions = g.Positions[:0]
	g.Tokens = g.Tokens[:0]
	g.RootNode = 0
}

func (g *Graph) NewNode(n ast.SimpleNode) graph.NodeID {
	g.Nodes = append(g.Nodes, n)
	id := uint(len(g.Nodes) - 1)
	return g.Graph.NewNode(id, NodeTypeNode)
}

func (g *Graph) GetNode(nodeID graph.NodeID) *ast.SimpleNode {
	node := g.Graph.GetNode(nodeID)
	return &g.Nodes[node.ID]
}

func (g *Graph) NewPosition(n ast.Position) graph.NodeID {
	g.Positions = append(g.Positions, n)
	id := uint(len(g.Positions) - 1)
	return g.Graph.NewNode(id, NodeTypePosition)
}

func (g *Graph) GetPosition(nodeID graph.NodeID) *ast.Position {
	node := g.Graph.GetNode(nodeID)
	return &g.Positions[node.ID]
}

func (g *Graph) NewToken(n ast.Token) graph.NodeID {
	g.Tokens = append(g.Tokens, n)
	id := uint(len(g.Tokens) - 1)
	return g.Graph.NewNode(id, NodeTypeNode)
}

func (g *Graph) GetToken(nodeID graph.NodeID) *ast.Token {
	node := g.Graph.GetNode(nodeID)
	return &g.Tokens[node.ID]
}

func (g *Graph) Link(from graph.NodeID, to graph.NodeID, pos graph.LinkPosFunc) graph.EdgeID {
	return g.Graph.Link(from, to, pos)
}

func (g *Graph) Traverse(v traverser.Visitor) {
	g.Graph.TraverseDFS(g.RootNode, func(n graph.Node, depth int) bool {
		switch n.Type {
		case NodeTypeNode:
			v.VisitNode(g.Nodes[n.ID], depth)
		case NodeTypePosition:
			v.VisitPosition(g.Positions[n.ID], depth)
		case NodeTypeToken:
			v.VisitToken(g.Tokens[n.ID], depth)
		}

		return true
	})
}
