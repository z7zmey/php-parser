package parser

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/ast/linear"
)

func TestList(t *testing.T) {
	s := stackedNodeList{}

	s.Push()

	s.Add(linear.NodeID(1))
	s.Add(linear.NodeID(2))
	s.Add(linear.NodeID(3))

	expected := []linear.NodeID{linear.NodeID(1), linear.NodeID(2), linear.NodeID(3)}
	actual := s.Pop()
	assert.DeepEqual(t, expected, actual)

	expected = []linear.NodeID{}
	actual = s.Pop()
	assert.DeepEqual(t, expected, actual)
}

func TestListNested(t *testing.T) {
	s := stackedNodeList{}

	s.Push()

	s.Add(linear.NodeID(1))
	s.Add(linear.NodeID(2))
	s.Add(linear.NodeID(3))

	s.Push()
	s.Add(linear.NodeID(4))
	s.Add(linear.NodeID(5))

	expected := []linear.NodeID{linear.NodeID(4), linear.NodeID(5)}
	actual := s.Pop()
	assert.DeepEqual(t, expected, actual)

	expected = []linear.NodeID{linear.NodeID(1), linear.NodeID(2), linear.NodeID(3)}
	actual = s.Pop()
	assert.DeepEqual(t, expected, actual)

	expected = []linear.NodeID{}
	actual = s.Pop()
	assert.DeepEqual(t, expected, actual)
}
