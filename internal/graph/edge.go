package graph

type EdgeID uint

type Edge struct {
	From NodeID
	To   NodeID
	Prev EdgeID
	Next EdgeID
}
