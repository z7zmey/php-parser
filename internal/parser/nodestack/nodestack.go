package nodestack

import (
	"github.com/z7zmey/php-parser/pkg/ast"
)

type NodeStack struct {
	list  []ast.Node
	stack []int
}

func (s *NodeStack) Reset() {
	s.list = s.list[:0]
	s.stack = s.stack[:0]
}

func (s *NodeStack) Add(n ...ast.Node) {
	s.list = append(s.list, n...)
}

func (s *NodeStack) Push(n ...ast.Node) {
	s.stack = append(s.stack, len(s.list))
	s.Add(n...)
}

func (s *NodeStack) Last() *ast.Node {
	return &s.list[len(s.list)-1]
}

func (s *NodeStack) Len() int {
	p := 0
	if len(s.stack) > 0 {
		p = s.stack[len(s.stack)-1]
	}
	return len(s.list[p:])
}

func (s *NodeStack) Pop(groups ...ast.NodeGroup) []ast.Node {
	e := len(s.list)
	p := 0

	for _, g := range groups {
		if len(s.stack) > 0 {
			p = s.stack[len(s.stack)-1]
		}

		for i := p; i < e; i++ {
			s.list[i].Group = g
		}

		if len(s.stack) > 0 {
			s.stack = s.stack[:len(s.stack)-1]
		}

		e = p
	}

	list := s.list[e:]
	s.list = s.list[:e]

	return list
}
