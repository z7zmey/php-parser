package graph

type AdjacencyList struct {
	First EdgeID
	Last  EdgeID
}

type NodeID uint

type Node struct {
	ID      uint
	Type    uint
	AdjList AdjacencyList
}
