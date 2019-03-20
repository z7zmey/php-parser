package ast

type AST struct {
	FileData  []byte
	Positions *PositionStorage
	Nodes     *NodeStorage
	Edges     *EdgeStorage
	RootNode  NodeID
}

func (t *AST) Reset() {
	t.FileData = t.FileData[:0]
	t.Positions.Reset()
	t.Nodes.Reset()
	t.Edges.Reset()
	t.RootNode = 0
}
