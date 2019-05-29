package graph

import (
	"github.com/z7zmey/php-parser/ast"
)

type AST struct {
	FileData  []byte
	Positions *PositionStorage
	Nodes     *NodeStorage
	Edges     *EdgeStorage
	Tokens    *TokenStorage
	RootNode  NodeID

	queue []queueItem
}

type queueItem struct {
	id    NodeID
	depth int
}

func (a *AST) Reset() {
	a.FileData = a.FileData[:0]
	a.Nodes.Reset()
	a.Edges.Reset()
	a.Positions.Reset()
	a.Tokens.Reset()
	a.RootNode = 0
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
			edges := stxtree.Edges.Get(node.Edge, EdgeTypeNode)

			for i := len(edges) - 1; i >= 0; i-- {
				stxtree.queue = append(stxtree.queue, queueItem{
					id:    NodeID(edges[i].Target),
					depth: depth,
				})
			}
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

		posEdges := stxtree.Edges.Get(node.Edge, EdgeTypePosition)
		if len(posEdges) > 0 {
			posID := PositionID(posEdges[0].Target)
			stack[depth].Position = stxtree.Positions.Get(posID)
		}

		tknEdges := stxtree.Edges.Get(node.Edge, EdgeTypeToken)
		for _, tknEdge := range tknEdges {
			tokenID := TokenID(tknEdge.Target)

			token := stxtree.Tokens.Get(tokenID)
			tokenPos := stxtree.Positions.Get(token.Pos)

			nestedToken := ast.Token{
				Type:  token.Type,
				Value: string(stxtree.FileData[tokenPos.PS:tokenPos.PE]),
			}

			stack[depth].Tokens[token.Group] = append(stack[depth].Tokens[token.Group], nestedToken)
		}

		if node.Type.Is(ast.NodeClassTypeValue) && len(posEdges) > 0 {
			posID := PositionID(posEdges[0].Target)
			pos := stxtree.Positions.Get(posID)
			stack[depth].Value = string(stxtree.FileData[pos.PS:pos.PE])
		}

		if depth > 0 {
			stack[depth-1].Children[node.Group] = append(stack[depth-1].Children[node.Group], stack[depth])
		}

		depth++
		edges := stxtree.Edges.Get(node.Edge, EdgeTypeNode)

		for i := len(edges) - 1; i >= 0; i-- {
			stxtree.queue = append(stxtree.queue, queueItem{
				id:    NodeID(edges[i].Target),
				depth: depth,
			})
		}

	}

	return stack[0]
}
