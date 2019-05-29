package graph

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
	Type   EdgeType
	Next   EdgeID
	Target uint32
}

type edge struct {
	Type   string
	Next   EdgeID
	Target uint32
}

func (e Edge) MarshalJSON() ([]byte, error) {
	out := edge{
		Type:   e.Type.String(),
		Next:   e.Next,
		Target: e.Target,
	}

	return json.Marshal(out)
}

// EdgeStorage store edges
type EdgeStorage struct {
	edges []Edge
	buf   []Edge
}

// Reset storage
func (s *EdgeStorage) Reset() {
	s.edges = s.edges[:0]
}

// Put new Edge in storage
func (s *EdgeStorage) Put(e Edge) EdgeID {
	s.edges = append(s.edges, e)
	return EdgeID(len(s.edges))
}

// Set Edge by EdgeID
func (s *EdgeStorage) Set(id EdgeID, e Edge) {
	s.edges[id.Idx()] = e
}

// GetOne Edge by EdgeID
func (s *EdgeStorage) GetOne(id EdgeID) Edge {
	return s.edges[id.Idx()]
}

// Get Edges by EdgeID
func (s *EdgeStorage) Get(id EdgeID, t EdgeType) []Edge {
	s.buf = s.buf[:0]
	if id == 0 {
		return s.buf
	}

	edge := s.edges[id.Idx()]
	if t == EdgeTypeNil || t == edge.Type {
		s.buf = append(s.buf, edge)
	}

	for {
		if edge.Next == 0 {
			break
		}

		id = edge.Next
		edge = s.edges[id.Idx()]
		if t == EdgeTypeNil || t == edge.Type {
			s.buf = append(s.buf, edge)
		}
	}

	return s.buf
}

func (s EdgeStorage) GetLastID(id EdgeID) EdgeID {
	if id == 0 {
		return 0
	}

	for {
		edge := s.edges[id.Idx()]

		if edge.Next == 0 {
			break
		}

		id = edge.Next
	}

	return id
}

func (s EdgeStorage) GetAll() []Edge {
	return s.edges
}
