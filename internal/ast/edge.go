package ast

import "encoding/json"

type EdgeType uint8

var edgeTypeRegistry []string

func NewEdgeType(name string) EdgeType {
	edgeTypeRegistry = append(edgeTypeRegistry, name)
	return EdgeType(len(edgeTypeRegistry) - 1)
}

func (t EdgeType) String() string {
	return edgeTypeRegistry[t]
}

var EdgeTypeNil = NewEdgeType("nil")

type EdgeID uint32

func (id EdgeID) Idx() int {
	return int(id) - 1
}

type Edge struct {
	id     EdgeID
	next   EdgeID
	Type   EdgeType
	Target uint32
}

type EdgeList struct {
	First EdgeID
	Last  EdgeID
}

type EdgeFilter func(e Edge) bool

type edge struct {
	ID     EdgeID
	Next   EdgeID
	Type   string
	Target uint32
}

func (e Edge) MarshalJSON() ([]byte, error) {
	return json.Marshal(edge{
		ID:     e.id,
		Next:   e.next,
		Type:   e.Type.String(),
		Target: e.Target,
	})
}

// EdgeStorage store edgesEdges
type EdgeStorage []Edge

// Put adds new Edge in storage
func (s *EdgeStorage) Put(e Edge) EdgeID {
	*s = append(*s, e)
	return EdgeID(len(*s))
}

// Get returns Edge by EdgeID
func (s EdgeStorage) Get(id EdgeID) Edge {
	return s[id.Idx()]
}

// Set updates Edge by EdgeID
func (s EdgeStorage) Set(id EdgeID, e Edge) {
	s[id.Idx()] = e
}
