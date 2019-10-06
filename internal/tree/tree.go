package tree

import (
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/traverser"
)

type Tree struct {
	nodes  []ast.Node
	tokens []ast.Token
}

func NewTree(size int) *Tree {
	return &Tree{
		nodes:  make([]ast.Node, 0, size),
		tokens: make([]ast.Token, 0, size),
	}
}

func (t *Tree) Reset() {
	t.nodes = t.nodes[:0]
	t.tokens = t.tokens[:0]
}

func (t *Tree) RootNode() *ast.Node {
	return &t.nodes[len(t.nodes)-1]
}

func (t *Tree) AddNodes(nodes ...[]ast.Node) []ast.Node {
	off := len(t.nodes)
	for _, nn := range nodes {
		t.nodes = append(t.nodes, nn...)
	}

	return t.nodes[off:]
}

func (t *Tree) AppendNodes(l int, nodes ...[]ast.Node) []ast.Node {
	off := len(t.tokens) - l

	for _, nn := range nodes {
		t.nodes = append(t.nodes, nn...)
	}

	return t.nodes[off:]
}

func (t *Tree) AddTokens(tokens ...[]ast.Token) []ast.Token {
	off := len(t.tokens)
	for _, tt := range tokens {
		t.tokens = append(t.tokens, tt...)
	}

	return t.tokens[off:]
}

func (t *Tree) AppendTokens(l int, tokens ...[]ast.Token) []ast.Token {
	off := len(t.tokens) - l

	for _, tt := range tokens {
		t.tokens = append(t.tokens, tt...)
	}

	return t.tokens[off:]
}

func (t *Tree) Traverse(v traverser.Visitor) {
	s := Stack{}
	s.Push(t.RootNode(), 0)
	for {
		if s.IsEmpty() {
			break
		}
		node, depth := s.Pop()

		if !v.VisitNode(*node, depth) {
			continue
		}

		for i := len(node.Children) - 1; i >= 0; i-- {
			n := &node.Children[i]
			s.Push(n, depth+1)
		}
	}
}
