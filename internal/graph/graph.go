package graph

type Graph struct {
	// todo: edges and nodes pool
	nodes []Node
	edges []Edge
}

type LinkPosFunc func(id NodeID) EdgeID

func (g *Graph) Append(id NodeID) EdgeID {
	return g.nodes[int(id)-1].AdjList.Last
}

func (g *Graph) Prepend(id NodeID) EdgeID {
	return 0
}

func (g *Graph) NewNode(id uint, t uint) NodeID {
	g.nodes = append(g.nodes, Node{ID: id, Type: t})
	return NodeID(len(g.nodes))
}

func (g *Graph) GetNode(nodeID NodeID) Node {
	return g.nodes[nodeID-1]
}

func (g *Graph) Link(from NodeID, to NodeID, pos LinkPosFunc) EdgeID {
	if pos == nil {
		pos = g.Append
	}

	e := Edge{
		From: from,
		To:   to,
	}
	g.edges = append(g.edges, e)
	curID := EdgeID(len(g.edges))
	curEdge := &g.edges[curID-1]

	prevID := pos(from)
	if prevID == 0 {
		nextID := g.nodes[from-1].AdjList.First
		if nextID != 0 {
			nextEdge := &g.edges[nextID-1]
			nextEdge.Prev = curID
			curEdge.Next = nextID
		}
	} else {
		prevEdge := &g.edges[prevID-1]
		nextID := prevEdge.Next
		if nextID != 0 {
			nextEdge := &g.edges[nextID-1]
			nextEdge.Prev = curID
			curEdge.Next = nextID
		}

		prevEdge.Next = curID
		curEdge.Prev = prevID
	}

	if curEdge.Prev == 0 {
		g.nodes[from-1].AdjList.First = curID
	}

	if curEdge.Next == 0 {
		g.nodes[from-1].AdjList.Last = curID
	}

	return curID
}

func (g *Graph) Unlink(id EdgeID) {
	curEdge := &g.edges[id-1]
	prevEdgeID := curEdge.Prev
	nextEdgeID := curEdge.Next

	node := &g.nodes[curEdge.From-1]

	if prevEdgeID == 0 {
		node.AdjList.First = curEdge.Next
	} else {
		prevEdge := &g.edges[prevEdgeID-1]
		prevEdge.Next = curEdge.Next
	}

	if nextEdgeID == 0 {
		node.AdjList.Last = curEdge.Prev
	} else {
		nextEdge := &g.edges[nextEdgeID-1]
		nextEdge.Prev = curEdge.Prev
	}

	curEdge.Next = 0
	curEdge.Prev = 0
}

func (g *Graph) Foreach(r NodeID, f func(Edge, Node) bool) {
	edgeID := g.nodes[r-1].AdjList.First
	for {
		if edgeID == 0 {
			break
		}

		edge := g.edges[edgeID-1]
		node := g.nodes[edge.To-1]
		if !f(edge, node) {
			return
		}

		edgeID = g.edges[edgeID-1].Next
	}
}

func (g *Graph) TraverseBFS(r NodeID, f func(Node, int) bool) {
	if f == nil {
		return
	}

	q := Queue{}
	q.Enqueue(r, 0)
	visited := make(map[NodeID]bool)
	for {
		if q.IsEmpty() {
			break
		}
		nodeID, depth := q.Dequeue()
		node := g.nodes[nodeID-1]
		visited[nodeID] = true

		if !f(node, depth) {
			continue
		}

		edgeID := node.AdjList.First
		for {
			if edgeID == 0 {
				break
			}
			n := g.edges[edgeID-1].To
			if !visited[n] {
				q.Enqueue(n, depth+1)
				visited[n] = true
			}

			edgeID = g.edges[edgeID-1].Next
		}
	}
}

func (g *Graph) TraverseDFS(r NodeID, f func(Node, int) bool) {
	if f == nil {
		return
	}

	s := Stack{}
	s.Push(r, 0)
	visited := make(map[NodeID]bool)
	for {
		if s.IsEmpty() {
			break
		}
		nodeID, depth := s.Pop()
		node := g.nodes[nodeID-1]

		visited[nodeID] = true
		if !f(node, depth) {
			continue
		}

		edgeID := node.AdjList.Last
		for {
			if edgeID == 0 {
				break
			}
			n := g.edges[edgeID-1].To
			if !visited[n] {
				s.Push(n, depth+1)
				visited[n] = true
			}

			edgeID = g.edges[edgeID-1].Prev
		}
	}
}
