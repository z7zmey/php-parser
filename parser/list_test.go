package parser

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/graph"
)

func TestList(t *testing.T) {
	s := stackedNodeList{}

	s.Push()

	s.Add(graph.NodeID(1))
	s.Add(graph.NodeID(2))
	s.Add(graph.NodeID(3))

	expected := []graph.NodeID{graph.NodeID(1), graph.NodeID(2), graph.NodeID(3)}
	actual := s.Pop()
	assert.DeepEqual(t, expected, actual)

	expected = []graph.NodeID{}
	actual = s.Pop()
	assert.DeepEqual(t, expected, actual)
}

func TestListNested(t *testing.T) {
	s := stackedNodeList{}

	s.Push()

	s.Add(graph.NodeID(1))
	s.Add(graph.NodeID(2))
	s.Add(graph.NodeID(3))

	s.Push()
	s.Add(graph.NodeID(4))
	s.Add(graph.NodeID(5))

	expected := []graph.NodeID{graph.NodeID(4), graph.NodeID(5)}
	actual := s.Pop()
	assert.DeepEqual(t, expected, actual)

	expected = []graph.NodeID{graph.NodeID(1), graph.NodeID(2), graph.NodeID(3)}
	actual = s.Pop()
	assert.DeepEqual(t, expected, actual)

	expected = []graph.NodeID{}
	actual = s.Pop()
	assert.DeepEqual(t, expected, actual)
}
