package graph

import (
	"github.com/z7zmey/php-parser/pkg/ast"
)

type AST struct {
	FileData  []byte
	Positions PositionStorage
	Nodes     NodeStorage
	Edges     EdgeStorage
	Tokens    TokenStorage
	RootNode  NodeID

	queue []queueItem
}

type queueItem struct {
	id    NodeID
	depth int
}

func (a *AST) Reset() {
	a.FileData = a.FileData[:0]
	a.Nodes = a.Nodes[:0]
	a.Edges = a.Edges[:0]
	a.Positions = a.Positions[:0]
	a.Tokens = a.Tokens[:0]
	a.RootNode = 0
}

func (a *AST) Link(nodeID NodeID, edgeType EdgeType, target uint32) {
	edge := Edge{
		Type:   edgeType,
		Target: target,
	}

	edgeID := a.Edges.Put(edge)

	nodeEdges := &a.Nodes[nodeID-1].Edges

	if nodeEdges.First == 0 {
		nodeEdges.First = edgeID
		nodeEdges.Last = edgeID
	} else {
		a.Edges[nodeEdges.Last-1].next = edgeID
		nodeEdges.Last = edgeID
	}
}

func (a *AST) LinkFront(nodeID NodeID, edgeType EdgeType, target uint32) {
	edge := Edge{
		Type:   edgeType,
		Target: target,
	}

	edgeID := a.Edges.Put(edge)

	nodeEdges := &a.Nodes[nodeID-1].Edges

	if nodeEdges.First == 0 {
		nodeEdges.First = edgeID
		nodeEdges.Last = edgeID
	} else {
		a.Edges[edgeID-1].next = nodeEdges.First
		nodeEdges.First = edgeID
	}
}

func (a *AST) RemoveEdges(nodeID NodeID, f EdgeFilter) EdgeList {
	nodeEdges := &a.Nodes[nodeID-1].Edges

	var removedEdges EdgeList
	var prevEdgeID EdgeID

	edgeID := nodeEdges.First
	for edgeID != 0 {
		edge := &a.Edges[edgeID-1]

		if f(*edge) {
			if prevEdgeID == 0 {
				nodeEdges.First = edge.next
			} else {
				a.Edges[prevEdgeID-1].next = edge.next
			}

			removedEdges = a.AppendEdges(removedEdges, EdgeList{edgeID, edgeID})
		} else {
			prevEdgeID = edgeID
		}

		edgeID = edge.next
	}

	return removedEdges
}

func (a *AST) AppendEdges(src EdgeList, edges EdgeList) EdgeList {
	if edges.First == 0 {
		return src
	}

	if src.First == 0 {
		return edges
	}

	a.Edges[src.Last-1].next = edges.First
	src.Last = edges.Last

	return src
}

func (a *AST) AppendNodeEdges(nodeID NodeID, edges EdgeList) {
	nodeEdges := a.Nodes[nodeID-1].Edges
	a.Nodes[nodeID-1].Edges = a.AppendEdges(nodeEdges, edges)
}

func (a *AST) PrependEdges(src EdgeList, edges EdgeList) EdgeList {
	if edges.First == 0 {
		return src
	}

	if src.First == 0 {
		return edges
	}

	a.Edges[edges.Last-1].next = src.First
	src.First = edges.First

	return src
}

func (a *AST) EachEdge(edges EdgeList, callback func(e Edge) bool) {
	edgeID := edges.First
	for edgeID != 0 {
		edge := a.Edges[edgeID-1]

		if callback(edge) {
			return
		}

		edgeID = edge.next
	}
}

func (a *AST) PrependNodeEdges(nodeID NodeID, edges EdgeList) {
	nodeEdges := a.Nodes[nodeID-1].Edges
	a.Nodes[nodeID-1].Edges = a.PrependEdges(nodeEdges, edges)
}

func (stxtree *AST) Traverse(v Visitor) {
	stxtree.queue = stxtree.queue[:0]
	stxtree.queue = append(stxtree.queue, queueItem{
		id:    stxtree.RootNode,
		depth: 0,
	})

	for {
		if len(stxtree.queue) == 0 {
			break
		}

		item := stxtree.queue[len(stxtree.queue)-1]
		stxtree.queue = stxtree.queue[:len(stxtree.queue)-1]

		node := stxtree.Nodes.Get(item.id)
		depth := item.depth

		visitChild := v.VisitNode(stxtree, node, depth)

		if visitChild {
			depth++
			stxtree.EachEdge(node.Edges, func(e Edge) bool {
				if e.Type != EdgeTypeNode {
					return false
				}

				stxtree.queue = append(stxtree.queue, queueItem{
					id:    NodeID(e.Target),
					depth: depth,
				})

				return false
			})
		}

	}
}

func (stxtree *AST) Nested() ast.Node {
	stack := []ast.Node{}

	stxtree.queue = stxtree.queue[:0]
	stxtree.queue = append(stxtree.queue, queueItem{
		id:    stxtree.RootNode,
		depth: 0,
	})

	for {
		if len(stxtree.queue) == 0 {
			break
		}

		item := stxtree.queue[len(stxtree.queue)-1]
		stxtree.queue = stxtree.queue[:len(stxtree.queue)-1]

		node := stxtree.Nodes.Get(item.id)
		depth := item.depth

		if len(stack) <= depth+1 {
			stack = append(stack, ast.Node{})
		}

		stack[depth] = ast.Node{
			Type:     node.Type,
			Flags:    node.Flag,
			Tokens:   make(map[ast.TokenGroup][]ast.Token),
			Children: make(map[ast.NodeGroup][]ast.Node),
		}

		var posID PositionID
		stxtree.EachEdge(node.Edges, func(e Edge) bool {
			if e.Type != EdgeTypePosition {
				return false
			}

			posID = PositionID(e.Target)
			stack[depth].Position = stxtree.Positions.Get(posID)

			return true
		})

		stxtree.EachEdge(node.Edges, func(e Edge) bool {
			if e.Type != EdgeTypeToken {
				return false
			}

			tokenID := TokenID(e.Target)

			token := stxtree.Tokens.Get(tokenID)
			tokenPos := stxtree.Positions.Get(token.Pos)

			nestedToken := ast.Token{
				Type:  token.Type,
				Value: string(stxtree.FileData[tokenPos.PS:tokenPos.PE]),
			}

			stack[depth].Tokens[token.Group] = append(stack[depth].Tokens[token.Group], nestedToken)

			return false
		})

		if node.Type.Is(ast.NodeClassTypeValue) && posID > 0 {
			pos := stxtree.Positions.Get(posID)
			stack[depth].Value = string(stxtree.FileData[pos.PS:pos.PE])
		}

		if depth > 0 {
			stack[depth-1].Children[node.Group] = append(stack[depth-1].Children[node.Group], stack[depth])
		}

		depth++

		stxtree.EachEdge(node.Edges, func(e Edge) bool {
			if e.Type != EdgeTypeNode {
				return false
			}

			stxtree.queue = append(stxtree.queue, queueItem{
				id:    NodeID(e.Target),
				depth: depth,
			})

			return false
		})

	}

	return stack[0]
}
