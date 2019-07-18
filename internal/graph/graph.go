package graph

import (
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/traverser"
)

type Graph struct {
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

func (g *Graph) Reset() {
	g.FileData = g.FileData[:0]
	g.Nodes = g.Nodes[:0]
	g.Edges = g.Edges[:0]
	g.Positions = g.Positions[:0]
	g.Tokens = g.Tokens[:0]
	g.RootNode = 0
}

func (g *Graph) Link(nodeID NodeID, edgeType EdgeType, target uint32) {
	edge := Edge{
		Type:   edgeType,
		Target: target,
	}

	edgeID := g.Edges.Put(edge)

	nodeEdges := &g.Nodes[nodeID-1].Edges

	if nodeEdges.First == 0 {
		nodeEdges.First = edgeID
		nodeEdges.Last = edgeID
	} else {
		g.Edges[nodeEdges.Last-1].next = edgeID
		nodeEdges.Last = edgeID
	}
}

func (g *Graph) AppendEdges(src EdgeList, edges EdgeList) EdgeList {
	if edges.First == 0 {
		return src
	}

	if src.First == 0 {
		return edges
	}

	g.Edges[src.Last-1].next = edges.First
	src.Last = edges.Last

	return src
}

func (g *Graph) RemoveEdges(nodeID NodeID, f EdgeFilter) EdgeList {
	nodeEdges := &g.Nodes[nodeID-1].Edges

	var removedEdges EdgeList
	var prevEdgeID EdgeID

	edgeID := nodeEdges.First
	for edgeID != 0 {
		edge := &g.Edges[edgeID-1]

		if f(*edge) {
			if prevEdgeID == 0 {
				nodeEdges.First = edge.next
			} else {
				g.Edges[prevEdgeID-1].next = edge.next
			}

			removedEdges = g.AppendEdges(removedEdges, EdgeList{edgeID, edgeID})
		} else {
			prevEdgeID = edgeID
		}

		edgeID = edge.next
	}

	return removedEdges
}

func (g *Graph) EachEdge(edges EdgeList, callback func(e Edge) bool) {
	edgeID := edges.First
	for edgeID != 0 {
		edge := g.Edges[edgeID-1]

		if callback(edge) {
			return
		}

		edgeID = edge.next
	}
}

func (g *Graph) Traverse(v traverser.Visitor) {
	g.queue = g.queue[:0]
	g.queue = append(g.queue, queueItem{
		id:    g.RootNode,
		depth: 0,
	})

	for {
		if len(g.queue) == 0 {
			break
		}

		item := g.queue[len(g.queue)-1]
		g.queue = g.queue[:len(g.queue)-1]

		graphNode := g.Nodes.Get(item.id)
		depth := item.depth

		visitChild := v.VisitNode(graphNode.SimpleNode, depth)

		g.EachEdge(graphNode.Edges, func(e Edge) bool {
			if e.Type == EdgeTypeToken {
				tokenID := TokenID(e.Target)

				token := g.Tokens.Get(tokenID)
				tokenPos := g.Positions.Get(token.Pos)

				astToken := ast.Token{
					Type:  token.Type,
					Group: token.Group,
					Value: string(g.FileData[tokenPos.PS:tokenPos.PE]),
				}

				visitChild = v.VisitToken(astToken, depth)
			}

			if e.Type == EdgeTypeToken {
				posID := PositionID(e.Target)

				visitChild = v.VisitPosition(g.Positions.Get(posID), depth)
			}

			return false
		})

		if visitChild {
			depth++
			g.EachEdge(graphNode.Edges, func(e Edge) bool {
				if e.Type != EdgeTypeNode {
					return false
				}

				g.queue = append(g.queue, queueItem{
					id:    NodeID(e.Target),
					depth: depth,
				})

				return false
			})
		}

	}
}
