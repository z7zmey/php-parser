package parser

import (
	"github.com/z7zmey/php-parser/internal/graph"
)

type stackedNodeList struct {
	list  []graph.NodeID
	stack []int
}

func (s *stackedNodeList) Reset() {
	s.list = s.list[:0]
	s.stack = s.stack[:0]
}

func (s *stackedNodeList) Add(n graph.NodeID) {
	s.list = append(s.list, n)
}

func (s *stackedNodeList) Push() {
	s.stack = append(s.stack, len(s.list))
}

func (s *stackedNodeList) Last() graph.NodeID {
	return s.list[len(s.list)-1]
}

func (s *stackedNodeList) Len() int {
	p := 0
	if len(s.stack) > 0 {
		p = s.stack[len(s.stack)-1]
	}
	return len(s.list[p:])
}

func (s *stackedNodeList) Pop() []graph.NodeID {
	p := 0
	if len(s.stack) > 0 {
		p = s.stack[len(s.stack)-1]
	}
	list := s.list[p:]

	s.list = s.list[:p]
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}

	return list
}
