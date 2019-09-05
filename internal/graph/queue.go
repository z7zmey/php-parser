package graph

type queueItem struct {
	nodeID NodeID
	depth  int
}

type Queue struct {
	items []queueItem
}

func (s *Queue) Enqueue(id NodeID, depth int) {
	s.items = append(s.items, queueItem{id, depth})
}

func (s *Queue) Dequeue() (NodeID, int) {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return item.nodeID, item.depth
}

func (s *Queue) IsEmpty() bool {
	return len(s.items) == 0
}
