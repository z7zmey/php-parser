package graph

type stackItem struct {
	nodeID NodeID
	depth  int
}

type Stack struct {
	items []stackItem
}

func (s *Stack) Push(id NodeID, depth int) {
	s.items = append(s.items, stackItem{id, depth})
}

func (s *Stack) Pop() (NodeID, int) {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item.nodeID, item.depth
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
