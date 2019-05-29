package graph

type Visitor interface {
	VisitNode(stxtree *AST, n Node, depth int) bool
}
