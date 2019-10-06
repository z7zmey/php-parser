package tree

import "github.com/z7zmey/php-parser/pkg/ast"

type stackItem struct {
	node  *ast.Node
	depth int
}

type Stack struct {
	items []stackItem
}

func (s *Stack) Push(n *ast.Node, depth int) {
	s.items = append(s.items, stackItem{n, depth})
}

func (s *Stack) Pop() (*ast.Node, int) {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item.node, item.depth
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
